package resolver

import (
	"github.com/dalloriam/synthia/modular"
	"github.com/dalloriam/websynth/app/audio"
)

type ModuleResolver struct {
	sys *audio.System
	mod interface{}
}

func (m *ModuleResolver) ToOscillator() (*OscillatorResolver, bool) {
	if osc, ok := m.mod.(*modular.Oscillator); ok {
		return &OscillatorResolver{m.sys, osc}, true
	}
	return nil, false
}
