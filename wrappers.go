package main

import (
	"fmt"
	"github.com/mhconradt/personal_website/pages"
	"github.com/mhconradt/personal_website/pages/article"
	"github.com/mhconradt/personal_website/pages/blog"
	"html/template"
	"net/http"
)

func WrapGetHandler(h func(r *http.Request) (article.Article, int), subtemplates ...string) func(w http.ResponseWriter, r *http.Request) {
	withMain := append([]string{"templates/main.gohtml"}, subtemplates...)
	t, err := template.ParseFiles(withMain...)
	if err != nil {
		fmt.Println("template error", err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		params, code := h(r)
		// Not success
		if (code / 100) != 2 {
			// return error page
		}
		err = t.Execute(w, params)
	}
}

func WrapHandler(h func(r *http.Request) (interface{}, int), subtemplates ...string) func(w http.ResponseWriter, r *http.Request) {
	withMain := append([]string{"templates/main.gohtml"}, subtemplates...)
	t, err := template.ParseFiles(withMain...)
	if err != nil {
		fmt.Println(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		params, code := h(r)
		// Not success
		if (code / 100) != 2 {
			// return error page
		}
		err = t.Execute(w, params)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func WrapAboutHandler(h func(r *http.Request) (pages.AboutInfo, int), subtemplates ...string) func(w http.ResponseWriter, r *http.Request) {
	withMain := append([]string{"templates/main.gohtml"}, subtemplates...)
	t, err := template.ParseFiles(withMain...)
	if err != nil {
		fmt.Println(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		params, code := h(r)
		// Not success
		if (code / 100) != 2 {
			// return error page
		}
		err = t.Execute(w, params)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func WrapSearchHandler(h func(r *http.Request) (blog.SearchResults, int), subtemplates ...string) func(w http.ResponseWriter, r *http.Request) {
	withMain := append([]string{"templates/main.gohtml"}, subtemplates...)
	t, err := template.ParseFiles(withMain...)
	return func(w http.ResponseWriter, r *http.Request) {
		params, code := h(r)
		// Not success
		if (code / 100) != 2 {
			// return error page
		}
		err = t.Execute(w, params)
	}
}
