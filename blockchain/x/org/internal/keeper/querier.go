package keeper

import (
	"fmt"
	"strings"
	
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	
	"github.com/saiSunkari19/aicumen/blockchain/x/org/internal/types"
)

const (
	QueryEmployeeByID      = "id"
	QueryActiveEmployees   = "active_employees"
	QueryDeActiveEmployees = "deactive_employees"
	QueryByDepartment      = "dept"
	QueryByName            = "name"
	QueryBySearch          = "search"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abcitypes.RequestQuery) (bytes []byte, err error) {
		switch path[0] {
		
		case QueryEmployeeByID:
			return queryEmployeeById(ctx, path[1:], k)
		case QueryActiveEmployees:
			return queryActiveEmployees(ctx, k)
		case QueryDeActiveEmployees:
			return queryDeActiveEmployees(ctx, k)
		case QueryByDepartment:
			return queryByDepartment(ctx, path[1:], k)
		case QueryByName:
			return queryByName(ctx, path[1:], k)
		case QueryBySearch:
			return queryBySearch(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
			
		}
	}
}

func queryEmployeeById(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	employee, found := k.GetEmployee(ctx, path[0])
	if !found {
		return nil, sdkerrors.Wrap(types.ErrEmployee, fmt.Sprintf("employee details not found %s ", path[0]))
	}
	
	res, err := codec.MarshalJSONIndent(k.cdc, employee)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	
	return res, nil
}

func queryActiveEmployees(ctx sdk.Context, k Keeper) ([]byte, error) {
	employees := k.GetActiveEmployees(ctx)
	if len(employees) == 0 {
		return nil, nil
	}
	
	res, err := codec.MarshalJSONIndent(k.cdc, employees)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	
	return res, nil
}

func queryDeActiveEmployees(ctx sdk.Context, k Keeper) ([]byte, error) {
	employees := k.GetDeActiveEmployees(ctx)
	if len(employees) == 0 {
		return nil, nil
	}
	
	res, err := codec.MarshalJSONIndent(k.cdc, employees)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	
	return res, nil
}

func queryByDepartment(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	employees := k.GetActiveEmployees(ctx)
	if len(employees) == 0 {
		return nil, nil
	}
	var byDept types.Employess
	
	for _, employee := range employees {
		if strings.EqualFold(employee.Department, path[0]) {
			byDept = append(byDept, employee)
		}
	}
	
	res, err := codec.MarshalJSONIndent(k.cdc, byDept)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	
	return res, nil
}

func queryByName(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	employees := k.GetActiveEmployees(ctx)
	if len(employees) == 0 {
		return nil, nil
	}
	var byName types.Employess
	
	name := path[0]
	for _, employee := range employees {
		if strings.EqualFold(employee.Person.Name, name) {
			byName = append(byName, employee)
		}
	}
	
	res, err := codec.MarshalJSONIndent(k.cdc, byName)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	
	return res, nil
}

func queryBySearch(ctx sdk.Context, path []string, keeper Keeper) ([]byte, error) {
	employeees := keeper.GetEmployees(ctx)
	var updatedEmployees types.Employess
	
	switch path[0] {
	case "dept":
		for _, employee := range employeees {
			if strings.EqualFold(employee.Department, path[1]) {
				updatedEmployees = append(updatedEmployees, employee)
			}
		}
		return marshalBinary(keeper.cdc, updatedEmployees)
	case "name":
		for _, employee := range employeees {
			if strings.EqualFold(employee.Person.Name, path[1]) {
				updatedEmployees = append(updatedEmployees, employee)
			}
		}
		return marshalBinary(keeper.cdc, updatedEmployees)
	case "addr":
		for _, employee := range employeees {
			if strings.EqualFold(employee.Person.Address, path[1]) {
				updatedEmployees = append(updatedEmployees, employee)
			}
		}
		return marshalBinary(keeper.cdc, updatedEmployees)
	case "skills":
		skills := strings.Split(path[1], ",")
		if len(skills) == 0 {
			return nil, nil
		}
		
		m := make(map[string]bool)
		for _, skill := range skills {
			m[skill] = true
		}
		
		for _, employee := range employeees {
			c := 0
			for _, s := range employee.Person.Skills {
				if _, ok := m[s]; ok {
					c += 1
				}
			}
			if c >= len(skills) {
				updatedEmployees = append(updatedEmployees, employee)
			}
		}
		return marshalBinary(keeper.cdc, updatedEmployees)
	}
	
	return nil, nil
}

func marshalBinary(cdc *codec.Codec, employees types.Employess) ([]byte, error) {
	res, err := codec.MarshalJSONIndent(cdc, employees)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	
	return res, nil
	
}
