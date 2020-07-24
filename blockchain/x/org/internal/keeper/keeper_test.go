package keeper

import (
	"testing"
	
	"github.com/stretchr/testify/require"
)

func TestKeeper_EmployeeCycle(t *testing.T) {
	keeper, ctx := CreateTestInput(t)
	
	var employee = TestEmployee
	// Add employee details
	require.Nil(t, keeper.AddEmployeeInfo(ctx, employee))
	keeper.SetActiveEmployee(ctx, employee.ID)
	
	employee.ID = "employee2"
	require.Nil(t, keeper.AddEmployeeInfo(ctx, employee))
	keeper.SetActiveEmployee(ctx, employee.ID)
	
	employee.ID = "employee3"
	employee.Department = "network"
	require.Nil(t, keeper.AddEmployeeInfo(ctx, employee))
	keeper.SetActiveEmployee(ctx, employee.ID)
	
	// get employees details
	require.Equal(t, 3, len(keeper.GetEmployees(ctx)))
	
	// update employee info
	res, found := keeper.GetEmployee(ctx, "employee3")
	require.True(t, found)
	require.NotNil(t, res)
	require.Equal(t, "network", res.Department)
	
	res.Department = "management"
	require.Nil(t, keeper.UpdateEmployeeinfo(ctx, res))
	res, found = keeper.GetEmployee(ctx, "employee3")
	require.True(t, found)
	require.Equal(t, "management", res.Department)
	
	// delete employee
	res, found = keeper.GetEmployee(ctx, "employee2")
	require.True(t, found)
	require.Nil(t, keeper.DeleteEmployeeInfo(ctx, res.ID))
	
	// get employees details
	require.Equal(t, 2, len(keeper.GetEmployees(ctx)))
	
}
