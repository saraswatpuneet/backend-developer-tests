package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stackpath/backend-developer-tests/rest-service/constants"
	"github.com/stackpath/backend-developer-tests/rest-service/contracts"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
)

//HealthCheckHandler ...
func HealthCheckHandler(c *gin.Context) {
	log.Debugf("HealthCheckHandler Method --> %s", c.Request.Method)

	switch c.Request.Method {
	case http.MethodGet:
		GetHealthStatus(c.Writer)
	default:
		err := errors.New("method not supported")
		ResponseError(c.Writer, http.StatusBadRequest, err)
	}

}

// AllPeopleHandler ... single handler to either get all people to based on provided filters
func AllPeopleHandler(c *gin.Context) {
	log.Info("Request is made to get list of people")
	firstName := c.Query("first_name")
	lastName := c.Query("last_name")
	phoneNumber := c.Query("phone_number")
	currentFilter := contracts.FilterContract{
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
	}
	filterProvided := currentFilter.FirstName != "" || currentFilter.LastName != "" || currentFilter.PhoneNumber != ""
	allPeople := models.AllPeople()
	outputPeople := make(map[string]*models.Person, 0)
	// support if  first or last name are provided or phone number is provided
	// permutations of first and last name and phone number are supported
	for _, person := range allPeople {
		if filterProvided {
			if currentFilter.PhoneNumber != "" && person.PhoneNumber == currentFilter.PhoneNumber {
				outputPeople[person.ID.String()] = person
			}
			if (currentFilter.FirstName != "" && person.FirstName == currentFilter.FirstName) || (currentFilter.LastName != "" && person.LastName == lastName) {
				outputPeople[person.ID.String()] = person
			}
		} else {
			outputPeople[person.ID.String()] = person
		}
	}
	resultingList := make([]*models.Person, 0)
	for _, person := range outputPeople {
		resultingList = append(resultingList, person)
	}
	returnStatus := http.StatusOK
	if len(outputPeople) == 0 {
		returnStatus = http.StatusNotFound
	}
	c.JSON(returnStatus, gin.H{
		constants.RESPONSE_JSON_DATA:   resultingList,
		constants.RESPONSDE_JSON_ERROR: nil,
	})
}

// PersonWithIDHandler ... handle when specific person with id is requested
func PersonWithIDHandler(c *gin.Context) {
	log.Info("Request is made to get a specific person")
	outputPeople := make([]*models.Person, 0)
	personId := c.Param("id")
	// uuid will panic if wrong id is sent as uuid string
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("PersonWithIDHandler ... panicked, wrong id specified: %v", r)
			c.JSON(http.StatusNotFound, gin.H{
				constants.RESPONSE_JSON_DATA:   outputPeople,
				constants.RESPONSDE_JSON_ERROR: fmt.Errorf("invalid person id").Error(),
			})
			return
		}
	}()
	// convert person string uuid to uuid
	personUUID := uuid.Must(uuid.FromString(personId))
	allPeople := models.AllPeople()
	// only support if both first and last name are provided or phone number is provided
	for _, person := range allPeople {
		if person.ID == personUUID {
			outputPeople = append(outputPeople, person)
		}
	}
	returnStatus := http.StatusOK
	if len(outputPeople) == 0 {
		returnStatus = http.StatusNotFound
	}
	c.JSON(returnStatus, gin.H{
		constants.RESPONSE_JSON_DATA:   outputPeople,
		constants.RESPONSDE_JSON_ERROR: nil,
	})
}

//GetHealthStatus ...
func GetHealthStatus(w http.ResponseWriter) {
	healthStatus := "backend is healthy"
	response, _ := json.Marshal(healthStatus)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(response); err != nil {
		log.Errorf("GetHealthStatus ... unable to write JSON response: %v", err)
	}
}

// ResponseError ... essentially a single point of sending some error to route back
func ResponseError(w http.ResponseWriter, httpStatusCode int, err error) {
	log.Errorf("Response error %s", err.Error())
	response, _ := json.Marshal(err)
	w.Header().Add("Status", strconv.Itoa(httpStatusCode)+" "+err.Error())
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(httpStatusCode)

	if _, err := w.Write(response); err != nil {
		log.Errorf("ResponseError ... unable to write JSON response: %v", err)
	}
}
