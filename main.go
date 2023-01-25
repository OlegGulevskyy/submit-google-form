package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	PREF_RESP_ID = "1FAIpQLSe2GMA1Tj1viLoslhmyPILzxN_aqTvKxEHzY-RUrE6_CrPAQQ"
)

var responses = [][]string{
	{"entry.491128001", "ProductA"}, // Product
	{"entry.558317", "100"},         // Price
}

func main() {
	baseUrl := fmt.Sprint(
		"https://docs.google.com/forms/d/e/",
		PREF_RESP_ID,
		"/formResponse",
		"?usp=pp_url",
	)

	url := baseUrl
	for _, response := range responses {
		url = fmt.Sprint(
			url,
			"&",
			response[0], "=", response[1],
		)
	}

	resp, err := http.Post(url, "application/x-www-form-urlencoded", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response created with status => ", resp.Status)
}
