# test-fiber

Testing out Go Lang Fiber framework

## Setup

- create file `books.db`
- run `npm i && npm run dev`
- run `go run main.go`

### Routes

| Method | Route             | Parameters                                 | Function                           | Description                          |
| ------ | ----------------- | ------------------------------------------ | ---------------------------------- | ------------------------------------ |
| GET    | /static           | N/A                                        | N/A                                | Used to access assets (CSS, JS, etc) |
| GET    | /                 | N/A                                        | modules/home/home.go = Index       |                                      |
| GET    | /api/v1/book      | N/A                                        | modules/book/books.go = GetBooks   |                                      |
| GET    | /api/v1/book/{id} | N/A                                        | modules/book/books.go = GetBook    |                                      |
| POST   | /api/v1/book      | title: string, author: string, rating: int | modules/book/books.go = NewBook    |                                      |
| DELETE | /api/v1/book/{id} | N/A                                        | modules/book/books.go = DeleteBook |                                      |

### Resources

- https://tutorialedge.net/golang/basic-rest-api-go-fiber/
- https://github.com/gofiber/fiber
