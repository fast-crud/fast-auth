package res

import "github.com/fast-crud/fast-auth/app/model/system"

type AccessTokenRes struct {
	User      *system.User `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
}
