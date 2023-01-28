package util

import (
	"math/rand"
	"strings"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

const alphabet="abcdefghijklmnopqrstuvwxyz"

func RandonInt(min int64,max int64)int64{
	return min+rand.Int63n(max-min+1)
}

func RandonString(n int)string {
	var sb strings.Builder
	k:=len(alphabet)

	for i:=0; i<n;i++{
		c:=alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName()string{
	return RandonString(8)
}

func RandomPhoneNumber()string{
	return string(RandonInt(1000000,2000000))
}

func RandomCurrency()string{
	currencies:=[]string{"EUR","USD","CAD"}
	n:=len(currencies)
	return currencies[rand.Intn((n))]
}