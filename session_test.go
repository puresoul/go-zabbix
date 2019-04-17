package zabbix

import (
	"os"
	"strconv"
	"testing"
	"time"
)

var session *Session

func GetTestCredentials(t *testing.T) (username string, password string, url string, timeout time.Duration) {
	url = os.Getenv("ZBX_URL")
	if url == "" {
		url = "http://192.168.33.11:8040/api_jsonrpc.php"
	}

	username = os.Getenv("ZBX_USERNAME")
	if username == "" {
		username = "Admin"
	}

	password = os.Getenv("ZBX_PASSWORD")
	if password == "" {
		password = "zabbix"
	}

	timeoutS := os.Getenv("ZBX_TIMEOUT")
	if timeoutS == "" {
		timeout = 5 * time.Second
	} else {
		timeoutI, err := strconv.Atoi(timeoutS)
		if err != nil {
			t.Fatalf("Error converting ZBX_TIMEOUT to int: %v", err)
		}
		timeout = time.Duration(timeoutI) * time.Second
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
