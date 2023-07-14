package datahandler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"template.com/restapi/internal/app/httpserver"
	"template.com/restapi/internal/pkg/apperrors"
	"template.com/restapi/internal/pkg/types/apihttp"
)

type datahandler struct{}

// GetData @Summary 	Get Data
// @Description Get data.
// @Tags 		Data
// @Accept 		json
// @Param		dataId		path 	string	true 	"The data id"
// @Success 	200  {object}  data.Data
// @Failure     400  {object}  apperrors.HttpError
// @Failure     500  {object}  apperrors.HttpError
// @Router 		/data/{dataId} [get]
func GetDataById(s *httpserver.Server) *echo.Route {
	dh := datahandler{}
	return s.Router.Root.GET("/data/:dataId", dh.getDataById(s))
}

func (dh *datahandler) getDataById(s *httpserver.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		dataId, err := uuid.Parse(c.Param("dataId"))
		if err != nil || dataId == uuid.Nil {
			return apperrors.Builder().Status(apperrors.BadRequest).Message("invalid platformUserId param").Error(err).Build()
		}

		target, err := s.DataService.GetDataById(dataId)
		if err != nil {
			return apperrors.Builder().
				Status(apperrors.NotFound).
				Message(fmt.Sprintf("id %s not found", dataId)).
				Error(err).Build()
		}

		return c.JSON(http.StatusOK, target)
	}
}

// CreateData @Summary 	Create Data
// @Description Create data.
// @Tags 		Data
// @Accept 		json
// @Param 		Body  body	   apihttp.CreateDataRequest 		true 	"Data name"
// @Success 	200  {object}  data.Data
// @Failure     400  {object}  apperrors.HttpError
// @Failure     500  {object}  apperrors.HttpError
// @Router 		/data [post]
func CreateData(s *httpserver.Server) *echo.Route {
	dh := datahandler{}
	return s.Router.Root.POST("/data", dh.createData(s))
}

func (dh *datahandler) createData(s *httpserver.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := apihttp.CreateDataRequest{}
		if err := c.Bind(&req); err != nil {
			return apperrors.Builder().Status(apperrors.BadRequest).Error(err).Build()
		}
		target, err := s.DataService.CreateData(req.Name, req.Value)
		if err != nil {
			return apperrors.Builder().Status(apperrors.BadRequest).Error(err).Build()
		}
		return c.JSON(http.StatusCreated, target)
	}
}
