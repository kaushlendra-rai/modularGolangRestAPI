package util

import (
	"fmt"

	"github.com/spf13/viper"
)

// Get Viper for Configuration management
// go get github.com/spf13/viper

func init() {
	viper.SetConfigName("application")
	viper.AddConfigPath(".")
	viper.AddConfigPath("meta/resources/")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Config not found")
	}
}
func main() {
	fmt.Println("In main")
	viper.SetConfigName("application")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config not found")
	} else {
		name := viper.GetString("name")
		fmt.Print("Name configuration:", name)
	}
}

func GetProperty(key string) string {
	value := viper.GetString(key)

	return value
}
