package hello

import (
	"appengine"
	"appengine/urlfetch"
	"appengine/user"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func init() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/root", rootHandler)
	http.HandleFunc("/sign", signHandler)
}

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
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get("http://t66y.com/htm_data/8/1306/917581.html")
	//resp, err := http.Get("http://www.baidu.com")
	// resp, err := http.Get("http://www.163.com")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
	}

	src := string(body)

	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	input := re.ReplaceAllString(src, "\n")

	//out := make([]byte, len(input))
	//out = out[:]
	//iconv.Convert(input, out, "gb2312", "utf-8")

	fmt.Fprint(w, strings.TrimSpace(input))

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
