package config

type db struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func GetDbConnection() *db {
	return &db{}
}
