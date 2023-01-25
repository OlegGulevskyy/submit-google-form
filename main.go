package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	PREF_RESP_ID            = "1FAIpQLSe2GMA1Tj1viLoslhmyPILzxN_aqTvKxEHzY-RUrE6_CrPAQQ"
	PRODUCT_QUESTION_ID     = "entry.491128001"
	PRODUCT_QUESTION_ANSWER = "ProductA"
	PRICE_QUESTION_ID       = "entry.558317"
	PRICE_QUESTION_ANSWER   = "100"
)

func main() {
	baseUrl := fmt.Sprint(
		"https://docs.google.com/forms/d/e/",
		PREF_RESP_ID,
		"/formResponse",
		"?usp=pp_url",
	)
	url := fmt.Sprint(
		baseUrl,
		"&",
		PRODUCT_QUESTION_ID, "=", PRODUCT_QUESTION_ANSWER,
		"&",
		PRICE_QUESTION_ID, "=", PRICE_QUESTION_ANSWER,
	)

	resp, err := http.Post(url, "application/x-www-form-urlencoded", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response created with status => ", resp.Status)
}
