package utils

import (
	"strings"
	"errors"
)

type TreeSet struct {
	key string
	values []interface{}
	child []*TreeSet
}

func NewTreeSet() *TreeSet{
	return &TreeSet{}
}

func getValue(t *TreeSet) []interface{} {
	var vs []interface{}
	if t.values != nil {
		vs = append(vs, t.values...)
	}
	for _, v := range t.child {
		vv := getValue(v)
		if vv != nil {
			vs = append(vs, vv...)
		}
	}

	return vs
}

func getChild(treeSet *TreeSet, key string) *TreeSet{
	for _, v := range treeSet.child {
		if v.key == key {
			return v
		}
	}
	return nil
}

func (t *TreeSet) Set(key string, value interface{}) error {
	if strings.HasPrefix(key, "/") == false {
		return errors.New("invalid key (key must start with '/')")
	}
	if key == "/" {
		t.key = key
		t.values = append(t.values, value)
		return nil
	}
	kSet := strings.Split(key, "/")
	treeSet := t
	for _, v:= range kSet {
		if len(v) > 0 {
			ts := getChild(treeSet, v)
			if ts != nil {
				treeSet = ts
				continue
			} else {
				child := &TreeSet{key:v}
				treeSet.child = append(treeSet.child, child)
				treeSet = child
			}
		}
	}
	treeSet.values = append(treeSet.values, value)
	return nil
}

func (t *TreeSet) Get(key string) ([]interface{}, error) {
	if strings.HasPrefix(key, "/") == false {
		return nil, errors.New("invalid key (key must start with '/')")
	}
	if key == "/" {
		return getValue(t), nil
	}
	kSet := strings.Split(key, "/")
	treeSet := t
	for _, v:= range kSet {
		if len(v) > 0 {
			ts := getChild(treeSet, v)
			if ts != nil {
				treeSet = ts
				continue
			} else {
				return nil, errors.New("not found")
			}
		}
	}
	values := getValue(treeSet)

	return values, nil
}

func (t *TreeSet) Replace(key string, value interface{}) error {
	if strings.HasPrefix(key, "/") == false {
		return errors.New("invalid key (key must start with '/')")
	}
	if key == "/" {
		t.values = make([]interface{}, 1)
		t.values[0] = value
		return nil
	}
	kSet := strings.Split(key, "/")
	treeSet := t
	for _, v:= range kSet {
		if len(v) > 0 {
			ts := getChild(treeSet, v)
			if ts != nil {
				treeSet = ts
				continue
			} else {
				return errors.New("not found")
			}
		}
	}
	treeSet.values = make([]interface{}, 1)
	treeSet.values[0] = value
	return nil
}

func (t *TreeSet) Remove(key string) error {
	if strings.HasPrefix(key, "/") == false {
		return errors.New("invalid key (key must start with '/')")
	}
	if key == "/" {
		t.values = nil
		return nil
	}
	kSet := strings.Split(key, "/")
	treeSet := t
	for _, v:= range kSet {
		if len(v) > 0 {
			ts := getChild(treeSet, v)
			if ts != nil {
				treeSet = ts
				continue
			} else {
				return errors.New("not found")
			}
		}
	}
	treeSet.values = nil
	return nil
}