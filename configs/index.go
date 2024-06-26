package configs

import "github.com/spf13/viper"

type Config struct {
	CDNCepPath string `mapstructure:"CDN_CEP_PATH"`
	VIACepPath string `mapstructure:"VIA_CEP_PATH"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config
	viper.SetConfigName("channel-go")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil

}
