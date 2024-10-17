package services

import (
	"encoding/json"
	"net/http"
	"pc3r/matchesResult"
	db "pc3r/prisma"
)


func GetMatchesResult(res http.ResponseWriter, req *http.Request) {
	matches, err := matchesResult.GetMatchesResultFromAPI()
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		message := "Bad request to MATCHES/RESULT API"
		json.NewEncoder(res).Encode(CustomError(message, BAD_REQUEST))
		return
	}
	type responseGetMatchesResult struct {
		MatchesResult []matchesResult.MatchesResult `json:"matchesResult"`
	}
	response := responseGetMatchesResult{
		MatchesResult: matches,
	}
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(response)
}

type responseGetMatchesResults struct {
	MatchesResults []matchesResult.MatchesResult `json:"matchesResults"`
}

func GetMatchesResultsFromDB(res http.ResponseWriter, req *http.Request) { 
	prisma, ctx := db.GetPrisma()

	result, err := prisma.MatchResult.FindMany().Exec(ctx)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		message := "Match Not Found"
		json.NewEncoder(res).Encode(CustomError(message, NOT_FOUND))
		return
	}

	// Convert result to the appropriate type
    var matchesResults [] matchesResult.MatchesResult
    for _, r := range result {
        matchesResults = append(matchesResults, matchesResult.MatchesResult{
			Team1: r.Team1,
			Team2: r.Team2,
			Score1: r.Score1,
			Score2: r.Score2,
			Flag1: r.Flag1,
			Flag2: r.Flag2,
			Time_completed: r.TimeCompleted,
			Round_info: r.RoundInfo,
			Tournament_name: r.TournamentName,
			Match_page: r.MatchPage,
			Tournament_icon: r.TournamentIcon,
			ChatId: r.ChatID,
			Id: r.ID,
        })

    }

	response := responseGetMatchesResults {
		MatchesResults: matchesResults,
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(response)
}
