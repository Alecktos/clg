package dev_utils

import (
	"fmt"
	"github.com/Alecktos/clg/game/config"
)

func Log(message string) {
	if config.DevMode {
		fmt.Println("Log: ", message)
	}
}
