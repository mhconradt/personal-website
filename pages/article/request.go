package article

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhconradt/personal_website/cfg"
	"github.com/mhconradt/personal_website/pages"
	"net/http"
)

type Article struct {
	ID        int      `json:"id"`
	Timestamp pages.Timestamp    `json:"timestamp,omitempty"`
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Topics    []string `json:"topics,omitempty"`
	Views     int      `json:"views,omitempty"`
}

func GetArticle(r *http.Request) (Article, int) {
	a := Article{}
	v := mux.Vars(r)
	articleId := v["id"]
	subpath := fmt.Sprintf("/articles/%v", articleId)
	url := cfg.ApiEndpoint + subpath
	res, err := http.Get(url)
	if err != nil {
		return a, 404
	}
	err = json.NewDecoder(res.Body).Decode(&a)
	fmt.Println(a)
	if err != nil {
		return a, 500
	}
	return a, 200
}
