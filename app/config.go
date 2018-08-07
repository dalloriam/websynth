package app

import "github.com/dalloriam/websynth/app/audio"

// Config wraps all synthesizer configuration.
type Config struct {
	Audio audio.Config `mapstructure:"AUDIO"`

	GQLRoute        string
	Host            string
	SchemaDirectory string
}
