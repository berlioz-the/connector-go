package berlioz

import (
	"fmt"
	"log"
	"strings"
	// "github.com/orcaman/concurrent-map"
)

type sectionT struct {
	items       map[string]interface{}
	subscribers map[string][]registryChangeHandler
}

type sectionsT map[string]sectionT

type registryT struct {
	sections sectionsT
}

type registryChangeHandler func(interface{})

var registry = *newRegistry()
var registrySectionNames = [...]string{
	"endpoints",
	"policies",
	"service",
	"cluster",
	"database",
	"queue",
}

func newRegistry() *registryT {
	x := registryT{}
	x.sections = make(sectionsT)
	for _, name := range registrySectionNames {
		x.sections[name] = sectionT{
			items:       make(map[string]interface{}),
			subscribers: make(map[string][]registryChangeHandler),
		}
	}
	keys := make([]string, 0, len(x.sections))
	for k := range x.sections {
		keys = append(keys, k)
	}

	return &x
}

func (x registryT) set(name string, path []string, value interface{}) {
	log.Printf("[REGISTRY] SET. Section: %s, Path: %v \n", name, path)
	// log.Printf("[REGISTRY] SET. Section: %s, Path: %v, Value: %v \n", name, path, value)

	if section, err := x._getSection(name); err == nil {
		fullname := _makeFullName(path)
		section.items[fullname] = value
		x._notifyToSubscribers(section, fullname, value)
	}
}

func (x registryT) subscribe(name string, path []string, callback registryChangeHandler) {
	log.Printf("[REGISTRY] Subscribe. Section: %s, Path: %v \n", name, path)

	if section, err := x._getSection(name); err == nil {
		fullname := _makeFullName(path)
		section.subscribers[fullname] = append(section.subscribers[fullname], callback)

		val := section.items[fullname]
		if val != nil {
			callback(val)
		}
	}
}

func (x registryT) get(name string, path []string) interface{} {
	if section, err := x._getSection(name); err == nil {
		fullname := _makeFullName(path)
		if val, ok := section.items[fullname]; ok {
			return val
		}
	}
	return nil
}

func (x registryT) getAsIndexedMap(name string, path []string) indexedMap {
	value := x.get(name, path)
	if value == nil {
		return indexedMap{}
	}
	return value.(indexedMap)
}

func (x registryT) setAsIndexedMap(name string, path []string, value interface{}) {
	x.set(name, path, newIndexedMap(value))
}

func (x registryT) _getSection(name string) (*sectionT, error) {
	if val, ok := x.sections[name]; ok {
		return &val, nil
	}
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
