package main

import (
	"currencyconverter/lib"
	"fmt"
	"os"
)

func main() {
	appId := os.Getenv("OPENEXCHANGE_APP_ID")

	if appId == "" {
		fmt.Println("OPENEXCHANGE_APP_ID environment variable not set")
		return
	}

	currencies := lib.GetCurrencies()

	if currencies == nil {
		return
	}

	amount, from, to := lib.CreateForm(currencies)

	rates := lib.GetLatestRates()

	if rates == nil {
		return
	}

	rateFrom := rates[from]
	rateTo := rates[to]

	converted := amount * (rateTo / rateFrom)

	fmt.Printf("%.2f %s", converted, to)
}
