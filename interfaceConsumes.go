package berlioz

/*
 * GLOBAL ACCESSORS
 */

/*
 * Consumes Accessor
 */
type ConsumesAccessor struct {
}

func (x ConsumesAccessor) massage(value interface{}) []ConsumesModel {
	return value.([]ConsumesModel)
}

// TBD
func (x ConsumesAccessor) All() interface{} {
	return x.massage(registry.get("consumes", []string{}))
}

// TBD
func (x ConsumesAccessor) MonitorAll(callback func([]ConsumesModel)) SubscribeInfo {
	return registry.subscribe("consumes", []string{}, func(value interface{}) {
		callback(x.massage(value))
	})
}
