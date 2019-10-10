# forex

A cli tool to check the foreign exchange rates against Nepal's currency using Nepal Rastra Bank's [API](https://www.nrb.org.np/exportForexJSON.php).

## Installation

You can install `forex` by running:

```sh
$ go get github.com/aviskarkc10/forex
```

Make sure your PATH includes the $GOPATH/bin directory so your commands can be easily used:

```
export PATH=$PATH:$GOPATH/bin
```

## Usage

To view the available commands, run:

```sh
$ forex
```

To view the foreign exchange rates of all the available countries, run:

```sh
$ forex list or forex l
```

To convert a foreign currency to NPR, run:

```sh
$ forex convert <amount> <currency>
```

For example:

```sh
# Convert 1 American dollar to NPR
$ forex convert 1 USD or forex c 1 USD
1.00 USD -> 113.34 NPR

# Convert 1 Indian Rupee to NPR
$ forex convert 1 INR or forex c 1 INR
1.00 INR -> 1.60 NPR
```

## LICENSE

[MIT](LICENSE)
