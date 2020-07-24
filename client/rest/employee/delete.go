package employee

import (
	"net/http"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/rs/zerolog/log"
	validator "gopkg.in/go-playground/validator.v9"
	
	"github.com/saiSunkari19/aicumen/blockchain/x/org"
	"github.com/saiSunkari19/aicumen/client/config"
)

type DeleteEmployeeBody struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	ID       string       `json:"id" validate:"required"`
	Remove   bool         `json:"remove"`
	Password string       `json:"password"`
}

func DeleteEmployee(cli *config.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("METHOD", r.Method).Str("URL", r.RequestURI).Str("HOST", r.RemoteAddr).Msg("Delete Employee info")
		
		v := validator.New()
		
		var req DeleteEmployeeBody
		if !rest.ReadRESTReq(w, r, cli.CliCtx.Codec, &req) {
			return
		}
		
		err := v.Struct(req)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		msg := org.NewMsgDeleteEmployeeInfo(req.ID, req.Remove, cli.CliCtx.GetFromAddress())
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}
		
		config.BuildSignAndBroadCast(w, cli, req.BaseReq, []sdk.Msg{msg}, req.Password)
		
	}
}
