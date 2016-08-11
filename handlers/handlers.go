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
  products := []Product{
    Product{Id:1, Name:"QualityTurbo", Description:"When a generic turbo won't do, go for the QualityTurbo!", Price:23000, Weight:20},
    Product{Id:2, Name:"StarDynamic", Description: "This is the best telescope that money can buy. Find new stars even in the midst of light pollution", Price: 7700, Weight: 33},
    Product{Id:3, Name:"Magictronics", Description: "Magic + Electronics = Magictronics. Enough said.", Price:10000, Weight:12},
    Product{Id:4, Name:"CentreWork", Description: "Find the absolute center of any object. Perfect for the modern machine shop.", Price: 9030, Weight: 55555},
    Product{Id:5, Name:"WorldLogix", Description: "World shipping logistics: solved. Save a bundle on transporting your goods anywhere in the world", Price: 82010, Weight: 18},
  }
  output, _ := json.MarshalIndent(&products, "", "\t")
  w.Header().Set("Content-Type", "application/json")

  time.Sleep(3 * time.Second)
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
