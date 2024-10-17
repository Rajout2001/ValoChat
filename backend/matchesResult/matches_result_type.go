package matchesResult

type MatchesResult struct {
	Team1 string    `json:"team1"`
	Team2 string    `json:"team2"`
	Score1 string    `json:"score1"`
	Score2 string    `json:"score2"`
	Flag1 string    `json:"flag1"`
	Flag2 string    `json:"flag2"`
	Time_completed string    `json:"time_completed"`
	Round_info string    `json:"round_info"`
	Tournament_name string    `json:"tournament_name"`
	Match_page string    `json:"match_page"`
	Tournament_icon string    `json:"tournament_icon"`
	ChatId string    `json:"chatId"`
	Id string    `json:"id"`
}

type MatchesResultData struct {
	Status 		int		`json:"status"`
	Segments 	[]MatchesResult `json:"segments"`
}

type MatchesResultResponse struct {
	Data	MatchesResultData 	`json:"data"`
}

