package githubInfo

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", r)
}

func Search(w http.ResponseWriter, r *http.Request) {

}