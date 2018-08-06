package resolver

import (
	"fmt"
	"reflect"

	"github.com/dalloriam/synthia/core"
	"github.com/dalloriam/websynth/app/audio"
)

type SignalResolver struct {
	sys *audio.System
	sgn *core.Signal
}

func (s *SignalResolver) Attach(args struct {
	ModuleIdx   int32
	ModuleField *string
}) (bool, error) {

	i := int(args.ModuleIdx)

	if i >= len(s.sys.Modules) {
		return false, fmt.Errorf("cannot connect module %d, only %d modules are defined", i, len(s.sys.Modules))
	}

	if args.ModuleField == nil {
		if newSignal, ok := s.sys.Modules[i].(core.Signal); ok {
			*s.sgn = newSignal
			return true, nil
		}
		return false, fmt.Errorf("module %d cannot be cast to signal, maybe you are missing an attribute?", i)
	}

	// Need to use reflection
	v := reflect.ValueOf(s.sys.Modules[i]).Elem()
	f := v.FieldByName(*args.ModuleField)

	if !f.IsValid() {
		return false, fmt.Errorf("field %s doesnt exist on module %d", *args.ModuleField, i)
	}

	if newSignal, ok := f.Interface().(core.Signal); ok {
		*s.sgn = newSignal
		return true, nil
	}

	return false, fmt.Errorf("field %s of module %d cannot be cast to signal", *args.ModuleField, i)
}

func (s *SignalResolver) Detach() bool {
	*s.sgn = nil
	return true
}
