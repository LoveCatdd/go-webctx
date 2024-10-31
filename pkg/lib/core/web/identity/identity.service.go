package identity

import (
	"github.com/LoveCatdd/webctx/pkg/lib/core/context"
	"github.com/LoveCatdd/webctx/pkg/lib/core/goroutine"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	// 获取 custonContextKey
	custonContextKey = context.CustonContextKey
)

type IdentityService interface {

	// 从 goroutine.IDENTITY_CONTEXT_INFO 提取身份信息
	IdentityInfo(*gin.Context) (*IdentityInfo, bool)

	// 从 提取信息
	ClaimMap(*gin.Context) (jwt.MapClaims, bool)

	// get uid
	UserId(*gin.Context) string

	// get username
	UserName(*gin.Context) string

	// get user identity name
	UserIdentityName(*gin.Context) string

	// get client id
	ClientId(*gin.Context) string
}

type Impl struct{}

func parse(c *gin.Context, key string) (any, bool) {
	customContext := c.Request.Context().Value(custonContextKey).(*context.CustomContext)
	if info, ok := customContext.ContextHolder().ContenxtMap().Get(key); ok {
		return info, true
	}
	return nil, false
}

func (Impl) identityInfo(c *gin.Context) (*IdentityInfo, bool) {
	if info, ok := parse(c, goroutine.IDENTITY_CONTEXT_INFO_KEY); ok {
		return info.(*IdentityInfo), true
	}
	return nil, false
}

func (Impl) claimMap(c *gin.Context) (jwt.MapClaims, bool) {
	if info, ok := parse(c, goroutine.JWT_MAP_CLAIM); ok {
		return info.(jwt.MapClaims), true
	}
	return nil, false
}

func (i Impl) UserId(c *gin.Context) string {
	if info, ok := i.identityInfo(c); ok && info != nil {
		return info.UserId
	}
	if info, ok := i.claimMap(c); ok && info != nil {
		return info[IDENTITY_USERTID_KEY].(string)
	}
	return ""
}

func (i Impl) UserName(c *gin.Context) string {
	if info, ok := i.identityInfo(c); ok && info != nil {
		return info.UserName
	}
	if info, ok := i.claimMap(c); ok && info != nil {
		return info[IDENTITY_USERTNAEM_KEY].(string)
	}
	return ""
}

func (i Impl) UserIdentityName(c *gin.Context) string {
	if info, ok := i.identityInfo(c); ok && info != nil {
		return info.UserIdentityName
	}
	if info, ok := i.claimMap(c); ok && info != nil {
		return info[IDENTITY_USERIDENTITYNAME_KEY].(string)
	}
	return ""
}

func (i Impl) ClientId(c *gin.Context) string {
	if info, ok := i.identityInfo(c); ok && info != nil {
		return info.ClientId
	}
	if info, ok := i.claimMap(c); ok && info != nil {
		return info[IDENTITY_CLIENTID_KEY].(string)
	}
	return ""
}
