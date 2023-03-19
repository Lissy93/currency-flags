package main

import (
	"encoding/json"
	"io"
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
	yamlFile, err := ioutil.ReadFile("../data/currencies.yaml")
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
	vars := mux.Vars(r)
	code := vars["code"]

	for _, c := range currencies {
		if c.Code == code {
			resp, err := http.Get(c.Flag)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()

			// Set the content type of the response to image/png
			w.Header().Set("Content-Type", "image/png")

			// Write the image data to the response writer
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			return
		}
	}

	http.Error(w, "Currency not found", http.StatusNotFound)
}
