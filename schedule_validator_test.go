package main

import (
	"fmt"
	"testing"
	"time"
)

func TestValidateScheduleNoTeams(t *testing.T) {
	schedule := Schedule{}
	err := validateSchedule(schedule, map[TeamName]Team{}, ScheduleRestrictions{teamsInLeague: 1})
	if err.Error() != "expected 1 teams in schedule" {
		t.Fatal("Expected failure due to no teams")
	}
}

func TestValidateScheduleTeamsWithNotEnoughGames(t *testing.T) {
	schedule := Schedule{
		games: []Game{
			{
				homeTeam: BOS,
				awayTeam: LAL,
				location: GeoData{},
				date:     time.Now(),
			},
		},
	}

	err := validateSchedule(schedule, map[TeamName]Team{}, ScheduleRestrictions{teamsInLeague: 2, gamesPerSeason: 2})
	if err.Error() != fmt.Sprintf(`expected schedule to have %d games`, 2) {
		t.Fatal("Expected failure due to not enough games for a team")
	}
}

func TestValidateScheduleTeamsWithNotEnoughIntraConfGames(t *testing.T) {
	schedule := Schedule{
		games: []Game{
			{
				homeTeam: DEN,
				awayTeam: LAL,
				location: GeoData{},
				date:     time.Now(),
			},
		},
	}

	den := Team{
		name:       DEN,
		conference: WESTERN,
		division:   NORTHWEST,
	}
	lal := Team{
		name:       LAL,
		conference: WESTERN,
		division:   PACIFIC,
	}
	teamMap := map[TeamName]Team{}
	teamMap[DEN] = den
	teamMap[LAL] = lal

	err := validateSchedule(schedule, teamMap, ScheduleRestrictions{
		teamsInLeague:        2,
		gamesPerSeason:       1,
		gamesPerIntraConfOpp: 2,
	})
	if (err.Error() != fmt.Sprintf(`expected %d to have 2 games against %d`, DEN, LAL)) && (err.Error() != fmt.Sprintf(`expected %d to have 2 games against %d`, LAL, DEN)) {
		t.Fatal("Expected failure due to not enough games against intra-conference team")
	}
}

func TestValidateScheduleTeamsWithNotEnoughInterConfGames(t *testing.T) {
	schedule := Schedule{
		games: []Game{
			{
				homeTeam: BOS,
				awayTeam: LAL,
				location: GeoData{},
				date:     time.Now(),
			},
		},
	}

	bos := Team{
		name:       BOS,
		conference: EASTERN,
		division:   ATLANTIC,
	}
	lal := Team{
		name:       LAL,
		conference: WESTERN,
		division:   PACIFIC,
	}
	teamMap := map[TeamName]Team{}
	teamMap[BOS] = bos
	teamMap[LAL] = lal

	err := validateSchedule(schedule, teamMap, ScheduleRestrictions{
		teamsInLeague:        2,
		gamesPerSeason:       1,
		gamesPerInterConfOpp: 2,
	})
	if (err.Error() != fmt.Sprintf(`expected %d to have 2 games against %d`, BOS, LAL)) && (err.Error() != fmt.Sprintf(`expected %d to have 2 games against %d`, LAL, BOS)) {
		t.Fatal("Expected failure due to not enough games against inter-conference team")
	}
}

func TestValidateScheduleTeamsWithNotEnoughDivisionGames(t *testing.T) {
	schedule := Schedule{
		games: []Game{
			{
				homeTeam: GSW,
				awayTeam: LAL,
				location: GeoData{},
				date:     time.Now(),
			},
		},
	}

	gsw := Team{
		name:       GSW,
		conference: WESTERN,
		division:   PACIFIC,
	}
	lal := Team{
		name:       LAL,
		conference: WESTERN,
		division:   PACIFIC,
	}
	teamMap := map[TeamName]Team{}
	teamMap[GSW] = gsw
	teamMap[LAL] = lal

	err := validateSchedule(schedule, teamMap, ScheduleRestrictions{
		teamsInLeague:  2,
		gamesPerSeason: 1,
		gamesPerDivOpp: 2,
	})
	if (err.Error() != fmt.Sprintf(`expected %d to have 2 games against %d`, GSW, LAL)) && (err.Error() != fmt.Sprintf(`expected %d to have 2 games against %d`, LAL, GSW)) {
		t.Fatal("Expected failure due to not enough games against inter-division team")
	}
}
