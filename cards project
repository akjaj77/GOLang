package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValue := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardsSuits {
		for _, value := range cardsValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	slicesOfString := []string(d)
	s := strings.Join(slicesOfString, ",")
	return s
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	str := string(bs)
	slicesOfString := strings.Split(str, ",")
	return deck(slicesOfString)
}

func (d deck) shuffle() {
	t := time.Now().UnixNano()
	source := rand.NewSource(t)
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
