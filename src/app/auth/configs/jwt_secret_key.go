package configs

import "config"

func GetJWTSecretKey() string {
	conf := config.GetProjectConf()
	return conf.AuthSecretkey
}
