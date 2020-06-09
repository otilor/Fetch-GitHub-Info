package githubInfo

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", r)
}

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
	}
	if r.Method == "POST" {
		_ = r.ParseForm()
		username := r.FormValue("githubUsername")
		logrus.Println("Getting details for ", username)
		getDetails(username)
	}
}