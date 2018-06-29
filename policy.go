package berlioz

import (
	"log"
)

var policyDefaults = map[string]interface{}{
	"enable-zipkin":          true,
	"zipkin-endpoint":        "",
	"timeout":                5000,
	"no-peer-retry":          true,
	"retry-count":            3,
	"retry-initial-delay":    500,
	"retry-delay-multiplier": 2,
	"retry-max-delay":        5000,
}

func resolvePolicyBool(name string, target []string) bool {
	val := resolvePolicy(name, target)
	return toBool(val)
}

func resolvePolicyString(name string, target []string) string {
	val := resolvePolicy(name, target)
	return toString(val)
}

func resolvePolicyInt(name string, target []string) int {
	val := resolvePolicy(name, target)
	return toInt(val)
}

func resolvePolicyFloat(name string, target []string) float64 {
	val := resolvePolicy(name, target)
	return toFloat(val)
}

func monitorPolicy(name string, target []string, callback func(interface{})) {
	currVal := resolvePolicy(name, target)
	registry.subscribe("policies", nil, func(interface{}) {
		newVal := resolvePolicy(name, target)
		if currVal != newVal {
			currVal = newVal
			callback(currVal)
		}
	})
	callback(currVal)
}

func monitorBool(name string, target []string, callback func(bool)) {
	monitorPolicy(name, target, func(val interface{}) {
		callback(toBool(val))
	})
}

func monitorString(name string, target []string, callback func(string)) {
	monitorPolicy(name, target, func(val interface{}) {
		callback(toString(val))
	})
}

func monitorInt(name string, target []string, callback func(int)) {
	monitorPolicy(name, target, func(val interface{}) {
		callback(toInt(val))
	})
}

func monitorFloat(name string, target []string, callback func(float64)) {
	monitorPolicy(name, target, func(val interface{}) {
		callback(toFloat(val))
	})
}

func resolvePolicy(name string, target []string) interface{} {
	// log.Printf("[resolvePolicy] %s, %v\n", name, target)

	var res interface{}
	root := registry.get("policies", nil)
	if root != nil {
		res = _resolvePolicy(root.(*policyModel), name, target)
	}
	if res == nil {
		if defVal, ok := policyDefaults[name]; ok {
			res = defVal
		} else {
			log.Printf("Default for policy %s not set!", name)
		}
	}
	return res
}

func _resolvePolicy(root *policyModel, name string, target []string) interface{} {
	// log.Printf("[_resolvePolicy] %s, %v\n", name, target)
	if len(target) > 0 {
		if root.Children != nil {
			if child, ok := root.Children[target[0]]; ok {
				childTarget := target[1:]
				res := _resolvePolicy(&child, name, childTarget)
				if res != nil {
					return res
				}
			}
		}
	}
	if root.Values != nil {
		if value, ok := root.Values[name]; ok {
			return value
		}
	}
	return nil
}

func toBool(x interface{}) bool {
	if safeVal, ok := x.(bool); ok {
		return safeVal
	}
	return true
}

func toString(x interface{}) string {
	if safeVal, ok := x.(string); ok {
		return safeVal
	}
	return ""
}

func toInt(x interface{}) int {
	if safeVal, ok := x.(int); ok {
		return safeVal
	}
	return 0
}

func toFloat(x interface{}) float64 {
	if safeVal, ok := x.(float64); ok {
		return safeVal
	}
	return 0.0
}
