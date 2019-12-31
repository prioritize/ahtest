package main

import (
	ah "auctionhouse"
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

	// requestChannel := make(chan auctionhouse.AuctionHandler, 300)
	// dbInfo, err := auctionhouse.GetDBInfo()
	// if !err {
	// 	return
	// }
	// d, check := auctionhouse.NewDaemon("us", "en_US")
	dumb := make(chan int, 0)
	d, ok := ah.NewDaemon("us", "en_US")
	if !ok {
		fmt.Println("Daemon was not generated")
		return
	}
	realms, ok := d.GetRealms()
	if !ok {
		fmt.Println("Realms were not generated")
		return
	}
	// if !check {
	// 	fmt.Println(d.Token.Token)
	// }
	// _, check = d.GetRealms()
	// if !check {
	// 	fmt.Println("Bad thangs popping")
	// }
	// ah := make([]auctionhouse.AuctionHandler, 0)
	// for i := range d.Realms {
	// 	ah = append(ah, auctionhouse.NewAuctionHandler(d.Token, d.Realms[i], dbInfo))
	// }
	// server := ah[0]
	// go server.SendAuctionToDB()
	// server.RequestAuctionData()
	// for _, v := range ah {
	// 	requestChannel <- v
	// }

	t := ah.NewToken()
	fmt.Println(t)
	t.ValidateToken()
	itemMan := ah.NewItemManager()
	a := ah.ServerHandler(realms, t.Token(), itemMan.DBInfo, &itemMan)
	for _, v := range a {
		v.RequestAuctionData()
	}
	// item := itemMan.NewItem(169299, t.Token())
	// fmt.Println(item.Icon.Asset[0].HREF)
	itemMan.CheckItem(152511, t.Token())
	// // item, check := itemMan.QueryItemInformation(169299, t)
	// if !check {
	// 	fmt.Println("Item not found")
	// } else {
	// 	fmt.Println(item)
	// }

	<-dumb
}
