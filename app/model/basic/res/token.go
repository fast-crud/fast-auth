package res

import "github.com/flipped-aurora/gf-vue-admin/app/model/system"

type AccessToken struct {
	User      *system.User `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
}
