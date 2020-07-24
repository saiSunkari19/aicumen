package keeper

import (
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/saiSunkari19/aicumen/blockchain/types/store"
	"github.com/saiSunkari19/aicumen/blockchain/x/org/internal/types"
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
	store.Set(ctx, keeper.storeKey, keeper.cdc, types.GetGlobalEmployeeCountKey(), count)
	return nil
}

func (keeper Keeper) AddEmployeeInfo(ctx sdk.Context, employee types.Employee) error {
	err := store.SetNotExists(ctx, keeper.storeKey, keeper.cdc, types.GetEmployeeKey([]byte(employee.ID)), employee)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrEmployee, "unable to add employee %s id", employee.ID)
	}
	return nil
}

func (keeper Keeper) UpdateEmployeeinfo(ctx sdk.Context, employee types.Employee) error {
	err := store.SetExists(ctx, keeper.storeKey, keeper.cdc, types.GetEmployeeKey([]byte(employee.ID)), employee)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrEmployee, "unable to add employee %s id", employee.ID)
	}
	return nil
}

func (keeper Keeper) DeleteEmployeeInfo(ctx sdk.Context, id string) error {
	if err := store.Delete(ctx, keeper.storeKey, types.GetEmployeeKey([]byte(id))); err != nil {
		return err
	}
	
	prevIDs := keeper.GetActiveEmployeeList(ctx)
	if findElement(prevIDs, id) {
		activeIDs := keeper.RemoveEmployeeIDFromActiveList(ctx, id)
		if err := keeper.UpdateActiveEmployeesIDsList(ctx, activeIDs); err != nil {
			return err
		}
	} else {
		deActiveIDs := keeper.RemoveEmployeeIDFromDeActiveList(ctx, id)
		if err := keeper.UpdateDeActiveEmployeesIDsList(ctx, deActiveIDs); err != nil {
			return err
		}
		
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

func (keeper Keeper) UpdateActiveEmployeesIDsList(ctx sdk.Context, ids []string) error {
	if err := store.SetExists(ctx, keeper.storeKey, keeper.cdc, types.GetActivatedEmployeeKey(), ids); err != nil {
		return err
	}
	
	return nil
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

func (keeper Keeper) SetDeActiveEmployee(ctx sdk.Context, id string) {
	ids := keeper.GetDeActiveEmployeeList(ctx)
	ids = append(ids, id)
	
	store.Set(ctx, keeper.storeKey, keeper.cdc, types.GetDeActivatedEmployeeKey(), ids)
}

func (keeper Keeper) UpdateDeActiveEmployeesIDsList(ctx sdk.Context, ids []string) error {
	if err := store.SetExists(ctx, keeper.storeKey, keeper.cdc, types.GetDeActivatedEmployeeKey(), ids); err != nil {
		return err
	}
	
	return nil
}

func (keeper Keeper) GetDeActiveEmployeeList(ctx sdk.Context) []string {
	var ids []string
	
	db := ctx.KVStore(keeper.storeKey)
	bz := db.Get(types.GetDeActivatedEmployeeKey())
	if bz == nil {
		return []string{}
	}
	
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &ids)
	sort.Strings(ids)
	
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

func (keeper Keeper) RemoveEmployeeIDFromActiveList(ctx sdk.Context, eid string) []string {
	existedIDs := keeper.GetActiveEmployeeList(ctx)
	var index uint64
	
	for i, id := range existedIDs {
		if strings.EqualFold(id, eid) {
			index = uint64(i)
		}
	}
	
	if len(existedIDs) == 1 {
		return []string{}
	}
	return append(existedIDs[:index], existedIDs[index+1:]...)
}

func (keeper Keeper) RemoveEmployeeIDFromDeActiveList(ctx sdk.Context, eid string) []string {
	existedIDs := keeper.GetDeActiveEmployeeList(ctx)
	var index uint64
	
	for i, id := range existedIDs {
		if strings.EqualFold(id, eid) {
			index = uint64(i)
		}
	}
	
	if len(existedIDs) == 1 {
		return []string{}
	}
	
	return append(existedIDs[:index], existedIDs[index+1:]...)
}

func findElement(ids []string, elm string) bool {
	for _, id := range ids {
		if strings.EqualFold(id, elm) {
			return true
		}
	}
	return false
}
