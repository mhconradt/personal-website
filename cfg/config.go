package cfg

import "os"

var ApiEndpoint string

func init() {
	if endpt, ok := os.LookupEnv("API_ENDPOINT"); !ok {
		ApiEndpoint = "http://localhost:3000"
	} else {
		ApiEndpoint = endpt
	}
}
