package berlioz

import (
	"fmt"
	"strings"
	// "github.com/orcaman/concurrent-map"

	"github.com/davecgh/go-spew/spew"
)

type sectionT map[string]interface{}

type sectionsT map[string]sectionT

type registryT struct {
	sections sectionsT
}

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
		x.sections[name] = sectionT{}
	}
	keys := make([]string, 0, len(x.sections))
	for k := range x.sections {
		keys = append(keys, k)
	}
	fmt.Printf("[REGISTRY] KEYS!!!!!!!: %v\n", keys)

	return &x
}

func (x registryT) set(name string, path []string, value interface{}) {
	fmt.Printf("[REGISTRY] SET. Section: %s, Path: %v \n", name, path)

	if section, err := x._getSection(name); err == nil {
		fullname := _makeFullName(path)
		section[fullname] = value
	}
}

func (x registryT) get(name string, path []string) interface{} {
	if section, err := x._getSection(name); err == nil {
		fullname := _makeFullName(path)
		if val, ok := section[fullname]; ok {
			return val
		}
	}
	return nil
}

func (x registryT) _getSection(name string) (sectionT, error) {
	if val, ok := x.sections[name]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("Unknown section: %s", name)
}

func (x registryT) debugOutput() {
	fmt.Printf("[REGISTRY] DEBUG OUTPUT\n")
	spew.Dump(x)
}

func _makeFullName(arr []string) string {
	return strings.Join(arr, "-")
}
