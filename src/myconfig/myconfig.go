package myconfig

type MyConfig struct {
	DBName   string
	DBEngine string
}

var myconfig *MyConfig

func init() {
	myconfig = &MyConfig{
		DBName:   "main_1.db",
		DBEngine: "sqlite",
	}
}

func GetMyConfig() *MyConfig {
	return myconfig
}
