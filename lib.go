package berlioz

import (
	"os"
)

func init() {
	wsURL := os.Getenv("BERLIOZ_AGENT_PATH")
	initClient(wsURL, processMessage)
}
