package main

import "time"

type Game struct {
	homeTeam TeamName
	awayTeam TeamName
	location GeoData
	date     time.Time
}

type Schedule struct {
	games []Game
}

func (s Schedule) gamesByDate() map[time.Time][]Game {
	res := map[time.Time][]Game{}

	for _, game := range s.games {
		if _, ok := res[game.date]; ok {
			res[game.date] = append(res[game.date], game)
		} else {
			res[game.date] = []Game{game}
		}
	}
	return res
}

func (s Schedule) gamesByTeam() map[TeamName][]Game {
	res := map[TeamName][]Game{}

	for _, game := range s.games {
		if _, ok := res[game.homeTeam]; ok {
			res[game.homeTeam] = append(res[game.homeTeam], game)
		} else {
			res[game.homeTeam] = []Game{game}
		}

		if _, ok := res[game.awayTeam]; ok {
			res[game.awayTeam] = append(res[game.awayTeam], game)
		} else {
			res[game.awayTeam] = []Game{game}
		}
	}
	return res
}
