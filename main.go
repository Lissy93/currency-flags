package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type Currency struct {
	Name   string `json:"name" yaml:"name"`
	Symbol string `json:"symbol" yaml:"symbol"`
	Code   string `json:"code" yaml:"code"`
	Flag   string `json:"flag" yaml:"flag"`
}

var currencies []Currency

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/currencies", getAllCurrencies).Methods("GET")
	router.HandleFunc("/currency/{code}", getCurrencyByCode).Methods("GET")
	router.HandleFunc("/currency/{code}/flag", getCurrencyFlagByCode).Methods("GET")

	loadCurrenciesFromYaml()

	log.Fatal(http.ListenAndServe(":8000", router))
}

func loadCurrenciesFromYaml() {
	yamlFile, err := ioutil.ReadFile("currencies.yaml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &currencies)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}
}

func getAllCurrencies(w http.ResponseWriter, r *http.Request) {
	jsonCurrencies, err := json.Marshal(currencies)
	if err != nil {
		log.Fatalf("Failed to marshal currencies to JSON: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonCurrencies)
}

func getCurrencyByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := strings.ToUpper(vars["code"])

	for _, currency := range currencies {
		if currency.Code == code {
			jsonCurrency, err := json.Marshal(currency)
			if err != nil {
				log.Fatalf("Failed to marshal currency to JSON: %v", err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonCurrency)
			return
		}
	}

	http.NotFound(w, r)
}

func getCurrencyFlagByCode(w http.ResponseWriter, r *http.Request) {
	// Get currency code from request URL
	params := mux.Vars(r)
	code := strings.ToUpper(params["code"])

	// Find currency with matching code
	var currency Currency
	for _, c := range currencies {
		if c.Code == code {
			currency = c
			break
		}
	}

	// Check if currency was found
	if currency.Code == "" {
		http.NotFound(w, r)
		return
	}

	// Get image file
	resp, err := http.Get(currency.Flag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Set response header
	w.Header().Set("Content-Type", "image/png")

	// Write image file to response
	_, err = fmt.Fprint(w, resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
