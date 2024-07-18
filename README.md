# Go Restful API with Gin, GORM - MVP Shop

Welcome to the **MVP Shop** project! This repository contains a fully functional RESTful API for an online shop, built with Go, Gin, and GORM. We appreciate any feedback and contributions. Feel free to open issues for comments and discussions.

## Table of Contents

- [Features](#features)
- [Schema Design](#schema-design)
- [Getting Started](#getting-started)
  - [Local Setup](#local-setup)
- [API Documentation](#api-documentation)

## Features

This project includes the following features to meet the MVP criteria:

- **Product Management**:
  - View product list by category
  - Add products to the shopping cart
  - View items in the shopping cart
  - Remove items from the shopping cart
  - Checkout and process payment transactions

- **User Authentication**:
  - Customer login and registration
  - JWT-based authorization

- **Database and Logging**:
  - Automatic schema migration using GORM
  - Custom error and info logging with Logrus

## Schema Design

For detailed schema design, please visit our [Wiki](./migration/wiki/).

## Getting Started

Follow these steps to set up and run the application:

### Local Setup

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/mvp-shop.git
   cd mvp-shop
   ```
2. **Database Setup**:
   - Create a PostgreSQL database.
   - Set environment variables for your database as shown in the .env.example file.
3. **Run the Application**:
   - Directly with Go:
   ```bash
   go run main.go
   ```
   - Using Docker Compose:
   ```bash
   docker-compose up -d --build --force-recreate
   ```
4. *(Optional)* Update Swagger Documentation:
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest && swag init
   ```
   
## Api Documentation
Access the comprehensive API documentation through the following links:
- [Postman Collection](https://documenter.getpostman.com/view/11257503/2sA3dviBtA)
- [Swagger UI](http://localhost:3001/swagger/index.html)
