package employee

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/rs/zerolog/log"
	"github.com/saiSunkari19/aicumen/blockchain/x/org"
	"github.com/saiSunkari19/aicumen/client/config"
	validator "gopkg.in/go-playground/validator.v9"
	"net/http"
)

type UpdateEmployeeBody struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	ID         string       `json:"id" validate:"required"`
	Department string       `json:"department"`
	Skills     []string     `json:"skills"`
	Address    string       `json:"address"`
	Password   string       `json:"password"`
}

func UpdateEmployee(cli *config.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("METHOD", r.Method).Str("URL", r.RequestURI).Str("HOST", r.RemoteAddr).Msg("Update Employee info")

		v := validator.New()

		var req UpdateEmployeeBody
		if !rest.ReadRESTReq(w, r, cli.CliCtx.Codec, &req) {
			return
		}

		err := v.Struct(req)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := org.NewMsgUpdateEmployeeInfo(req.ID, req.Address, req.Department, req.Skills, cli.CliCtx.GetFromAddress())
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		config.BuildSignAndBroadCast(w, cli, req.BaseReq, []sdk.Msg{msg}, req.Password)

	}
}
