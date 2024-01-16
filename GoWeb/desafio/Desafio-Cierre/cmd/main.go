package main

import (
	"app/internal/application"
	"fmt"
)

func main() {
	// env
	// ...

	// application
	// - config
	// cfg := &ConfigAppDefault{
	// 	ServerAddr: os.Getenv("SERVER_ADDR"),
	// 	DbFile:     os.Getenv("DB_FILE"),
	// }
	cfg := &application.ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "./docs/db/tickets.csv",
	}
	app := application.NewApplicationDefault(cfg)

	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
