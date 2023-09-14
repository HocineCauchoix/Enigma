package pnj

import "fmt"

func main() {

	type Pnj struct {
		name   string
		Metier string
		sex    string
	}

	Ilyes := Pnj{
		name:   "Ilyes",
		Metier: "curé du village",
		sex:    "man",
	}

	fmt.Println(Ilyes)

	var pnj string
	fmt.Printf("\n							Ilyes(pnj) :\nJe t'ai remarquer combattre tout ces monstres je ne sais pour quel raison evidemment mais tu es un guerrier de Vony pour cela et tant donné que je suis curé du village je    tien a te proposer trois objet choisi l'un des trois pour continuer à nous défendre \nTien regard :	1 pour Epée longue (5pts dmg) \n 		2 pour Potion de vie (restaure 10pv) \n 		3 pour le Bouclier (5 def)\n")
	fmt.Scan(&pnj)

	switch pnj {
	case "1":
		fmt.Println("Si c'est ce que tu désire l'épée longue est a toi")
	case "2":
		fmt.Println("Si c'est ce que tu désire la potion est a toi ")
	case "3":
		fmt.Println("Si c'est ce que tu désire le bouclier est a toi")
	}

}
