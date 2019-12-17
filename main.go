package main

import (
	"auctionhouse"
	"fmt"
)

func main() {
	// token, check := auctionauth.GetNewToken()
	// if check {
	// 	fmt.Println("error in GetNewToken()")
	// }
	// address, check := auctionhouse.GetRealmAddress("us", "tichondrius", token.Token)
	// if check {
	// 	fmt.Println("error in auctionhouse.GetRealmAddress()")
	// }
	// fmt.Println(address)
	// out, check := auctionhouse.CallRealmAPI(address)
	// if !check {
	// 	fmt.Println(out)
	// }
	// fmt.Println(auctionhouse.GetRealms())
	d, check := auctionhouse.NewDaemon("us", "en_US")
	if !check {
		fmt.Println(d.Token.Token)
	}
	out, check := d.GetRealms()
	if !check {
		fmt.Println("Bad thangs popping")
	}
	fmt.Println(out)
}
