# Panchang

Panchang Term is a command-line interface (CLI) calendar application that allows you to view the month, store events, and retrieve events for the month.

## Features

- Display the current month calendar
- Add events to specific dates
- Retrieve events for the month

## Installation

To install Panchang Term, clone the repository and navigate to the project directory:

```sh
git clone https://github.com/yourusername/panchang_term.git
cd panchang_term
```

## Usage

### Display the Month

To display the current month's calendar, use the following command:

```sh
./panchang
```

### Add an Event

To add an event to current date, use the following command:
```sh
./panchang event add 
```

To add an event to a specific date, use the following command:
```sh
./panchang event add --date DATE
```

### Get Events for the Month

To retrieve all events for the current month, use the following command:

```sh
./panchang event get
```

To retrieve all events for the selected month, use the following command:

```sh
./panchang event get --month MONTH
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or suggestions, please open an issue on GitHub.
