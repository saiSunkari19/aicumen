package rest

import (
	"github.com/gorilla/mux"
	"github.com/saiSunkari19/aicumen/client/config"
	"github.com/saiSunkari19/aicumen/client/rest/employee"
)

func RegisterRoutes(cliCtx *config.CLI, r *mux.Router) {
	e := r.PathPrefix("/employee").Subrouter()
	e.HandleFunc("/add", employee.AddEmployee(cliCtx)).Methods("POST")

	e.HandleFunc("/list", employee.QueryEmployees(cliCtx.CliCtx, r)).Methods("GET")

}
