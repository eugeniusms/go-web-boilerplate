# Go Web Boilerplate
This boilerplate is for applications using go with Postgres as the database.

## Getting started
First, clone this project with the following command:
```
git clone https://github.com/eugeniusms/go-web-boilerplate.git YourAppName
```
Then, if you want to rename your module, you can edit this first line of your go.mod  
```
module go-web-boilerplate >>> module YourAppName
``` 
⚠ If you change it, some files will show new issues because they depend on this package. For example in ```application/auth/service.go```, you have to change ```go-web-boilerplate/application/auth``` to ```YourAppName/application/auth```
Then download all the dependencies using
```bash
go mod download
```
After this, you can run your app using 
```bash
go run main.go
```
or build it using 
```bash
go build
```
This will generate an executable file named `<YourAppName>`.

# References
## STACKS
- Test : Testify
- Database : Postgres
- Authentication : JWT & OAuth2 (Google)
## GO WIKI
- SQL Drivers : https://github.com/golang/go/wiki/SQLDrivers

## DEPENDENCIES
- HTTP Router : github.com/julienschmidt/httprouter
- PostgreSQL Driver : github.com/lib/pq
- Testing : github.com/stretchr/testify
- Validator : github.com/go-playground/validator

## Contribute
Feel free to contribute, this project will be greater for anyone.

## Authors
[Nofaldi Atmam](https://github.com/nofamex) | [Haikal Susanto](https://github.com/haikalSusanto) | [Eugenius Mario](https://github.com/eugeniusms)