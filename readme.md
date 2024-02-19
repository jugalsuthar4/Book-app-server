# Book App Server

This is a basic GoLang web server for a book application. It provides RESTful API endpoints for managing books, including functionalities such as creating, retrieving, updating, and deleting books.

## Getting Started

1. Clone the repository:

```bash
git clone git@github.com:jugalsuthar4/Book-app-server.git
cd Book-app-server
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the server:

```bash
go run main.go
```

The server will be running on `http://localhost:5001`.

## API Endpoints

- **Create Book:**

  `POST /books`

- **Get All Books:**

  `GET /books`

- **Get Book by ID:**

  `GET /books/{id}`

- **Update Book:**

  `PUT /books/{id}`

- **Delete Book:**

  `DELETE /books/{id}`

## CORS Configuration

Cross-Origin Resource Sharing (CORS) is configured to allow the following methods:

- GET
- POST
- PUT
- DELETE

## Dependencies

- [gorilla/mux](https://github.com/gorilla/mux) - HTTP router and dispatcher for Go.
- [rs/cors](https://github.com/rs/cors) - Go package for handling CORS (Cross-Origin Resource Sharing).

## Server Information

- **Port:** 5001

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use and modify the code according to your needs.
