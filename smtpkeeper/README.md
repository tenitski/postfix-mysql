# SMTP Keeper

Maintains a list of SASL users and *from* addresses they are allowed to use.

## Usage

Build:
```bash 
go build
```

Start the server:

```bash
LOG_LEVEL=debug ./smtpkeeper DB_USER:DB_PASSWORD@(DB_HOST:DB_PORT)/DB_NAME 
```

Make requests:

```bash
# Create a new user
curl --request POST --header "Content-Type: application/json" --data '{"login":"someone@example.com","password":"somepass"}' http://127.0.0.1:8080/users

# Update existing user
curl --request PUT --header "Content-Type: application/json" --data '{"password":"newpass"}' http://127.0.0.1:8080/user/someone%40example.com

# Get a single user
curl http://127.0.0.1:8080/user/someone%40example.com

# Get all users
curl http://127.0.0.1:8080/users

# Add sender address
curl --request POST --header "Content-Type: application/json" --data '"@example.com"' http://127.0.0.1:8080/user/someone%40example.com/senders
curl --request POST --header "Content-Type: application/json" --data '"@mail.example.com"' http://127.0.0.1:8080/user/someone%40example.com/senders
curl --request POST --header "Content-Type: application/json" --data '"somebody@example.com"' http://127.0.0.1:8080/user/someone%40example.com/senders

# Get all sender addresses for user
curl http://127.0.0.1:8080/user/someone%40example.com/senders

# Delete sender address
curl --request DELETE -v http://127.0.0.1:8080/user/someone%40example.com/sender/%40example.com
curl --request DELETE -v http://127.0.0.1:8080/user/someone%40example.com/sender/%40mail.example.com
curl --request DELETE -v http://127.0.0.1:8080/user/someone%40example.com/sender/somebody%40example.com

# Delete user
curl --request DELETE -v http://127.0.0.1:8080/user/someone%40example.com

```

## Testing

```bash
go test ./...
```

## Todo

- Add data validation
- Add tests
