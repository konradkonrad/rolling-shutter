package retry

import (
	"context"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func multDuration(a time.Duration, m float64) time.Duration {
	return time.Duration(float64(a.Nanoseconds()) * m)
}

type retrier struct {
	clock            clock.Clock
	numRetries       int
	infiniteRetries  bool
	interval         time.Duration
	maxInterval      time.Duration
	cancelingErrors  []error
	multiplier       float64
	executionContext string
	identifier       string
}

func newRetrier() *retrier {
	return &retrier{
		clock:           clock.New(),
		numRetries:      3,
		interval:        2 * time.Second,
		maxInterval:     60 * time.Second,
		multiplier:      1.,
		cancelingErrors: []error{},
	}
}

func (r *retrier) option(opts []Option) {
	for _, opt := range opts {
		opt(r)
	}
}

func (r *retrier) logWithContext(e *zerolog.Event) *zerolog.Event {
	if r.identifier != "" {
		e = e.Str("id", r.identifier)
	}

	if r.executionContext != "" {
		e = e.Str("context", r.executionContext)
	}
	return e
}

func (r *retrier) logError(err error, msg string) {
	e := log.Debug().Err(err)
	r.logWithContext(e).Msg(msg)
}

func (r *retrier) iterator(next <-chan time.Time) <-chan time.Time {
	iter := make(chan time.Time, 1)
	go func() {
		defer close(iter)
		interval := r.interval
		// emit the time once, this is the initial action
		// (e.g. the initial function call)
		iter <- r.clock.Now()
		i := 0
		// next receives information when the last
		// action (e.g. function call) was executed
		for lastExecutionFinished := range next {
			if i >= r.numRetries && !r.infiniteRetries {
				return
			}
			nextTick := r.clock.Until(lastExecutionFinished.Add(interval))
			// emit the time, this is a consecutive retry event
			iter <- <-r.clock.After(nextTick)
			i++
			if float64(interval) >= float64(r.maxInterval)/r.multiplier {
				interval = r.maxInterval
			} else {
				interval = multDuration(interval, r.multiplier)
			}
		}
	}()
	return iter
}

type (
	Option                   func(*retrier)
	RetriableFunction[T any] func(ctx context.Context) (T, error)
)

// FunctionCall calls the given function multiple times until it doesn't return an error
// or one of any optional, user-defined specific errors is returned.
func FunctionCall[T any](ctx context.Context, fn RetriableFunction[T], opts ...Option) (T, error) {
	retrier := newRetrier()
	retrier.option(opts)

	next := make(chan time.Time, 1)
	defer close(next)

	retry := retrier.iterator(next)

	var err error
	var null T

	callCount := 0

	for {
		select {
		case _, ok := <-retry:
			if !ok {
				retrier.logWithContext(log.Debug()).Msg("retry limit reached")
				return null, err
			}
			var result T
			if callCount > 0 {
				retrier.logWithContext(log.Debug().Int("count", callCount)).Msg("retrying")
			}
			start := retrier.clock.Now()
			result, err = fn(ctx)
			stopped := retrier.clock.Now()

			retrier.logWithContext(
				log.Debug().TimeDiff("took (ms)", time.Now(), start),
			).Msg("called retriable function")
			callCount++
			if err == nil {
				return result, nil
			}
			if ctx.Err() != nil {
				return null, ctx.Err()
			}
			for _, cErr := range retrier.cancelingErrors {
				if err == cErr {
					retrier.logError(err, "request errored")
					return null, err
				}
			}
			retrier.logError(err, "request errored")
			next <- stopped
		case <-ctx.Done():
			return null, ctx.Err()
		}
	}
}
