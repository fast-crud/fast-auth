package config

type Casdoor struct {
	Endpoint     string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	ClientId     string `mapstructure:"clientId" json:"clientId" yaml:"clientId"`
	ClientSecret string `mapstructure:"clientSecret" json:"clientSecret" yaml:"clientSecret"`
	Organization string `mapstructure:"organization" json:"organization" yaml:"organization"`
	Application  string `mapstructure:"application" json:"application" yaml:"application"`
	JwtPublicKey string `mapstructure:"jwtPublicKey" json:"jwtPublicKey" yaml:"jwtPublicKey"`
}
