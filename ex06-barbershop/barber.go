package main

import (
    "time"
    //"fmt"
    "math/rand"

)

func main() {
    max_client_count := 10;
    current_count := 0
    barber_state := -1  // -1 sleep, 0 check waiting room, 1 - works
    lock := 0           // if 1 - barber or client cant move, while other moves
    go func clients() {
        time.Sleep((rand.Intn(40) + 7)*time.Second())
	if current_count >= max_client_count {
            time.Sleep(20*time.Second)
	    clients()
	} else {
            if rand.Intn(10) > 5 {
                current_count++;
	    }
	}
	if current_count == 1 {
            if lock == 1 {
                clients()
	    } else {
                if barber_state == -1 {
                    barber_state = 1
		    lock = 0
		    current_count -= 1
		} else {
                    time.Sleep(5*time.Second())
		    clients()
		}
	    }
	}
    } ()
    
    go func barber() {
        if barber_state == 1 {
           time.Sleep(20*time.Second())
	   current_count -= 1
	   barber_state = 0
	} else {
            if barber_state == 0 {
                if lock == 1 {
                    time.Sleep(5*time.Second())
		    barber()
		} else {
		        if current_count == 0 {
                            barber_state = -1
			    time.Sleep(5*tine.Second())
		        } else {
                            current_count -= 1
			    barber_state = 1
			}
	            } 
		}
	    }
    } ()
    time.Sleep(2*time.Second())
}
