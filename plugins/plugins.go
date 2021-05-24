package plugins

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// CheckTokenMiddleware 校验令牌中间件
func CheckTokenMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// return model.GetApiResult(code.ErrParam), nil
			return next(ctx, request)
		}
	}
}

// CheckAuthMiddleware 校验权限中间件
func CheckAuthMiddleware(authority string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return next(ctx, request)
		}
	}
}

// AccessMiddleware 增加访问日志及panic恢复处理中间件
func AccessMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			//defer func() {
			//	cc := ctx.Value(iCtxKey).(*Context)
			//	rc := recover()
			//	if rc != nil {
			//		logging.CrashLogf("recover panic info:%q,raw request method:%s,uri:%s\n", rc, cc.Request.Method, cc.Request.URL.String())
			//	}
			//	if s.options.logger.A() == nil { // logging disable
			//		if rc != nil {
			//			panic(rc)
			//		}
			//		return
			//	}
			//	code := cc.Response.Status()
			//	if err := flow.Err(); err != nil {
			//		code = ecode.ConvertHttpStatus(err)
			//	}
			//
			//	logItems := []interface{}{
			//		"start", cc.startTime.Format(utils.TimeFormat),
			//		"cost", math.Ceil(float64(time.Since(cc.startTime).Nanoseconds()) / 1e6),
			//		"trace_id", cc.traceId,
			//		"peer_name", cc.Peer,
			//		"req_method", cc.Request.Method,
			//		"req_uri", cc.Request.URL.String(),
			//		"real_ip", getRemoteIP(cc.Request),
			//		"http_code", code,
			//		"busi_code", cc.BusiCode(),
			//		"namespace", cc.Namespace,
			//	}
			//
			//	_ = cc.ForeachBaggage(func(key, val string) error {
			//		logItems = append(logItems, key[len(utils.DaeBaggageHeaderPrefix):], val)
			//		return nil
			//	})
			//
			//	if flow.Err() != nil {
			//		logItems = append(logItems, "error", fmt.Sprintf("%q", flow.Err().Error()))
			//	}
			//
			//	// request body 全局打印开关与单个uri接口开关
			//	if !s.options.reqBodyLogOff && cc.printReqBody {
			//		if _, ok := cc.loggingExtra[internalReqBodyLogTag]; !ok {
			//			logItems = append(logItems, "req_body", fmt.Sprintf("%q", cc.bodyBuff.Bytes()))
			//		}
			//	}
			//	// response body
			//	if cc.printRespBody {
			//		if _, ok := cc.loggingExtra[internalRespBodyLogTag]; !ok {
			//			logItems = append(logItems, "resp_body", fmt.Sprintf("%q", cc.Response.ByteBody()))
			//		}
			//	}
			//
			//	if len(cc.loggingExtra) > 0 {
			//		extraList := make([]interface{}, 0)
			//		for k, v := range cc.loggingExtra {
			//			extraList = append(extraList, k, v)
			//		}
			//		if len(extraList) > 0 {
			//			logItems = append(logItems, extraList...)
			//		}
			//	}
			//	s.options.logger.A().Debugw("httpserver", logItems...)
			//	if rc != nil {
			//		panic(rc)
			//	}
			//}()
			return next(ctx, request)
		}
	}
}
