package resolver

import (
	"github.com/dalloriam/synthia"
)

type RootResolver struct {
	synth *synthia.Synthia
}

func New(synth *synthia.Synthia) *RootResolver {
	return &RootResolver{synth}
}

func (r *RootResolver) Mixer() *MixerResolver {
	return &MixerResolver{r.synth.Mixer}
}
