package dia

import "context"

type service struct {
	getter solutionGetter
}

func NewService(getter solutionGetter) *service {
	return &service{
		getter: getter,
	}
}

type (
	solutionGetter interface {
		GetAllSolutions(ctx context.Context) ([]solution, error)
	}
)

func (s *service) GetAllSolutions(ctx context.Context) ([]solution, error) {
	return s.getter.GetAllSolutions(ctx)
}
