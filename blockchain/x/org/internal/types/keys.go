package types

import "strconv"

const (
	ModuleName   = "org"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	
	EmployeePrefix = "employee"
)

var (
	EmployeeKey = []byte{0x01}
	
	ActivatedPrefixKey   = []byte{0x02}
	DeActivatedPrefixKey = []byte{0x03}
	
	GlobalEmployeeNumberKey = []byte("globalEmployeeNumber")
)

func GetGlobalEmployeeCountKey() []byte {
	return GlobalEmployeeNumberKey
}

func GetEmployeeKey(id []byte) []byte {
	return append(EmployeeKey, id...)
}

func GetDeActivatedEmployeeKey() []byte {
	return DeActivatedPrefixKey
}

func GetActivatedEmployeeKey() []byte {
	return ActivatedPrefixKey
}

func GetEmployeePrefixKey(id uint64) string {
	return EmployeePrefix + strconv.Itoa(int(id))
}
