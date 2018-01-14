package main

import (
  "fmt"
  "math/rand"
  "strconv"
  "time"
)

const NUM_CARDS = 52

func main() {
  //create and seed a psuedo-random number stream
  shuffler := rand.New(rand.NewSource(time.Now().UnixNano()))

  permutation := shuffler.Perm(NUM_CARDS)

  stdDeck := newStdDeck()
  shuffledDeck := shuffle(stdDeck, permutation)

  fmt.Println(stdDeck, "\n")
  fmt.Println(permutation, "\n")
  fmt.Println(shuffledDeck, "\n")

  fmt.Println(dealNHands(2, 7, shuffledDeck))
}

func newStdDeck() [NUM_CARDS]string {
    var card string
    var deck [NUM_CARDS]string
    var a, b int
    var suit []string

    //construct each suit j, 0 - 3 in a slice and build an ordered array of cards:
    for j := 0; j < 4; j++ {
      //compute slice endpoints for suit j:
      a = 13 * j //start at index 0 and shift 13 cards every suit
      b = a + 13

      //construct a slice to reference the current suit in deck array:
      suit = deck[a:b]

      //construct appropriate string for each card in suit j:
      for i := 1; i <= NUM_CARDS / 4; i++ { //i cards in each suit
        //give face cards proper names:
        switch i {
        case 1:
          card = "A"
        case 11:
          card = "J"
        case 12:
          card = "Q"
        case 13:
          card = "K"
        default:
          card = strconv.Itoa(i)
        }

        //add suit:
        switch j {
        case 0:
          card += "s"
        case 1:
          card += "c"
        case 2:
          card += "h"
        case 3:
          card += "d"
        }

        //write card to appropriate place in deck via suit:
        suit[i-1] = card
      }
    }

    return deck
}

//shuffle([]string, []int []string)
// cards []string:    list of cards in an arbitrary order
// permutation []int: precalculated permutation of {0, ... NUM_CARDS}
// returns list of cards in the permuted order
func shuffle(cards [NUM_CARDS]string, permutation []int) [NUM_CARDS]string {
  var shuffledDeck [NUM_CARDS]string

  for i := 0; i < NUM_CARDS; i++ {
    shuffledDeck[i] = cards[permutation[i]]
  }

  return shuffledDeck
}

//TODO: Finish dealNHands()
// n int: number of hands to deal
func dealNHands(n int, numCards int, cards [NUM_CARDS]string) [][]string {
  hands := make([][]string, n)

  //sanity check that there are enough cards to be dealt
  if (n * numCards) > NUM_CARDS {
    fmt.Println("Error: requesed more cards to be dealt then exist")
    return hands
  }

  //Use increment slice to push numCards cards into n different hands
  var inc []string
  inc = cards[0:]

  for i := range hands {
    hands[i] = inc[0:numCards]
    inc = inc[numCards:]
  }

  return hands
}
