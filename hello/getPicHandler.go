package hello

import (
	"appengine"
	"appengine/datastore"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/get", getPicHandler)
	http.HandleFunc("/getPicJson", getPicJsonHandler)
}

func getPicHandler(w http.ResponseWriter, r *http.Request) {
	pictures := getPictures(w, r)
	if err := getPicTemplate.Execute(w, pictures); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getPicJsonHandler(w http.ResponseWriter, r *http.Request) {
	pictures := getPictures(w, r)
	strJson, err := json.Marshal(pictures)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	fmt.Fprintf(w, "%s", strJson)
}

var getPicTemplate = template.Must(template.New("picture").Parse(getPicTemplateHTML))

const getPicTemplateHTML = `
<html>
  <body>
    <form action="/add" method="post">
      <div><label>picture title</label><input type="text" name="title"></div>
	  <div><label>picture url</label><input type="text" name="url"></textarea></div>
      <div><input type="submit" value="新图"></div>
    </form>

    {{range .}}
      <p>{{.Title}}</p>
      <img src="{{.Url}}" alt="{{.Title}}" >
    {{end}}

  </body>
</html>
`
const PerRowPictures = 4

func getPictures(w http.ResponseWriter, r *http.Request) []Picture {
	c := appengine.NewContext(r)
	page, err := strconv.ParseInt(r.FormValue("page"), 10, 0)
	if err != nil {
		page = 0
	}
	q := datastore.NewQuery("Picture").Order("-Date").Offset(int(page * PerRowPictures)).Limit(PerRowPictures)
	pictures := make([]Picture, 0, PerRowPictures)
	if _, err := q.GetAll(c, &pictures); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return pictures
}
