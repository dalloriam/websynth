package resolver

import (
	"github.com/dalloriam/synthia/modular"
	"github.com/dalloriam/websynth/app/audio"
)

type ModuleResolver struct {
	sys *audio.System
	mod interface{}
}

func (m *ModuleResolver) ToClock() (*ClockResolver, bool) {
	if clock, ok := m.mod.(*modular.Clock); ok {
		return &ClockResolver{m.sys, clock}, true
	}
	return nil, false
}

func (m *ModuleResolver) ToClockMutation() (*ClockResolver, bool) {
	return m.ToClock()
}

func (m *ModuleResolver) ToOscillator() (*OscillatorResolver, bool) {
	if osc, ok := m.mod.(*modular.Oscillator); ok {
		return &OscillatorResolver{m.sys, osc}, true
	}
	return nil, false
}

func (m *ModuleResolver) ToOscillatorMutation() (*OscillatorResolver, bool) {
	return m.ToOscillator()
}
