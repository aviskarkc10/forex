package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/urfave/cli"
)

const defaultCurrency = "usd"
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
	app.Name = "exchange"
	app.Usage = "Check the value of NPR against foreign currencies."
	app.Commands = []cli.Command{
		{
			Name:    "convert",
			Aliases: []string{"c"},
			Usage:   "Convert foreign currency to NPR and vice versa.",
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
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Could not read data. Please try again later.")
	}

	var res APIResponse

	err = json.Unmarshal([]byte(string(data)), &res)

	if err != nil {
		fmt.Printf("Could not read data. Please try again later.")
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

	amt, err := strconv.Atoi(amount)

	if err != nil {
		fmt.Printf("Please input a valid number for amount.\n")

		return
	}

	fmt.Printf("This is a work in progress.")
}
