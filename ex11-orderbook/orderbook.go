package main

import "fmt"

type Orderbook struct {
	Asks []*Order
	Bids []*Order
}

type Trade struct {
	Bid    *Order
	Ask    *Order
	Volume uint64
	Price  uint64
}

type Order struct {
	ID   int
	Side int8
	Kind int8

	Volume uint64
	Price  uint64
}

//////////////////////////////////////////
func New() *Orderbook {
	Orderbook := &Orderbook{}
	Orderbook.Asks = []*Order{}
	Orderbook.Bids = []*Order{}
	return Orderbook
}

func (ob *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	if order.Side == 1 {
		return ob.MatchBid(order)
	}
	return ob.MatchAsk(order)
}

func (ob *Orderbook) MatchBid(order *Order) ([]*Trade, *Order) {
	Trades := make([]*Trade, 0)
	for {
		index := ob.BestOffer(order)
		if index == -1 {
			ob.Bids = append(ob.Bids, order)
			return Trades, order
		}
		fmt.Println("LOLLLL")
		ask := ob.Asks[index]
		fmt.Println("Index", index)
		if ask.Volume >= order.Volume {
			trade := Trade{order, ask, order.Volume, ask.Price}
			Trades = append(Trades, &trade)
			ask.Volume = ask.Volume - order.Volume
			if ask.Volume == 0 {
				if index == len(ob.Asks)-1 {
					ob.Asks = ob.Asks[:index]
				} else {
					ob.Asks = append(ob.Asks[:index], ob.Asks[index+1:]...)
				}
			}
			return Trades, nil
		}
		trade := Trade{order, ask, ask.Volume, ask.Price}
		Trades = append(Trades, &trade)
		order.Volume -= ask.Volume
		if index == len(ob.Asks)-1 {
			ob.Asks = ob.Asks[:index]
		} else {
			ob.Asks = append(ob.Asks[:index], ob.Asks[index+1:]...)
		}
	}
}

func (ob *Orderbook) MatchAsk(order *Order) ([]*Trade, *Order) {
	Trades := make([]*Trade, 0)
	for {
		index := ob.BestOffer(order)
		if index == -1 {
			ob.Asks = append(ob.Asks, order)
			return Trades, order
		}
		bid := ob.Bids[index]
		if bid.Volume >= order.Volume {
			trade := Trade{bid, order, order.Volume, bid.Price}
			Trades = append(Trades, &trade)
			bid.Volume = bid.Volume - order.Volume
			if bid.Volume == 0 {
				if index == len(ob.Bids)-1 {
					ob.Bids = ob.Bids[:index]
				} else {
					ob.Bids = append(ob.Bids[:index], ob.Bids[index+1:]...)
				}
			}
			return Trades, nil
		}
		trade := Trade{bid, order, bid.Volume, bid.Price}
		Trades = append(Trades, &trade)
		order.Volume = order.Volume - bid.Volume
		if index == len(ob.Bids)-1 {
			ob.Bids = ob.Bids[:index]
		} else {
			ob.Bids = append(ob.Bids[:index], ob.Bids[index+1:]...)
		}
	}
}

func (ob *Orderbook) BestOffer(order *Order) int {
	catalog := ob.Bids
	var sign int64
	sign = 1
	isLimit := order.Kind - 1
	if order.Side == 1 {
		catalog = ob.Asks
		sign = -1
	}
	if len(catalog) == 0 {
		return -1
	}
	index := 0
	bestPrice := catalog[0].Price
	flag := 0
	var cond bool
	for i := 1; i < len(catalog); i++ {
		offer := catalog[i]
		if sign == 1 {
			cond = offer.Price > order.Price
		} else {
			cond = offer.Price < order.Price
		}
		if (isLimit == 1) && !cond {
			continue
		}
		if sign == 1 {
			cond = ((bestPrice - offer.Price) > 0)
		} else {
			cond = ((bestPrice - offer.Price) < 0)
		}
		if cond {
			index = i
			bestPrice = offer.Price
			flag = 1
		}
	}
	if flag == 1 || (isLimit == 0) {
		return index
	} else {
		return -1
	}
}

func (ob *Orderbook) Print(text string) {
	fmt.Println("\nPrinting Orderbook", text)
	fmt.Println("ASKS :")
	var i int
	var or Order
	for i = 0; i < len(ob.Asks); i++ {
		or = *ob.Asks[i]
		fmt.Println(or.ID, or.Side, or.Kind, "Volume: ", or.Volume, "Price: ", or.Price)
	}
	fmt.Println("Bids :")
	for i = 0; i < len(ob.Bids); i++ {
		or = *ob.Bids[i]
		fmt.Println(or.ID, or.Side, or.Kind, "Volume: ", or.Volume, "Price: ", or.Price)
	}
}

func main_test() {
	ob := New()
	var or Order
	trades, rejects := ob.Match(&or)
	var i uint64
	for i = 0; i < 10; i++ {
		or = Order{int(i), (1 + int8(i%2)), 2, 5 * i, 20 * i}
		trades, rejects = ob.Match(&or)
	}

	if trades != nil && rejects != nil {
		fmt.Println(1)
	}
}

func main() {
	ob := New()
	or := Order{1, 2, 2, 5, 20}
	trades, rejects := ob.Match(&or)

	trades, rejects = ob.Match(&Order{2, 2, 2, 10, 10})
	trades, rejects = ob.Match(&Order{3, 2, 2, 15, 30})
	trades, rejects = ob.Match(&Order{4, 2, 2, 25, 25})
	trades, rejects = ob.Match(&Order{5, 2, 2, 35, 40})

	ob.Print("")

	trades, rejects = ob.Match(&Order{6, 1, 2, 5, 55})
	//trades, rejects = ob.Match(&Order{7, 2, 2, 1, 5})

	ob.Print("")

	if trades == nil && rejects == nil {
		fmt.Println("")
	}
}
