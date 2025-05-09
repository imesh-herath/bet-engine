# Bet Engine

Bet Engine is a simple and efficient betting system designed to manage and process bets. This project provides a foundation for building custom betting applications.

## Features

- Manage user bets.
- Calculate winnings and losses.
- Support for multiple bet types.
- Easy-to-extend architecture.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/imesh-herath/bet-engine
    ```
2. Navigate to the project directory:
    ```bash
    cd bet-engine
    ```
3. Build the project:
    ```bash
    go build
    ```

## Usage

1. Run the application:
    ```bash
    ./bet-engine
    ```
2. Access the application at `http://localhost:8080`.

## Sample API Requests

### Place a Bet
```bash
curl -X POST http://localhost:8080/bets -H 'Content-Type: application/json' \
-d '{"user_id":1,"event_id":101,"odds":2.0,"amount":100}'
```

### Get Balance
```bash
curl -X POST http://localhost:8080/settle -H 'Content-Type: application/json' \
-d '{"event_id":101,"result":0}'
```

### Settle Bet
```bash
curl -X POST http://localhost:8080/settle -H 'Content-Type: application/json' \
-d '{"event_id":101,"result":0}' 
```

### Create User
```bash
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" \
-d '{"user_id":1, "balance":2000}'
```

## Contact

For questions or feedback, please contact [imeshnipun2000@gmail.com].