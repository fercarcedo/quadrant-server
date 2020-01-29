package main

import (
	"fmt"
	"github.com/fercarcedo/quadrant-server/internal/config"
)

func main() {
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
}