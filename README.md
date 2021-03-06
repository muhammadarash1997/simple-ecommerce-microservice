# Simple E-Commerce Microservice
> This is a backend service for Simple E-Commerce Microservice. Simple E-Commerce Microservice is used by customer to be able to go shopping by online.

## Requirements
You need [Docker](https://www.docker.com) and [Docker Compose](https://docs.docker.com)

## Running App

```bash
docker compose up
```

## Technologies Used
- Go 1.18
- PostgreSQL
- Docker
- Gin Web Framework
- GORM
- JSON Web Token
- Microservices

## Usecases
1. Customer were be able to register and login.
2. Customer were be able to see all product.
3. Customer were be able to search product by category.
4. Customer were be able to view product with information.
5. Customer were be able to view their payment.
6. Customer were be able to view cart.
7. Customer were be able to view order.
8. Customer were be able to update quantity item in cart.
9. Customer were be able to delete item from cart.
10. Customer were be able to order.
11. Customer were be able to pay the order.

## Code Structure
The design contains several layers and components and very much similar to onion or clean architecture attempt.

### Components
1. Controllers
2. Services
3. Repositories

#### Controllers
Controllers is where all the http handlers exist. This layer is responsible to hold all the http handlers and request validation.

#### Services
Services mediates communication between a controller and repository layer. The service layer contains business logic.

#### Repositories
Repositories is for accessing the database and helps to extend the CRUD operations on the database.

### Architecture Microservice
![alt text](https://github.com/muhammadarash1997/simple-ecommerce-microservice/blob/master/ARCHITECTURE.png?raw=true)

### Entity Relationship Diagram
![alt text](https://github.com/muhammadarash1997/simple-ecommerce-microservice/blob/master/ERD.png?raw=true)

### Flow Chart
![alt text](https://github.com/muhammadarash1997/simple-ecommerce-microservice/blob/master/FLOWCHART.png?raw=true)
