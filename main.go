package main

import (
	"enigma/jeu"
	"enigma/personnage"
	"fmt"
)

func main() {
	jeu.Game()
	perso := personnage.Getguerrier()
	fmt.Println(perso)
	perso1 := personnage.Getmage()
	fmt.Println(perso1)
	perso2 := personnage.Getassassin()
	fmt.Println(perso2)
	perso3 := personnage.Getarcher()
	fmt.Println(perso3)
}
