package middleware

import (
	gcontext "context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/LoveCatdd/webctx/pkg/lib/core/context"
	"github.com/LoveCatdd/webctx/pkg/lib/core/goroutine"
	"github.com/LoveCatdd/webctx/pkg/lib/core/web/auth"
	"github.com/LoveCatdd/webctx/pkg/lib/core/web/identity"
	"github.com/LoveCatdd/webctx/pkg/lib/core/web/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// todo

// init context.WebContextHolder & goroutine.GoroutineContextHolder
func ContextMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		tokenStr := ctx.Request.Header.Get(auth.JWT_AUTHORIZATION_KEY)

		mapClaims, err := auth.ExtractMapClaims(tokenStr)

		// 身份拦截
		if err != nil {

			// log todo
			ctx.JSON(
				http.StatusNetworkAuthenticationRequired,
				response.FailWithCodeAndMessage(
					response.AUTHORIZE_FAIL,
					fmt.Sprintf("/%s%s", ctx.Request.Host, ctx.Request.URL.String()),
					"Authorization err",
				),
			)

			ctx.Abort()
			return
		}

		contextHolder := new(goroutine.GoroutineContextHolder)
		contextHolder.Initialization()

		// 插入header信息
		userId := withInfo(ctx.Request.Header, mapClaims, identity.IDENTITY_USERTID_KEY)
		username := withInfo(ctx.Request.Header, mapClaims, identity.IDENTITY_USERTNAEM_KEY)
		userIdentityName := withInfo(ctx.Request.Header, mapClaims, identity.IDENTITY_USERIDENTITYNAME_KEY)
		clientId := withInfo(ctx.Request.Header, mapClaims, identity.IDENTITY_CLIENTID_KEY)

		identityInfo := identity.NewIdetityInfo(userId, username, userIdentityName, clientId)

		// identityInfo & mapClaims 维护到 contextHolder
		contextHolder.Change(goroutine.IDENTITY_CONTEXT_INFO_KEY, &identityInfo)
		contextHolder.Change(goroutine.JWT_MAP_CLAIM, mapClaims)

		// 持久化到 request_context 中
		customContext := context.NewCustomContext(contextHolder)

		c := gcontext.WithValue(ctx.Request.Context(), context.CustonContextKey, customContext)
		ctx.Request = ctx.Request.WithContext(c)

		ctx.Next()
	}
}

func withHeader(header http.Header, key string) string {

	//获取请求头并转义
	info, _ := url.QueryUnescape(header.Get(key))
	return info
}

func withClaim(claim jwt.MapClaims, key string) string {

	// 通过解析claim获取
	return claim[key].(string)
}

func withInfo(header http.Header, claim jwt.MapClaims, key string) string {
	info := withHeader(header, key)
	if info == "" {
		info = withClaim(claim, key)
	}
	return info
}
