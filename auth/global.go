package auth

import "github.com/gonyyi/atype"

const (
	ERR_USERID_EXIST            atype.ErrorStr = "user exist"
	ERR_USERID_NOT_EXIST        atype.ErrorStr = "user not exist"
	ERR_MISSING_REQUIRED_FIELDS atype.ErrorStr = "missing required field"
	ERR_BAD_CREDENTIAL          atype.ErrorStr = "bad credential"
	ERR_DISABLED_ID             atype.ErrorStr = "disabled user account"
	ERR_USER_NOT_IN_GROUP       atype.ErrorStr = "user not in group"
	ERR_ENCRYPTION_CIPHER_SHORT atype.ErrorStr = "cipher text too short"
	ERR_KEY_NOT_EXIST           atype.ErrorStr = "key not exists"
)
