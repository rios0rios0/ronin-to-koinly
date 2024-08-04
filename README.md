# Ronin Wallet to Koinly CSV Exporter

This project is a Go application that reads transaction data from the Ronin wallet API and exports it to a CSV file in the Koinly format.

## Features

- Fetches transaction data from the Ronin wallet API.
- Parses the transaction data.
- Exports the data to a CSV file compatible with Koinly.

## Requirements

- Go 1.16 or higher
- Internet connection to access the Ronin wallet API

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/ronin-to-koinly.git
    cd ronin-to-koinly
    ```

2. Install the required Go packages:

    ```sh
    go get -u github.com/go-resty/resty/v2
    go get -u github.com/gocarina/gocsv
    ```

## Usage

1. Update the API endpoint in the `main.go` file to point to your Ronin wallet API endpoint.

    ```go
    resp, err := client.R().Get("https://api.roninchain.com/wallet/transactions")
    ```

2. Build and run the application:

    ```sh
    go build -o ronin-to-koinly
    ./ronin-to-koinly
    ```

3. The application will create a `koinly_transactions.csv` file in the current directory with the transaction data.

## Example

Here is an example of the CSV output:

```csv
Date,Sent Amount,Sent Currency,Received Amount,Received Currency,Fee Amount,Fee Currency,Tag
2024-08-04 08:00:00,0.5,ETH,100,USD,0.01,ETH,Trade
2024-08-04 09:00:00,1.0,ETH,200,USD,0.02,ETH,Trade
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Resty](https://github.com/go-resty/resty) - Simple HTTP and REST client for Go inspired by Ruby rest-client.
- [GoCSV](https://github.com/gocarina/gocsv) - The Go library for CSV serialization and deserialization.

## Contact

For any questions or suggestions, feel free to open an issue or contact me at [rios0rios0@outlook.com].
