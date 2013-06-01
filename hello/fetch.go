package hello

import (
	"appengine"
	"appengine/urlfetch"
	"net/http"
)

func init() {
	http.HandleFunc("/fetch", fetchHandler)
}

func fetchHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get(r.FormValue("url"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Write(w)

}
