package resolver

import "github.com/dalloriam/websynth/app/audio"

type RootResolver struct {
	sys *audio.System
}

func New(system *audio.System) *RootResolver {
	return &RootResolver{system}
}

func (r *RootResolver) Mixer() *MixerResolver {
	return &MixerResolver{r.sys, r.sys.Synth.Mixer}
}

func (r *RootResolver) Modules() *ModuleCollectionResolver {
	return &ModuleCollectionResolver{r.sys}
}
