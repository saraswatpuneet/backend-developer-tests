// this file tests handler functions in handlers.go
package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	uuid "github.com/kevinburke/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_AllPeopleHandler(t *testing.T) {
	routeEngine := gin.Default()
	routeEngine.GET("/v1/people", AllPeopleHandler)
	httpTester := httptest.NewRecorder()
	buf := bytes.NewBuffer([]byte(`{}`))
	req, _ := http.NewRequest("GET", "/v1/people", buf)
	routeEngine.ServeHTTP(httpTester, req)
	assert.Equal(t, http.StatusOK, httpTester.Code)
	decoder := json.NewDecoder(httpTester.Body)
	var listPerson map[string]interface{}
	err := decoder.Decode(&listPerson)
	if err != nil {
		t.Errorf("Error getting list of people: %v", err)
	}
	assert.NotNil(t, listPerson["data"])
	assert.Empty(t, listPerson["error"])
	listOutput := listPerson["data"].([]interface{})
	if listOutput == nil {
		t.Errorf("Error unmarshalling list of people")
	}
	assert.Greater(t, len(listOutput), 0)
}

func Test_AllPeopleHandlerWithFilters(t *testing.T) {
	routeEngine := gin.Default()
	routeEngine.GET("/v1/people", AllPeopleHandler)
	httpTester := httptest.NewRecorder()
	buf := bytes.NewBuffer([]byte(`{}`))
	req, _ := http.NewRequest("GET", "/v1/people", buf)
	req.URL.RawQuery = "first_name=John&last_name=Doe"
	routeEngine.ServeHTTP(httpTester, req)
	assert.Equal(t, http.StatusOK, httpTester.Code)
	decoder := json.NewDecoder(httpTester.Body)
	var listPerson map[string]interface{}
	err := decoder.Decode(&listPerson)
	if err != nil {
		t.Errorf("Error getting list of people: %v", err)
	}
	assert.NotNil(t, listPerson["data"])
	assert.Empty(t, listPerson["error"])
	listOutput := listPerson["data"].([]interface{})
	if listOutput == nil {
		t.Errorf("Error unmarshalling list of people")
	}
	assert.Greater(t, len(listOutput), 0)
	for _, person := range listOutput {
		personMap := person.(map[string]interface{})
		assert.Equal(t, "John", personMap["first_name"])
		assert.Equal(t, "Doe", personMap["last_name"])
	}
}

func Test_AllPeopleHandlerWithFiltersPhone(t *testing.T) {
	routeEngine := gin.Default()
	routeEngine.GET("/v1/people", AllPeopleHandler)
	httpTester := httptest.NewRecorder()
	buf := bytes.NewBuffer([]byte(`{}`))
	req, _ := http.NewRequest("GET", "/v1/people", buf)
	req.URL.RawQuery = "phone_number=+1 (800) 555-1212"
	routeEngine.ServeHTTP(httpTester, req)
	assert.Equal(t, http.StatusOK, httpTester.Code)
	decoder := json.NewDecoder(httpTester.Body)
	var listPerson map[string]interface{}
	err := decoder.Decode(&listPerson)
	if err != nil {
		t.Errorf("Error getting list of people: %v", err)
	}
	assert.NotNil(t, listPerson["data"])
	assert.Empty(t, listPerson["error"])
	listOutput := listPerson["data"].([]interface{})
	if listOutput == nil {
		t.Errorf("Error unmarshalling list of people")
	}
	assert.Greater(t, len(listOutput), 0)
	for _, person := range listOutput {
		personMap := person.(map[string]interface{})
		assert.Equal(t, "John", personMap["first_name"])
		assert.Equal(t, "Doe", personMap["last_name"])
	}
}

