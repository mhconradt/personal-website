package main

import (
	"github.com/gorilla/mux"
	"github.com/mhconradt/personal_website/pages"
	"github.com/mhconradt/personal_website/pages/article"
	"github.com/mhconradt/personal_website/pages/blog"
	"log"
	"net/http"
	"path/filepath"
)

/*
/
/about
/resume
/personal_website
	/articles - do same thing as /recent
		/by
			/recent - list articles by date
			/topic - allow searching by topic
			/search - allow searching by word/phrase
		/<articleId> - display one article
*/

func PopulateFolder(dir, ext string, files ...string) []string {
	fullpaths := make([]string, len(files))
	for i, name := range files {
		fullpaths[i] = filepath.Join(dir, name) + ext
	}
	return fullpaths
}

func main() {
	r := mux.NewRouter()
	// PathPrefix is responsible for routing
	/*
		Blog list template params:
		SelectedIndex
		Results
		Cursor
	*/
	aboutTemplates := PopulateFolder("templates/about", ".gohtml", "about", "language", "languages")
	articleTemplates := PopulateFolder("templates/article", ".gohtml", "body", "header", "article")
	blogListTemplates := PopulateFolder("templates/blog", ".gohtml", "article_list", "indices", "pagination", "search_bar", "search_results", "search_result", "empty")
	fs, err := NewFileServer("./static")
	if err != nil {
		log.Fatalf("failed to launch file server: %v", err)
	}
	r.PathPrefix("/static/").Handler(fs)
	r.HandleFunc("/", WrapHandler(pages.EmptyHandler, "templates/landing_page.gohtml"))
	r.HandleFunc("/about", WrapAboutHandler(pages.FetchAbout, aboutTemplates...))
	r.HandleFunc("/resume", WrapHandler(pages.EmptyHandler, "templates/resume.gohtml"))
	r.HandleFunc("/articles", WrapSearchHandler(blog.RequestArticles, blogListTemplates...))
	r.HandleFunc("/articles/{id}", WrapGetHandler(article.GetArticle, articleTemplates...))
	log.Fatal(http.ListenAndServe(":8080", r))
}
