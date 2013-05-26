package hello

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"net/http"
	"time"
)

type Picture struct {
	Author string
	Title  string
	Url    string
	Date   time.Time
}

func init() {
	http.HandleFunc("/add", addPicHandler)
}

func addPicHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	//add pic must login
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	title := r.FormValue("title")
	url := r.FormValue("url")
	if title == "" || url == "" {
		http.Redirect(w, r, "/get", http.StatusFound)
	}
	g := Picture{
		Author: u.String(),
		Title:  title,
		Url:    url,
		Date:   time.Now(),
	}

	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Picture", nil), &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/get", http.StatusFound)
}
