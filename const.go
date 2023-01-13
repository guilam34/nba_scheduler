package main

type TeamName int64

const (
	ATL TeamName = iota
	BOS
	BKN
	CHA
	CHI
	CLE
	DAL
	DEN
	DET
	GSW
	HOU
	IND
	LAC
	LAL
	MEM
	MIA
	MIL
	MIN
	NOH
	NYK
	OKC
	ORL
	PHI
	PHX
	POR
	SAC
	SAS
	TOR
	UTA
	WSH
	TEAM_NAME_LIMIT
)

type Conference int64

const (
	WESTERN Conference = iota
	EASTERN
)

type Division int64

const (
	ATLANTIC Division = iota
	CENTRAL
	SOUTHEAST
	NORTHWEST
	PACIFIC
	SOUTHWEST
)
