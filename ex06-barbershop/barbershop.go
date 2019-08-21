package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup // Amount of potentional customers

type Barber struct {
	name string
	sync.Mutex
	state    int    // 0 sleeping, 1 cutting, 2 checking 
	customer *Customer
}

func (b *Barber) String() string {
    if b.state == 0 {
         return "sleeping"
    }
    if b.state == 1 {
         return "cutting"
    }
    return "checking"
}

type Customer struct {
	num int
}

func (c *Customer) String() string {
	return fmt.Sprintf("Client N%d",c.num) 
}

func NewBarber() (b *Barber) {
	return &Barber{
		name:  "Sam",
		state: 0,
	}
}

// Barber goroutine
// Checks for customers
// Sleeps - wait for wakers to wake him up
func barber(b *Barber, wr chan *Customer, wakers chan *Customer) {
	for {
		b.Lock()
		defer b.Unlock()
		b.state = 2
		b.customer = nil

		// checking the waiting room
		fmt.Printf("Checking waiting room: %d clients waiting\n", len(wr))
		time.Sleep(time.Millisecond * 100)
		select {
		case c := <-wr:
			fmt.Printf("Barber takes %s\n",c)
			HairCut(c, b)
			b.Unlock()
		default: // Waiting room is empty
			fmt.Printf("Barber going to sleep\n")
			b.state = 0
			b.customer = nil
			b.Unlock()
			c := <-wakers
			b.Lock()
			fmt.Printf("Barber woken by %s\n", c)
			HairCut(c, b)
			b.Unlock()
		}
	}
}

func HairCut(c *Customer, b *Barber) {
	b.state = 1
	b.customer = c
	b.Unlock()
	fmt.Printf("Cutting hair of %s\n", c)
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("Barber finished serving of %s\n",c)
	b.Lock()
	wg.Done()
	b.customer = nil
}

// customer goroutine
// just fizzles out if it's full, otherwise the customer
// is passed along to the channel handling it's haircut etc
func run_customer(c *Customer, b *Barber, wr chan<- *Customer, wakers chan<- *Customer,mode) {
	// arrive
	time.Sleep(time.Millisecond * 50)
	s:= "came"
	if mode == 1 {
             s = "returned"
        }     
	fmt.Printf("%s %s into barbershop\n",c,s)
	flag := 0
	if len(wr) > 0 {
             flag = 1
	}
	// Check on barber
	b.Lock()
        if flag == 0 {
	    fmt.Printf("%s checks, barber %s | in room: %d\n",c, b , len(wr))
        }
	switch b.state {
	case 0:
		select {
		case wakers <- c:	
		//default:
		//	select {
		//	case wr <- c:
		//	     fmt.Println("%s seats in chair\n",c)
		//	default:
		//		wg.Done()
		//	}
		}
	case 1:
		select {
		case wr <- c:
		    fmt.Println("%s seats in chair\n",c)
		default: // Full waiting room, leave shop
		    time.Sleep(10*time.Second)
		    run_customer(c,b,we,wakers,1)
		    //wg.Done()
		}
	case 2:
		panic("Customer shouldn't check for the Barber when Barber is Checking the waiting room")
	}
	b.Unlock()
}

func main() {
	b := NewBarber()
	b.name = "Rocky"
	room_size := 5  // 5 chairs in waiting room
	WaitingRoom := make(chan *Customer, room_size)
	Wakers := make(chan *Customer, 1)      // barber can cut only 1 person at time
	go barber(b, WaitingRoom, Wakers)

	time.Sleep(time.Millisecond * 100)
	wg = new(sync.WaitGroup)
	n := 10
	wg.Add(10)
	// Spawn customers
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond * 50)
		c := new(Customer)
		c.num = i
		go run_customer(c, b, WaitingRoom, Wakers,0)
	}

	wg.Wait()
	fmt.Println("No more customers for the day")
}
