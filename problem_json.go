package zabbix

import (
	"fmt"
	"strconv"
	"time"
)

type jProblem struct {
	EventID       string `json:"eventid"`
	Source        string `json:"source"`
	Object        string `json:"object"`
	ObjectID      string `json:"objectid"`
	Clock         string `json:"clock"`
	Nanoseconds   string `json:"ns"`
	REventID      string `json:"r_eventid"`
	RClock        string `json:"r_clock"`
	RNs           string `json:"r_ns"`
	CorrelationID string `json:"correlationid"`
	UserID        string `json:"userid"`
	Name          string `json:"name"`
	Ack           string `json:"acknowledged"`
	Severity      string `json:"severity"`
	Suppressed    int    `json:"suppressed"`
}

func (c *jProblem) Problem() (*Problem, error) {
	problem := &Problem{}
	problem.EventID = c.EventID
	//	problem.Object = c.Object
	problem.ObjectID = c.ObjectID

	sec, err := strconv.ParseInt(c.Clock, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Error parsing Problem clock: %v", err)
	}
	ns, err := strconv.ParseInt(c.Nanoseconds, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Error parsing nanoseconds: %v", err)
	}

	problem.Clock = c.Clock

	problem.Timestamp = time.Unix(sec, ns)

	/*	problem.Source = c.Source
		problem.REventID, err = strconv.Atoi(c.REventID)
		rsec, err := strconv.ParseInt(c.RClock, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Error parsing recovery clock: %v", err)
		}
		rns, err := strconv.ParseInt(c.RNs, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Error while parsing recovery ns: %v", err)
		}
		problem.RTimestamp = time.Unix(rsec, rns)
		problem.CorrelationID = c.CorrelationID
		problem.UserID, err = strconv.Atoi(c.UserID)
		if err != nil {
			return nil, fmt.Errorf("Error while parsing userID: %v", err)
		}*/
	problem.Name = c.Name
	//	problem.Ack = c.Ack
	problem.Severity, err = strconv.Atoi(c.Severity)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing Severity: %v", err)
	}
	//	problem.Suppressed = c.Suppressed

	return problem, nil
}
