package internal

import (
	"2-B/internal/contract"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Big2 struct {
	deck       *Deck
	players    [4]*Player
	curPlayerI int
	topPattern contract.CardPattern
	topPlayer  *Player
}

func NewBig2() *Big2 {
	return &Big2{}
}

func (b *Big2) SetDeck(deck *Deck) {
	b.deck = deck
}

func (b *Big2) SetPlayers(players [4]*Player) {
	b.players = players
}

func (b *Big2) Start(ch chan string) {
	// 發牌
	var firstPlayerI int
	for {
		if b.deck.IsEmpty() {
			break
		}
		card := b.deck.Deal()
		if card.String() == "C[3]" {
			// 梅花三先出
			firstPlayerI = b.curPlayerI
		}
		b.players[b.curPlayerI].AddCard(card)
		b.curPlayerI++
		b.curPlayerI = b.curPlayerI % len(b.players)
	}
	b.curPlayerI = firstPlayerI
	b.topPlayer = b.players[firstPlayerI]
	// 開始打牌
	for {
		player := b.players[b.curPlayerI]
		if player == b.topPlayer {
			log.Println("新的回合開始了。")
			b.topPlayer = nil
			b.topPattern = nil
		}
		log.Printf("輪到%s了\n", player.Name)
		pattern := b.playerRound(ch, player)
		if pattern == nil {
			log.Printf("玩家 %s PASS.", player.Name)
		} else {
			log.Printf("玩家 %s 打出了 %s", player.Name, pattern)
			b.topPlayer = player
			b.topPattern = pattern
		}
		if player.IsHandEmpty() {
			log.Printf("遊戲結束，遊戲的勝利者為 %s", player.Name)
			return
		}
		b.curPlayerI++
		b.curPlayerI = b.curPlayerI % len(b.players)
	}
}

func (b *Big2) playerRound(ch <-chan string, player *Player) contract.CardPattern {
	showCards(player.Cards)
	var cardIndexes []int
	command := <-ch
	split := strings.Split(command, " ")
	for _, s := range split {
		i, _ := strconv.Atoi(s)
		cardIndexes = append(cardIndexes, i)
	}
	if cardIndexes[0] == -1 {
		if b.topPlayer == nil {
			log.Printf("你不能在新的回合中喊 PASS")
			return b.playerRound(ch, player)
		}
		// pass
		return nil
	} else {
		pattern := player.Play(cardIndexes)
		if !b.isPatternValid(pattern) {
			// 放回手牌
			if pattern != nil {
				for _, card := range pattern.GetCards() {
					player.AddCard(card)
				}
			}
			log.Println("此牌型不合法，請再嘗試一次。")
			return b.playerRound(ch, player)
		}
		return pattern
	}
}

func (b *Big2) isPatternValid(pattern contract.CardPattern) bool {
	if pattern == nil {
		return false
	}
	if b.topPattern == nil {
		return true
	}
	return pattern.IsGreaterThan(b.topPattern)
}

func showCards(cards []contract.Card) {
	var s string
	lastIndex := len(cards) - 1
	for i, card := range cards {
		var format string
		if i != lastIndex {
			l := 4 + len(card.Rank)
			format = "%-" + strconv.Itoa(l) + "s"
		} else {
			format = "%s"
		}
		s += fmt.Sprintf(format, strconv.Itoa(i))
	}
	log.Printf(s)
	s = ""
	for i, card := range cards {
		s += card.String()
		if i != lastIndex {
			s += " "
		}
	}
	log.Printf(s)
}
