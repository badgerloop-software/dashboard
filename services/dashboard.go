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

	service.Route(service.GET("/lastupdate").To(lastUpdate))

	return service
}

func lastUpdate(request *restful.Request, response *restful.Response) {
	// TODO Fill in dummy function
	response.WriteEntity("Dummy Value")
}
