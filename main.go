package main

import (
	"log"
	"ui/pkg/backend"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	b, err := backend.NewBackend()
	if err != nil {
		log.Fatal(err)
	}

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "ui",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(b.Kubernetes.Clusters)
	// app.Bind(b.Kubernetes.Machines)
	app.Run()
}