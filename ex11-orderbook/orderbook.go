package main

import "fmt"

type Orderbook struct {
	Asks   []*Order
	Bids   []*Order
	Trades []*Trade
}

type Trade struct {
	Bid    *Order
	Ask    *Order
	Volume uint64
	Price  uint64
}

type Side int8

const (
	SideBid Side = 1
	SideAsk Side = 2
)

func (side Side) String() string {
	switch side {
	case SideBid:
		return "BID"
	case SideAsk:
		return "ASK"
	}

	return "UNKNOWN"
}

type Kind int8

const (
	KindMarket Kind = 1
	KindLimit  Kind = 2
)

func (kind Kind) String() string {
	switch kind {
	case KindMarket:
		return "MARKET"
	case KindLimit:
		return "LIMIT"
	}

	return "UNKNOWN"
}

type Order struct {
	ID int

	Side Side
	Kind Kind

	Volume uint64
	Price  uint64
}
//////////////////////////////////////////
func New() *Orderbook {
	Orderbook := &Orderbook{}
	Orderbook.Asks   = []*Order{}
	Orderbook.Bids   = []*Order{}
	Orderbook.Trades = []*Trade{}
	return Orderbook
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	if order.Side == SideAsk {
		return orderbook.HandleAsk(order)
	}
	return orderbook.HandleBid(order)
}

func (ob *Orderbook) HandleAsk(order *Order) ([]*Trade, *Order) {
	if len(ob.Bids) == 0 {
		ob.Asks = append(ob.Asks, order)
		return nil, order
	}
	var bid *Order
	var bestPrice uint64
	bestPrice = 0
	bid_num := -1
	for i := 0; i < len(ob.Bids); i++ {
		bid = ob.Bids[i]
		if bid.Price < order.Price {
			continue
		}
		if bid.Price > bestPrice {
			bestPrice = bid.Price
			bid_num = i
		}
	}
	if bid_num == -1 {
		ob.Asks = append(ob.Asks, order)
		return nil,order
	}
	bid = ob.Bids[bid_num]
	delta := order.Volume - bid.Volume
	if delta <= 0 { // if we can satisfy order
		bid.Volume = -delta
		if bid.Volume == 0 {  // if no items left - delete order from orderbook
			ob.Bids[bid_num] = ob.Bids[len(ob.Bids)- 1]
			ob.Bids[bid_num] = nil
			ob.Bids = ob.Bids[:len(ob.Bids) - 1]
		}
		return []*Trade{&Trade{bid, order, order.Volume, bid.Price}}, nil
	}
	//ob.Bids[bid_num] = ob.Bids[len(ob.Bids)- 1]
	//ob.Bids[bid_num] = nil
	//ob.Bids = ob.Bids[:len(ob.Bids) - 1]
	trades, reject:= ob.HandleAsk(&Order{order.ID,order.Side,order.Kind, delta, order.Price})
	if trades == nil {
		return nil, reject
	}
	trades2, reject2 := ob.HandleAsk(reject)
	return append(trades, trades2...), reject2
}


func (ob *Orderbook) HandleBid(order *Order) ([]*Trade, *Order) {
	if len(ob.Asks) == 0 {
		ob.Bids = append(ob.Bids, order)
		return nil, order
	}
	var ask *Order
	var bestPrice uint64
	bestPrice = 1000
	ask_num := -1
	for i := 0; i < len(ob.Asks); i++ {
		ask = ob.Asks[i]
		if ask.Price > order.Price {
			continue
		}
		if ask.Price < bestPrice {
			bestPrice = ask.Price
			ask_num = i
		}
	}
	//fmt.Println("Debug point 1")
	if ask_num == -1 {
		ob.Bids = append(ob.Bids, order)
		return nil,order
	}
	//fmt.Println("Debug point 2")
	ask = ob.Asks[ask_num]
	if ask.Volume > order.Volume { // if we can satisfy order
		ask.Volume = ask.Volume - order.Volume
		//if ask.Volume == 0 {  // if no items left - delete order from orderbook
			//fmt.Println("Debug point 3")
		//	ob.Asks[ask_num] = ob.Asks[len(ob.Asks)- 1]
			//fmt.Println("Debug point 33")
		//	ob.Asks[ask_num] = nil
			//fmt.Println("Debug point 333")
		//	ob.Asks = ob.Asks[:len(ob.Asks) - 1]
		//}
		fmt.Println("Debug point iii4")
		return []*Trade{&Trade{ask, order, order.Volume, ask.Price}}, nil
	}
	//ob.Asks[ask_num] = ob.Asks[len(ob.Asks)- 1]
	//ob.Asks[ask_num] = nil
	//ob.Asks = ob.Asks[:len(ob.Asks) - 1]
	ask.Volume = 0
	trades, reject:= ob.HandleAsk(&Order{order.ID,order.Side,order.Kind, order.Volume - ask.Volume, order.Price})
	fmt.Println("Debug point 5")
	if trades == nil {
		return nil, reject
	}
	trades2, reject2 := ob.HandleAsk(reject)
	return append(trades, trades2...), reject2
}

func (ob *Orderbook) Print () {
	fmt.Println("\nPrinting Orderbook")
	fmt.Println("ASKS :")
	var i int;
	var or Order
	for i = 0; i < len(ob.Asks); i++ {
		or = *ob.Asks[i]
		fmt.Println(or.ID, or.Side, or.Kind, "Volume: ",or.Volume, "Price: ",or.Price)
	}
	fmt.Println("Bids :")
	for i = 0; i < len(ob.Bids); i++ {
		or = *ob.Bids[i]
		fmt.Println(or.ID, or.Side, or.Kind, "Volume: ",or.Volume, "Price: ",or.Price)
	}


}

func main_test() {
	ob := New()
	var or Order
	trades, rejects := ob.Match(&or)
	var i uint64
	for i = 0; i < 10; i++ {
		or = Order{int(i), Side(1 + i % 2) , 2, 5*i, 20 * i}
		trades, rejects = ob.Match(&or)
	}

	if trades != nil && rejects != nil {
		fmt.Println(1)
	}
}

func main() {
	ob := New()
	or := Order{1, 2 , 2, 5, 20}
	trades, rejects := ob.Match(&or )
	
	trades, rejects = ob.Match( &Order{2,2,2,5,120} )
	trades, rejects = ob.Match( &Order{3,2,2,15,220} )
	trades, rejects = ob.Match( &Order{4,2,2,25,320} )
	trades, rejects = ob.Match( &Order{5,2,2,35,420} )

	ob.Print()


	trades, rejects = ob.Match( &Order{6,1,2,5,140} )
	//trades, rejects = ob.Match( &Order{7,1,2,2,170} )

	ob.Print()


	if trades == nil && rejects == nil {
		fmt.Println("")
	}
}
