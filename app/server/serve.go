package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

type Manifest struct {
	MainCSS struct {
		File string `json:"file"`
		Src  string `json:"src"`
	} `json:"src/main.css"`
	MainTS struct {
		File string `json:"file"`
		Src  string `json:"src"`
	} `json:"src/main.ts"`
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	indexHTML, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := os.ReadFile("static/manifest.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var manifest Manifest
	err = json.Unmarshal(jsonData, &manifest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var html bytes.Buffer
	indexHTML.Execute(&html, struct {
		CSS string
		JS  string
	}{
		CSS: manifest.MainCSS.File,
		JS:  manifest.MainTS.File,
	})
	w.Write(html.Bytes())
}
