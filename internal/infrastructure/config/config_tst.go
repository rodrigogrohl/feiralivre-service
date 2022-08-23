package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitializeTest() {
	viper.AddConfigPath("../../configs/")
	viper.AddConfigPath("../../../configs/")
	viper.AddConfigPath("../../../../configs/")

	Initialize("dev")
	// DbConnection = "../../../../feiralivre.db"

	DbConnection = fmt.Sprintf("%sfeiralivre.db", BasePath)
}
