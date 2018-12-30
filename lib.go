package berlioz

import (
	"os"
)

func init() {
	initIdentity()
	wsURL := os.Getenv("BERLIOZ_AGENT_PATH")
	initClient(wsURL, processMessage)
	initZipkin()
}
