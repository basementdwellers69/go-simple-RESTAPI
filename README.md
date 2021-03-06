# Golang Based Simple REST-API
only for PostgresSQL Database

## Connection Setup
to setup change the dsn variable inside conn/conn.go

## Getting Started
below are some syntax to use this thing

**Getting Api Key**
navigate to http://localhost:8080/account/login and send a POST REQUEST with the format below
```
{
    "username": "myname,
    "password": "myname123"
}
```
Then make sure to copy the key and add a header request with a key of `Authorization` and the key as the value

**http://localhost:8080/comment/**
