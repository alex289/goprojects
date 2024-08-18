package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type CalculatorClient struct {
    client http.Client
	baseUrl string
}

func NewCalculatorClient(baseUrl string) CalculatorClient {
    return CalculatorClient{
        client: http.Client{},
		baseUrl: baseUrl,
    }
}

func (c CalculatorClient) basicRequest(url string, a float64, b float64) (int, error) {
	body := map[string]float64{
		"number1": a,
		"number2": b,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

    res, err := c.client.Post(url, "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	var result struct {
		Result int `json:"result"`
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	return result.Result, nil
}

func (c CalculatorClient) Add(a float64, b float64) (int, error) {
    return c.basicRequest(c.baseUrl + "add", a, b)
}

func (c CalculatorClient) Divide(a float64, b float64) (int, error) {
    return c.basicRequest(c.baseUrl + "divide", a, b)
}

func (c CalculatorClient) Multiply(a float64, b float64) (int, error) {
    return c.basicRequest(c.baseUrl + "multiply", a, b)
}

func (c CalculatorClient) Subtract(a float64, b float64) (int, error) {
    return c.basicRequest(c.baseUrl + "subtract", a, b)
}

func (c CalculatorClient) Sum(nums []float64) (int, error) {
    body := map[string][]float64{
		"numbers": nums,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

    res, err := c.client.Post(c.baseUrl + "sum", "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	var result struct {
		Result int `json:"result"`
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	return result.Result, nil
}