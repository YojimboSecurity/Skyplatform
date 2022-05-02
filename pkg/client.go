package pkg

import (
	"os"

	"github.com/YojimboSecurity/skytap-sdk-go/api"
)

func NewClient() *api.SkytapClient {
	token := os.Getenv("SKYTAP_TOKEN")
	user := os.Getenv("SKYTAP_USER")
	if token == "" || user == "" {
		panic("SKYTAP_TOKEN and SKYTAP_USER must be set")
	}
	client := api.NewSkytapClient(user, token)
	return client
}
