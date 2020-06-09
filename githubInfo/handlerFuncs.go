package githubInfo

import (
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
		details := getDetails(username)
		renderWithData(w, "search.html", details, r)

	}
}