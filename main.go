package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/urfave/cli"
)

const defaultCurrency = "USD"
const defaultAmount = "1"
const apiBaseURL = "https://www.nrb.org.np/exportForexJSON.php"

type APIResponse struct {
	Conversion Conversion `json:"Conversion"`
}

type Conversion struct {
	Currency []Currency `json:"Currency"`
}

type Currency struct {
	Date           string `json:"Date"`
	BaseCurrency   string `json:"BaseCurrency"`
	TargetCurrency string `json:"TargetCurrency"`
	BaseValue      string `json:"BaseValue"`
	TargetBuy      string `json:"TargetBuy"`
	TargetSell     string `json:"TargetSell"`
}

func main() {
	app := cli.NewApp()
	app.Name = "forex"
	app.Usage = "Check the value of NPR against foreign currencies."
	app.Commands = []cli.Command{
		{
			Name:    "convert",
			Aliases: []string{"c"},
			Usage:   "Convert foreign currency to NPR and vice versa(WIP).",
			Action:  convert,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "Display the foreign exchange rates of all the available currencies.",
			Action:  list,
		},
	}

	app.Run(os.Args)
}

func fetchExchangeRate() []Currency {
	response, err := http.Get(apiBaseURL)

	if err != nil {
		fmt.Printf("Could not fetch exchange rates. Please try again later.")
		os.Exit(1)
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Could not read data. Please try again later.")
		os.Exit(1)
	}

	var res APIResponse

	err = json.Unmarshal([]byte(string(data)), &res)

	if err != nil {
		fmt.Printf("Could not read data. Please try again later.")
		os.Exit(1)
	}

	return res.Conversion.Currency
}

func list(c *cli.Context) {
	rates := fetchExchangeRate()
	var ratesLength = len(rates)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', 0)

	fmt.Fprintln(w, "Currency \tBuying Rate\tSelling rate")

	for i := 0; i < ratesLength; i++ {
		fmt.Fprintf(w, "%v %v\t%v\t%v\t\n", rates[i].BaseValue, rates[i].BaseCurrency, rates[i].TargetBuy, rates[i].TargetSell)
	}

	w.Flush()
}

func convert(c *cli.Context) {
	amount := c.Args().Get(0)
	currency := c.Args().Get(1)

	validateArgs(currency, amount)

	amt, err := strconv.Atoi(amount)

	if err != nil {
		fmt.Printf("Please input a valid number for amount.\n")
		os.Exit(1)
	}

	currency = strings.ToUpper(currency)

	rates := fetchExchangeRate()

	selectedCurrency := getSelectedCurrency(rates, currency)

	if (selectedCurrency == Currency{}) {
		fmt.Println("Could not find the currency you are looking for.\nRun 'forex l' to view all the available currencies.")
		os.Exit(1)
	}

	buyingValue, _ := strconv.ParseFloat(selectedCurrency.TargetBuy, 64)
	a := float64(amt)

	fmt.Printf("%.2f %s -> %.2f NPR\n", a, currency, a*buyingValue)
}

func validateArgs(currency string, amount string) {
	if currency == "" && amount == "" {
		fmt.Printf("Using default values '%s' and '%s' for currency and amount.\n", defaultCurrency, defaultAmount)

		amount = defaultAmount
		currency = defaultCurrency
	} else if currency == "" {
		fmt.Printf("Using default value '%s' for currency.\n", defaultCurrency)

		currency = defaultCurrency
	} else if amount == "" {
		fmt.Printf("Using default value '%s' for amount.\n", defaultAmount)

		amount = defaultAmount
	}
}

func getSelectedCurrency(rates []Currency, currency string) Currency {
	var selectedCurrency Currency

	for i := 0; i < len(rates); i++ {
		if rates[i].BaseCurrency == currency {
			selectedCurrency = rates[i]
		}
	}

	return selectedCurrency
}
