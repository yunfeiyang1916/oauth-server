package http

import (
	"context"
	"encoding/json"
	"net/http"

	httpserver "github.com/yunfeiyang1916/toolkit/framework/http/server"

	"github.com/go-kit/kit/transport"

	"github.com/go-kit/kit/log"
	"github.com/yunfeiyang1916/oauth-server/api/code"
	"github.com/yunfeiyang1916/oauth-server/model"
	"github.com/yunfeiyang1916/toolkit/logging"

	"github.com/gorilla/mux"

	"github.com/yunfeiyang1916/oauth-server/service"

	"github.com/yunfeiyang1916/oauth-server/endpoint"

	kithttp "github.com/go-kit/kit/transport/http"
)

var (
	options      []kithttp.ServerOption
	eptContainer *endpoint.Container
	srv          *service.Container
)

// 成功json响应
func encodeJsonResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// 编码错误时的响应
func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(model.GetResp(code.ErrParam, nil))
}
func errorHandler(ctx context.Context, err error) {
	logging.Error("err", err)
}

func newHttpHandler(srv *service.Container) http.Handler {
	eptContainer = endpoint.New(srv)
	r := mux.NewRouter()
	kitLog := log.NewLogfmtLogger(logging.GetLogger().Writer())
	kitLog = log.With(kitLog, "ts", log.DefaultTimestampUTC)
	kitLog = log.With(kitLog, "caller", log.DefaultCaller)

	options = []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.ErrorHandlerFunc(errorHandler)),
		kithttp.ServerErrorEncoder(errorEncoder),
	}
	r.Methods(http.MethodGet).Path("/api/v1/oauth/client/get").Handler(getClient())
	return r
}

func initRouter(s httpserver.Server) {
	s.GET("/api/v1/oauth/client/get", qywxGetCallback)
}

// 企业微信get回调
func qywxGetCallback(c *httpserver.Context) {
	resp := srv.ClientService.Get(c.Ctx, model.ClientReq{})
	c.Raw(resp, resp.Code)
}
