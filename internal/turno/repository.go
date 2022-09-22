package turno

import (
	"context"

	"github.com/santiagokessy/finalBackIII/internal/domain"
)

type Repository interface {
	Get(cntxt context.Context, id int) (domain.Turno, error)
}
