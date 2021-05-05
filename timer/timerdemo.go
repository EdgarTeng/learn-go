package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case <-ticker.C:
			log.Println("line: ", time.Now().Unix())
		default:
			//log.Println("continue: ", time.Now().Unix())
			continue
		}
	}
	//
	//for now := range time.Tick(time.Second) {
	//	log.Println("", now)
	//}

	//ticker := time.NewTicker(time.Second)
	//defer ticker.Stop()
	//done := make(chan bool)
	//go func() {
	//	time.Sleep(100 * time.Second)
	//	done <- true
	//}()
	//for {
	//	select {
	//	case <-done:
	//		fmt.Println("Done!")
	//		return
	//	case t := <-ticker.C:
	//		fmt.Println("Current time: ", t)
	//	}
	//}
}

var (
	shutdown     = make(chan interface{})
	shutdownBool = false
)

func isShutdown() bool {
	select {
	case <-shutdown:
		return true
	default:
	}
	return false
}

func boolShutdown() bool {
	select {
	case <-shutdown:
		return true
	default:
	}
	return false
}
