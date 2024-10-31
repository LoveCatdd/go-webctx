package identity

const (
	IDENTITY_USERTID_KEY          = "USER_ID"
	IDENTITY_USERTNAEM_KEY        = "USER_NAME"
	IDENTITY_USERIDENTITYNAME_KEY = "USER_IDENTITYNAME"
	IDENTITY_CLIENTID_KEY         = "CLIENT_ID"
)

type IdentityInfo struct {
	UserId           string
	UserName         string
	UserIdentityName string
	ClientId         string
}

func NewIdetityInfo(userId, username, userIdentityName, clientId string) IdentityInfo {
	return IdentityInfo{
		UserId:           userId,
		UserName:         username,
		UserIdentityName: userIdentityName,
		ClientId:         clientId,
	}
}
