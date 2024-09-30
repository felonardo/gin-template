// usecases/hello_usecase.go
package usecases

import (
	"gin-template/domain"
	"gin-template/repositories"
)

type HelloUseCase interface {
	GetHelloMessage() domain.Hello
}

type helloUseCase struct {
	helloRepo repositories.HelloRepository
}

func NewHelloUseCase(repo repositories.HelloRepository) HelloUseCase {
	return &helloUseCase{
		helloRepo: repo,
	}
}

func (u *helloUseCase) GetHelloMessage() domain.Hello {
	return u.helloRepo.FetchMessage()
}
