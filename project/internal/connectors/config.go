package connectors
import (
env "github.com/Netflix/go-env"
)
type databaseCfg struct {
Host string `env:"DB_HOST,default=localhost,required=true"`
Port int `env:"DB_PORT,default=5432,required=true"`
Username string `env:"DB_USERNAME,required=true"`
Password string `env:"DB_PASSWORD,required=true"`
Database string `env:"DB_DATABASE,required=true"`
}

func newDatabaseCfg() *databaseCfg {
	var cfg databaseCfg
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
			panic(err)
	}
	return &cfg
}