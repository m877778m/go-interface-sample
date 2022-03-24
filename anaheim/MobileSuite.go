package anaheim

import (
	mobilesuite "github.com/m877778m/go-interface-sample/MobileSuite"
)

type AeMobileSuite struct {
	Spec    *mobilesuite.Spec
	code    string
	price   int
	forEfsf bool
	forZeon bool
}

// 宙域利用が可能かを判定
func (aems *AeMobileSuite) AvailableInSpace() bool {
	return aems.Spec.AvailavleInSpace()
}

// 海中利用が可能かを判定
func (aems *AeMobileSuite) AvailableInSea() bool {
	return aems.Spec.AvailavleInSea()
}

// ビーム武器の保有を判定
// 判定対象はメイン武器のみ、また遠距離・近接の判定は実施しない
func (aems *AeMobileSuite) WithBeamWepon() bool {
	if aems.Spec.MainWepon.Output != "beam" {
		return false
	}
	return true
}

// 遠隔ビーム武器の保有を判定
// 判定対象はメイン武器のみ
func (aems *AeMobileSuite) WithBeamGun() bool {
	if !aems.WithBeamWepon() {
		return false
	}
	if aems.Spec.MainWepon.Type != "gun" {
		return false
	}
	return true
}

// MSコード（識別子）の取得
func (aems *AeMobileSuite) GetCode() string {
	return aems.code
}