func Test_AllPeopleHandlerWithID(t *testing.T) {
	routeEngine := gin.Default()
	routeEngine.GET("/v1/people/:id", PersonWithIDHandler)
	httpTester := httptest.NewRecorder()
	buf := bytes.NewBuffer([]byte(`{}`))
	req, _ := http.NewRequest("GET", "/v1/people/81eb745b-3aae-400b-959f-748fcafafd81", buf)
	routeEngine.ServeHTTP(httpTester, req)
	assert.Equal(t, http.StatusOK, httpTester.Code)
	decoder := json.NewDecoder(httpTester.Body)
	var listPerson map[string]interface{}
	err := decoder.Decode(&listPerson)
	if err != nil {
		t.Errorf("Error getting list of people: %v", err)
	}
	assert.NotNil(t, listPerson["data"])
	assert.Empty(t, listPerson["error"])
}
func Test_HealthCheckHandler(t *testing.T) {
	routeEngine := gin.Default()
	routeEngine.GET("/v1/healthz", HealthCheckHandler)
	httpTester := httptest.NewRecorder()
	buf := bytes.NewBuffer([]byte(`{}`))
	req, _ := http.NewRequest("GET", "/v1/healthz", buf)
	routeEngine.ServeHTTP(httpTester, req)
	assert.Equal(t, http.StatusOK, httpTester.Code)
}

func Test_AllPeopleHandlerWithFilters_negative1(t *testing.T) {
	routeEngine := gin.Default()
	routeEngine.GET("/v1/people", AllPeopleHandler)
	httpTester := httptest.NewRecorder()
	buf := bytes.NewBuffer([]byte(`{}`))
	req, _ := http.NewRequest("GET", "/v1/people", buf)
	req.URL.RawQuery = "first_name=Jonny&last_name=Does"
	routeEngine.ServeHTTP(httpTester, req)
	assert.Equal(t, http.StatusNotFound, httpTester.Code)
	decoder := json.NewDecoder(httpTester.Body)
	var listPerson map[string]interface{}
	err := decoder.Decode(&listPerson)
	if err != nil {
		t.Errorf("Error getting list of people: %v", err)
	}
	assert.NotEmpty(t, listPerson["error"])
	listOutput := listPerson["data"].([]interface{})
	if listOutput == nil {
		t.Errorf("Error unmarshalling list of people")
	}
	assert.Equal(t, len(listOutput), 0)
}

func Test_AllPeopleHandlerWithFilters_negative2(t *testing.T) {
	routeEngine := gin.Default()
	routeEngine.GET("/v1/people", AllPeopleHandler)
	httpTester := httptest.NewRecorder()
	buf := bytes.NewBuffer([]byte(`{}`))
	req, _ := http.NewRequest("GET", "/v1/people", buf)
	req.URL.RawQuery = "first_name=Jonny"
	routeEngine.ServeHTTP(httpTester, req)
	assert.Equal(t, http.StatusBadRequest, httpTester.Code)
}
func Test_AllPeopleHandlerWithFilters_negative3(t *testing.T) {
	routeEngine := gin.Default()
	routeEngine.GET("/v1/people", AllPeopleHandler)
	httpTester := httptest.NewRecorder()
	buf := bytes.NewBuffer([]byte(`{}`))
	req, _ := http.NewRequest("GET", "/v1/people", buf)
	req.URL.RawQuery = "first_name=Jonny&last_name=Does&phone_number=1234"
	routeEngine.ServeHTTP(httpTester, req)
	assert.Equal(t, http.StatusBadRequest, httpTester.Code)
}

func Test_AllPeopleHandlerWithID_negative1(t *testing.T) {
	routeEngine := gin.Default()
	routeEngine.GET("/v1/people/:id", PersonWithIDHandler)
	httpTester := httptest.NewRecorder()
	buf := bytes.NewBuffer([]byte(`{}`))
	req, _ := http.NewRequest("GET", "/v1/people/something", buf)
	routeEngine.ServeHTTP(httpTester, req)
	assert.Equal(t, http.StatusBadRequest, httpTester.Code)
}

func Test_AllPeopleHandlerWithID_negative2(t *testing.T) {
	routeEngine := gin.Default()
	routeEngine.GET("/v1/people/:id", PersonWithIDHandler)
	httpTester := httptest.NewRecorder()
	buf := bytes.NewBuffer([]byte(`{}`))
	currentUUID := uuid.NewV4().String()
	pathWithFakeUUID := fmt.Sprintf("/v1/people/%s", currentUUID)
	req, _ := http.NewRequest("GET", pathWithFakeUUID, buf)
	routeEngine.ServeHTTP(httpTester, req)
	assert.Equal(t, http.StatusNotFound, httpTester.Code)
}
