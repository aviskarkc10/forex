# forex

A cli tool to check the foreign exchange rates against Nepal's currency using Nepal Rastra Bank's [api](https://www.nrb.org.np/exportForexJSON.php).

## Installation

You can install `forex` by running:

```
$ go get github.com/aviskarkc10/forex
```

## Usage

To view the available commands run:

```
$ forex
```

To view the foreign exchange rates of all the available countries run

```
$ forex list # or forex l
```

To convert a foreign currency to NPR run:

```
$ forex convert <amount> <currency>
```


For example:

```
# Convert 1 American dollar to NPR
$ forex convert 1 USD # or forex c 1 USD
```

# LICENSE

[MIT](LICENSE)
