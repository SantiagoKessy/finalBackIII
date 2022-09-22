package paciente

import (
	"context"
)

type Repository interface {
	Get(context context.Context, id int) (domain.Turno, error)
}
