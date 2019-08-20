package main

import (
   "fmt"
   "sync"
   "time"
)

var wg sync.WaitGroup
var cutting_time = 5 // time of hair cutting
var walk_time = 3    // time of walk of client that can't seat on chair
var check_time = 1   // time for what barber checks waiting room

type Barber struct {
    client chan Client
    waiters chan Client
    state int   // 0 -sleeps, 1 - cuts, 2 - checks room
    mutex int   // 1 if client checks barber
}

type Client struct {
    id int
}

func CreateClient(id int) Client {  
    return Client{id}
}

func CreateBarber (room_size int) Barber {
    barber := Barber{make (chan Client,1), make (chan Client,room_size),0,0}
    return barber
}

func (b Barber) run_barber () { 
    for {
         select {
	 case c:= <-b.client:
	     b.haircut(c)
	     b.check_room()
	 default:    
	     select  {		 
             case c:= <-b.waiters:
       		 b.client <- c
		 fmt.Printf("Barber takes client %d\n",c.id)
	     default:
		 fmt.Printf("Barber is going to sleep\n")
		 b.state = 0
	     }
         }
    }
}

func (b Barber) check_room() {
    b.state = 2
    time.Sleep(time.Second)
}

func (b Barber) haircut (client Client) {
   b.state = 1	
   fmt.Printf("Barber serves client %d\n",client.id)
   time.Sleep(time.Second)
   <-b.client
   fmt.Printf("Barber cuts client client %d\n",client.id)
   wg.Done()
}

func (c Client) run_client (b *Barber, flag int) {
    var way string
    if flag == 0 {
        way = "came"
    } else {
        way = "return"
    }	
    fmt.Printf("Client %d %s to barbershop\n",c.id,way)
    for { 
         select {
	 case <-b.client : 
	     select {
	     case  b.waiters <- c:
	         fmt.Printf("Client %d seats in chair\n",c.id)
	     default:
	         fmt.Printf("Client %d went for walk from barbershop\n",c.id)
		 time.Sleep(time.Second)
		 c.run_client(b,1)
	     }
         default:
	     b.client <- c
	     fmt.Printf("Client %d woke up barber\n",c.id)
        
     }
    }
}

func main() {
    room_size := 1
    client_count:= 1
    barber := CreateBarber(room_size)
    go barber.run_barber()
    time.Sleep(time.Second)
    wg.Add(client_count)
    for i:= 0; i < client_count; i++ {
        go CreateClient(i).run_client(&barber,0)
    }
    wg.Wait()
 }
