// repositories/hello_repository.go
package repositories

import "gin-template/domain"

type HelloRepository interface {
	FetchMessage() domain.Hello
}
