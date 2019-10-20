package pattay_test

import (
	"testing"

	"github.com/minhajuddinkhan/pattay"
	"github.com/stretchr/testify/assert"
)

func TestCard_ShouldHaveValidNumberAndHouse(t *testing.T) {

	house := "Spades"
	cardNumber := 1
	card := pattay.NewCard(house, cardNumber)
	assert.Equal(t, card.House(), house)
	assert.Equal(t, card.Number(), cardNumber)
}
