package main

import (
	"github.com/McCdama/govue/router"
)

func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":8080"))
}
