package berlioz

import (
	"os"
	"strconv"
	"strings"
)

var myIdentity = ""

func initIdentity() {
	identity := os.Getenv("BERLIOZ_IDENTITY")

	prefix := os.Getenv("BERLIOZ_IDENTITY_PREFIX")
	if prefix != "" {
		identity = strings.TrimPrefix(identity, prefix)
	}

	proc := os.Getenv("BERLIOZ_IDENTITY_PROCESS")
	if proc == "plus_one" {
		if i, err := strconv.ParseInt(identity, 10, 64); err == nil {
			identity = strconv.FormatInt(i+1, 10)
		}
	}

	myIdentity = identity
}

// TBD
func Identity() string {
	return myIdentity
}
