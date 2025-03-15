# CryptoBro

![image](https://github.com/user-attachments/assets/809796e9-af82-4065-8575-50e4c5213483)


A lightweight CPU stress tool that simulates cryptocurrency mining activity. 

Useful as a proof of concept tool red teams can use to show impact, without actually mining cryptocurrency.

## Usage

```bash
# Run with default settings (5 minutes)
./cryptobro

# Run for a specific duration
./cryptobro -t 10    # Run for 10 minutes

```

## Building

```bash

git clone https://github.com/grahamhelton/cryptobro

go build -o cryptobro
```
