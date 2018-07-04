package zabbix

import (
	"os"
	"testing"
	"time"
	"strconv"
)

var session *Session

func GetTestCredentials(t *testing.T) (username string, password string, url string, timeout time.Duration) {
	url = os.Getenv("ZBX_URL")
	if url == "" {
		url = "http://localhost:8080/api_jsonrpc.php"
	}

	username = os.Getenv("ZBX_USERNAME")
	if username == "" {
		username = "Admin"
	}

	password = os.Getenv("ZBX_PASSWORD")
	if password == "" {
		password = "zabbix"
	}

	timeout_s := os.Getenv("ZBX_TIMEOUT")
	if timeout_s == "" {
		timeout = 5 * time.Second
	} else {
		timeout_i, err := strconv.Atoi(timeout_s)
		if err != nil {
			t.Fatalf("Error converting ZBX_TIMEOUT to int: %v", err)
		}
		timeout = time.Duration(timeout_i) * time.Second
	}

	return username, password, url, timeout
}

func GetTestSession(t *testing.T) *Session {
	var err error
	if session == nil {
		username, password, url, timeout := GetTestCredentials(t)

		session, err = NewSession(url, username, password, timeout)
		if err != nil {
			t.Fatalf("Error creating a session: %v", err)
		}
	}

	return session
}

func TestSession(t *testing.T) {
	s := GetTestSession(t)

	if s.Version() == "" {
		t.Errorf("No API version found for session")
	}
}
