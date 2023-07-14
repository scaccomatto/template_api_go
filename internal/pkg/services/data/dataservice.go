package data

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"template.com/restapi/internal/pkg/apperrors"
)

type Service struct {
	DataList map[uuid.UUID]Data
}

func (sv *Service) CreateData(name string, value int) (Data, error) {
	newData := Data{
		Id:    uuid.New(),
		Name:  name,
		Value: value,
	}

	_, ok := sv.DataList[newData.Id]
	if ok {
		log.Error().Msgf(" collision %s is already in the list", newData.Id)
		return Data{}, apperrors.Builder().Message("internal error").Build()
	} else {
		sv.DataList[newData.Id] = newData
		return newData, nil
	}
}

func (sv *Service) GetDataById(id uuid.UUID) (Data, error) {
	target, ok := sv.DataList[id]
	if !ok {
		return Data{}, apperrors.Builder().Message(fmt.Sprintf("%s is not in the list", id)).Status(apperrors.NotFound).Build()
	} else {
		return target, nil
	}
}

func NewService() *Service {
	return &Service{
		DataList: make(map[uuid.UUID]Data),
	}
}
