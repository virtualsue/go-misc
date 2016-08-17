 package main

 import (
	"io/ioutil"
	"log"
 	"net/http"
 )

 func main() {
	 resp, err := http.Get("http://www.radiotimes.com/news/2016-07-27/radio-times-tv-champion-2016-battle-3-aidan-turner-v-mary-beard")
	 if err != nil {
	 	log.Println(err)
	 }
	 defer resp.Body.Close()
	 body, err := ioutil.ReadAll(resp.Body)
	 log.Println(body)
 }
