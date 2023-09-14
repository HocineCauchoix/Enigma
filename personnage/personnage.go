package personnage

type personnage struct {
	Nom       string
	hp        int
	att       int
	def       int
	stam      int
	mana      int
	intel     int
	str       int
	dex       int
	SkilPoint int
	skill     []string
}

func Getguerrier() personnage {
	Guerrier := personnage{
		hp:        0,
		att:       12,
		def:       15,
		stam:      1000000,
		mana:      10,
		intel:     6,
		str:       12,
		dex:       6,
		SkilPoint: 20,
		skill:     []string{"charge"},
	}
	return Guerrier
}

func Getarcher() personnage {
	Archer := personnage{
		hp:        15,
		att:       12,
		def:       15,
		stam:      1000000,
		mana:      10,
		intel:     6,
		str:       12,
		dex:       6,
		SkilPoint: 20,
		skill:     []string{"charge"},
	}
	return Archer
}

func Getassassin() personnage {
	Assassin := personnage{
		hp:        5,
		att:       12,
		def:       15,
		stam:      1000000,
		mana:      10,
		intel:     6,
		str:       12,
		dex:       6,
		SkilPoint: 20,
		skill:     []string{"charge"},
	}
	return Assassin
}

func Getmage() personnage {
	Mage := personnage{
		hp:        20,
		att:       12,
		def:       15,
		stam:      1000000,
		mana:      10,
		intel:     6,
		str:       12,
		dex:       6,
		SkilPoint: 20,
		skill:     []string{"charge"},
	}
	return Mage
}

// func Personnage(name string) personnage {

// 	Assassin := personnage{
// 		hp:        10,
// 		att:       10,
// 		def:       10,
// 		stam:      100000000,
// 		mana:      0,
// 		intel:     12,
// 		str:       8,
// 		dex:       15,
// 		SkilPoint: 20,
// 		skill:     []string{"evisceration"},
// 	}

// 	Mage := personnage{
// 		hp:        10,
// 		att:       12,
// 		def:       10,
// 		stam:      10000,
// 		mana:      20,
// 		intel:     15,
// 		str:       10,
// 		dex:       12,
// 		SkilPoint: 20,
// 		skill:     []string{"fire ball"},
// 	}

// 	Archer := personnage{
// 		hp:        0,
// 		att:       10,
// 		def:       10,
// 		stam:      100000,
// 		mana:      15,
// 		intel:     12,
// 		str:       8,
// 		dex:       15,
// 		SkilPoint: 20,
// 		skill:     []string{"poison arrow"},
// 	}

// 	return name
