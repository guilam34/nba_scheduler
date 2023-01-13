package main

import "fmt"

// TODO: Move into class that calls the validation code
const NUM_TEAMS_IN_LEAGUE = 30
const NUM_GAMES_PER_SEASON = 82
const NUM_GAMES_PER_DIVISION_OPP = 4
const MIN_GAMES_PER_CONF_OPP = 3
const GAMES_PER_INTER_CONF_OPP = 2

type ScheduleRestrictions struct {
	gamesPerSeason       int
	teamsInLeague        int
	gamesPerDivOpp       int
	gamesPerIntraConfOpp int
	gamesPerInterConfOpp int
}

func validateSchedule(s Schedule, teamMetadata map[TeamName]Team, sr ScheduleRestrictions) error {
	gamesByTeam := s.gamesByTeam()
	if len(gamesByTeam) != sr.teamsInLeague {
		return fmt.Errorf(`expected %d teams in schedule`, sr.teamsInLeague)
	}

	for team, games := range gamesByTeam {
		err := validateTeamSchedule(team, games, teamMetadata, sr)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateTeamSchedule(teamName TeamName, games []Game, teamMetadata map[TeamName]Team, sr ScheduleRestrictions) error {
	if len(games) != sr.gamesPerSeason {
		return fmt.Errorf(`expected schedule to have %d games`, sr.gamesPerSeason)
	}

	gamesPerOpp := map[TeamName]int{}
	for _, game := range games {
		oppName := game.homeTeam
		if oppName == teamName {
			oppName = game.awayTeam
		}
		if _, ok := gamesPerOpp[oppName]; ok {
			gamesPerOpp[oppName]++
		} else {
			gamesPerOpp[oppName] = 1
		}
	}

	for oppName, numGames := range gamesPerOpp {
		team := teamMetadata[teamName]
		opp := teamMetadata[oppName]

		minGames := getMinGames(team, opp, sr)
		if numGames < minGames {
			return fmt.Errorf(`expected %d to have %d games against %d`, teamName, minGames, oppName)
		}

	}

	return nil
}

// TODO: Account for yearly rotation of intra-conf opponents played 3 times
func getMinGames(team Team, opp Team, sr ScheduleRestrictions) int {
	isSameConference := team.conference == opp.conference
	isSameDivision := team.division == opp.division

	if isSameDivision {
		return sr.gamesPerDivOpp
	} else if isSameConference {
		return sr.gamesPerIntraConfOpp
	}
	return sr.gamesPerInterConfOpp
}
