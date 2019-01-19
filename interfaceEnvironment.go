package berlioz

import (
	"os"
	"regexp"
	"strings"
)

var environment map[string]string

func initEnvironment() {

	rawEnvironment := make(map[string]string)
	for _, item := range os.Environ() {
		splits := strings.Split(item, "=")
		key := splits[0]
		val := splits[1]
		rawEnvironment[key] = val
	}

	environment = make(map[string]string)
	for key, value := range rawEnvironment {
		substitutions := [][2]int{}

		newValue := value
		r := regexp.MustCompile("\\${(\\w*)}")
		matches := r.FindAllStringSubmatchIndex(value, -1)
		for _, match := range matches {
			indexes := [2]int{match[0], match[1]}
			substitutions = append([][2]int{indexes}, substitutions...)
		}

		for _, substitution := range substitutions {
			otherVar := newValue[substitution[0]+2 : substitution[1]-1]
			if otherVal, ok := rawEnvironment[otherVar]; ok {
				newValue = newValue[0:substitution[0]] +
					otherVal +
					newValue[substitution[1]:]
			}
		}

		environment[key] = newValue
	}
}

// TBD
func GetEnvironmentVariable(name string) string {
	if val, ok := environment[name]; ok {
		return val
	}
	return ""
}
