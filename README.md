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
./obfuscation-cell -s=10.1.1.9:1600 -c=10.1.1.10:1700 -n=mary -x=17 -f=10
```

###### example Docker container
```
docker build .

docker run -p 9000:9000 -e 'SERVER_ADDRESS=:9000' -e 'CLIENT_ADDRESS=10.1.0.100:9001' -e 'CELL_NAME=cella' -e 'SEED_VALUE=1898' -e 'FREQUENCY=12' -d --restart always 492d3fa0dcd4
docker run -p 9001:9001 -e 'SERVER_ADDRESS=:9001' -e 'CLIENT_ADDRESS=10.1.0.100:9000' -e 'CELL_NAME=cellb' -e 'SEED_VALUE=198' -e 'FREQUENCY=12' -d --restart always 492d3fa0dcd4
```

###### TODO
* Document better
* Refactor