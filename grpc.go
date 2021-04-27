package dia

import (
	"context"

	"github.com/audrenbdb/dia/pb"
)

type diaServer struct {
	pb.UnimplementedSolutionsServer
	solutionGetter solutionGetter
}

func NewServer(getter solutionGetter) *diaServer {
	return &diaServer{
		solutionGetter: getter,
	}
}

func (s *diaServer) GetAllSolutions(ctx context.Context, req *pb.GetAllSolutionsRequest) (*pb.GetAllSolutionsResponse, error) {
	solutions, err := s.solutionGetter.GetAllSolutions(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetAllSolutionsResponse{
		Solutions: convertSolutionArrayToProtobuf(solutions, []*pb.Solution{}),
	}, nil
}

func convertSolutionArrayToProtobuf(solutions []solution, pbSolutions []*pb.Solution) []*pb.Solution {
	if len(solutions) == 0 {
		return pbSolutions
	}
	nextSolution := convertSolutionToProtobuf(solutions[0])
	return convertSolutionArrayToProtobuf(solutions[1:], append(pbSolutions, nextSolution))
}

func convertSolutionToProtobuf(solution solution) *pb.Solution {
	return &pb.Solution{
		Title:       solution.title,
		Description: solution.description,
	}
}
