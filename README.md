<h1 align="center">Ronin to Koinly</h1>
<p align="center">
    <a href="https://github.com/rios0rios0/ronin-to-koinly/releases/latest">
        <img src="https://img.shields.io/github/release/rios0rios0/ronin-to-koinly.svg?style=for-the-badge&logo=github" alt="Latest Release"/></a>
    <a href="https://github.com/rios0rios0/ronin-to-koinly/blob/main/LICENSE">
        <img src="https://img.shields.io/github/license/rios0rios0/ronin-to-koinly.svg?style=for-the-badge&logo=github" alt="License"/></a>
    <a href="https://github.com/rios0rios0/ronin-to-koinly/actions/workflows/default.yaml">
        <img src="https://img.shields.io/github/actions/workflow/status/rios0rios0/ronin-to-koinly/default.yaml?branch=main&style=for-the-badge&logo=github" alt="Build Status"/></a>
</p>

A Go application that reads transaction data from the Ronin wallet API and exports it to a CSV file in the Koinly format.

## Features

- Fetches transaction data from the Ronin wallet API
- Parses the transaction data
- Exports the data to a CSV file compatible with Koinly

## Requirements

- Go 1.16 or higher
- Internet connection to access the Ronin wallet API

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/rios0rios0/ronin-to-koinly.git
    cd ronin-to-koinly
    ```

2. Install the required Go packages:

    ```bash
    go get -u github.com/go-resty/resty/v2
    go get -u github.com/gocarina/gocsv
    ```

## Usage

1. Update the API endpoint in the `main.go` file to point to your Ronin wallet API endpoint.

    ```go
    resp, err := client.R().Get("https://api.roninchain.com/wallet/transactions")
    ```

2. Build and run the application:

    ```bash
    go build -o ronin-to-koinly
    ./ronin-to-koinly
    ```

3. The application will create a `koinly_transactions.csv` file in the current directory with the transaction data.

## Example Output

Here is an example of the CSV output:

```csv
Date,Sent Amount,Sent Currency,Received Amount,Received Currency,Fee Amount,Fee Currency,Tag
2024-08-04 08:00:00,0.5,ETH,100,USD,0.01,ETH,Trade
2024-08-04 09:00:00,1.0,ETH,200,USD,0.02,ETH,Trade
```

## Contributing

Contributions are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
