package main

import (
  log "github.com/Sirupsen/logrus"
  "net/http"
  "github.com/matryer/respond"
  "github.com/gorilla/mux"
)

func main() {
  log.WithFields(log.Fields{
    "Test": "Hello it worked",
  }).Info("Hello")

  s := http.Server{
    Addr: ":8080",
  }

  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  http.Handle("/", r)

  s.ListenAndServe()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  log.WithFields(log.Fields{
    "Request": r,
  }).Debug("New Request")
  respond.With(w, r, http.StatusOK, "Test")
}
