package anaheim

import (
	"fmt"

	mobilesuite "github.com/m877778m/go-interface-sample/MobileSuite"
)

func PrintAnaheim() {
	fmt.Println("We are not 'Merchants of Death'.")
}

func Catalog() *[]AeMobileSuite {
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
	return &[]AeMobileSuite{
		{
			Spec:    inSpaceOk,
			code:    "Anaheim-ms-0001",
			price:   1000000000,
			forEfsf: true,
			forZeon: false,
		},
		{
			Spec:    inSpaceNg,
			code:    "Anaheim-ms-0002",
			price:   1000000000,
			forEfsf: true,
			forZeon: false,
		},
		{
			Spec:    inSpaceNg,
			code:    "Anaheim-ms-0003",
			price:   1000000000,
			forEfsf: true,
			forZeon: false,
		},
		{
			Spec:    inSpaceOk,
			code:    "Anaheim-ms-0004",
			price:   1000000000,
			forEfsf: true,
			forZeon: false,
		},
	}
}
