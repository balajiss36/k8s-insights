package misc

import "github.com/spf13/viper"

type Config struct {
	HTTPAddress   string `mapstructure:"HTTP_ADDR"`
	MongoAddr     string `mapstructure:"MONGO_ADDR"`
	MongoUser     string `mapstructure:"MONGO_USER"`
	MongoPassword string `mapstructure:"MONGO_PASS"`
	MongoService  string `mapstructure:"MONGO_SRV"`
	Namespace     string `mapstructure:"NAMESPACE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
