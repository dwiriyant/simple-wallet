# Simple Wallet REST API

## Tech Stack

- Golang v1.23 : https://github.com/golang/go
- MySQL (Database) : https://github.com/mysql/mysql-server

## Installation

Clone the repository then switch to the repo folder

```
$ git clone https://github.com/dwiriyant/simple-wallet.git
$ cd simple-wallet
```

Copy the example env file and make the required configuration changes in the .env file

```
cp .env.example .env
```

Run migration

```
go run main.go migrate up
```

Run the golang development server

```
go run main.go
```

The REST API can be accessed at

```
http://localhost:8001
```
