package berlioz

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"

	"github.com/google/uuid"
	// "github.com/orcaman/concurrent-map"
)

type sectionT struct {
	items       map[string]interface{}
	subscribers map[string]map[string]registryChangeHandler
}

type sectionsT map[string]sectionT

type registryT struct {
	lock     sync.Mutex
	sections sectionsT
}

type SubscribeInfo struct {
	id       string
	name     string
	path     []string
	registry *registryT
}

type registryChangeHandler func(interface{})

var registry = *newRegistry()
var registrySectionNames = [...]string{
	"endpoints",
	"policies",
	"peers",
	"consumes",
	"service",
	"cluster",
	"database",
	"queue",
	"secret_public_key",
	"secret_private_key",
}

func newRegistry() *registryT {
	x := registryT{}
	x.sections = make(sectionsT)
	for _, name := range registrySectionNames {
		x.sections[name] = sectionT{
			items:       make(map[string]interface{}),
			subscribers: make(map[string]map[string]registryChangeHandler),
		}
	}
	keys := make([]string, 0, len(x.sections))
	for k := range x.sections {
		keys = append(keys, k)
	}

	return &x
}

func (x registryT) set(name string, path []string, value interface{}) {
	x.lock.Lock()
	defer x.lock.Unlock()

	log.Printf("[REGISTRY] SET. Section: %s, Path: %v \n", name, path)
	// log.Printf("[REGISTRY] SET. Section: %s, Path: %v, Value: %v \n", name, path, value)

	if section, err := x._getSection(name); err == nil {
		fullname := _makeFullName(path)
		section.items[fullname] = value
		x._notifyToSubscribers(section, fullname, value)
	}
}

func (x registryT) subscribe(name string, path []string, callback registryChangeHandler) SubscribeInfo {
	x.lock.Lock()
	defer x.lock.Unlock()

	log.Printf("[REGISTRY] Subscribe. Section: %s, Path: %v \n", name, path)

	subscribeID := uuid.New()
	subscribeIDStr := subscribeID.String()

	if section, err := x._getSection(name); err == nil {
		fullname := _makeFullName(path)
		if _, ok := section.subscribers[fullname]; !ok {
			section.subscribers[fullname] = make(map[string]registryChangeHandler)
		}
		section.subscribers[fullname][subscribeIDStr] = callback

		val := section.items[fullname]
		if val != nil {
			callback(val)
		}

		return SubscribeInfo{id: subscribeIDStr, name: name, path: path, registry: &x}
	}

	return SubscribeInfo{}
}

func (x SubscribeInfo) Stop() {
	if x.registry == nil {
		return
	}
	x.registry.unsubscribe(x.name, x.path, x.id)
}

func (x registryT) unsubscribe(name string, path []string, subscribeID string) {
	x.lock.Lock()
	defer x.lock.Unlock()

	log.Printf("[REGISTRY] Unubscribe. Section: %s, Path: %v \n", name, path)

	if section, err := x._getSection(name); err == nil {
		fullname := _makeFullName(path)
		if _, ok := section.subscribers[fullname]; ok {
			delete(section.subscribers[fullname], subscribeID)
		}
	}
}

func removeFromArray(vs []registryChangeHandler, value registryChangeHandler) []registryChangeHandler {
	vsf := make([]registryChangeHandler, 0)
	for _, v := range vs {
		if reflect.ValueOf(v).Pointer() != reflect.ValueOf(value).Pointer() {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (x registryT) get(name string, path []string) interface{} {
	x.lock.Lock()
	defer x.lock.Unlock()

	// log.Printf("[REGISTRY] GET. Section: %s, Path: %v \n", name, path)

	if section, err := x._getSection(name); err == nil {
		fullname := _makeFullName(path)
		if val, ok := section.items[fullname]; ok {
			// log.Printf("[REGISTRY] GET. Res: %v \n", val)
			return val
		}
	}
	return nil
}

func (x registryT) getAsIndexedMap(name string, path []string) IndexedMap {
	value := x.get(name, path)
	if value == nil {
		return IndexedMap{}
	}
	return value.(IndexedMap)
}

func (x registryT) setAsIndexedMap(name string, path []string, value interface{}) {
	// log.Printf("[REGISTRY] setAsIndexedMap. Section: %s, Path: %v, Value: %v \n", name, path, value)
	x.set(name, path, newIndexedMap(value))
}

func (x registryT) _getSection(name string) (*sectionT, error) {
	if val, ok := x.sections[name]; ok {
		return &val, nil
	}
	log.Printf("[REGISTRY] Unknown section: %s.", name)
	return nil, fmt.Errorf("Unknown section: %s", name)
}

func (x registryT) _notifyToSubscribers(section *sectionT, fullname string, value interface{}) {
	for _, callback := range section.subscribers[fullname] {
		callback(value)
	}
}

func (x registryT) debugOutput() {
	// log.Printf("[REGISTRY] DEBUG OUTPUT: %s\n", spew.Sdump(x))
}

func _makeFullName(arr []string) string {
	if arr == nil {
		return ""
	}
	return strings.Join(arr, "-")
}
