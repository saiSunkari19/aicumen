package org

import (
	"testing"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestAddEmployee(t *testing.T) {
	keeper, ctx := CreateTestInput(t)
	handler := NewHandler(keeper)
	
	employee := TestEmployee
	
	msg := NewMsgAddEmployee(employee.Person.Name, employee.Department, employee.Person.Address, employee.Person.Skills, TestAddress1)
	res, err := handler(ctx, msg)
	require.Nil(t, err)
	require.NotNil(t, res)
	
	require.Equal(t, 1, len(keeper.GetActiveEmployees(ctx)))
	
	employee.Department = "managment"
	employee.Person.Address = "US"
	msg = NewMsgAddEmployee(employee.Person.Name, employee.Department, employee.Person.Address, employee.Person.Skills, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	require.NotNil(t, res)
	
	employee.Department = "finance"
	msg = NewMsgAddEmployee(employee.Person.Name, employee.Department, employee.Person.Address, employee.Person.Skills, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	require.NotNil(t, res)
	
	require.Equal(t, 3, len(keeper.GetActiveEmployees(ctx)))
}

func TestUpdateEmployee(t *testing.T) {
	keeper, ctx := CreateTestInput(t)
	handler := NewHandler(keeper)
	var msg sdk.Msg
	
	employee := TestEmployee
	msg = NewMsgAddEmployee(employee.Person.Name, employee.Department, employee.Person.Address, employee.Person.Skills, TestAddress1)
	res, err := handler(ctx, msg)
	require.Nil(t, err)
	require.NotNil(t, res)
	
	employee.Department = "managment"
	employee.Person.Address = "US"
	msg = NewMsgAddEmployee(employee.Person.Name, employee.Department, employee.Person.Address, employee.Person.Skills, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	
	employee.Department = "finance"
	msg = NewMsgAddEmployee(employee.Person.Name, employee.Department, employee.Person.Address, employee.Person.Skills, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	
	require.Equal(t, 3, len(keeper.GetActiveEmployees(ctx)))
	
	employee, found := keeper.GetEmployee(ctx, employee.ID)
	require.True(t, found)
	employee.Person.Address = "India"
	
	msg = NewMsgUpdateEmployeeInfo("randomid", employee.Person.Address, "", []string{}, TestAddress1)
	res, err = handler(ctx, msg)
	require.NotNil(t, err)
	
	// updating address
	msg = NewMsgUpdateEmployeeInfo(employee.ID, employee.Person.Address, "", []string{}, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	
	employee, _ = keeper.GetEmployee(ctx, employee.ID)
	require.Equal(t, "India", employee.Person.Address)
	require.Equal(t, 2, len(employee.Person.Skills))
	
	// updating skills
	msg = NewMsgUpdateEmployeeInfo(employee.ID, "", "", Skills{"rust", "python"}, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	
	employee, _ = keeper.GetEmployee(ctx, employee.ID)
	require.Equal(t, 4, len(employee.Person.Skills))
	require.Equal(t, "managment", employee.Department)
	
	// update department
	employee.Department = "finance"
	msg = NewMsgUpdateEmployeeInfo(employee.ID, "", employee.Department, []string{}, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	
	employee, _ = keeper.GetEmployee(ctx, employee.ID)
	require.Equal(t, "finance", employee.Department)
	
	// Update InActive employee
	employee.Status = StatusInactive
	require.Nil(t, keeper.UpdateEmployeeinfo(ctx, employee))
	
	msg = NewMsgUpdateEmployeeInfo(employee.ID, "", "", Skills{"rust", "python"}, TestAddress1)
	res, err = handler(ctx, msg)
	require.NotNil(t, err)
	
}

func TestDeleteAndRestoreEmployee(t *testing.T) {
	keeper, ctx := CreateTestInput(t)
	handler := NewHandler(keeper)
	var msg sdk.Msg
	
	employee := TestEmployee
	msg = NewMsgAddEmployee(employee.Person.Name, employee.Department, employee.Person.Address, employee.Person.Skills, TestAddress1)
	res, err := handler(ctx, msg)
	require.Nil(t, err)
	require.NotNil(t, res)
	
	employee.Department = "managment"
	employee.Person.Address = "US"
	msg = NewMsgAddEmployee(employee.Person.Name, employee.Department, employee.Person.Address, employee.Person.Skills, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	
	employee.Department = "finance"
	msg = NewMsgAddEmployee(employee.Person.Name, employee.Department, employee.Person.Address, employee.Person.Skills, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	
	require.Equal(t, 3, len(keeper.GetActiveEmployees(ctx)))
	
	msg = NewMsgDeleteEmployeeInfo("randomid", true, TestAddress1)
	res, err = handler(ctx, msg)
	require.NotNil(t, err)
	
	// delete employee
	msg = NewMsgDeleteEmployeeInfo("employee0", false, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	
	id := "employee0"
	employee, _ = keeper.GetEmployee(ctx, id)
	require.True(t, employee.Status.Equal(StatusInactive))
	
	require.Equal(t, 3, len(keeper.GetEmployees(ctx)))
	require.Equal(t, 2, len(keeper.GetActiveEmployees(ctx)))
	require.Equal(t, 1, len(keeper.GetDeActiveEmployees(ctx)))
	
	// restore employee
	msg = NewMsgRestoreEmployeeInfo(id, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	
	msg = NewMsgRestoreEmployeeInfo("randomid", TestAddress1)
	res, err = handler(ctx, msg)
	require.NotNil(t, err)
	
	employee, _ = keeper.GetEmployee(ctx, id)
	require.True(t, employee.Status.Equal(StatusActive))
	
	require.Equal(t, 3, len(keeper.GetEmployees(ctx)))
	require.Equal(t, 3, len(keeper.GetActiveEmployees(ctx)))
	require.Equal(t, 0, len(keeper.GetDeActiveEmployees(ctx)))
	
	// delete permanently when employee in active state
	msg = NewMsgDeleteEmployeeInfo(id, true, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	
	require.Equal(t, 2, len(keeper.GetEmployees(ctx)))
	require.Equal(t, 2, len(keeper.GetActiveEmployees(ctx)))
	require.Equal(t, 0, len(keeper.GetDeActiveEmployees(ctx)))
	
	// delete permanently when employee in deactive state
	id = "employee1"
	msg = NewMsgDeleteEmployeeInfo(id, false, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	
	require.Equal(t, 2, len(keeper.GetEmployees(ctx)))
	require.Equal(t, 1, len(keeper.GetActiveEmployees(ctx)))
	require.Equal(t, 1, len(keeper.GetDeActiveEmployees(ctx)))
	
	msg = NewMsgDeleteEmployeeInfo(id, true, TestAddress1)
	res, err = handler(ctx, msg)
	require.Nil(t, err)
	
	require.Equal(t, 1, len(keeper.GetEmployees(ctx)))
	require.Equal(t, 1, len(keeper.GetActiveEmployees(ctx)))
	require.Equal(t, 0, len(keeper.GetDeActiveEmployees(ctx)))
	
}
