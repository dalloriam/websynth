package resolver

import (
	"github.com/dalloriam/synthia/modular"
	"github.com/dalloriam/websynth/app/audio"
)

type OscillatorResolver struct {
	sys *audio.System
	osc *modular.Oscillator
}

func (r *OscillatorResolver) Volume() *KnobResolver {
	return &KnobResolver{r.sys, r.osc.Volume}
}

func (r *OscillatorResolver) Frequency() *KnobResolver {
	return &KnobResolver{r.sys, r.osc.Frequency}
}
