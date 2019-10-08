package main

import (
  "fmt"
  "os"
  "strconv"

  "github.com/urfave/cli"
);

const defaultCurrency = "usd";
const defaultAmount = "1";


func main() {
  app := cli.NewApp()
  app.Name = "exchange"
  app.Usage = "Check the value of NPR against other currencies."
  app.Action = exchange;

  app.Run(os.Args);
}

func exchange(c *cli.Context) {
  amount := c.Args().Get(0);
  currency := c.Args().Get(1);

  if (currency == "" && amount == "") {
    fmt.Printf("Using default values '%s' and '%s' for currency and amount.\n", defaultCurrency, defaultAmount)

    amount = defaultAmount;
    currency = defaultCurrency;
  } else if (currency == "") {
    fmt.Printf("Using default value '%s' for currency.\n", defaultCurrency)

    currency = defaultCurrency;
  } else if (amount == "") {
    fmt.Printf("Using default value '%s' for amount.\n", defaultAmount)

    amount = defaultAmount;
  }

  amt, err := strconv.Atoi(amount);

  if(err != nil) {
    fmt.Printf("Please input a valid number for amount.\n");

    return;
  }

  fmt.Printf("You want to view value of '%d'.\n", amt);
}
