# go-crypto-cli

Go Crypto CLI is a small CLI utility that can fetch the spot, buy and sell prices of cryptocurrencies from the Coinbase API.
This is a small project I created to learn Golang.

## Installation
**Created using go 1.16**

`go install crypto.go`

## Usage
`crypto get <cryptocurrency> [currency] [price type]`


To get the current Bitcoin market price in USD:
- `crypto get btc`

To get the current Ethereum sell price in GBP:
- `crypto get eth gbp sell`

To get the current Litecoin buy price in GBP:
- `crypto get ltc gbp buy`