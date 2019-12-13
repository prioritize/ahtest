package main

import (
	"auctionauth"
	"auctionhouse"
	"fmt"
)

func main() {
	token, check := auctionauth.GetNewToken()
	if check {
		fmt.Println("error in GetNewToken()")
	}
	address, check := auctionhouse.GetRealmAddress("us", "tichondrius", token.Token)
	if check {
		fmt.Println("error in auctionhouse.GetRealmAddress()")
	}
	fmt.Println(address)
	out, check := auctionhouse.CallRealmAPI(address)
	if !check {
		fmt.Println(out)
	}
	fmt.Println(auctionhouse.GetRealms())
}
