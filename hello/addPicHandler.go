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
	g := Picture{
		Title: r.FormValue("title"),
		Url:   r.FormValue("url"),
		Date:  time.Now(),
	}
	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Picture", nil), &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/get", http.StatusFound)
}
