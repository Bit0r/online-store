package conf

import "github.com/spf13/viper"

func Unmarshal(table string, prefix string, rawVal interface{}) {
	if readErr != nil {
		viper.SetEnvPrefix(prefix)
		viper.Unmarshal(rawVal)
	} else {
		viper.Sub(table).Unmarshal(rawVal)
	}
}
