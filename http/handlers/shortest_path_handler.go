package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	apphttp "github.com/modiamir/shortest-path/http"
	"github.com/modiamir/shortest-path/service"
	"github.com/modiamir/shortest-path/storage"
	"net/http"
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
	shortestPathFinder := service.NewShortestPathWithMaxEdgeFinder(vertices, 4)
	path, _ := shortestPathFinder.Find(form.From, form.To)

	writer.Header().Add("Content-Type", "application/json")
	jsonBody, err := json.Marshal(path)
	fmt.Fprintf(writer, string(jsonBody))

}

func IsValidAirport(fl validator.FieldLevel) bool {
	_, ok := storage.GetDefaultStorage().GetVertices()[fl.Field().String()]

	return ok
}
