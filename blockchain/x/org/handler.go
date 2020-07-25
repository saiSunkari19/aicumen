package org

import (
	"fmt"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (result *sdk.Result, err error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		
		switch msg := msg.(type) {
		case MsgAddEmployee:
			return handleMsgAddEmployeeInfo(ctx, keeper, msg)
		case MsgUpdateEmployeeInfo:
			return handleMsgUpdateEmployeeInfo(ctx, keeper, msg)
		case MsgDeleteEmployeeInfo:
			return handleMsgDeleteEmployeeInfo(ctx, keeper, msg)
		case MsgRestoreEmployeeInfo:
			return handleMsgRestoreEmployeeInfo(ctx, keeper, msg)
		
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized org	 message type: %T", msg)
		}
		
	}
}

func handleMsgAddEmployeeInfo(ctx sdk.Context, keeper Keeper, msg MsgAddEmployee) (*sdk.Result, error) {
	count := keeper.GetGlobalEmployeeCount(ctx)
	id := GetEmployeePrefixKey(count)
	
	employe := Employee{
		Person: Person{
			Name:    msg.Name,
			Address: msg.Address,
			Skills:  msg.Skills,
			Owner:   msg.From,
		},
		ID:         id,
		Department: msg.Department,
		Status:     StatusActive,
	}
	
	if err := keeper.AddEmployeeInfo(ctx, employe); err != nil {
		return nil, err
	}
	
	if err := keeper.SetGlobalEmployeeCount(ctx, count+1); err != nil {
		return nil, err
	}
	keeper.SetActiveEmployee(ctx, id)
	
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		EventTypeMsgAddEmployes,
		sdk.NewAttribute(AttributeKeyEmployeeID, employe.ID),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.From.String()),
	))
	
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateEmployeeInfo(ctx sdk.Context, k Keeper, msg MsgUpdateEmployeeInfo) (*sdk.Result, error) {
	
	employee, found := k.GetEmployee(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(ErrEmployee, fmt.Sprintf("employee %s not found", msg.Id))
	}
	
	if !employee.Person.Owner.Equals(msg.From) {
		return nil, sdkerrors.Wrap(ErrEmployee, fmt.Sprintf("employee unauthorised %s", msg.From))
	}
	
	if employee.Status != StatusActive {
		return nil, sdkerrors.Wrap(ErrEmployee, fmt.Sprintf("invalid employee status"))
	}
	
	if len(msg.Department) > 0 {
		employee.Department = msg.Department
	}
	
	if len(msg.Address) > 0 {
		employee.Person.Address = msg.Address
	}
	
	skills := msg.Skills.Sort()
	if len(skills) > 0 {
		for _, skill := range skills {
			_, found := employee.Person.Skills.Find(skill)
			if !found {
				employee.Person.Skills = append(employee.Person.Skills, skill)
			}
		}
	}
	
	if err := k.UpdateEmployeeinfo(ctx, employee); err != nil {
		return nil, err
	}
	
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteEmployeeInfo(ctx sdk.Context, keeper Keeper, msg MsgDeleteEmployeeInfo) (*sdk.Result, error) {
	employee, found := keeper.GetEmployee(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(ErrEmployee, fmt.Sprintf("employee %s not found", msg.Id))
	}
	
	if !employee.Person.Owner.Equals(msg.From) {
		return nil, sdkerrors.Wrap(ErrEmployee, fmt.Sprintf("employee unauthorised %s", msg.From))
	}
	
	if msg.Remove {
		if err := keeper.DeleteEmployeeInfo(ctx, msg.Id); err != nil {
			return nil, err
		}
	} else {
		employee.Status = StatusInactive
		if err := keeper.UpdateEmployeeinfo(ctx, employee); err != nil {
			return nil, err
		}
		
		updatedIDs := keeper.RemoveEmployeeIDFromActiveList(ctx, employee.ID)
		if err := keeper.UpdateActiveEmployeesIDsList(ctx, updatedIDs); err != nil {
			return nil, err
		}
		
		keeper.SetDeActiveEmployee(ctx, employee.ID)
	}
	
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgRestoreEmployeeInfo(ctx sdk.Context, keeper Keeper, msg MsgRestoreEmployeeInfo) (*sdk.Result, error) {
	employee, found := keeper.GetEmployee(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(ErrEmployee, fmt.Sprintf("employee %s not found", msg.Id))
	}
	
	if !employee.Person.Owner.Equals(msg.From) {
		return nil, sdkerrors.Wrap(ErrEmployee, fmt.Sprintf("employee unauthorised %s", msg.From))
	}
	
	if employee.Status != StatusInactive {
		return nil, sdkerrors.Wrap(ErrEmployee, fmt.Sprintf("invalid employee status "))
	}
	
	employee.Status = StatusActive
	if err := keeper.UpdateEmployeeinfo(ctx, employee); err != nil {
		return nil, err
	}
	
	updatedIDs := keeper.RemoveEmployeeIDFromDeActiveList(ctx, employee.ID)
	if err := keeper.UpdateDeActiveEmployeesIDsList(ctx, updatedIDs); err != nil {
		return nil, err
	}
	keeper.SetActiveEmployee(ctx, employee.ID)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
