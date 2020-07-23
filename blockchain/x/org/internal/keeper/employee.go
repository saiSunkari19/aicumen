package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/saiSunkari19/aicumen/blockchain/types/store"
	"github.com/saiSunkari19/aicumen/blockchain/x/org/internal/types"
	"sort"
)

func (keeper Keeper) GetGlobalEmployeeCount(ctx sdk.Context) uint64 {

	db := ctx.KVStore(keeper.storeKey)
	key := types.GetGlobalEmployeeCountKey()
	bz := db.Get(key)
	if bz == nil {
		return 0
	}

	var count uint64
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &count)
	return count
}

func (keeper Keeper) SetGlobalEmployeeCount(ctx sdk.Context, count uint64) error {
	err := store.SetExists(ctx, keeper.storeKey, keeper.cdc, types.GetGlobalEmployeeCountKey(), count)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrEmployee, "unable to add employee count %s id", count)
	}

	return nil
}

func (keeper Keeper) AddEmployee(ctx sdk.Context, employee types.Employee) error {
	err := store.SetNotExists(ctx, keeper.storeKey, keeper.cdc, types.GetEmployeeKey([]byte(employee.ID)), employee)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrEmployee, "unable to add employee %s id", employee.ID)
	}
	return nil
}

func (keeper Keeper) GetEmployee(ctx sdk.Context, id string) (info types.Employee, found bool) {
	if err := store.Get(ctx, keeper.storeKey, keeper.cdc, types.GetEmployeeKey([]byte(id)), &info); err != nil {
		return types.Employee{}, false
	}

	return info, true
}

func (keeper Keeper) iterateEmployees(ctx sdk.Context, cb func(employees types.Employee) (stop bool)) {

	iterator := store.Iterator(ctx, keeper.storeKey, types.EmployeeKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var employee types.Employee
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &employee)

		if cb(employee) {
			break
		}
	}
}

func (keeper Keeper) GetEmployees(ctx sdk.Context) (employees types.Employess) {
	keeper.iterateEmployees(ctx, func(employee types.Employee) (stop bool) {
		employees = append(employees, employee)
		return false
	})

	return employees.Sort()
}

// -------------------------------

func (keeper Keeper) SetActiveEmployee(ctx sdk.Context, id string) {
	ids := keeper.GetActiveEmployeeList(ctx)
	ids = append(ids, id)

	store.Set(ctx, keeper.storeKey, keeper.cdc, types.GetActivatedEmployeeKey(), ids)
}

func (keeper Keeper) GetActiveEmployeeList(ctx sdk.Context) []string {
	var ids []string

	db := ctx.KVStore(keeper.storeKey)
	bz := db.Get(types.GetActivatedEmployeeKey())
	if bz == nil {
		return []string{}
	}

	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &ids)

	sort.Strings(ids)
	return ids
}

func (keeper Keeper) GetActiveEmployees(ctx sdk.Context) (employees types.Employess) {

	ids := keeper.GetActiveEmployeeList(ctx)
	if len(ids) == 0 {
		return types.Employess{}
	}

	for _, id := range ids {
		employee, _ := keeper.GetEmployee(ctx, id)
		employees = append(employees, employee)
	}

	return employees.Sort()
}

// -------------------------------

func (keeper Keeper) SetDectiveEmployee(ctx sdk.Context, id string) {
	ids := keeper.GetDeActiveEmployeeList(ctx)
	ids = append(ids, id)

	store.Set(ctx, keeper.storeKey, keeper.cdc, types.GetDeActivatedEmployeeKey(), ids)
}

func (keeper Keeper) GetDeActiveEmployeeList(ctx sdk.Context) []string {
	var ids []string

	db := ctx.KVStore(keeper.storeKey)
	bz := db.Get(types.GetDeActivatedEmployeeKey())
	if bz == nil {
		return []string{}
	}

	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &ids)
	return ids
}

func (keeper Keeper) GetDeActiveEmployees(ctx sdk.Context) (employees types.Employess) {

	ids := keeper.GetDeActiveEmployeeList(ctx)
	if len(ids) == 0 {
		return types.Employess{}
	}

	for _, id := range ids {
		employee, _ := keeper.GetEmployee(ctx, id)
		employees = append(employees, employee)
	}

	return employees.Sort()
}
