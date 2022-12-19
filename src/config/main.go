package config

type ProjectConf struct {
	DBName   string
	DBEngine string
}

var conf *ProjectConf

func init() {
	conf = &ProjectConf{
		DBName:   "main_1",
		DBEngine: "postgres",
	}
}

func GetProjectConf() *ProjectConf {
	return conf
}
