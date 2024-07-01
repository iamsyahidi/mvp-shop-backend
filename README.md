# Go Restful API with Gin, GORM - MVP Shop

Any feedback and pull requests are welcome and highly appreciated. Feel free to open issues just for comments and discussions.

## Features

The following feature is to fulfill the criteria in mvp constraint :

- Customer can view product list by product category
- Customer can add product to shopping cart
- Customer can see a list of products that have been added to the shopping cart
- Customer can delete product list in shopping cart
- Customer can checkout and make payment transactions
- Login and register customer

additional :
- Migrate with gorm for creating our schemas on app start
- Jwt Authorization Header
- Logrus for custom internal error/info

## Schema design

Please visit [wiki](./migration/wiki/) 

## Start Application

- Clone and change into this repository

### Local

- Create a postgres database and set environment variables for your database following the example env file
- Run the application with command : `go run main.go`
- or Run via docker compose : `docker-compose up -d --build --force-recreate`
- (optional) to update swagger docs : `go install github.com/swaggo/swag/cmd/swag@latest && swag init`

## API documentations/collections

Please visit [postman](https://documenter.getpostman.com/view/11257503/2sA3dviBtA)
or visit [swagger](http://localhost:3001/swagger/index.html)
