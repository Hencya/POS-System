# Golang Point Of Sale

Example Golang API backend rest about a simple case Point Of Sale using Echo Framework and Gorm ORM Database.

## User Dummy Login
  Username = username
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
|                 | */api/v1/auth/login*        | *POST*     |
| **Transaction** |                            |            |
|                 | */api/v1/transaction*      | *POST*     |
|                 | */api/v1/transaction*      | *GET*      |
|                 | */api/v1/transaction/:id*  | *DELETE*   |
|                 | */api/v1/transaction/:id*  | *PUT*      |
