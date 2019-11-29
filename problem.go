package zabbix

import (
	"fmt"
	"time"
)

const (
	ProblemNotAcknowledged = iota
	ProblemAcknowledged
)

const (
	ProblemNotClassified = iota
	ProblemInformation
	ProblemWarning
	ProblemAverage
	ProblemHigh
	ProblemDisaster
)

const (
	ProblemNormal = iota
	ProblemSuppressed
)

type Problem struct {
	EventID       string
	//Source        string
	//Object        string
	ObjectID      int
	Clock         string
	Timestamp     time.Time
	//REventID      int
	//RTimestamp    time.Time
	//CorrelationID string
	//UserID        int
	Name          string
	//Ack           string
	Severity      int
	//Suppressed    int
}

type ProblemGetParams struct {
	GetParameters

	EventIDs           []string    `json:"eventids,omitempty"`
	GroupIDs           []string    `json:"groupids,omitempty"`
	HostIDs            []string    `json:"hostids,omitempty"`
	ObjectIDs          []string    `json:"objectids,omitempty"`
	ApplicationIDs     []string    `json:"applicationids,omitempty"`
	Source             string      `json:"source,omitempty"`
	Object             string      `json:"object,omitempty"`
	Acknowledged       bool        `json:"acknowledged,omitempty"`
	Suppressed         bool        `json:"suppressed,omitempty"`
	Severities         []int       `json:"severities,omitempty"`
	EvalTypeTags       int         `json:"evaltype,omitempty"`
	Tags               []string    `json:"tags,omitempty"`
	Recent             bool        `json:"recent,omitempty"`
	EventID_from       string      `json:"eventid_from,omitempty"`
	EventID_till       string      `json:"eventid_till,omitempty"`
	Time_from          string      `json:"time_from,omitempty"`
	Time_till          string      `json:"time_till,omitempty"`
	SelectAcknowledges SelectQuery `json:"SelectAcknowledges,omitempty"`
}

func (c *Session) GetProblems(params ProblemGetParams) ([]Problem, error) {
	problems := make([]jProblem, 0)
	err := c.Get("problem.get", params, &problems)
	if err != nil {
		return nil, err
	}

	if len(problems) == 0 {
		return nil, ErrNotFound
	}

	// map JSON Problems to Go Problems
	out := make([]Problem, len(problems))
	for i, jproblem := range problems {
		problem, err := jproblem.Problem()
		if err != nil {
			return nil, fmt.Errorf("Error mapping Problem %d in response: %v", i, err)
		}
		out[i] = *problem
	}
	return out, nil
}
