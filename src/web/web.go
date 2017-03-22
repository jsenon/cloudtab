package web

import (
	"html/template"
	"net/http"
	// "strconv"
	"db"
	// "encoding/json"
	"fmt"
)

func Index(res http.ResponseWriter, req *http.Request) {
	// var rs Server
	rs, err := db.GetAll()

	// t := template.Must(template.New("tmpl").Parse(tmpl))
	t, _ := template.ParseFiles("templates/index.html")

	t.Execute(res, rs)

	fmt.Println("err", err)
	fmt.Println("rs", rs)

}
