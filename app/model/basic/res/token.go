package res

import "github.com/fast-crud/fast-auth/app/model/system"

type AccessTokenRes struct {
	User      *system.User        `json:"user"`
	App       *system.Application `json:"app"`
	Token     string              `json:"token"`
	ExpiresAt int64               `json:"expiresAt"`
}
