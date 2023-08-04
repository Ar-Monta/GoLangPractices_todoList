A simple CRUD backend for a Todo App written in GoLang using Domain-Driven Design (DDD) principles.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project provides a backend implementation for a Todo App using GoLang. The codebase follows Domain-Driven Design (DDD) principles, making it easy to understand, maintain, and extend.

## Features

- Create, read, update, and delete todo items.
- Get a list of all todos or fetch a specific todo by its ID.
- Mark a todo item as completed or mark it as incomplete.

## Installation

1. Clone the repository to your local machine:

```bash
git clone https://github.com/Ar-Monta/go-todo-app.git
cd go-todo-app
```
## Initialize Go modules and download dependencies:
```bash
go mod init
go mod download
```
## Usage
### Start the server:
```bash
go run cmd/server/main.go
```
The server will be running at http://localhost:8080.

## API Endpoints

The following API endpoints are available for interacting with the Todo App:

### Get All Todos

- **Endpoint:** `/todos`
- **Method:** GET
- **Description:** Get a list of all todo items.
- **Response:**
  - Status Code: 200 (OK)
  - Body: JSON array containing todo items.

### Get Todo by ID

- **Endpoint:** `/todos/{id}`
- **Method:** GET
- **Description:** Get a specific todo item by its ID.
- **Parameters:**
  - `{id}`: The unique ID of the todo item.
- **Response:**
  - Status Code: 200 (OK) - if the todo item is found.
  - Status Code: 404 (Not Found) - if the todo item is not found.
  - Body: JSON object representing the todo item.

### Create Todo

- **Endpoint:** `/todos`
- **Method:** POST
- **Description:** Create a new todo item.
- **Request Body:** JSON object with the following properties:
  - `title` (string, required): The title of the todo item.
  - `description` (string, required): The description of the todo item.
- **Response:**
  - Status Code: 201 (Created) - if the todo item is created successfully.
  - Body: JSON object representing the created todo item.

### Update Todo Completed Status

- **Endpoint:** `/todos/{id}/completed`
- **Method:** PUT
- **Description:** Update the completed status of a todo item.
- **Parameters:**
  - `{id}`: The unique ID of the todo item to update.
- **Request Body:** JSON object with the following property:
  - `completed` (boolean, required): The new completed status (true for completed, false for incomplete).
- **Response:**
  - Status Code: 200 (OK) - if the todo item is updated successfully.
  - Status Code: 404 (Not Found) - if the todo item is not found.

### Update Todo Details

- **Endpoint:** `/todos/{id}`
- **Method:** PUT
- **Description:** Update the details of a todo item.
- **Parameters:**
  - `{id}`: The unique ID of the todo item to update.
- **Request Body:** JSON object with the following optional properties:
  - `title` (string): The updated title of the todo item.
  - `description` (string): The updated description of the todo item.
  - `completed` (boolean): The updated completed status (true for completed, false for incomplete).
- **Response:**
  - Status Code: 200 (OK) - if the todo item is updated successfully.
  - Status Code: 404 (Not Found) - if the todo item is not found.

### Delete Todo

- **Endpoint:** `/todos/{id}`
- **Method:** DELETE
- **Description:** Delete a todo item.
- **Parameters:**
  - `{id}`: The unique ID of the todo item to delete.
- **Response:**
  - Status Code: 200 (OK) - if the todo item is deleted successfully.
  - Status Code: 404 (Not Found) - if the todo item is not found.

## Contributing
Contributions are welcome! If you find any bugs or want to add new features, please feel free to open an issue or submit a pull request.

## Generate mocks using mockery
```bash
./bin/mockery --name InterfaceName
```

For example

```bash
./bin/mockery --name TodoRepository
```
