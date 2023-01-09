package main

import (
	"log"
)

func main() {
	log.Println(LOS_ANGELES_LAKERS.homeGeo.distanceBetween(DENVER_NUGGETS.homeGeo))
}
