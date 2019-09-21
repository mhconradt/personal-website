package article

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/mhconradt/personal_website/cfg"
	"github.com/mhconradt/personal_website/pages"
	"io/ioutil"
	"net/http"
	"time"
)

func (m Article) Since() string {
	return pages.Since(m.Timestamp)
}

func GetArticle(r *http.Request) (*Article, int) {
	a := new (Article)
	v := mux.Vars(r)
	articleId := v["id"]
	route := fmt.Sprintf("/articles/%v", articleId)
	url := cfg.ApiEndpoint + route
	res, err := http.Get(url)
	if err != nil {
		return a, 404
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return a, 500
	}
	start := time.Now().UnixNano()
	err = proto.Unmarshal(b, a)
	end := time.Now().UnixNano()
	if err != nil {
		return a, 500
	}
	fmt.Println("article decoding took: ", end - start)
	err = res.Body.Close()
	if err != nil {
		return a, 500
	}
	return a, 200
}
