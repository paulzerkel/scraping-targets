package handlers

import (
  "encoding/json"
  "fmt"
  "html/template"
  "math/rand"
  "net/http"
  "time"
)

type PageOpts struct {
  Page string
  OnSale bool
}

type Product struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Description string `json:"description"`
  Price int `json:"price"`
  Weight int `json:"weight"`
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("%s\t%s\n", time.Now().String(), r.URL.String())
  if r.URL.Path != "/" {
    Error(w, r, http.StatusNotFound)
    return
  }

  opts := PageOpts{ Page: "index" }
  if rand.Intn(5) == 0 {
    opts.OnSale = true
  }
  processLayout(w, r, opts)
}

func About(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("%s\t%s\n", time.Now().String(), r.URL.String())
  opts := PageOpts{ Page: "about" }
  processLayout(w, r, opts)
}

func Products(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("%s\t%s\n", time.Now().String(), r.URL.String())
  opts := PageOpts{ Page: "products" }
  processLayout(w, r, opts)
}

func ProductData(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("%s\t%s\n", time.Now().String(), r.URL.String())
  product := Product{Id:1, Name:"asdf", Description:"cool thing", Price:1000, Weight:10}
  output, _ := json.MarshalIndent(&product, "", "\t")
  w.Header().Set("Content-Type", "application/json")
  w.Write(output)
}

func Error(w http.ResponseWriter, r *http.Request, status int) {
  fmt.Printf("%s\tERROR\t%s\n", time.Now().String(), r.URL.String())
  w.WriteHeader(status)

  opts := PageOpts{ Page: "error" }
  processLayout(w, r, opts)
}

func processLayout(w http.ResponseWriter, r *http.Request, pageOpts PageOpts) {
  extraFile := fmt.Sprintf("templates/%s.html", pageOpts.Page)
  t, _ := template.ParseFiles("templates/layout.html", extraFile)
  t.ExecuteTemplate(w, "layout", pageOpts)
}
