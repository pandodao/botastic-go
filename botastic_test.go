package botastic

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	client *Client
}

func TestSuite(t *testing.T) {
	client := New(os.Getenv("BOTASTIC_APP_ID"), os.Getenv("BOTASTIC_APP_SECRET"), WithDebug(true), WithHost(os.Getenv("BOTASTIC_HOST")))
	if client.appID == "" || client.appSecret == "" || client.host == "" {
		t.SkipNow()
	}

	suite.Run(t, &Suite{client: client})
}
