# test-fiber

Testing out Go Lang Fiber framework

## Setup

- create file `books.db`
- run `go run main.go`

### Routes

| Method | Route             | Parameters                                 | Function                   |
| ------ | ----------------- | ------------------------------------------ | -------------------------- |
| GET    | /api/v1/book      | N/A                                        | book/books.go = GetBooks   |
| GET    | /api/v1/book/{id} | N/A                                        | book/books.go = GetBook    |
| POST   | /api/v1/book      | title: string, author: string, rating: int | book/books.go = NewBook    |
| DELETE | /api/v1/book/{id} | N/A                                        | book/books.go = DeleteBook |

### Resources

- https://tutorialedge.net/golang/basic-rest-api-go-fiber/
- https://github.com/gofiber/fiber
