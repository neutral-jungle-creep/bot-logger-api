package configs

import "github.com/spf13/viper"

func InitConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config_template")
	return viper.ReadInConfig()
}
