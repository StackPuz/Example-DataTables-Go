# Example-DataTables-Go
- The example shows how to create an API for the DataTables with Go and using MySQL as a database.
- The article of this repo https://blog.stackpuz.com/create-an-api-for-datatables-with-go/
- To find more resources, please visit https://stackpuz.com

## Prerequisites
- Go 1.21
- MySQL

## Installation
- Clone this repository `git clone https://github.com/stackpuz/Example-DataTables-Go`
- Install the dependencies `go mod tidy`
- Create a new database and run [/database.sql](/database.sql) script to create tables and import data.
- Edit the database configuration in [/.env](/.env) file.

## Run project

- Run project `go run main.go`
- Navigate to http://localhost:8080