package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	apphttp "shortest-path/http"
	"shortest-path/service"
	"shortest-path/storage"
)

type Form struct {
	From string `validate:"required,is-valid-airport"`
	To   string `validate:"required,is-valid-airport"`
}

func (f *Form) handleRequest(request http.Request) {
	keys, ok := request.URL.Query()["from"]
	if ok && len(keys[0]) >= 1 {
		f.From = keys[0]
	}

	keys, ok = request.URL.Query()["to"]
	if ok && len(keys[0]) >= 1 {
		f.To = keys[0]
	}
}

type ShortestPathHandler struct{}

func (h ShortestPathHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	form := Form{}
	form.handleRequest(*request)

	validate := validator.New()
	validate.RegisterValidation("is-valid-airport", IsValidAirport)

	err := validate.Struct(form)

	if err != nil {
		http.Error(writer, string(apphttp.EncodeError(err)), http.StatusUnprocessableEntity)
		return
	}

	vertices := storage.GetDefaultStorage().GetVertices()
	path := service.FindShortestPathWithMaxEdge(vertices, *vertices[form.From], *vertices[form.To], 4)

	jsonBody, err := json.Marshal(path)
	fmt.Fprintf(writer, string(jsonBody))

}

func IsValidAirport(fl validator.FieldLevel) bool {
	_, ok := storage.GetDefaultStorage().GetVertices()[fl.Field().String()]

	return ok
}
