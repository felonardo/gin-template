// infrastructure/db_hello_repository.go
package infrastructure

import "gin-template/domain"

type DbHelloRepository struct{}

func NewDbHelloRepository() *DbHelloRepository {
	return &DbHelloRepository{}
}

func (r *DbHelloRepository) FetchMessage() domain.Hello {
	// In a real-world application, this could be a DB query.
	return domain.Hello{Message: "Hello, Clean Architecture!"}
}
