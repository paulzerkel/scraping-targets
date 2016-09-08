package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/paulzerkel/scraping-targets/handlers"
)


func main() {
  fmt.Println("starting server")

  mux := http.NewServeMux()
  files := http.FileServer(http.Dir("public"))
  mux.Handle("/static/", http.StripPrefix("/static/", files))

  mux.HandleFunc("/about", handlers.About)
  mux.HandleFunc("/products", handlers.Products)
  mux.HandleFunc("/product-data", handlers.ProductData)
  mux.HandleFunc("/", handlers.Index)

  server := http.Server {
    Addr: "127.0.0.1:8080",
    Handler: mux,
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 5 * time.Second,
  }
  server.ListenAndServe()
}
