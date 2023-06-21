package repository

import "context"

type FileRepository interface {
	Exist(c context.Context, path string) (bool, error)
	Update(c context.Context) error
	Create(c context.Context) error
	Delete(c context.Context, path string) error
	Get(c context.Context, path string) error
}
