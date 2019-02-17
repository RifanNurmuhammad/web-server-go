# Simple Web Application Using Golang
Build web application integrated with postgres


# How to Start:
    - Install [dep](https://github.com/golang/dep)
    - Run this command `dep ensure`
    - Run this command `go build` 
    

# Database Application
    CREATE DATABASE db_book_store;

    CREATE TABLE books (
        id SERIAL PRIMARY KEY,
        title VARCHAR(256),
        description VARCHAR(1024)
    );
    
`created by rifannurmuhammad`
