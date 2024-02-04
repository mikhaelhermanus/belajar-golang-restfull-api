package helper

import (
	"belajar-golang-restful-api/model/web"
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.WriteHeader(response.(web.WebResponse).Code)     // parsing inteface struct into object
	writer.Header().Add("Content-Type", "application/json") // tell the system it json
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
