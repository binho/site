package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"os"
	_ "bytes"
	_ "fmt"
	_ "log"
)

func main() {

	m := martini.Classic()
	m.Use(martini.Static("assets"))

	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	m.Get("/about", func(r render.Render) {
		r.HTML(200, "about", nil)
	})

	m.Get("/work", func(r render.Render) {
		r.HTML(200, "work", nil)
	})

	m.Get("/work/:name", func(r render.Render, params martini.Params, req *http.Request) {
		pwd, _ := os.Getwd()
		filename := params["name"]
		filepath := (pwd + "/templates/work/" + filename + ".tmpl")

		if _, err := os.Stat(filepath); err == nil {
			r.HTML(200, ("work/" + filename), nil)
		} else {
			htmlParams := map[string]interface{}{"project_name": params["name"]}
			r.HTML(404, "not-found", htmlParams)
		}
	})

	// m.Get("/contact", func(r render.Render) {
	// 	r.HTML(200, "contact", nil)
	// })

	// m.Post("/contact/send", func(r render.Render, req *http.Request) {
	// 	e := email.NewEmail()
	// 	e.From = "Site <from@email.com>"
	// 	e.To = []string{"to@email.com"}
	// 	e.Subject = "Contato Site"

	// 	message := "Name: " + req.FormValue("name") +
	// 		"\nEmail: " + req.FormValue("email") +
	// 		"\nMessage: " + req.FormValue("message") +
	// 		"\nIP: " + req.RemoteAddr

	// 	e.Text = []byte(message)
	// 	//e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	// 	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "your@email.com", "yourpass", "smtp.gmail.com"))

	// 	htmlParams := map[string]interface{}{"success": true}
	// 	r.HTML(200, "contact", htmlParams)
	// })

	m.RunOnAddr(":8080")
}
