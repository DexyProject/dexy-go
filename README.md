# DEXY Go

[![Build Status](https://travis-ci.com/DexyProject/dexy-go.svg?token=SGE7GHsjEHmsR4VosLJx&branch=development)](https://travis-ci.com/DexyProject/dexy-go)

Open source orderbook code for the DEXY exchange. The code contains basic orderbook functionality as well as various chain watching code which ensures the orderbook stays up to date. Additionally this repository also contains APIs for the various endpoints required for the exchange, this includes price tickers as well as trading history.

The API documentation can be found under the [docs](docs) directory.

## Documentation

The dexy go server is made up of several components.

### Setting up the database

```bash
mongo < /configs/mongo_indexes.js

````

### API

This command is used to run the rest API. This includes all endpoints documented [here](docs).

```
Usage of main.go
  -ethnode string
    	ethereum node address
  -mongo string
    	mongodb connection string
  -vault string
    	vault address
```

### Trade Watcher

This command is responsible for monitoring and indexing trades, this will update the orderbook as well as insert transactions into the history.

```
Usage of main.go
  -addr string
    	exchange address
  -ethnode string
    	ethereum node address
  -mongo string
    	mongodb connection string
```

### Cancelled Watcher

This command is responsible for monitoring which orders have been cancelled, these will then be removed from the orderbook.

```
Usage of main.go
  -addr string
    	exchange address
  -ethnode string
    	ethereum node address
  -mongo string
    	mongodb connection string
```

### Ticker worker

This command monitors the blocks, and generates new ticks for every block.

```
Usage of main.go
  -ethnode string
    	ethereum node address
  -mongo string
    	mongodb connection string
```

### Markets worker

This command monitors the blocks, and generates new market statistics for every block.

```
Usage of main.go
  -ethnode string
    	ethereum node address
  -mongo string
    	mongodb connection string
  -path string
    	path to tokens file
```

### Balance Monitor

## Authors

* **Dean Eigenmann** - [decanus](https://github.com/decanus)
* **Niranjan Ravichandra** - [nravic](https://github.com/nravic)

See also the list of [contributors](https://github.com/DexyProject/dexy-go/contributors) who participated in this project.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/DexyProject/dexy-go/tags).

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details
