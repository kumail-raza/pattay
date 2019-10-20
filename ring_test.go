package pattay_test

import (
	"testing"

	"github.com/minhajuddinkhan/pattay"
	"github.com/stretchr/testify/assert"
)

var (
	NorthPlayer = "North Player"
	SouthPlayer = "South Player"
	WestPlayer  = "WestPlayer"
	EastPlayer  = "East Player"
)

var defaultPlayersNo = 4

func beforeEach(numberOfPlayers int) (pattay.Ring, error) {
	playerNames := []string{NorthPlayer, EastPlayer, SouthPlayer, WestPlayer}
	var players []pattay.RingPlayer
	for i := 0; i < numberOfPlayers; i++ {
		pl := pattay.NewRingPlayer(playerNames[i]).(pattay.RingPlayer)
		players = append(players, pl)
	}
	return pattay.NewRing(players...)

}

func TestRing_HasFourPlayers(t *testing.T) {
	ring, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(ring.Players()))
}

func TestRing_NextAfterSouthPlayerNextPlayerIsWest(t *testing.T) {

	ring, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	players := ring.Players()
	ring.SetCurrentPlayer(players[0])
	p, err := ring.Next()
	assert.Nil(t, err)
	assert.Equal(t, p.Name(), EastPlayer)

	p, err = ring.Next()
	assert.Nil(t, err)
	assert.Equal(t, p.Name(), SouthPlayer)

	p, err = ring.Next()
	assert.Nil(t, err)
	assert.Equal(t, p.Name(), WestPlayer)

	p, err = ring.Next()
	assert.Nil(t, err)
	assert.Equal(t, p.Name(), NorthPlayer)

}

func TestRing_IsCurrentPlayerSetInRing(t *testing.T) {
	ring, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	assert.False(t, false, ring.HasCurrentPlayer())

}

func TestRing_NextWithoutSettingCurrentShouldError(t *testing.T) {
	r, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	_, err = r.Next()
	assert.Error(t, err)
}

func TestRing_CreateRingWithThreePlayers(t *testing.T) {
	_, err := beforeEach(3)
	assert.Error(t, err)
}

func TestRing_GetAndSetPlayers(t *testing.T) {
	r, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	players := r.Players()
	p1 := players[0]
	assert.Nil(t, r.GetCurrentPlayer())
	r.SetCurrentPlayer(p1)
	assert.Equal(t, p1.Name(), r.GetCurrentPlayer().Name())
}
