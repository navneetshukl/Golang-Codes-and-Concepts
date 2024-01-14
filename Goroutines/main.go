package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func getStatus(url string) {
	status, err := http.Get(url)

	if err != nil {
		log.Println(fmt.Sprintf("Error in getting the status code of url %s ", url))
	} else{
		log.Println(fmt.Sprintf("Status code of %s is %d ",url,status.StatusCode))
	}
}
func main() {
	defer wg.Done()
	urls:=[]string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.twitter.com",
		"https://www.amazon.com",
		"https://www.reddit.com",


	}
	for _,val:=range urls{
		wg.Add(1)
		go getStatus(val)
	}
	wg.Wait()

}
