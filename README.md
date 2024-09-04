# Student Management API

This repository contains a simple Student Management API built with Go and the Gin web framework. The API allows you to manage student records, including creating, reading, updating, and deleting student information.

## Features

- Create a new student record
- Retrieve a list of students
- Update an existing student record
- Delete a student record

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://golang.org/dl/) (version 1.16 or later)
- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/themobileprof/students-yhj57.git
    cd students-yhj57
    ```

2. Install the dependencies:

    ```sh
    go mod tidy
    ```

## Running the Application Locally

To run the application locally, use the following command:

```sh
go run main.go
```


## Building and Running the Docker Container

To build the Docker image, use the following command:

```sh
docker build --build-arg MY_ENV_VARIABLE=custom_value -t student-management-api .
```

To run the Docker container, use the following command:

```sh
docker run -p 8080:8080 student-management-api
```

The API will be available at http://localhost:8080.

## Running Tests
To run the tests, use the following command:

```sh
go test ./...
```

This will execute all the tests in the repository and provide a summary of the results.

## Contributing
If you would like to contribute to this project, please fork the repository and submit a pull request. We welcome all contributions!

## License
This project is licensed under the MIT License. See the LICENSE file for details.

