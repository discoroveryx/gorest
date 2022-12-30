package config

type ProjectConf struct {
	DBHost   string
	DBUser   string
	DBName   string
	DBEngine string
}

var conf *ProjectConf

// Singleton
func init() {
	conf = &ProjectConf{
		DBHost:   "gorest-pg",
		DBUser:   "postgres",
		DBName:   "main_1",
		DBEngine: "postgres",
	}
}

func GetProjectConf() *ProjectConf {
	return conf
}
