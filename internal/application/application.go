package application

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/arbuzick57/calc_go/pkg/calc"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type ResponseCorrect struct {
	Result float64 `json:"result"`
}

type ResponseError struct {
	Error string `json:"error"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	errExpression := calc.CheckExpression(request.Expression)
	if errExpression != nil {
		response := ResponseError{
			Error: errExpression.Error(),
		}
		responseJson, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(responseJson)
		return
	}

	result, errCalc := calc.Calc(request.Expression)
	if errCalc != nil {
		response := ResponseError{
			Error: errCalc.Error(),
		}
		responseJson, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(responseJson)
		return
	}
	response := ResponseCorrect{
		Result: result,
	}
	responseJson, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (a *Application) RunServer() error {
	http.HandleFunc("/", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
