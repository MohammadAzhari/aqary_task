# Aqary Task

This project is a solution to the Aqary Task, which involves setting up and running a server.

## Getting Started

### Prerequisites

To run this project, you need to have the following installed:

- [Docker](https://www.docker.com/)
- [Go](https://golang.org/)

### Installation

1. Clone this repository to your local machine

2. Change into the project directory

3. Run the following commands to set up and initialize the database:

   ```bash
   make rundb
   make createdb
   make migrateup
   ```

## Usage

To start the server, run the following command:

```bash
go run main.go
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
