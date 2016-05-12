package main

import (
  log "github.com/Sirupsen/logrus"
)

func main() {
  log.WithFields(log.Fields{
    "Test": "Hello it worked",
  }).Info("Hello")
}
