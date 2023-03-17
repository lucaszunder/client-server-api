package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const requestTimeout = 300 * time.Millisecond
const currencyServer = "http://localhost:8080"

type usdToReal struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type Currency struct {
	ID        string `json:"id"`
	usdToReal `json:"USDBRL"`
}

func main() {
	currencyValue := getCurrency()
	saveToFile(currencyValue)
}

func getCurrency() string {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", currencyServer+"/cotacao", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	currency := Currency{}

	err = json.Unmarshal(body, &currency)
	return currency.Bid
}

func saveToFile(currencyValue string) {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Write([]byte(fmt.Sprintf("Dólar: %v\n", currencyValue)))
	if err != nil {
		panic(err)
	}
}
