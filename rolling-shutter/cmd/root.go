// Package cmd implements the shuttermint subcommands
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/shutter-network/rolling-shutter/rolling-shutter/cmd/bootstrap"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/cmd/chain"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/cmd/collator"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/cmd/completion"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/cmd/cryptocmd"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/cmd/keyper"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/cmd/mocknode"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/cmd/mocksequencer"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/cmd/proxy"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/cmd/shversion"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/cmd/snapshot"
	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley"
)

var (
	logFormatArg  string
	logFormatName string = "logformat"
	logLevelArg   string
	logLevelName  string = "loglevel"
)

func errorForFlag(flg *flag.Flag) error {
	if flg != nil {
		return errors.Errorf("failed to parse '%s' option. usage: '%s'", flg.Name, flg.Usage)
	}
	return errors.Errorf("'%s' option not specified", logFormatName)
}

func configureCaller(l zerolog.Logger, short bool) zerolog.Logger {
	if short {
		pathsep := string(os.PathSeparator)
		// default is long filename
		zerolog.CallerMarshalFunc = func(_ uintptr, file string, line int) string {
			return fmt.Sprintf("%s:%d", file[1+strings.LastIndex(file, pathsep):], line)
		}
	}
	return l.With().Caller().Logger()
}

func configureTime(l zerolog.Logger) zerolog.Logger {
	zerolog.TimeFieldFormat = "2006/01/02 15:04:05.000000"
	return l.With().Timestamp().Logger()
}

func setupLogging(cmd *cobra.Command) (zerolog.Logger, error) {
	// create a basic logger with stdout writer
	// we will change the writer later

	l := zerolog.New(os.Stdout)
	exclude := []string{}
	// change the "message" field name, so that
	// it doesn't collide with e.g. logging of
	// shutter "message"
	zerolog.MessageFieldName = "log"

	logFormatFlag := cmd.PersistentFlags().Lookup(logFormatName)
	if logFormatFlag == nil {
		// this should not happen due to user error,
		// since the flag should have a default value attached
		return l, errors.Errorf("flag '%s' not found", logFormatName)
	}
	switch logFormatFlag.Value.String() {
	case "max", "long":
		l = configureTime(l)
		l = configureCaller(l, true)
	case "short":
		// no time/date logging
		l = configureCaller(l, true)
		exclude = []string{
			zerolog.TimestampFieldName,
		}
	case "min":
		// no time/date logging
		// no caller logging
		exclude = []string{
			zerolog.TimestampFieldName,
			zerolog.CallerFieldName,
		}
	default:
		return l, errorForFlag(logFormatFlag)
	}

	logLevelFlag := cmd.PersistentFlags().Lookup(logLevelName)
	if logFormatFlag == nil {
		// this should not happen due to user error,
		// since the flag should have a default value attached
		return l, errors.Errorf("flag '%s' not found", logLevelName)
	}
	switch logLevelFlag.Value.String() {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		return l, errorForFlag(logLevelFlag)
	}

	// reset the writer
	l = l.Output(zerolog.ConsoleWriter{
		NoColor:    true,
		Out:        os.Stderr,
		TimeFormat: zerolog.TimeFieldFormat,
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.LevelFieldName,
			zerolog.CallerFieldName,
			zerolog.MessageFieldName,
		},
		PartsExclude: exclude,
	})
	return l, nil
}

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rolling-shutter",
		Short:   "A collection of commands to run and interact with Rolling Shutter nodes",
		Version: shversion.Version(),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := medley.BindFlags(cmd, "ROLLING_SHUTTER")
			if err != nil {
				return err
			}

			logger, err := setupLogging(cmd.Root())
			if err != nil {
				return errors.Wrap(err, "failed to setup logging")
			}
			log.Logger = logger
			return nil
		},
	}
	cmd.PersistentFlags().StringVar(
		&logFormatArg,
		logFormatName,
		"long",
		"set log format, possible values:  min, short, long, max",
	)
	cmd.PersistentFlags().StringVar(
		&logLevelArg,
		logLevelName,
		"info",
		"set log level, possible values:  warn, info, debug",
	)
	cmd.AddCommand(bootstrap.Cmd())
	cmd.AddCommand(chain.Cmd())
	cmd.AddCommand(collator.Cmd())
	cmd.AddCommand(completion.Cmd())
	cmd.AddCommand(keyper.Cmd())
	cmd.AddCommand(mocknode.Cmd())
	cmd.AddCommand(snapshot.Cmd())
	cmd.AddCommand(cryptocmd.Cmd())
	cmd.AddCommand(proxy.Cmd())
	cmd.AddCommand(mocksequencer.Cmd())
	return cmd
}
