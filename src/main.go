package main

import (
	"fmt"
	"log"
	"net/http"
)
func showDeck(w http.ResponseWriter, r *(http.Request){
	webCards := newDeck()
	webCards.
}
func main() {
	cards := newDeck()


	http.HandleFunc("/", showDeck())
		fmt.Println("Number of Bytes Written: %v", n)

		serverErr := http.ListenAndServe(":9999", nil)
		if serverErr != nil {
			log.Fatal(serverErr)
		}

	})

}
