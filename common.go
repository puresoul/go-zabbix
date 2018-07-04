package zabbix

import (
	"regexp"
	"strconv"
	"github.com/pkg/errors"
)

var timeSuffixRegexp = regexp.MustCompile(`^(\d+)([smhdw])$`)

// getTime convert string with time (suffixes supported) to number of seconds
// https://www.zabbix.com/documentation/3.2/manual/config/triggers/suffixes
func getTime(tm string) (int, error) {
	sec, err := strconv.Atoi(tm)
	if err != nil {
		m := timeSuffixRegexp.FindStringSubmatch(tm)
		if len(m) == 3 {
			val, err := strconv.Atoi(m[1])
			if err != nil {
				return 0, errors.New("Can't convert " + tm + " to seconds")
			}
			switch m[2] {
				case "s": return val, nil
				case "m": return val*60, nil
				case "h": return val*3600, nil
				case "d": return val*86400, nil
				case "w": return val*604800, nil
			}
			return 0, errors.New("Can't convert " + tm + " to seconds")
		}
		return 0, nil
	}
	return sec, nil
}