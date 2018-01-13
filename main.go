package main

import (
  "fmt"
  "math/rand"
  "time"
)

const numCards = 52

func main() {
  stdDeck := [numCards]string{
    "As",
    "2s",
    "3s",
    "4s",
    "5s",
    "6s",
    "7s",
    "8s",
    "9s",
    "10s",
    "Js",
    "Qs",
    "Ks",
    "Ac",
    "2c",
    "3c",
    "4c",
    "5c",
    "6c",
    "7c",
    "8c",
    "9c",
    "10c",
    "Jc",
    "Qc",
    "Kc",
    "Ad",
    "2d",
    "3d",
    "4d",
    "5d",
    "6d",
    "7d",
    "8d",
    "9d",
    "10d",
    "Jd",
    "Qd",
    "Kd",
    "Ah",
    "2h",
    "3h",
    "4h",
    "5h",
    "6h",
    "7h",
    "8h",
    "9h",
    "10h",
    "Jh",
    "Qh",
    "Kh",
  }

  //create and seed a psuedo-random number stream
  shuffler := rand.New(rand.NewSource(time.Now().UnixNano()))

  permutation := shuffler.Perm(numCards-1)

  fmt.Println(permutation)
  fmt.Println("\n")
  fmt.Println(stdDeck)
}

//TODO:implement shuffle function
func shuffle(cards [numCards]string, permutation []int) int {

}
