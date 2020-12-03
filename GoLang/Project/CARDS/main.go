package main

func main() {

	//greeting := "Hi There!"
	//fmt.Print([]byte(greeting))
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile("my_cards")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

func newCard() {
	//return "Five of Diamonds"

	//cards := []string{}
	//cards := deck{"Three of Diamond", newCard()}
	//cards = append(cards, "Six of Spades")
	//cards = append(cards, "Ace of Diamond")
	//card := newCard()
	//fmt.Println(card)
	/*for i, card := range cards {
		fmt.Println(i, card)
	}*/
	//cards.print()

	//cards := newDeck()
	/*cards.print()
	newCards := cards[:3]
	fmt.Println(newCards)*/
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
}
