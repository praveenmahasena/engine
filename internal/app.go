package internal

import (
	"github.com/praveenmahasena/engine/internal/server"
	"github.com/spf13/viper"
)

func Start() error {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	s := server.New()

	return s.Run()
}
