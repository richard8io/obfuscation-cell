# obfuscation-cell
Run your own network of cells that randomly transmit and receive encrypted chunks of data.

#### Usage

###### dependancies (installed in the docker image below)
* Golang ~1.4

###### example build and usage
```
go build
# run it without flags to see the needed settings
./cell
Usage:
  -s=10.1.1.9:1600: ip and port to run the server on
  -c=10.1.1.10:1700: ip and port of client to connect to
  -n=john: cell name
  -x=8978: value to seed cell with
  -f=10: how many seconds in between data transmission
```

###### example binary runs
```
./cell -s=10.1.1.9:1600 -c=10.1.1.10:1700 -n=mary -x=17 -f=10
```

###### TODO
* Document better
* Refactor