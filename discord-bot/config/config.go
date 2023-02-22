package config

import (
	"os"
)

var (
	Token string
)

func ReadConfig() {
	Token = os.Getenv("TOKEN")
}
