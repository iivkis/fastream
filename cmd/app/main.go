package main

import (
	"github.com/iivkis/fastream/internal/config"
	"github.com/iivkis/fastream/internal/launcher"
)

func main() {
	config.LoadEnv(".env")
	launcher.Launch()
}
