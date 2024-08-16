package main

import (
	"currencyconverter/lib"
	"fmt"
	"os"

	"github.com/charmbracelet/huh/spinner"
)

func main() {
	appId := os.Getenv("OPENEXCHANGE_APP_ID")

	if appId == "" {
		fmt.Println("OPENEXCHANGE_APP_ID environment variable not set")
		return
	}

	var currencies map[string]string
	currencyAction := func() {
		currencies = lib.GetCurrencies()
	}

	_ = spinner.New().
		Title("Loading currencies...").
		Action(currencyAction).
		Run()

	if currencies == nil {
		return
	}

	amount, from, to := lib.RunForm(currencies)

	var rates map[string]float64
	ratesAction := func() {
		rates = lib.GetLatestRates()
	}

	_ = spinner.New().
		Title("Loading latest rates...").
		Action(ratesAction).
		Run()

	if rates == nil {
		return
	}

	rateFrom := rates[from]
	rateTo := rates[to]

	converted := amount * (rateTo / rateFrom)

	output := fmt.Sprintf("%.2f %s", converted, to)
	fmt.Println(output)
}
