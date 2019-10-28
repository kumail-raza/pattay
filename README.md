# Pattay ![alt text](https://image.flaticon.com/icons/png/128/1801/1801118.png)
Integrable deck of cards, and 4 player ring 

## Ring
```go
//Ring ring of players in a game of cards
type Ring interface {
	
        // Players returns
	Players() []RingPlayer
	
        //Next returns player to play next
	Next() (RingPlayer, error)

	//SetCurrentPlayer sets current player in the ring
	SetCurrentPlayer(player RingPlayer)

	//GetCurrentPlayer gets current player in the ring
	GetCurrentPlayer() (player RingPlayer)

	//HasCurrentPlayer return whether current player is set for the ring
	HasCurrentPlayer() bool
}
```

## Deck
```go
//Deck Deck of cards
type Deck interface {
	
        //Cards returns all cards
	CardsInDeck() []Card

	//IsCardPresent  Checks if a given card is present in deck
	IsCardPresent(c Card) bool

	//DrawCard draws a card from the deck
	DrawCard(i int) (Card, error)

	//DrawCards draws multiple cards from the deck
	DrawCards(m, n int) ([]Card, error)

	//PutCard puts card in th deck
	PutCard(c Card) error

	//PutCard puts multiple card in th deck
	PutCards(cards []Card) error

	//Shuffle shuffles the deck n times
	Shuffle(n int) error
}
```

