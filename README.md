# Simple Web Application Using Golang
Build web application integrated with postgres

[![Build Status](https://travis-ci.com/RifanNurmuhammad/web-server-go.svg?token=WwCECKiVYLEzNkExr6jz&branch=master)](https://travis-ci.com/RifanNurmuhammad/web-server-go)
[![codecov](https://codecov.io/gh/RifanNurmuhammad/web-server-go/branch/master/graph/badge.svg?token=vc4t5FAS3N)](https://codecov.io/gh/RifanNurmuhammad/web-server-go)

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
