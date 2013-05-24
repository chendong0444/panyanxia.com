package hello

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Picture struct {
	Author  string
	Title   string
	Url	string
	Date    time.Time
}

func init() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/root", rootHandler)
	http.HandleFunc("/sign", signHandler)
	http.HandleFunc("/get", getPicHandler)
	http.HandleFunc("/add", addPicHandler)
}

func addPicHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	g := Picture{
		Title:   r.FormValue("title"),
		Url: r.FormValue("url"),
		Date:    time.Now(),
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

func getPicHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Picture").Order("-Date").Limit(10)
	pictures := make([]Picture, 0, 10)
	if _, err := q.GetAll(c, &pictures); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := guestbookTemplate.Execute(w, pictures); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var guestbookTemplate = template.Must(template.New("picture").Parse(guestbookTemplateHTML))

const guestbookTemplateHTML = `
<html>
  <body>
    {{range .}}
      <pre>{{.Title}}</pre>
      <img src="{{.Url}}" alt="{{.Title}}" >
    {{end}}
    <form action="/add" method="post">
      <div><textarea name="title" rows="3" cols="60"></textarea></div>
	  <div><textarea name="url" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="submit"></div>
    </form>
  </body>
</html>
`

func loginHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
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
	fmt.Fprintf(w, "Hello, %v!", u)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, guestbookForm)
}

func signHandler(w http.ResponseWriter, r *http.Request) {
	err := signTemplate.Execute(w, r.FormValue("content"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

const guestbookForm = `
<html>
  <body>
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`

var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML))

const signTemplateHTML = `
<html>
  <body>
    <p>You wrote:</p>
    <pre>{{.}}</pre>
  </body>
</html>
`
