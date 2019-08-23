package orderbook

type Orderbook struct {
	Asks   []*Order
	Bids   []*Order
	Trades []*Trade 
}

func New() *Orderbook {
	Orderbook := &Orderbook{}
	Orderbook.Asks   = []*Order{}
	Orderbook.Bids   = []*Order{}
	Orderbook.Trades = []*Trade{}
	return Orderbook
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	if order.Side == SideBid {
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
	for i := 0; i < len(ob.Bids); i++ {
		bid = ob.Bids[i]
		if bid.Price < order.Price {
			continue
		}
	}
	return nil,nil
}

func (ob *Orderbook) HandleBid(order *Order) ([]*Trade, *Order) {
	if len(ob.Asks) == 0 {
		ob.Bids = append(ob.Asks, order)
		return nil, order
	}
	var ask *Order
	for i := 0; i < len(ob.Asks); i++ {
		ask = ob.Asks[i]
		if ask.Price > order.Price {
			continue
		}
	}
	return nil,nil
}

func (ob *Orderbook) AddBid(order *Order) {
	ob.Bids = append(ob.Bids,order)
}
