package config

type ProjectConf struct {
	DBName   string
	DBEngine string
}

var conf *ProjectConf

func init() {
	conf = &ProjectConf{
		DBName:   "main_1.db",
		DBEngine: "sqlite",
	}
}

func GetProjectConf() *ProjectConf {
	return conf
}
