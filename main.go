package main

import (
	"github.com/m877778m/go-interface-sample/anaheim"
	"github.com/m877778m/go-interface-sample/zeon"
	"github.com/m877778m/go-interface-sample/zeonic"
)

func main() {

	// ジオニック社からモビルスーツ（MS）一覧をもらう
	// 気遣いのできないジオニックの営業担当は、カタログをそのまま送りつけてくる
	zCatalog := zeonic.Catalog()

	// モビルスーツリストをチェックし、対象をピックアップする
	pickup := []string{}

	// ジオンの新担当者はカタログのスペックを見ても理解できない
	// そのため、自分たちの定義したチェックポイントのみで判断を行う
	for _, zMs := range *zCatalog {
		// ジオン側は、`type ZeonMS interface`でMSの仕様を抽象化している
		// ジオニック社のカタログは、その`interface{}`を満たすため、
		// ジオンはジオニック社の詳細仕様が理解できなくても期待のMSをピックアップする事ができる
		// -> つまり、`package zeon`は`package anaheim`に依存する必要がない！！！
		if zeon.Check(&zMs) {
			pickup = append(pickup, zMs.GetCode())
		}
	}

	// ジオンは自分たちのフォーマットで提案依頼書を作成する
	rfp := &zeon.Rfp{
		Pickup:  pickup,
		Comment: "Sieg Zeon!",
	}

	// ジオニック側は提案依頼書を`type Rfp interface`で抽象化している
	// ジオンの提案依頼書は、その`interface{}`を満たすため、
	// ジオニックの営業担当は提案依頼書の細かい部分を確認せずとも、提案するMSを取り出すことができる
	// -> 先程と同様に、`package anaheim`は`package zeon`に依存する必要がない！！！
	AeSales := new(anaheim.Sales)
	AeSales.CheckRfp(rfp)

	// 次回に続く！！！
	zeon.PrintZeon()
}
