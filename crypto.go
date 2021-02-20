package main

import (
	"encoding/json"
	"fmt"
	"github.com/finlaysawyer/go-crypto-cli/models"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	app := &cli.App{
		Usage: "A CLI utility to get the price of cryptocurrencies",
		Commands: []*cli.Command{
			{
				Name:      "get",
				Usage:     "Returns the price of a cryptocurrency",
				ArgsUsage: "<cryptocurrency> [currency] [price type]",
				Action: func(context *cli.Context) error {
					crypto := context.Args().Get(0)
					if crypto == "" {
						fmt.Println("Please specify a cryptocurrency, e.g. crypto get BTC")
						return nil
					}

					currency := context.Args().Get(1)
					if currency == "" {
						currency = "USD"
					}

					priceType := context.Args().Get(2)
					if priceType == "" {
						priceType = "spot"
					}

					price := getCryptoPrice(crypto, currency, priceType)
					if price.Data.Amount == "" {
						fmt.Println("You specified an invalid argument. Usage: <cryptocurrency> [currency] [price type]")
						return nil
					}

					fmt.Printf("%v is worth %v %v", strings.ToUpper(crypto), price.Data.Amount, strings.ToUpper(currency))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getCryptoPrice(crypto string, currency string, priceType string) models.Cryptocurrency {
	resp, err := http.Get(fmt.Sprintf("https://api.coinbase.com/v2/prices/%v-%v/%v", crypto, currency, priceType))
	if err != nil {
		log.Fatalln(err)
	}

	var mappedCrypto models.Cryptocurrency

	err = json.NewDecoder(resp.Body).Decode(&mappedCrypto)
	if err != nil {
		log.Fatalln(err)
	}

	return mappedCrypto
}
