package main

import (
	"github.com/BaianorASR/solid-go/cmd"
	"github.com/BaianorASR/solid-go/config/environment"
)

func main() {
	config.LoadEnv()
	cmd.Server()
}
