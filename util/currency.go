package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var currencyRates map[string]float64

func init() {
	updateCurrencyRates()
}

func updateCurrencyRates() {
	resp, err := http.Get("https://api.exchangerate-api.com/v4/latest/INR")
	if err != nil {
		fmt.Println("Error fetching currency rates:", err)
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding currency rates:", err)
		return
	}

	rates := data["rates"].(map[string]interface{})
	currencyRates = make(map[string]float64)
	for key, value := range rates {
		currencyRates[key] = value.(float64)
	}
}

func ConvertCurrency(amount float64, from, to string) float64 {
	if from == to {
		return amount
	}

	rate, ok := currencyRates[to]
	if !ok {
		fmt.Println("Unsupported currency:", to)
		return amount
	}

	return amount * rate
}
