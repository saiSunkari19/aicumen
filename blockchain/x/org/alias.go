package org

import (
	"github.com/saiSunkari19/aicumen/blockchain/x/org/internal/keeper"
	"github.com/saiSunkari19/aicumen/blockchain/x/org/internal/types"
)

const (
	ModuleName             = types.ModuleName
	StoreKey               = types.StoreKey
	RouterKey              = types.RouterKey
	QuerierRoute           = types.QuerierRoute
	StatusActive           = types.StatusActive
	StatusInactive         = types.StatusInactive
	QueryEmployeeByID      = keeper.QueryEmployeeByID
	QueryByDepartment      = keeper.QueryByDepartment
	QueryByName            = keeper.QueryByName
	QueryActiveEmployees   = keeper.QueryActiveEmployees
	QueryDeActiveEmployees = keeper.QueryDeActiveEmployees
)

type (
	MsgAddEmployee = types.MsgAddEmployee
	Person         = types.Person
	Employee       = types.Employee
	Employess      = types.Employess
	Skills         = types.Skills
	Status         = types.Status
	GenesisState   = types.GenesisState
	Keeper         = keeper.Keeper
)

var (
	NewKeeper         = keeper.NewKeeper
	NewMsgAddEmployee = types.NewMsgAddEmployee

	RegisterCodec = types.RegisterCodec
	ModuleCdc     = types.ModuleCdc

	NewQuerier = keeper.NewQuerier

	GetEmployeePrefixKey = types.GetEmployeePrefixKey
)
