package main

import (
	"strings"

	"github.com/dalloriam/websynth/app"
	"github.com/dalloriam/websynth/app/audio"
	"github.com/spf13/viper"
)

var defaults = map[string]interface{}{
	"AUDIO.BUFFERSIZE":     512,
	"AUDIO.SAMPLERATE":     44100,
	"AUDIO.MIXERCHANCOUNT": 2,
	"AUDIO.INPUTDEVICE":    1,
	"AUDIO.OUTPUTDEVICE":   2,

	"GQLROUTE":        "/synth",
	"HOST":            "0.0.0.0:8080",
	"SCHEMADIRECTORY": "./schema",
}

func initConfig() {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetEnvPrefix("synth")
	viper.AutomaticEnv()

	for k, v := range defaults {
		viper.BindEnv(k)
		viper.SetDefault(k, v)
	}
}

func main() {
	initConfig()

	config := &app.Config{
		Audio: audio.Config{},
	}
	viper.Unmarshal(config)

	app := app.New(*config)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
