package berlioz

import (
	"os"
)

// TBD
func MyEndpoint(name string) MyEndpointAccessor {
	return MyEndpointAccessor{name: name}
}

// TBD
func Consumes() ConsumesAccessor {
	return ConsumesAccessor{}
}

//TBD
func Cluster(name string) ClusterAccessor {
	id := "cluster://" + name
	return ClusterAccessor{id: id, peers: NewEndpointPeers(id, "default")}
}

// TBD
func Sector(name string) SectorAccessor {
	return SectorAccessor{name: name}
}

// TBD
func Service(name string) ServiceAccessor {
	return Sector(os.Getenv("BERLIOZ_SECTOR")).Service(name)
}

// TBD
func Database(name string) NativeResourceAccessor {
	return Sector(os.Getenv("BERLIOZ_SECTOR")).Database(name)
}

// TBD
func Queue(name string) NativeResourceAccessor {
	return Sector(os.Getenv("BERLIOZ_SECTOR")).Queue(name)
}
