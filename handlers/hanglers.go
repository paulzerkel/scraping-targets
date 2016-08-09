package handlers

import (
  "fmt"
  "html/template"
  "net/http"
)

type PageOpts struct {
  Page string
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Index: %s\n", r.URL.String())
  if r.URL.Path != "/" {
    Error(w, r, http.StatusNotFound)
    return
  }

  processLayout(w, r, PageOpts{"index"})
}

func About(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Products: %s\n", r.URL.String())
  processLayout(w, r, PageOpts{"about"})
}

func Products(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Products: %s\n", r.URL.String())
  processLayout(w, r, PageOpts{"products"})
}


func Error(w http.ResponseWriter, r *http.Request, status int) {
  fmt.Printf("Error: %s\t%d\n", r.URL.String(), status)
  w.WriteHeader(status)

  processLayout(w, r, PageOpts{"error"})
}

func processLayout(w http.ResponseWriter, r *http.Request, pageOpts PageOpts) {
  extraFile := fmt.Sprintf("templates/%s.html", pageOpts.Page)
  t, _ := template.ParseFiles("templates/layout.html", extraFile)
  t.ExecuteTemplate(w, "layout", pageOpts)
}
