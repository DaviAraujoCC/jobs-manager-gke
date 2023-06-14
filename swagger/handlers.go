// Package handlers for the RESTful Server
//
// Documentation for REST API
//
//				Schemes: http
//				BasePath: /
//				Version: 1.0.0
//
//				Consumes:
//				- application/json
//			 	- form-data
//
//				Produces:
//				- application/json
//
//		    Security:
//		    - Bearer:
//
//		    SecurityDefinitions:
//		    Bearer:
//		         type: apiKey
//		         name: Authorization
//		         in: header
//	          description: "Authorization header using the Bearer scheme. Example: \"Authorization: Bearer {token}\""
//
// swagger:meta
package handlers

import (
	"net/http"
)

// @termsOfService http://swagger.io/terms/

// swagger:route GET /cronjobs Cronjobs getCronjobs
// Get all cronjobs
//
// Security:
//   - Bearer: []
// responses:
// 	200: CronjobsResponse
// 	403: Unauthorized
// 	500: ErrorMessage

// swagger:route GET /cronjobs/{name} Cronjobs getCronjobByName
// Get a cronjob by name
//
// Security:
//  - Bearer: []
//
//
//  Parameters:
//   + name: name
//     in: path
//     description: CronJobs name
//     required: true
//     type: string
//
// responses:
// 	200: CronjobsResponse
// 	403: Unauthorized
// 	404: ErrorMessage
// 	500: ErrorMessage

// swagger:route PUT /cronjobs Cronjobs startCronjobByNamePut
// Start/Stop a cronjob by name
//
//   Security:
//  - Bearer: []
//
//  Parameters:
//   + name: cronjobName
//     in: query
//     description: Cronjob name
//     required: true
//     type: string
//   + name: action
//     in: query
//     description: Action to perform
//     required: true
//     type: string
//     enum: start,stop
//
// Deprecated: true
// responses:
// 	200: OK
// 	400: BadRequest
// 	403: Unauthorized
// 	404: ErrorMessage
// 	500: ErrorMessage

// swagger:route POST /cronjobs Cronjobs startCronjobByNamePost
// Start/Stop a cronjob by name
//
// Security:
//  - Bearer: []
//
//  Parameters:
//   + name: cronjobName
//     in: query
//     description: Cronjob name
//     required: true
//     type: string
//   + name: action
//     in: query
//     description: Action to perform
//     required: true
//     type: string
//     enum: start, stop
//
// responses:
// 	200: OK
// 	400: BadRequest
// 	403: Unauthorized
// 	404: ErrorMessage
// 	500: ErrorMessage

// CronjobsHandler is for handling everything that is not a match
func CronjobsHandler(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route GET /jobs/{name} Jobs getJobByName
// Get a job by name
//
// Security:
//  - Bearer: []
//
//  Parameters:
//   + name: name
//     in: path
//     description: Jobs name
//     required: true
//     type: string
//
// responses:
// 	200: JobsResponse
// 	400: BadRequest
// 	403: Unauthorized
// 	500: ErrorMessage

// swagger:route GET /jobs Jobs getJobs
// Get all jobs
//
// Security:
//  - Bearer: []
// responses:
// 	200: JobsResponse
// 	403: Unauthorized
// 	500: ErrorMessage

// JobsHandler is for handling everything that is not a match
func JobsHandler(rw http.ResponseWriter, r *http.Request) {

}

// Generic OK message returned as an HTTP Status Code
// swagger:response OK
type OK struct {
	// Description of the situation
	// in: body
	Body int
}

// Generic error message returned as an HTTP Status Code
// swagger:response ErrorMessage
type ErrorMessage struct {
	// Description of the situation
	// in: body
	Body int
}

// Generic BadRequest message returned as an HTTP Status Code
// swagger:response BadRequest
type BadRequest struct {
	// Description of the situation
	// in: body
	Body int
}

// Generic Unauthorized message returned as an HTTP Status Code
// swagger:response Unauthorized
type Unauthorized struct {
	// Description of the situation
	// in: body
	Body int
}

// Generic noContent message returned as an HTTP Status Code
// swagger:response noContent
type noContent struct {
	// Description of the situation
	// in: body
	Body int
}

// A list of Cronjobs
// swagger:response CronjobsResponse
type CronjobsResponseWrapper struct {
	// A list of cronjobs
	// in: body
	Body int
}

// A list of Jobs
// swagger:response JobsResponse
type JobsResponseWrapper struct {
	// A list of jobs
	// in: body
	Body int
}
