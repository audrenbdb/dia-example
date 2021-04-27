package dia

import (
	"context"
)

type solution struct {
	title       string
	description string
}

type inMemoryRepo struct {
	solutions []solution
}

func NewInMemoryRepo() *inMemoryRepo {
	return &inMemoryRepo{
		solutions: []solution{
			{
				title:       "Cyber range hybride",
				description: "Environnement virtuel pour simuler des entraînements aux combats cyber",
			},
			{
				title:       "Réponse à incident",
				description: "Intervention suite à rançongiciel ou une fuite de données",
			},
			{
				title:       "Capture the flag",
				description: "Operations de Capture the Flag pour une large palette de clients",
			},
		},
	}
}

func (r *inMemoryRepo) GetAllSolutions(ctx context.Context) ([]solution, error) {
	return r.solutions, nil
}
