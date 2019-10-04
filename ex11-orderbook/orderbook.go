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
		index := ob.BestMatch(order)
		if index == -1 {
			if order.Price == 0 {
				return Trades, order
			}
			ob.Bids = append(ob.Bids, order)
			return Trades, nil
		}
		ord := ob.Asks[index]
		if ord.Volume >= order.Volume {
			trade := Trade{order, ord, order.Volume, ord.Price}
			Trades = append(Trades, &trade)
			ord.Volume = ord.Volume - order.Volume
			if ord.Volume == 0 {
				if index == len(ob.Asks)-1 {
					ob.Asks = ob.Asks[:index]
				} else {
					ob.Asks = append(ob.Asks[:index], ob.Asks[index+1:]...)
				}
			}
			return Trades, nil
		}
		trade := Trade{order, ord, ord.Volume, ord.Price}
		Trades = append(Trades, &trade)
		order.Volume = order.Volume - ord.Volume
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
		index := ob.BestMatch(order)
		if index == -1 {
			if order.Price == 0 { // for market order
				return Trades, order
			}
			ob.Asks = append(ob.Asks, order)
			return Trades, nil
		}
		ord := ob.Bids[index]
		if ord.Volume >= order.Volume {
			trade := Trade{ord, order, order.Volume, ord.Price}
			Trades = append(Trades, &trade)
			ord.Volume = ord.Volume - order.Volume
			if ord.Volume == 0 {
				if index == len(ob.Bids)-1 {
					ob.Bids = ob.Bids[:index]
				} else {
					ob.Bids = append(ob.Bids[:index], ob.Bids[index+1:]...)
				}
			}
			return Trades, nil
		}
		trade := Trade{ord, order, ord.Volume, ord.Price}
		Trades = append(Trades, &trade)
		order.Volume = order.Volume - ord.Volume
		if index == len(ob.Bids)-1 {
			ob.Bids = ob.Bids[:index]
		} else {
			ob.Bids = append(ob.Bids[:index], ob.Bids[index+1:]...)
		}
	}
}

func (ob *Orderbook) BestMatch(order *Order) int {
	var catalog []*Order
	if order.Side == 2 {
		catalog = ob.Bids
	} else {
		catalog = ob.Asks
	}
	if len(catalog) == 0 {
		return -1
	}
	flag := 0
	index := 0
	bestPrice := catalog[0].Price
	isLimit := order.Kind == 2
	for i := 0; i < len(catalog); i++ {
		ord := catalog[i]
		if isLimit && (((order.Side == 2) && comp2(order.Price, ord.Price)) || ((order.Side == 1) && comp2(ord.Price, order.Price))) {
			continue
		}
		if ((order.Side == 1) && compPrice(bestPrice, ord.Price) == 1) || ((order.Side == 2) && compPrice(ord.Price, bestPrice) == 1) {
			index = i
			flag = 1
			bestPrice = ord.Price
		}
	}
	if flag == 1 {
		return index
	}
	return -1
}

func compPrice(p1, p2 uint64) int {
	if p1 >= p2 {
		return 1
	}
	return -1
}

func comp2(p1, p2 uint64) bool {
	return (p1 > p2)
}
