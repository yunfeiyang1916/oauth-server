package http

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	"github.com/yunfeiyang1916/oauth-server/api/code"
	"github.com/yunfeiyang1916/oauth-server/model"

	"github.com/gorilla/mux"

	"github.com/yunfeiyang1916/oauth-server/service"

	"github.com/yunfeiyang1916/oauth-server/endpoint"

	kithttp "github.com/go-kit/kit/transport/http"
)

var (
	options      []kithttp.ServerOption
	eptContainer *endpoint.Container
)

func encodeJsonResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(model.GetResp(code.ErrSystem, nil))
}

func newHttpHandler(srv *service.Container) http.Handler {
	eptContainer = endpoint.New(srv)
	r := mux.NewRouter()
	kitLog := log.NewLogfmtLogger(os.Stderr)
	kitLog = log.With(kitLog, "ts", log.DefaultTimestampUTC)
	kitLog = log.With(kitLog, "caller", log.DefaultCaller)
	options = []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(kitLog)),
		kithttp.ServerErrorEncoder(encodeError),
	}
	r.Methods(http.MethodGet).Path("/api/v1/oauth/client/get").Handler(getClient())
	return r
}
