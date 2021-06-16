package conf

import (
	"log"

	env "github.com/joeshaw/envdecode"
)

type Config struct {
    HTTP_PORT                string `env:"HTTP_PORT,default=3000"`
	ENVIRONMENT              string `env:"ENVIRONMENT,default=dev"`
}

func GetConfig() *Config {
	config := &Config{}
	err := env.Decode(config)	
	if err != nil {
		log.Printf("Error on decoding the environment variables:\n%v\n", err)
	}	
	return config
}

func (c *Config) Dump() {
	log.Printf("Config: HTTP_PORT: %s", c.HTTP_PORT)
	log.Printf("Config: ENVIRONMENT: %s", c.ENVIRONMENT)
}