package constant

import "time"

const UserKey = "user"

const TokenKey = "Authorization"

const TokenPrefix = "Bearer"

const TokenExpireTime = 1 * time.Minute

const CaptchaExpireTime = 5 * time.Minute

const CaptchaIdKey = "captcha_id-"
