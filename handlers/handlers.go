package handlers

import (
  "fmt"
  "html/template"
  "net/http"
)

type PageOpts struct {
  Page string
  OnSale bool
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Index: %s\n", r.URL.String())
  if r.URL.Path != "/" {
    Error(w, r, http.StatusNotFound)
    return
  }

  opts := PageOpts{ Page: "index", OnSale: true }
  processLayout(w, r, opts)
}

func About(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Products: %s\n", r.URL.String())
  opts := PageOpts{ Page: "about" }
  processLayout(w, r, opts)
}

func Products(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Products: %s\n", r.URL.String())
  opts := PageOpts{ Page: "products" }
  processLayout(w, r, opts)
}


func Error(w http.ResponseWriter, r *http.Request, status int) {
  fmt.Printf("Error: %s\t%d\n", r.URL.String(), status)
  w.WriteHeader(status)

  opts := PageOpts{ Page: "error" }
  processLayout(w, r, opts)
}

func processLayout(w http.ResponseWriter, r *http.Request, pageOpts PageOpts) {
  extraFile := fmt.Sprintf("templates/%s.html", pageOpts.Page)
  t, _ := template.ParseFiles("templates/layout.html", extraFile)
  t.ExecuteTemplate(w, "layout", pageOpts)
}
