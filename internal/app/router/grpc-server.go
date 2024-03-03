package router

import (
	"context"
	"github.com/AutoDivdik/enterprises-service/gen"
	"github.com/AutoDivdik/enterprises-service/internal/usecases"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var _ gen.EnterpriseServiceServer = (*EnterprisesGRPCServer)(nil)

type EnterprisesGRPCServer struct {
	gen.UnimplementedEnterpriseServiceServer
	uc usecases.UseCase
}

func NewGRPCServer(grpcServer *grpc.Server, uc usecases.UseCase) gen.EnterpriseServiceServer {
	srv := EnterprisesGRPCServer{
		uc: uc,
	}

	gen.RegisterEnterpriseServiceServer(grpcServer, &srv)
	reflection.Register(grpcServer)

	return &srv
}

func (g *EnterprisesGRPCServer) GetEnterprisesList(ctx context.Context, request *gen.GetEnterprisesListRequest) (*gen.GetEnterprisesListResponse, error) {
	res := gen.GetEnterprisesListResponse{}
	entities, err := g.uc.GetEnterprisesList(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get enterprises list :%v", err)
	}
	res.Enterprises = make([]*gen.EnterpriseDto, 0, len(entities))

	for _, entity := range entities {
		res.Enterprises = append(res.Enterprises, &gen.EnterpriseDto{
			Id:              entity.ID.String(),
			Name:            entity.Name,
			Country:         entity.Country,
			MaintenanceYear: uint32(entity.MaintenanceYear),
			Phone:           entity.Phone,
			Fax:             entity.Fax,
			TypeOfOwnership: &gen.TypeOfOwnershipDto{
				Id:   entity.TypeOfOwnership.ID.String(),
				Name: entity.TypeOfOwnership.Name,
			},
		})
	}

	return &res, nil
}

func (g *EnterprisesGRPCServer) GetEnterpriseByID(ctx context.Context, request *gen.GetEnterpriseByIDRequest) (*gen.GetEnterpriseByIDResponse, error) {
	res := &gen.GetEnterpriseByIDResponse{}
	enterpriseId, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot parse id")
	}

	entity, err := g.uc.GetEnterpriseByID(ctx, enterpriseId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot get enterprise with id: %s", request.Id)
	}

	res.Enterprise = &gen.EnterpriseDto{
		Id:              entity.ID.String(),
		Name:            entity.Name,
		Country:         entity.Country,
		MaintenanceYear: uint32(entity.MaintenanceYear),
		Phone:           entity.Phone,
		Fax:             entity.Fax,
		TypeOfOwnership: &gen.TypeOfOwnershipDto{
			Id:   entity.TypeOfOwnership.ID.String(),
			Name: entity.TypeOfOwnership.Name,
		},
	}

	return res, nil
}
