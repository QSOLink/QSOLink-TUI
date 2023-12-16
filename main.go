package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/QSOLink/QSOLink-TUI/qso"

	"github.com/joho/godotenv"
)

func main() {

	var data qso.ResponseStruct
	// use qso.ResponseStruct to unmarshal JSON response

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get API URL from environment variable
	apiURL := os.Getenv("API_URL")
	fmt.Println(apiURL)

	// Make GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Handle response
	if resp.StatusCode != http.StatusOK {
		log.Fatal("API error: ", resp.Status)
	}

	// Decode JSON response
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}
}
