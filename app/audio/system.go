package audio

import (
	"github.com/dalloriam/synthia"
)

// System wraps all interactions with the synthesizer.
type System struct {
	Synth            *synthia.Synthia
	Modules          []interface{}
	ModuleChannelMap map[int]int

	Config Config
}

// NewSystem initializes an audio system by creating an audio backend as well as a synthesizer.
func NewSystem(cfg Config) (*System, error) {
	backend, err := newBackend(cfg.SampleRate, cfg.BufferSize, cfg.InputDevice, cfg.OutputDevice)
	if err != nil {
		return nil, err
	}

	sys := &System{
		Modules:          []interface{}{},
		ModuleChannelMap: make(map[int]int),
	}

	sys.Synth = synthia.New(cfg.MixerChanCount, cfg.BufferSize, backend)

	return sys, nil
}
