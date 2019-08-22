package orderbook

type Orderbook struct {
	Asks   []*Order
	Bids   []*Order
	Trades []*Trade 
}

func New() *Orderbook {
	Orderbook = &Orderbook{}
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

func (ob *Orderbook) HandleAsk(order *Order) {
	var bid Order
	for i := 0; i < len(ob.Bids); i++ {
		bid = ob.Bids[i]
		if bid.Price < order.Price {
			continue
		}
		
	}
	return nil,nil
}


func (ob *Orderbook) HandleBid(order *Order) {
	var ask Order
	for i := 0; i < len(ob.Asks); i++ {
		bid = ob.Asks[i]
		if bid.Price > order.Price {
			continue
		}
	}
	return nil,nil
}


