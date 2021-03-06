# Golang Based Simple REST-API
only for PostgresSQL Database

## Connection Setup
to setup change the dsn variable inside conn/conn.go

## Getting Started
below are some syntax to use this thing

##### Getting Api Key
navigate to http://localhost:8080/account/login and send a POST REQUEST with the format below
```
{
    "username": "myname,
    "password": "myname123"
}
```
Then make sure to copy the key and add a header request with a key of `Authorization` and the key as the value


##### Account CRUD
-- http://localhost:8080/account/login      : [POST] getting the API KEY as explained above
-- http://localhost:8080/account/get        : [GET] get all accounts data from the database
-- http://localhost:8080/account/get/:id    : [GET] get accounts data for a certain ID
-- http://localhost:8080/account/add/       : [POST] add new data to the account table
-- http://localhost:8080/account/drop/:id   : [GET] delete a row from the account table for a certain ID
-- http://localhost:8080/account/edit/:type : [POST] Edit a row from the account table

##### Books CRUD
-- http://localhost:8080/book/get        : [GET] get all book data from the database
-- http://localhost:8080/book/get/:id    : [GET] get book data for a certain ID
-- http://localhost:8080/book/add/       : [POST] add new data to the book table
-- http://localhost:8080/book/drop/:id   : [GET] delete a row from the book table for a certain ID
-- http://localhost:8080/book/edit/:type : [POST] Edit a row from the book table

##### Comments CRUD
-- http://localhost:8080/comment/get              : [GET] get all comment data from the database
-- http://localhost:8080/comment/get/:id          : [GET] get comment data for a certain ID
-- http://localhost:8080/comment/getByPost/:id    : [GET] get comment data for a certain user_post ID
-- http://localhost:8080/comment/add/             : [POST] add new data to the user_comments table
-- http://localhost:8080/comment/drop/:id         : [GET] delete a row from the user_comments table for a certain ID
-- http://localhost:8080/comment/edit/:type       : [POST] Edit a row from the user_comments table

##### Post CRUD
-- http://localhost:8080/post/get        : [GET] get all post data from the database
-- http://localhost:8080/post/get/:id    : [GET] get post data for a certain ID
-- http://localhost:8080/post/add/       : [POST] add new data to the user_posts table
-- http://localhost:8080/post/drop/:id   : [GET] delete a row from the user_posts table for a certain ID
-- http://localhost:8080/post/edit/:type : [POST] Edit a row from the user_posts table
