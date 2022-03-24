package zeonic

import (
	"fmt"

	mobilesuite "github.com/m877778m/go-interface-sample/MobileSuite"
)

func PrintAnaheim() {
	fmt.Println("We are not 'Merchants of Death'.")
}

func Catalog() *[]zeonicMobileSuite {
	inSpaceOk := &mobilesuite.Spec{
		MainWepon: &mobilesuite.Wepon{
			Type:   "gun",
			Output: "beam",
		},
		InSpace: true,
	}
	inSpaceNg := &mobilesuite.Spec{
		InSpace: false,
	}
	return &[]zeonicMobileSuite{
		{
			Spec:    inSpaceOk,
			code:    "zeonic-ms-0001",
			price:   1000000000,
			forEfsf: true,
			forZeon: false,
		},
		{
			Spec:    inSpaceNg,
			code:    "zeonic-ms-0002",
			price:   1000000000,
			forEfsf: true,
			forZeon: false,
		},
		{
			Spec:    inSpaceNg,
			code:    "zeonic-ms-0003",
			price:   1000000000,
			forEfsf: true,
			forZeon: false,
		},
		{
			Spec:    inSpaceOk,
			code:    "zeonic-ms-0004",
			price:   1000000000,
			forEfsf: true,
			forZeon: false,
		},
	}
}
