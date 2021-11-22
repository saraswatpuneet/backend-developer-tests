package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AllPeopleHandler(t *testing.T) {
	routeEngine, err := BackendRouter()
	if err != nil {
		t.Errorf("Error creating route engine: %s", err.Error())
	}
	httpTester := httptest.NewRecorder()
	buf := bytes.NewBuffer([]byte(`{}`))
	req, _ := http.NewRequest("GET", "/v1/people", buf)
	routeEngine.ServeHTTP(httpTester, req)
	assert.Equal(t, http.StatusOK, httpTester.Code)
	decoder := json.NewDecoder(httpTester.Body)
	var listPerson map[string]interface{}
	err = decoder.Decode(&listPerson)
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
