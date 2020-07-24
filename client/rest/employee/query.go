package employee

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	
	"github.com/saiSunkari19/aicumen/blockchain/x/org"
)

func QueryEmployees(cliCtx context.CLIContext, r *mux.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var data org.Employess
		var res []byte
		var err error
		
		id := r.URL.Query().Get("id")
		dept := r.URL.Query().Get("dept")
		name := r.URL.Query().Get("name")
		
		if len(id) > 0 {
			log.Info().Str("METHOD", r.Method).Str("URL", r.RequestURI).Str("HOST", r.RemoteAddr).Msg("Query Employees By ID")
			
			res, _, err = cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", org.ModuleName, org.QueryEmployeeByID, id), nil)
			if rest.CheckInternalServerError(w, err) {
				return
			}
			
			if len(res) == 0 {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("res is empty"))
				return
			}
			
			var employee org.Employee
			cliCtx.Codec.MustUnmarshalJSON(res, &employee)
			data = append(data, employee)
			
		} else {
			if len(dept) > 0 {
				log.Info().Str("METHOD", r.Method).Str("URL", r.RequestURI).Str("HOST", r.RemoteAddr).Msg("Query Employees By Dept")
				
				res, _, err = cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", org.ModuleName, org.QueryByDepartment, dept), nil)
				if rest.CheckInternalServerError(w, err) {
					return
				}
			} else if len(name) > 0 {
				
				log.Info().Str("METHOD", r.Method).Str("URL", r.RequestURI).Str("HOST", r.RemoteAddr).Msg("Query Employees By Name")
				
				res, _, err = cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", org.ModuleName, org.QueryByName, name), nil)
				if rest.CheckInternalServerError(w, err) {
					return
				}
			} else {
				log.Info().Str("METHOD", r.Method).Str("URL", r.RequestURI).Str("HOST", r.RemoteAddr).Msg("Query Employees")
				
				res, _, err = cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", org.ModuleName, org.QueryActiveEmployees), nil)
				if rest.CheckInternalServerError(w, err) {
					return
				}
			}
			
			if len(res) == 0 {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("data is empty"))
				return
			}
			
			cliCtx.Codec.MustUnmarshalJSON(res, &data)
		}
		rest.PostProcessResponse(w, cliCtx, data)
	}
}

// -----------------------------------------------
func QueryInActiveEmployeesInfo(cliCtx context.CLIContext, r *mux.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		log.Info().Str("METHOD", r.Method).Str("URL", r.RequestURI).Str("HOST", r.RemoteAddr).Msg("Query InActive Employees ")
		
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", org.ModuleName, org.QueryDeActiveEmployees), nil)
		if rest.CheckInternalServerError(w, err) {
			return
		}
		
		if len(res) == 0 {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("data is empty"))
			return
		}
		
		var data org.Employess
		cliCtx.Codec.MustUnmarshalJSON(res, &data)
		rest.PostProcessResponse(w, cliCtx, data)
		
	}
}

// ------------------------------------------------------------------------------------

func QueryBySearch(cliCtx context.CLIContext, r *mux.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var employees org.Employess
		var res []byte
		var err error
		
		id := r.URL.Query().Get("id")
		dept := r.URL.Query().Get("dept")
		name := r.URL.Query().Get("name")
		addr := r.URL.Query().Get("addr")
		skills := r.URL.Query().Get("skills")
		
		if len(id) > 0 {
			log.Info().Str("METHOD", r.Method).Str("URL", r.RequestURI).Str("HOST", r.RemoteAddr).Msg("Search Employees By ID")
			
			res, _, err = cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", org.ModuleName, org.QueryEmployeeByID, id), nil)
			if rest.CheckInternalServerError(w, err) {
				return
			}
			
			if len(res) == 0 {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("res is empty"))
				return
			}
			
			var employee org.Employee
			cliCtx.Codec.MustUnmarshalJSON(res, &employee)
			employees = append(employees, employee)
		} else {
			if len(dept) > 0 {
				log.Info().Str("METHOD", r.Method).Str("URL", r.RequestURI).Str("HOST", r.RemoteAddr).Msg("Search Employees By Dept")
				
				res, _, err = cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s/%s", org.ModuleName, org.QueryBySearch, "dept", dept), nil)
				if rest.CheckInternalServerError(w, err) {
					return
				}
			} else if len(name) > 0 {
				
				log.Info().Str("METHOD", r.Method).Str("URL", r.RequestURI).Str("HOST", r.RemoteAddr).Msg("Search Employees By Name")
				
				res, _, err = cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s/%s", org.ModuleName, org.QueryBySearch, "name", name), nil)
				if rest.CheckInternalServerError(w, err) {
					return
				}
				
			} else if len(addr) > 0 {
				log.Info().Str("METHOD", r.Method).Str("URL", r.RequestURI).Str("HOST", r.RemoteAddr).Msg("Search Employees By Address")
				res, _, err = cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s/%s", org.ModuleName, org.QueryBySearch, "addr", addr), nil)
				if rest.CheckInternalServerError(w, err) {
					return
				}
				
			} else if len(skills) > 0 {
				log.Info().Str("METHOD", r.Method).Str("URL", r.RequestURI).Str("HOST", r.RemoteAddr).Msg("Search Employees By Skills")
				
				param := strings.TrimSpace(skills)
				data := strings.Split(param, ",")
				if findDup(data) {
					rest.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("duplicate skills data"))
					return
				}
				
				res, _, err = cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s/%s", org.ModuleName, org.QueryBySearch, "skills", skills), nil)
				if rest.CheckInternalServerError(w, err) {
					return
				}
			} else {
				rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("query params not provided/incorrect"))
				return
			}
			
			if len(res) == 0 {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("data is empty"))
				return
			}
			
			cliCtx.Codec.MustUnmarshalJSON(res, &employees)
		}
		rest.PostProcessResponse(w, cliCtx, employees)
	}
}

func findDup(data []string) bool {
	sort.Strings(data)
	
	prev := data[0]
	
	for i := 1; i < len(data); i++ {
		if strings.EqualFold(data[i], prev) {
			return true
		}
		prev = data[i]
	}
	return false
}
