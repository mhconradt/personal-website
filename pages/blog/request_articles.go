package blog

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/mhconradt/personal_website/cfg"
	"github.com/mhconradt/personal_website/pages"
	"io/ioutil"
	"net/http"
	"time"
)

func (m ArticleSnippet) Since() string {
	fmt.Println("called")
	return pages.Since(m.Timestamp)
}

func GetRoute(r *http.Request) string {
	q := r.URL.Query()
	_, ok := q["term"]
	if ok {
		return r.RequestURI
	}
	q.Set("index", "date")
	return "/articles?" + q.Encode()
}

func RequestArticles(r *http.Request) (sr SearchResults, code int) {
	fqp := cfg.ApiEndpoint + GetRoute(r)
	fmt.Println(fqp)
	res, err := http.Get(fqp)
	if err != nil {
		fmt.Println(err)
		return sr, 404
	}
	start := time.Now().UnixNano()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return sr, 500
	}
	if err = proto.Unmarshal(buf, &sr); err != nil {
		return sr, 500
	}
	end := time.Now().UnixNano()
	fmt.Println("Decode duration: ", end - start)
	// Word limit on index
	if err = res.Body.Close(); err != nil {
		return sr, 500
	}
	fmt.Println(sr.Cursor.Term)
	sr.Cursor.Index = func(m map[string][]string, k, d string) string {
		if v, ok := m[k]; ok && len(v) > 0 {
			return v[0]
		}
		return d
	}(r.URL.Query(), "index", "date")
	return sr, 200
}
