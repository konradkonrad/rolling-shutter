package snpjrpc

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bitwurx/jrpc2"
	"github.com/pkg/errors"

	"github.com/shutter-network/rolling-shutter/rolling-shutter/medley/service"
)

type SnpJRPC struct {
	Server *jrpc2.Server

	getDecryptionKeyCallback func(ctx context.Context, epochID []byte) error
	requestEonKeyCallback    func(ctx context.Context) error
}

type HexEncodedByteArray []byte

type GetDecryptionKeyParams struct {
	EonID   *uint64              `json:"eon_id,string"`
	EpochID *HexEncodedByteArray `json:"proposal"`
}

func (b HexEncodedByteArray) MarshalJSON() ([]byte, error) {
	hexString := hex.EncodeToString(b)
	return json.Marshal(hexString)
}

func (b *HexEncodedByteArray) UnmarshalJSON(data []byte) (err error) {
	var hexString string
	if err = json.Unmarshal(data, &hexString); err != nil {
		return
	}
	*b, err = hex.DecodeString(hexString)
	return
}

func (gdkp *GetDecryptionKeyParams) FromPositional(params []interface{}) error {
	if len(params) != 2 {
		return errors.Errorf("Two parameters required")
	}
	eonID, err := strconv.ParseUint(params[0].(string), 10, 64)
	if err != nil {
		return err
	}
	var epochID HexEncodedByteArray
	epochID, err = hex.DecodeString(params[1].(string))
	if err != nil {
		return err
	}
	gdkp.EonID = &eonID
	gdkp.EpochID = &epochID

	return nil
}

func (snpjrpc *SnpJRPC) GetDecryptionKey(ctx context.Context, params json.RawMessage) (
	interface{},
	*jrpc2.ErrorObject,
) {
	gdkParams := new(GetDecryptionKeyParams)
	if err := jrpc2.ParseParams(params, gdkParams); err != nil {
		return nil, err
	}

	if gdkParams.EonID == nil || gdkParams.EpochID == nil {
		return nil, &jrpc2.ErrorObject{
			Code:    jrpc2.InvalidParamsCode,
			Message: jrpc2.InvalidParamsMsg,
			Data:    "Two parameters required",
		}
	}

	err := snpjrpc.getDecryptionKeyCallback(ctx, *gdkParams.EpochID)
	if err != nil {
		return nil, &jrpc2.ErrorObject{
			Code:    jrpc2.InternalErrorCode,
			Message: jrpc2.InternalErrorMsg,
			Data: fmt.Sprintf(
				"Error requesting decryption key for proposal %s: %v",
				*gdkParams.EpochID,
				err,
			),
		}
	}

	return true, nil
}

func (snpjrpc *SnpJRPC) RequestEonKey(ctx context.Context, _ json.RawMessage) (
	interface{},
	*jrpc2.ErrorObject,
) {
	err := snpjrpc.requestEonKeyCallback(ctx)
	if err != nil {
		return nil, &jrpc2.ErrorObject{
			Code:    jrpc2.InternalErrorCode,
			Message: jrpc2.InternalErrorMsg,
			Data: fmt.Sprintf(
				"Error requesting eon key %v",
				err,
			),
		}
	}
	// FIXME: is this right?
	return true, nil
}

func New(
	jsonrpcHost string,
	jsonrpcPort uint16,
	getDecryptionKeyCallback func(ctx context.Context, epochID []byte) error,
	requestEonKeyCallback func(ctx context.Context) error,
) service.Service {
	host := fmt.Sprintf("%s:%d", jsonrpcHost, jsonrpcPort)
	server := jrpc2.NewServer(host, "/api/v1/rpc", nil)

	jrpc := SnpJRPC{
		Server: server,

		getDecryptionKeyCallback: getDecryptionKeyCallback,
		requestEonKeyCallback:    requestEonKeyCallback,
	}

	server.RegisterWithContext(
		"get_decryption_key",
		jrpc2.MethodWithContext{Method: jrpc.GetDecryptionKey},
	)
	server.RegisterWithContext(
		"request_eon_key",
		jrpc2.MethodWithContext{Method: jrpc.RequestEonKey},
	)

	return &jrpc
}

func (snpjrpc *SnpJRPC) Start(_ context.Context, _ service.Runner) error {
	// FIXME: this is probably not properly hooked into the service.Runner model
	snpjrpc.Server.Start()
	return nil
}
