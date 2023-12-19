package main

import (
	"encoding/json"
	"fmt"
	"github.com/QSOLink/QSOLink-TUI/qso"
	"log"
	"net/http"

	_ "github.com/QSOLink/QSOLink-TUI/qso"
)

func getQSOs(apiURL string) {

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var contacts []qso.Contact

	if err := json.NewDecoder(resp.Body).Decode(&contacts); err != nil {
		log.Fatal(err)
	}

	for _, contact := range contacts {
		fmt.Println(contact.ID, contact.Callsign, contact.QslR, contact.QslS)
	}

}
