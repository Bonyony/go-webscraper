# Go-webscraper

Ostensibly a webscraper CLI tool, I might just add some silly little features in as I go.

## How to Run

**Prerequisites**

Ensure you have the following installed on your system:

- Go (1.18 or later)
- Git (for cloning the repository)

### Clone the Repository

```
git clone https://github.com/Bonyony/go-webscraper.git
cd go-webscraper
```

### Install Dependencies

This project uses Cobra and Colly so ensure they are installed properly with:

```
go mod tidy
```

### Build the CLI

You can name it whatever you want, go-webscraper is just my example

```
go build -o go-webscraper
```

### Run the CLI

```
./go-webscraper [command] [flag]
```

To see all available commands:

```
./go-webscraper --help
```

### Development Mode

For quick testing without building into an executable:

```
go run main.go [command] [flag]
```

## Tech Stack

**Cobra:** for a nice, clean CLI setup

**Colly:** for webscraping with Golang

## Features

- Scrape music sites on the web!
- Simple password generator
