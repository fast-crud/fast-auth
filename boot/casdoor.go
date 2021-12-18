package boot

import (
	"github.com/casdoor/casdoor-go-sdk/auth"
	"github.com/fast-crud/fast-auth/library/global"
)

func InitCasdoor() {
	var Endpoint = global.Config.Casdoor.Endpoint
	var ClientId = global.Config.Casdoor.ClientId
	var ClientSecret = global.Config.Casdoor.ClientSecret
	var JwtPublicKey = global.Config.Casdoor.JwtPublicKey
	var Organization = global.Config.Casdoor.Organization
	var Application = global.Config.Casdoor.Application
	auth.InitConfig(Endpoint, ClientId, ClientSecret, JwtPublicKey, Organization, Application)
}
