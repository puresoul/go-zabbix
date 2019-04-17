package zabbix

import "testing"

func TestProblem(t *testing.T) {
	session := GetTestSession(t)

	params := ProblemGetParams{}

	problems, err := session.GetProblems(params)
	if err != nil {
		t.Fatalf("Error while getting problems: %v", err)
	}
	if len(problems) == 0 {
		t.Fatalf("No problems found")
	}
	for i, problem := range problems {
		t.Log(problem.EventID, problem.Name)
		if problem.EventID == "" {
			t.Fatalf("Problem %v without EventID found", i)
		}
	}
}
