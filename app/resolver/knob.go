package resolver

import "github.com/dalloriam/synthia/core"

type KnobResolver struct {
	knob *core.Knob
}

func (r *KnobResolver) Value() float64 {
	return r.knob.GetValue()
}

func (r *KnobResolver) Set(args struct{ Value float64 }) float64 {
	r.knob.SetValue(args.Value)
	return r.Value()
}
