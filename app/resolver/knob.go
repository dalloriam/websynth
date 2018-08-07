package resolver

import (
	"github.com/dalloriam/synthia/core"
	"github.com/dalloriam/websynth/app/audio"
)

type KnobResolver struct {
	sys  *audio.System
	knob *core.Knob
}

func (r *KnobResolver) Value() float64 {
	return r.knob.GetValue()
}

func (r *KnobResolver) Set(args struct{ Value float64 }) float64 {
	r.knob.SetValue(args.Value)
	return r.Value()
}

func (r *KnobResolver) Line() *SignalResolver {
	return &SignalResolver{r.sys, &r.knob.Line}
}
