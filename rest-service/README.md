# Overview
## rest-sevice is a RESTful API service for the primarily following APIs:
```
Note: currently, rest service do not implement authentication. Please only use it for dev purpose.

General: payload returns a data field (on successful completing) and error field (on error is not null)
1. GET /v1/people: get a list of people in the system. The response is a list of people with following fields:
- id: unique id of the person
- first_name: first name of the person
- last_name: last name of the person
- phone_number: phone number of the person

2. GET /v1/people/{id}: get a person by id. The response is a person with following fields:
- id: unique id of the person
- first_name: first name of the person
- last_name: last name of the person
- phone_number: phone number of the person
If the person does not exist, the response is 404 and the body is empty with a generic error message in error field.

3. GET /v1/people?phone_number={phone_number}: get a person by phone number. The response is a person with following fields:
- id: unique id of the person
- first_name: first name of the person
- last_name: last name of the person
- phone_number: phone number of the person
If the person does not exist, the response is 404 and the body is empty with a generic error message in error field.

4. GET /v1/people?first_name={first_name}&last_name={last_name}: get a person by first name and last name. The response is a person with following fields:
- id: unique id of the person
- first_name: first name of the person
- last_name: last name of the person
- phone_number: phone number of the person
If the person does not exist, the response is 404 and the body is empty with a generic error message in error field.
```
# Running the service
To run the service, you need to install the dependencies and run the following command:
```
go get : to get all the required dependencies
go run main.go : to run the service
By default the service will do pretty logs via logrus for auditing.
By default the port is 8090. To change it simply set env variable PORT to the desired port or
update the /options/default.go port field to the desired port.
```

# Testing the service
```
To test the service, you need to install the dependencies and run the following command:
go get : to get all the required dependencies
go test ./... : to run the tests, this command should be either run from root to run all test in project or to run specific tests for service run go test in rest-service folder

go test ./... -race -covermode=atomic --coverprofile=coverage.out : to check for code coverage 
```

# Layout of the rest-service project
```
1. main.go starts a graceful server from pkg/core/coreserver.go
2. gin-gonic is used for the rest service to define routers and simple middlewares such as cors, either auth checks can be implemented a middleware or separate sidecar proxy service can be used.
3. pkg/core/controller/ : defines the backend server controller which actually implements a graceful server
4. pkg/core/router/ : defines the routing mechanism using gin-gonic default router, more exhaustive routers can be defined in this package using gin-gonic non default router with additional features
5. /options/ : defines the options for the service this can be extended to include other options such as logging, metrics, etc.
6 /config/ : defines the configuration for the service, this can be extended to include other configuration such as database configuration, etc. via environment variables.
7. pkg/models/handlers/ : defines the handler for the service, this colocate with the models though handlers can be moved to outer scope of the project for better organization.
```

# Running tests via postman
```
To run tests via postman, import the collection (json) file provided in postman folder and run the requests after starting the server as detailed above in section Running the service.
```

# Logging
```
By default the service will do pretty logs via logrus for auditing. logrus is a standard library logger with exhausting messaging and traceability.
logrus is set to log json format for any level of logging. By default in debug mode the service will log at Debug Level which entailvery verbose logging.

In production mode, the service will log at TraceLevel Level which entail finer-grained informational events than the Debug.
```

# Ops
```
Logrus is used to enabled json formatted logs from info level, error level to simple debug level. In a ELK stack or open telemetry these json fields can be set as trigger point when specific level of error occurs. 
```

