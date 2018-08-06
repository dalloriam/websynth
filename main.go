package main

import "github.com/dalloriam/websynth/app"

func main() {
	app := app.New()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
