package persist

import "github.com/lunny/log"

func ItemServer() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Server: got item "+"#%d:%v", itemCount, item)
			itemCount++
		}
	}()
	return out
}