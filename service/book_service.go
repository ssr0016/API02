package service

import (
	"context"
	"practice01/main.go/data/request"
	"practice01/main.go/data/response"
)

type BookService interface {
	Create(ctx context.Context, request request.BookCreateRequest)
	Update(ctx context.Context, request request.BookUpdateRequest)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) response.BookResponse
	FindAll(ctx context.Context) []response.BookResponse
}
