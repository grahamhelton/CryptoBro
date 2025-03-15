# CryptoBro

A lightweight CPU stress tool that simulates cryptocurrency mining activity. 

Useful as a proof of concept tool red teams can use to show impact, without actually mining cryptocurrency.

## Usage

```bash
# Run with default settings (5 minutes)
./cryptobro

# Run for a specific duration
./cryptobro -t 10    # Run for 10 minutes

# Show help
./cryptobro -h
```

## Building

```bash

git clone https://github.com/grahamhelton/cryptobro

go build -o cryptobro
```
