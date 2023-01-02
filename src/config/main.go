package config

import "os"

type DataBaseConf struct {
	Engine   string
	Host     string
	Port     string
	User     string
	Name     string
	TimeZone string
}

type ProjectConf struct {
	AuthSecretkey string
	DataBase *DataBaseConf
}

var conf *ProjectConf

// Singleton
func init() {
	conf = &ProjectConf{
		AuthSecretkey: os.Getenv("AUTH_SECRET_KEY"),
		DataBase: &DataBaseConf{
			Engine:   os.Getenv("DB_ENGINE"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Name:     os.Getenv("DB_NAME"),
			TimeZone: os.Getenv("DB_TZ"),
		},
	}
}

func GetProjectConf() *ProjectConf {
	return conf
}
