package resolver

import (
	"github.com/dalloriam/synthia/modular"
	"github.com/dalloriam/websynth/app/audio"
)

type ModuleCollectionResolver struct {
	sys *audio.System
}

func (m *ModuleCollectionResolver) Create() *ModuleCreator {
	return &ModuleCreator{m.sys}
}

func (m *ModuleCollectionResolver) List() []*ModuleResolver {
	resolvers := make([]*ModuleResolver, len(m.sys.Modules))

	for i := 0; i < len(m.sys.Modules); i++ {
		resolvers[i] = &ModuleResolver{m.sys, m.sys.Modules[i]}
	}

	return resolvers
}

type ModuleCreator struct {
	sys *audio.System
}

func (c *ModuleCreator) Oscillator() bool {
	osc := modular.NewOscillator()
	c.sys.Modules = append(c.sys.Modules, osc)
	return true
}
