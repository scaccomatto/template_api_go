package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"template.com/restapi/internal/app/httpserver"
	"template.com/restapi/internal/app/httpserver/handlers"
	"template.com/restapi/internal/pkg/services/data"
	"template.com/restapi/internal/pkg/types/apihttp"
	"testing"
)

func WithServerInitialization(t *testing.T, executeTest func(s *httpserver.Server)) {
	t.Helper()

	serverInit(t, executeTest)
}

func serverInit(t *testing.T, executeTest func(s *httpserver.Server)) {
	t.Helper()
	dataServices := data.NewService()
	server := httpserver.New(dataServices)

	handlers.Init(server)

	executeTest(server)

	server.Shutdown(context.Background())

	//server = nil
}

func GetDataRequest(id string) *http.Request {
	url := fmt.Sprintf("/data/%s", id)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return req
}

func CreateDataRequest(name string) *http.Request {
	url := fmt.Sprintf("/data")
	b := apihttp.CreateDataRequest{
		Name: name,
	}
	body, _ := json.Marshal(b)
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return req
}
