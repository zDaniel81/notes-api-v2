# Notes API v2

This is a RESTful API written in Golang that allows you to manage notes. It provides endpoints for creating, updating, deleting, retrieving all notes, and retrieving notes by ID.

The API utilizes the following technologies:

- **Golang**: The programming language used to implement the API.
- **PostgreSQL**: The chosen DBMS (Database Management System) for storing notes.
- **httprouter**: A high-performance HTTP router library used for routing requests.

## Installation

To run the Notes API locally, make sure you have Golang and PostgreSQL installed on your system. Then, follow these steps:

1. **Clone this repository:**

   ```bash
   git clone https://github.com/zDaniel8103/notes-api-v2.git
   cd notes-api
   ```
2. Set up the database:
- Create a PostgreSQL database named notes then run the migrations.
```bash
migrate -path <path_to_migrations>/migrations -database "postgresql://user:passw@host:port/notes?sslmode=disable" -verbose up
```

3. **Install the dependencies:**
```
go mod download
```
4. **Build and run the API:**
```bash
go run main.go
```

**The API will be accessible at localhost:8080/**

## API Endpoints

`POST` `/notes/create/{title}/{content}`: **Create a new note.** <br>
`PUT` `/notes/{id}/title/{title}`: **Update an existing note's title.**<br>
`PUT` `/notes/{id}/title/{title}`: **Update an existing note's content.**<br>
`DELETE` `/notes/{id}`: **Delete a note.**<br>
`GET` `/notes/{id}`: **Retrieve a note by ID.**<br>
`GET` `/notes`: **Retrieve all notes.**<br>

*For detailed information on each endpoint, including request and response examples, please refer to the API documentation.*

## License

This project is licensed under the MIT License.