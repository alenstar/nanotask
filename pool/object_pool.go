package pool

import (
	"sync"
	"reflect"
	"errors"
)

type ObjectPool struct {
	sync.Mutex
	typItems map[string]reflect.Type
	objItems map[reflect.Type]map[interface{}]bool
}

func NewObjectPool() *ObjectPool {
	return &ObjectPool{
		typItems: make(map[string]reflect.Type),
		objItems: make(map[reflect.Type]map[interface{}]bool),
	}
}

func (o *ObjectPool) RegisterType(name string, obj interface{}) error {
	typ := reflect.Indirect(reflect.ValueOf(obj)).Type()
	o.Lock()
	defer o.Unlock()
	if t, ok := o.typItems[name]; ok {
		if t == typ {
			// was registered
			return nil
		}
		return errors.New(" the type name was registered")
	}

	o.typItems[name] = typ
	o.objItems[typ] = make(map[interface{}]bool)
	return nil
}

func (o *ObjectPool) Release(obj interface{}) {
	typ := reflect.Indirect(reflect.ValueOf(obj)).Type()
	o.Lock()
	defer o.Unlock()
	if used, ok := o.objItems[typ][obj]; ok {
		if used {
			o.objItems[typ][obj] = false
		} else {
			// the object not in this pool
			// FIXME
		}
		return
	}
	// not register type
	// TODO
}

func (o *ObjectPool) Obtain(name string) (interface{}, error) {
	if _, ok := o.typItems[name]; ok == false {
		return nil, errors.New(" the type was not registered")
	}
	typ := o.typItems[name]
	for value, used := range o.objItems[typ] {
		if used == false {
			o.objItems[typ][value] = true
			return value, nil
		}
	}

	// not found free object
	// new object and return

	value := reflect.New(typ).Interface()
	o.objItems[typ][value] = true
	return value, nil
}
