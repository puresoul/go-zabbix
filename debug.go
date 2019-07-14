package zabbix

import (
	"fmt"
	"os"
)

// debug caches the value of environment variable ZBX_DEBUG from program start.
var debug = true

// dprintf prints formatted debug message to STDERR if the ZBX_DEBUG environment
// variable is set to "1".
func dprintf(format string, a ...interface{}) {
	if debug == true {
		fmt.Fprintf(os.Stderr, format, a...)
	}
}
