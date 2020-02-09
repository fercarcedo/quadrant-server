package main

import (
	"fmt"
	"github.com/fercarcedo/quadrant-server/internal/config"
	"github.com/fercarcedo/quadrant-server/internal/db"
)

func main() {
	var err error
	if err = config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	config.Config.DB, err = db.ConnectToDb()
	if err != nil {
		panic(fmt.Errorf("error connecting to the db: %s", err))
	}
	defer config.Config.DB.Close()
}