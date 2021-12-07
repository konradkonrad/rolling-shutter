package collator

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/sha3"

	"github.com/shutter-network/shutter/shuttermint/collator/cltrdb"
	"github.com/shutter-network/shutter/shuttermint/collator/oapi"
	"github.com/shutter-network/shutter/shuttermint/shdb"
)

type server struct {
	c *collator
}

func sendError(w http.ResponseWriter, code int, message string) {
	e := oapi.Error{
		Code:    int32(code),
		Message: message,
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(e)
}

func (srv *server) Ping(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("pong"))
}

func (srv *server) GetNextEpoch(w http.ResponseWriter, req *http.Request) {
	epoch, err := getNextEpochID(req.Context(), srv.c.db)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(oapi.NextEpoch{
		Id: shdb.EncodeUint64(epoch),
	})
}

func (srv *server) SubmitTransaction(w http.ResponseWriter, r *http.Request) {
	var x oapi.SubmitTransactionJSONBody
	if err := json.NewDecoder(r.Body).Decode(&x); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid format for SubmitTransaction")
		return
	}
	ctx := r.Context()

	hash := sha3.New256()
	fmt.Fprintf(hash, "%d\n", len(x.Epoch))
	hash.Write(x.Epoch)
	hash.Write(x.EncryptedTx)
	txid := hash.Sum(nil)

	err := srv.c.db.InsertTx(ctx, cltrdb.InsertTxParams{
		TxID:        txid,
		EpochID:     x.Epoch,
		EncryptedTx: x.EncryptedTx,
	})
	if err != nil {
		sendError(w, http.StatusConflict, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(oapi.TransactionId{Id: txid})
}