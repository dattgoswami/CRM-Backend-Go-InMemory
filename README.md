# CRM Backend Project

## Getting Started

This repo contains a basic GO app to get started with constructing an API using GO. To get started, clone this repo and run `go run main.go` in your terminal at the project root.

# API Requirements

The company stakeholders want to create a CRM backend. You have been tasked with building the API that will support this application, and your coworker is building the frontend.

These are the notes from a meeting with the frontend developer that describe what endpoints the API needs to supply, as well as data shapes the frontend and backend have agreed meet the requirements of the application.

## API Endpoints

#### Customers

- A SHOW route: [GET] '/customers/{id}'
- An INDEX route: [GET] '/customers'
- A CREATE route: [POST] '/customers'
- An UPDATE route: [PATCH] '/customers/{id}'
- A DELETE route: [DELETE] '/customers/{id}'

## Data Shapes

#### Customers

- Id
- Name
- Role
- Email
- Phone
- Contacted

## Required Technologies

Your application must make use of the following libraries:

- gorrila mux

## Once the project is up and running we can test it using postman(/curl)

1. Send a GET request to url [http://0.0.0.0:3000/customers/]
2. Send a GET request to url [http://0.0.0.0:3000/customers/1222]
3. Send a POST request to url [http://0.0.0.0:3000/customers] with the body containing the following raw json

```
   {
   "id": 1456,
   "name": "Jack",
   "role": "Product Manager",
   "email": "jack@example.com",
   "phone": "67898989",
   "contacted": false
   }
```

4. Send a PATCH request to url [http://localhost:3000/customers/1456] with the body containing the following raw json

```
   {
   "name": "Jackie",
   "role": "Product Manager",
   "email": "jackie@example.com",
   "phone": "67898989",
   "contacted": true
   }
```

5. Send a DELETE request to url [http://localhost:3000/customers/1456]

## References:

1. https://drstearns.github.io/tutorials/gojson/
