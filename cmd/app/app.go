package main

import (
	"fmt"

	"github.com/raffleberry/sqlsheet/web"
)

func main() {

	dogs := []web.Pet{
		{
			Name:   "Jujube",
			Sex:    "Female",
			Intact: false,
			Age:    "10 months",
			Breed:  "German Shepherd/Pitbull",
		},
		{
			Name:   "Zephyr",
			Sex:    "Male",
			Intact: true,
			Age:    "13 years, 3 months",
			Breed:  "German Shepherd/Border Collie",
		},
	}
	web.PrintDogs(dogs)
	fmt.Println("Hello SqlSheet")

}
