# Golang Point Of Sale

Example Golang API backend rest about a simple case Point Of Sale using Echo Framework and Gorm ORM and PostgreSQL Database.

## User Dummy Login
  Username = username
  <br />
  Password = password

## Command

- ### Runnig App

```sh
$ go run main.go
```

## Endpoint

| **Nama**        | **Route**                  | **Method** |
| --------------- | -------------------------- | ---------- |
| **User**        |                            |            |
|                 | */api/v1/user/login*        | *POST*     |
| **Transaction** |                            |            |
|                 | */api/v1/transaction*      | *POST*     |
|                 | */api/v1/transaction*      | *GET*      |
|                 | */api/v1/transaction/:id*  | *DELETE*   |
|                 | */api/v1/transaction/:id*  | *PUT*      |

## Api Contract
### User API
1. Login
- Endpoint: /api/v1/user/login
- Method: POST
- Description: Perform user authentication with username and password
- Request Header:
  - Content-Type: application/json
- Request Body:
``` json
{
  "ID"      : "number"
  "username": "string",
  "password": "string",
  "CreatedAt" : "time",
	"UpdatedAt" : "time"
}
```
- Response:
- HTTP Status Code: 200 OK
- Body:
``` json
{
  "status": "number",
  "message": "string",
  "data" : {
     "token": "string"
  }
}
```
- HTTP Status Code: 404 Not Found
- Body:
``` json
{
  "message": "string",
}
```
### Transaction API
1. Create Transaction
- Endpoint: /api/v1/transaction
- Method: POST
- Description: Description: Create a new transaction with amount, notes, date, and type
- Request Header:
  - Content-Type: application/json
  - Authorization: Bearer <access_token>
- Request Body:
``` json
{
    "amount": "number",
    "notes": "string",
    "date": "date",
    "type": "string"
}
```
- Response:
- HTTP Status Code: 201 Created
- Body:
``` json
{
  "status": "number",
  "message": "string",
  "data" : {
    "ID":"number",
    "created_by": "string",
    "updated_by": "string",
    "amount": "number",
    "notes": "string",
    "date": "date",
    "type": "string",
    "CreatedAt" : "time",
	  "UpdatedAt" : "time",
  }
}
```
- HTTP Status Code: 400 Bad Request
- Body:
``` json
{
  "message": "string",
}
```
- HTTP Status Code: 401 Unauthorized
- Body:
``` json
{
  "message": "string",
}
```
- HTTP Status Code: 500 Internal Server Error
- Body:
``` json
{
  "message": "string",
}
```
2. Get Transactions
- Endpoint: /api/v1/transaction
- Method: GET
- Description: Get a list of transactions with sorting, filtering, and pagination options
- Request Header:
  - Authorization: Bearer <access_token>
- Query Parameters:
  - sort: optional, string, possible values: amount_asc, amount_desc, date_asc,date_desc
  - type: optional, string, possible values: income, expense
  - filter: Optional, string, possible values: min_amount, max_amount
  - min_amount: optional, number
  - max_amount: optional, number
  - page: optional, number, default: 1
  - offset: optional, number, default: 0
  - limit: optional, number, default: 10
- Response:
- HTTP Status Code: 200 OK
- Body:
``` json
{
  "status": "number",
  "message": "string",
  "data" : [
    {
      "ID":"number",
      "created_by": "string",
      "updated_by": "string",
      "amount": "number",
      "notes": "string",
      "date": "date",
      "type": "string",
      "CreatedAt" : "time",
      "UpdatedAt" : "time",
    }
  ]
  "page": {
    "offset": "number",
    "limit": "number",
    "total_data": "number"
   }
}
```
- HTTP Status Code: 401 Unauthorized
- Body:
``` json
{
  "message": "string",
}
```
- HTTP Status Code: 404 Not Found
- Body:
``` json
{
  "message": "string",
}
```
3. Update Transaction
- Endpoint: /api/v1/transaction/:id
- Method: PUT
- Description: Update a transaction by id
- Request Header:
  - Content-Type: application/json
  - Authorization: Bearer <access_token>
- Request Body:
``` json
{
    "amount": "number",
    "notes": "string",
    "date": "date",
    "type": "string"
}
```
- Response:
- HTTP Status Code: 200 OK
- Body:
``` json
{
  "status": "number",
  "message": "string",
  "data" : [
    {
      "ID":"number",
      "created_by": "string",
      "updated_by": "string",
      "amount": "number",
      "notes": "string",
      "date": "date",
      "type": "string",
      "CreatedAt" : "time",
      "UpdatedAt" : "time",
    }
  ]
  "page": {
    "offset": "number",
    "limit": "number",
    "total_data": "number"
   }
}
```
- HTTP Status Code: 404 Not Found
- Body:
``` json
{
  "message": "string",
}
```
- HTTP Status Code: 401 Unauthorized
- Body:
``` json
{
  "message": "string",
}
```
- HTTP Status Code: 500 Internal Server Error
- Body:
``` json
{
  "message": "string",
}
```
4. Get Transactions
- Endpoint: /api/v1/transaction/:id
- Method: DELETE
- Description: Delete a transaction by id
- Request Header:
  - Authorization: Bearer <access_token>
- Response:
- HTTP Status Code: 200 OK
- Body:
``` json
{
  "status": "number",
  "message": "string",
  "data": "domain"
}
```
- HTTP Status Code: 401 Unauthorized
- Body:
``` json
{
  "message": "string",
}
```
- HTTP Status Code: 404 Not Found
- Body:
``` json
{
  "message": "string",
}
```
        
  
  
  

