package main

import (
  "fmt"
  "log"
  "os"

	"github.com/joho/godotenv"
)

func main() {
	// init
	err := godotenv.Load()
	  if err != nil {
    log.Fatal("Error loading .env file")
  }

	qsolinkserver := os.Getenv("QSOLINK_SERVER")

  fmt.Println(qsolinkserver)
}


