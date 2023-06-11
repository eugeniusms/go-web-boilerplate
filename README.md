# Go Web Boilerplate
This boilerplate is for applications using go with Postgres.<br>
Lib used:<br>
- github.com/go-playground/validator/v10 
- github.com/goccy/go-json 
- github.com/gofiber/fiber/v2 
- github.com/gofiber/swagger 
- github.com/golang-jwt/jwt 
- github.com/golang-module/carbon 
- github.com/kelseyhightower/envconfig 
- github.com/pkg/errors 
- github.com/sendinblue/APIv3-go-library/v2 
- github.com/sirupsen/logrus 
- github.com/swaggo/swag 
- github.com/twharmon/gouid 
- go.uber.org/dig
- golang.org/x/crypto 
- gorm.io/driver/postgres 
- gorm.io/gorm 
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

## Wiki
- SQL Drivers : https://github.com/golang/go/wiki/SQLDrivers

## Contribute
Feel free to contribute, this project will be greater for anyone.

## Authors
[@Nofaldi Atmam](https://github.com/nofamex), [@Haikal Susanto](https://github.com/haikalSusanto), [@Eugenius Mario](https://github.com/eugeniusms)