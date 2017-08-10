package api

import (
	restful "github.com/emicklei/go-restful"
)

/*New creates a new REST endpoint
 */
func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/data").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)


	return service
}

/*
func lastUpdate(request *restful.Request, response *restful.Response) {
	// TODO Fill in dummy function
	theJSON = (*restful.Response).WriteAsJson(*restful.Request)
	response.WriteEntity(theJSON)
}
*/
