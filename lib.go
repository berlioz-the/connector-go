package berlioz

import ()

func init() {
	initEnvironment()
	initIdentity()
	initClient(processMessage)
	initZipkin()
}
