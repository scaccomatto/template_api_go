package datahandler_test

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"template.com/restapi/internal/app/httpserver"
	"template.com/restapi/internal/pkg/services/data"
	"template.com/restapi/internal/pkg/test"
	"testing"
)

func Test_GivenGetRequestWhenDataIdHasIncorrectFormatThenReturnBadRequest(t *testing.T) {
	test.WithServerInitialization(t, func(s *httpserver.Server) {
		req := test.GetDataRequest("123")
		resp := httptest.NewRecorder()

		s.ServeHTTP(resp, req)

		data, _ := ioutil.ReadAll(resp.Body)
		msg := string(data)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, msg, "invalid platformUserId param")
	})
}

func Test_GivenCorrectGetRequestWhenDataIsNotExistingThenReturnNotFound(t *testing.T) {
	test.WithServerInitialization(t, func(s *httpserver.Server) {
		//set up
		id := uuid.New().String()
		req := test.GetDataRequest(id)
		resp := httptest.NewRecorder()

		// act
		s.ServeHTTP(resp, req)

		data, _ := ioutil.ReadAll(resp.Body)
		msg := string(data)

		// assertions
		assert.Equal(t, http.StatusNotFound, resp.Code)
		assert.Contains(t, msg, "is not in the list")

	})
}

func Test_GivenCreateRequestWhenDataIsCorrectThenReturnCreated(t *testing.T) {
	test.WithServerInitialization(t, func(s *httpserver.Server) {
		//set up
		name := "test"
		req := test.CreateDataRequest(name)
		resp := httptest.NewRecorder()

		// act
		s.ServeHTTP(resp, req)

		// assertions
		data := data.Data{}
		b, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(b, &data)

		assert.Equal(t, http.StatusCreated, resp.Code)
		assert.Contains(t, data.Name, name)

	})
}
