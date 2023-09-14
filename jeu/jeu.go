package jeu

import (
	"enigma/personnage"
	"fmt"
)

func Game() {
	s := "->"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorBlue := "\033[34m"
	colorWhite := "\033[37m"
	colorPurple := "\033[35m"
	var nom string
	fmt.Println(string(colorWhite), "Avant de commencer tu dois inscrire ton nom d'Aventurier :")
	fmt.Scan(&nom)

	var imput string
	fmt.Printf("Enchanté %s avant de te lancer très important nous allons créer ton personnage :\n", nom)
	fmt.Println(string(colorRed), s, "Guerrier")
	fmt.Println(string(colorBlue), "  Mage")
	fmt.Println(string(colorPurple), "  Assassin")
	fmt.Println(string(colorGreen), "  Archer")
	fmt.Scan(&imput)

	switch imput {
	case "s":
		fmt.Println(string(colorRed), "  Guerrier")
		fmt.Println(string(colorBlue), s, "Mage")
		fmt.Println(string(colorPurple), "  Assassin")
		fmt.Println(string(colorGreen), "  Archer")
	case "ok":
		fmt.Print("Vous êtes desormais un guerrier")
		fmt.Print("\n")
		fmt.Print("voici vos statistiques de depart :", "\n")
		fmt.Print(personnage.Getguerrier())
	}
	var imput1 string
	fmt.Scan(&imput1)

	switch imput1 {
	case "s":
		fmt.Println(string(colorRed), "  Guerrier")
		fmt.Println(string(colorBlue), "  Mage")
		fmt.Println(string(colorPurple), s, "Assassin")
		fmt.Println(string(colorGreen), "  Archer")
	case "ok":
		fmt.Print("Vous etes desormais un guerrier")
		fmt.Print("\n")
		fmt.Print("voici vos statistiques de depart :", "\n")
		fmt.Print(personnage.Getmage())
	}

	var imput2 string
	fmt.Scan(&imput2)

	switch imput2 {
	case "s":
		fmt.Println(string(colorRed), "  Guerrier")
		fmt.Println(string(colorBlue), "  Mage")
		fmt.Println(string(colorPurple), "  Assassin")
		fmt.Println(string(colorGreen), s, "Archer")
	case "ok":
		fmt.Print("Vous etes desormais un Assassin")
		fmt.Print("\n")
		fmt.Print("voici vos statistiques de depart :", "\n")
		fmt.Print(personnage.Getassassin())

		var imput3 string
		fmt.Scan(&imput3)

		switch imput3 {
		case "ok":
			fmt.Print("Vous etes desormais un Archer")
			fmt.Print("\n")
			fmt.Print("voici vos statistiques de depart :", "\n")
			fmt.Print(personnage.Getarcher())
		}
	}
}
