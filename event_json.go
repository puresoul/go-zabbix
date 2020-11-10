package zabbix

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// jEvent is a private map for the Zabbix API Event object.
// See: https://www.zabbix.com/documentation/2.2/manual/api/reference/event/object
type jEvent struct {
	EventID      string `json:"eventid"`
	Acknowledged string `json:"acknowledged"`
	Clock        string `json:"clock"`
	Nanoseconds  string `json:"ns"`
	ObjectType   string `json:"object"`
	ObjectID     string `json:"objectid"`
	Source       string `json:"source"`
	Value        string `json:"value"`
	ValueChanged string `json:"value_changed"`
	Name         string `json:"name"`
	Severity     string `json:"severity"`
	Hosts        jHosts `json:"hosts"`
}

// Event returns a native Go Event struct mapped from the given JSON Event data.
func (c *jEvent) Event() (*Event, error) {
	event := &Event{}
	event.EventID = c.EventID
	//	event.Acknowledged = (c.Acknowledged == "1")

	// parse timestamp
	sec, err := strconv.ParseInt(c.Clock, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Event timestamp: %v", err)
	}

	nsec, err := strconv.ParseInt(c.Nanoseconds, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Event timestamp nanoseconds: %v", err)
	}

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)

	event.Timestamp = time.Unix(sec, nsec)

	tmp := strings.Split(fmt.Sprint(now.Sub(time.Unix(sec, nsec))), ".")

	event.Age = tmp[0] + "s"

	//	event.ObjectType, _ = strconv.Atoi(c.ObjectType)
	//	if err != nil {
	//		return nil, fmt.Errorf("Error parsing Event Object Type: %v", err)
	//	}

	event.ObjectID, err = strconv.Atoi(c.ObjectID)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Event Object ID: %v", err)
	}

	//	event.Source, _ = strconv.Atoi(c.Source)
	//	if err != nil {
	//		return nil, fmt.Errorf("Error parsing Event Source: %v", err)
	//	}

	event.Value, err = strconv.Atoi(c.Value)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Event Source: %v", err)
	}

	//	event.ValueChanged = (c.ValueChanged == "1")

	event.Name = c.Name
	event.Severity = c.Severity

	return event, nil
}
