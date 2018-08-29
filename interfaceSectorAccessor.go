package berlioz

import (
	"os"
	"strings"
)

/*
 * SECTOR ACCESSOR
 */
type SectorAccessor struct {
	name string
}

// TBD
func (x SectorAccessor) Service(name string) ServiceAccessor {
	values := []string{}
	values = append(values, os.Getenv("BERLIOZ_CLUSTER"))
	values = append(values, x.name)
	values = append(values, name)
	id := "service://" + strings.Join(values, "-")
	return ServiceAccessor{id: id, peers: NewEndpointPeers(id, "default")}
}

// TBD
func (x SectorAccessor) Database(name string) NativeResourceAccessor {
	return x.nativeResource("database", name)
}

// TBD
func (x SectorAccessor) Queue(name string) NativeResourceAccessor {
	return x.nativeResource("queue", name)
}

func (x SectorAccessor) nativeResource(kind string, name string) NativeResourceAccessor {
	values := []string{}
	values = append(values, os.Getenv("BERLIOZ_CLUSTER"))
	values = append(values, x.name)
	values = append(values, name)
	id := kind + "://" + strings.Join(values, "-")
	return NativeResourceAccessor{id: id, peers: NewResourcePeers(id)}
}
