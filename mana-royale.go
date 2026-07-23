package main

import (
    "fmt"
    "ManaRoyale/entity"
    "ManaRoyale/util"
)

func main() {
    turn := 10
    numCards := 15
    pick := 3
    ind := pick+1
    pickedCards := [3]int{}
    cards := []assets.Card{
            {Name: "-2xDMG", Category: "DMG", Value: -2},
            {Name: "-1xDMG", Category: "DMG", Value: -1},
            {Name: "0xDMG", Category: "DMG", Value: 0},
            {Name: "1xDMG", Category: "DMG", Value: 1},
            {Name: "2xDMG", Category: "DMG", Value: 2},
            {Name: "Arrow25", Category: "PROJECTILE", Value: 25},
            {Name: "Arrow50", Category: "PROJECTILE", Value: 50},
            {Name: "Arrow100", Category: "PROJECTILE", Value: 100},
            {Name: "Arrow200", Category: "PROJECTILE", Value: 200},
            {Name: "Arrow500", Category: "PROJECTILE", Value: 500},
            {Name: "Heal", Category: "EPIC", Value: 10},
            {Name: "Poison", Category: "EPIC", Value: 11},
            {Name: "Sleep", Category: "EPIC", Value: 12},
            {Name: "Hex", Category: "EPIC", Value: 13},
            {Name: "Emag Revo", Category: "EPIC", Value: 14},
            }
    players := []assets.Player{
            {Name: "Artims", DPT:10, HP:100},
            {Name: "Nomanad", DPT:10, HP: 250},
            }
    for i:=0; i<turn ; i++ {
        fmt.Printf("Turn %d for %s\n",i,players[0].Name)
        for j:=0; j<pick; j++ {
            pickedCards[j] = randGen.Generate(numCards)
        }    
        //var cardIndex = randGen.Generate(numCards)
        //fmt.Printf("Card came as %s\n", cards[cardIndex].Name)
        fmt.Printf("Please pick a card, %s now!!!\n", players[0].Name)
        ind = pick+1
        for ; ind > pick ; {
           fmt.Printf("Pick a number between 1-%d\n", pick)
           _, err := fmt.Scan(&ind)
           if err != nil {
               break
           }
        }
        fmt.Printf("You selected card %s\n", cards[pickedCards[ind-1]].Name)
    }
}
