package configs

import (
	"flag"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() error {
	var configPath, configFile string

	flag.StringVar(&configPath, "path", "configs", "Path to config file")
	flag.StringVar(&configFile, "config", "config", "Name of config file")
	flag.StringVar(&configPath, "p", "configs", "Path to config file")
	flag.StringVar(&configFile, "c", "config", "Name of config file")
	flag.Parse()

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFile)
	return viper.ReadInConfig()
}

func CreateOrOpenFileForLogs(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		file, err = os.Create(fileName)
		if err != nil {
			return file, err
		}
	}

	return file, nil
}
