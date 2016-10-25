package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Stock struct {
	T string `json:"t"`
	L string `json:"l"`
}

func getStockData(ticker string) {
	url := "https://www.google.com/finance/info?infotype=infoquoteall&q=" + ticker
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Cleaning up the json response a little bit, as it started with '//'
		body = body[bytes.IndexRune(body, '['):]

		// The json response is in an array at the top level, so the struct here needs to be one as well.
		var stock []Stock

		// assuming no error, populates stock struct as if calling: stock = Stock [ {T: "data", L: "moreData" } ]
		err = json.Unmarshal(body, &stock)
		if err != nil {
			log.Fatal(err)

			// stock struct in an array, so can't be accessed simply by stock.T.
		} else {
			fmt.Printf("Company name: %s\nToday's price: %s\n", stock[0].T, stock[0].L)
		}
	}
}

func main() {

	tickerCommand := flag.NewFlagSet("ticker", flag.ExitOnError)

	tickerPointer := tickerCommand.String("ticker", "", "Ticker symbol to look up")

	if len(os.Args) < 2 {
		fmt.Println("ticker subcommand is required")
		os.Exit(1)
	}

	// Leaving in case I want to add more detailed flag sets
	switch os.Args[1] {
	case "ticker":

		tickerCommand.Parse(os.Args[2:])

	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if tickerCommand.Parsed() {

		// Required flags
		if *tickerPointer == "" {
			tickerCommand.PrintDefaults()
			os.Exit(1)
		} else {
			getStockData(*tickerPointer)
		}
	}
}
