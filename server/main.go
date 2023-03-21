package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
	"time"
)

const databaseParams = "./currency.db"
const currencySupplyerUri = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
const requestTimeout = 20000 * time.Millisecond
const dbTimeout = 1000 * time.Millisecond

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
	http.HandleFunc("/cotacao", getCurrencyHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func getCurrencyHandler(w http.ResponseWriter, _ *http.Request) {
	currency, err := getCurrencyRequester()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("sqlite3", databaseParams)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = saveRecurrence(db, currency)
	if err != nil {
		panic(err)
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&currency)
}

func saveRecurrence(db *sql.DB, currency *Currency) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt, err := db.PrepareContext(ctx,
		`INSERT INTO currencies(
                       id, 
                       code, 
                       codein, 
                       name, 
                       high, 
                       low, 
                       varbid, 
                       pctchange, 
                       bid, 
                       ask, 
                       timestamp, 
                       create_date
                ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	values := []interface{}{
		currency.ID,
		currency.Code,
		currency.Codein,
		currency.Name,
		currency.High,
		currency.Low,
		currency.VarBid,
		currency.PctChange,
		currency.Bid,
		currency.Ask,
		currency.Timestamp,
		currency.CreateDate,
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, values...)
	if err != nil {
		return err
	}

	if ctx.Err() != nil {
		return ctx.Err()
	}

	return nil
}

func getCurrencyRequester() (*Currency, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", currencySupplyerUri, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("response: %v", res)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	currency := Currency{}

	err = json.Unmarshal(body, &currency)
	currency.ID = uuid.New().String()
	if err != nil {
		return nil, err
	}
	return &currency, nil
}
