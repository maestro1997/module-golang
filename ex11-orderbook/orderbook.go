package orderbook

type Orderbook struct {
	Asks []*Order
	Bids []*Order
}

func New() *Orderbook {
	Orderbook := &Orderbook{}
	Orderbook.Asks = []*Order{}
	Orderbook.Bids = []*Order{}
	return Orderbook
}

func (ob *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	var t1 []*Trade
	var t2 *Order
	if order.Side == 1 {
		t1, t2 = ob.MatchBid(order)
	} else {
		t1, t2 = ob.MatchAsk(order)
	}
	for i := 0; i < len(t1); i++ {
		for j := 0; j < len(t1)-1; j++ {
			if t1[j].Volume > t1[j+1].Volume {
				temp := t1[j]
				t1[j] = t1[j+1]
				t1[j+1] = temp
			}
		}
	}
	return t1, t2
}

func (ob *Orderbook) MatchBid(order *Order) ([]*Trade, *Order) {
	Trades := make([]*Trade, 0)
	for {
		index := ob.BestAsk(order)
		if index == -1 {
			if order.Price == 0 {
				return Trades, order
			}
			ob.Bids = append(ob.Bids, order)
			return Trades, nil
		}
		ask := ob.Asks[index]
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
		order.Volume = order.Volume - ask.Volume
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
		index := ob.BestBid(order)
		if index == -1 {
			if order.Price == 0 { // for market order
				return Trades, order
			}
			ob.Asks = append(ob.Asks, order)
			return Trades, nil
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

func (ob *Orderbook) BestBid(order *Order) int {
	if len(ob.Bids) == 0 {
		return -1
	}
	flag := 0
	index := 0
	bestPrice := ob.Bids[0].Price
	isLimit := order.Kind == 2
	for i := 0; i < len(ob.Bids); i++ {
		bid := ob.Bids[i]
		if isLimit && bid.Price < order.Price {
			continue
		}
		if bid.Price >= bestPrice {
			index = i
			flag = 1
			bestPrice = bid.Price
		}
	}
	if flag == 1 {
		return index
	}
	return -1
}

func (ob *Orderbook) BestAsk(order *Order) int {
	if len(ob.Asks) == 0 {
		return -1
	}
	flag := 0
	index := 0
	bestPrice := ob.Asks[0].Price
	isLimit := order.Kind == 2
	for i := 0; i < len(ob.Asks); i++ {
		ask := ob.Asks[i]
		if isLimit && ask.Price > order.Price {
			continue
		}
		if ask.Price <= bestPrice {
			index = i
			flag = 1
			bestPrice = ask.Price
		}
	}
	if flag == 1 {
		return index
	}
	return -1
}
