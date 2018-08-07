package resolver

import (
	"fmt"

	"github.com/dalloriam/websynth/app/audio"
)

type RootResolver struct {
	sys *audio.System
}

func New(system *audio.System) *RootResolver {
	return &RootResolver{system}
}

func (r *RootResolver) Mixer() *MixerResolver {
	return &MixerResolver{r.sys, r.sys.Synth.Mixer}
}

func (r *RootResolver) Module(args struct{ Idx int32 }) (*ModuleResolver, error) {
	i := int(args.Idx)

	if i >= len(r.sys.Modules) {
		return nil, fmt.Errorf("cannot get module %d, only got %d modules", i, len(r.sys.Modules))
	}

	return &ModuleResolver{r.sys, r.sys.Modules[i]}, nil
}

func (r *RootResolver) Modules() *ModuleCollectionResolver {
	return &ModuleCollectionResolver{r.sys}
}
