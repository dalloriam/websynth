package main

import "github.com/dalloriam/my_synth/app"

func main() {
	app := app.New()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
