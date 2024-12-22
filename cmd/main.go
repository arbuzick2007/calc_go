package main

import (
	"github.com/arbuzick57/calc_go/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}
