package berlioz

import (
	"math/rand"
	"reflect"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type IndexedMap struct {
	keys   []string
	values map[string]interface{}
}

func newIndexedMap(obj interface{}) IndexedMap {
	result := IndexedMap{
		keys:   make([]string, 0),
		values: make(map[string]interface{}),
	}
	robj := reflect.ValueOf(obj)
	rkeys := robj.MapKeys()
	for _, rkey := range rkeys {
		k := rkey.String()
		result.keys = append(result.keys, k)
		rv := robj.MapIndex(rkey)
		result.values[k] = rv.Interface()
	}
	return result
}

func (x IndexedMap) all() map[string]interface{} {
	return x.values
}

func (x IndexedMap) get(key string) interface{} {
	if value, ok := x.values[key]; ok {
		return value
	}
	return nil
}

func (x IndexedMap) random() interface{} {
	if len(x.keys) > 0 {
		key := x.keys[rand.Intn(len(x.keys))]
		return x.values[key]
	}
	return nil
}

func (x IndexedMap) first() interface{} {
	if len(x.keys) > 0 {
		key := x.keys[0]
		return x.values[key]
	}
	return nil
}

func firstKeySelector(x IndexedMap) interface{} {
	return x.first()
}
