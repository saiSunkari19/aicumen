package org

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (result *sdk.Result, err error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case MsgAddEmployee:
			return handleMsgAddEmployee(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized ORG message type: %T", msg)
		}

	}
}

func handleMsgAddEmployee(ctx sdk.Context, keeper Keeper, msg MsgAddEmployee) (*sdk.Result, error) {
	count := keeper.GetGlobalEmployeeCount(ctx)
	id := GetEmployeePrefixKey(count)

	employe := Employee{
		Person: Person{
			Name:    msg.Name,
			Address: msg.Address,
			Skills:  msg.Skills,
		},
		ID:         id,
		Department: msg.Department,
		Status:     StatusActive,
	}

	if err := keeper.AddEmployee(ctx, employe); err != nil {
		return nil, err
	}

	if err := keeper.SetGlobalEmployeeCount(ctx, count+1); err != nil {
		return nil, err
	}
	keeper.SetActiveEmployee(ctx, id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
