package skill

import "fmt"

func main() {

	type skill struct {
		att      int
		costmana int
		coststam int
	}

	charge := skill{
		att:      10,
		costmana: 4,
		coststam: 8,
	}

	evisceration := skill{
		att:      12,
		costmana: 8,
		coststam: 8,
	}

	fireball := skill{
		att:      12,
		costmana: 8,
		coststam: 8,
	}

	poisonarrow := skill{
		att:      12,
		costmana: 8,
		coststam: 8,
	}

	fmt.Println(charge)
	fmt.Println(evisceration)
	fmt.Println(fireball)
	fmt.Println(poisonarrow)
}
