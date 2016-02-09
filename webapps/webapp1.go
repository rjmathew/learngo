package main
import (
	"net/http"
	"text/template"
	"fmt"
)

func main() {
	fmt.Println("Starting up the server..")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		w.Header().Add("Content Type", "text/html")
		tmpl, err := template.New("test").Parse(doc)
		if (err == nil) {
			tmpl.Execute(w, req.URL.Path)
		} else {
			fmt.Println(err)
		}

	})
	http.ListenAndServe(":8080", nil)
}

const doc = `<html>
<head></head>
<body>
	<h1>Hello {{.}}</h1>
</body>
</html>`
