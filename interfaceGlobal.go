package berlioz

import (
	"os"
)

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