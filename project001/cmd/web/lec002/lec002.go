package lec002

import (
	"fmt"
	"math"
)

func LearnString() string {
	// string
	var kimcExplicit string = "code master1"
	var kimcImplicit = "code master2"
	var kimcDeclare string
	kimcShort := "code master4" // 함수 내에서만 가능
	return kimcExplicit + " " + kimcImplicit + " " + kimcDeclare + " " + kimcShort
}

func LearnNumber() int32 {
	// int
	var price int32 = 1000
	var price2 = 3000
	priceThree := 4000
	var pricefour uint32 = math.MaxUint32;
	var priceFive float64 = float64(pricefour) - 10000.0
	priceSix := 1.1

	return price + int32(price2) + int32(priceThree) + int32(priceFive + priceSix)
}

func Lec002() string {
	retStr := ""
	retStr += LearnString()
	retStr += fmt.Sprintf("%d", LearnNumber())
	return retStr
}