package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseUrl = "https://openexchangerates.org/api/"

func GetCurrencies() map[string]string {
	appId := os.Getenv("OPENEXCHANGE_APP_ID")
	resp, err := http.Get(baseUrl + "currencies.json" + "?app_id=" + appId)

	if err != nil {
		fmt.Println("Error getting currencies")
		return nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	var currencies map[string]string

	err = json.Unmarshal(body, &currencies)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil
	}

	return currencies
}

func GetLatestRates() map[string]float64 {
	appId := os.Getenv("OPENEXCHANGE_APP_ID")
	resp, err := http.Get(baseUrl + "latest.json" + "?app_id=" + appId)

	if err != nil {
		fmt.Println("Error getting rates")
		return nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	type ExchangeRatesResponse struct {
		Disclaimer string             `json:"disclaimer"`
		License    string             `json:"license"`
		Timestamp  int64              `json:"timestamp"`
		Base       string             `json:"base"`
		Rates      map[string]float64 `json:"rates"`
	}

	var rates ExchangeRatesResponse

	err = json.Unmarshal(body, &rates)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil
	}

	return rates.Rates
}
