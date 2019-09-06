package blog

import (
	"encoding/json"
	"fmt"
	"github.com/mhconradt/personal_website/cfg"
	"net/http"
)

func RequestArticles(r *http.Request) (SearchResults, int) {
	sr := SearchResults{}
	fqp := cfg.ApiEndpoint + r.RequestURI
	fmt.Println(fqp)
	req, err := http.Get(fqp)
	if err != nil {
		fmt.Println(err)
		return sr, 404
	}
	err = json.NewDecoder(req.Body).Decode(&sr)
	if err != nil {
		fmt.Println(err)
		return sr, 500
	}
	fmt.Println("results: ", sr)
	sr.Cursor.Index = func(m map[string][]string, k, d string) string {
		if v, ok := m[k]; ok && len(v) > 0 {
			return v[0]
		}
		return d
	}(r.URL.Query(), "index", "date")
	return sr, 200
}
