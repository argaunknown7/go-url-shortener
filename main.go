package main

import (
  "fmt"
  "log"
  "net/http"
  "net/url"
  "strings"
)

var urlMap = make(map[string]string)

func shortenHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    longURL := r.FormValue("url")
    shortKey := generateShortKey()
    urlMap[shortKey] = longURL
    fmt.Fprintf(w, "http://localhost:8080/%s", shortKey)
  }
}

func expandHandler(w http.ResponseWriter, r *http.Request) {
  shortKey := strings.TrimPrefix(r.URL.Path, "/")
  longURL, exists := urlMap[shortKey]
  if exists {
    http.Redirect(w, r, longURL, http.StatusMovedPermanently)
  } else {
    http.NotFound(w, r)
  }
}

func generateShortKey() string {
  return "short" // Simple placeholder
}

func main() {
  http.HandleFunc("/", expandHandler)
  http.HandleFunc("/shorten", shortenHandler)

  log.Println("URL shortener running on http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}