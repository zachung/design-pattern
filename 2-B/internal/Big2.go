package internal

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Big2 struct {
	deck       *Deck
	players    [4]*Player
	curPlayerI int
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
	// 開始打牌
	passCount := 3
	for {
		if passCount == 3 {
			log.Println("新的回合開始了。")
			passCount = 0
		}
		player := b.players[b.curPlayerI]
		log.Printf("輪到%s了\n", player.Name)
		showCards(player.Cards)
		var cardIndexes []int
		command := <-ch
		split := strings.Split(command, " ")
		for _, s := range split {
			i, _ := strconv.Atoi(s)
			cardIndexes = append(cardIndexes, i)
		}
		if cardIndexes[0] == -1 {
			// pass
			log.Printf("玩家 %s PASS.", player.Name)
			passCount++
		} else {
			pattern := player.Play(cardIndexes)
			log.Printf("玩家 %s 打出了 %s", player.Name, pattern)
			passCount = 0

			if player.IsHandEmpty() {
				log.Printf("遊戲結束，遊戲的勝利者為 %s", player.Name)
				break
			}
		}
		b.curPlayerI++
		b.curPlayerI = b.curPlayerI % len(b.players)
	}
}

func showCards(cards []Card) {
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
