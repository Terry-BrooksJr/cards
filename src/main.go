package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("**********Full deck**********")
	cards.print()
	fmt.Println("********************")

	hand1, hand2 := deal(cards, 14)

	fmt.Println("**********Player 1 Hand*********")
	hand1.print()
	fmt.Println("********************")
	fmt.Println("**********Player 1 Hand**********")
	hand2.print()
	fmt.Printf("********************")

	hand2.saveToFile("hand2.cards")
	hand1.saveToFile("hand1.cards")

}
