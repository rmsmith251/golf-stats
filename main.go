package main

import (
	"fmt"
	"time"

	"encoding/json"
	"io/ioutil"

	// "github.com/google/uuid"
)

type (
	Hole struct {
		Number int `json:"number"`
		Par int `json:"par"`
		Score int `json:"score"`
		Putts int `json:"putts"`
		HandicapStrokes int `json:"handicap_strokes"`
		AdjustedScore int `json:"adj_score"`
	}
	Round struct {
		UUID string `json:"uuid"`
		Number int `json:"number"`
		Course string `json:"course"`
		Slope int `json:"slope"`
		Rating float32 `json:"rating"`
		Tees string `json:"tees"`
		Holes []Hole `json:"holes"`
		Score int `json:"score"`
		Date time.Time `json:"date"`
		ScoreDiff float32 `json:"score_diff"`
		AdjGrossScore int `json:"adj_gross_score"`
		PlayingConditionsCalc float32 `json:"playing_conditions_calc"`
		PairedWith string `json:"pair"`
	}
	Handicap struct {
		Valid bool `json:"valid"`
		Handicap float32 `json:"handicap"`
		DateGenerated time.Time `json:"date_generated"`
		Rounds int `json:"rounds"`
	}
	Stats struct {
		Player string `json:"player"`
		Rounds []Round `json:"rounds"`
		NumRounds float32 `json:"num_rounds"`
		Handicaps []Handicap `json:"handicaps"`
		IndicesInUse []int `json:"rounds_in_use"`
		LowHandicapIndex float32 `json:"low_handicap_index"`
		LowHandicapDate time.Time `json:"low_handicap_date"`
	}
)



func (hole Hole) calculateAdjScore() Hole {
	adj := hole.Score + hole.HandicapStrokes
	max_score := hole.Par + hole.HandicapStrokes + 2
	hole.AdjustedScore = max_score
	if adj < max_score {
		hole.AdjustedScore = adj
	}
	return hole
}

func (round Round) scoreDifferentialCalculation() Round {
	if round.AdjGrossScore == 0 || round.Score == 0 {
		adj_score := 0
		gross_score := 0
		for idx, hole := range round.Holes {
			gross_score += hole.Score
			round.Holes[idx] = hole.calculateAdjScore()
			adj_score += round.Holes[idx].AdjustedScore
		}
		round.Score, round.AdjGrossScore = gross_score, adj_score
	}
	if len(round.Holes) < 18 {
		fmt.Println("Unable to calculate handicap for incomplete rounds")
	} else {
		round.ScoreDiff = (float32(round.AdjGrossScore) - round.Rating) * (float32(113) / float32(round.Slope))
	}
	return round
}

func (stats Stats) CalculateHandicap() Stats {
	if stats.NumRounds < 3 {
		fmt.Println("Not enough rounds. Need", 3 - stats.NumRounds, "more to calculate a handicap")
	}
	// dt := time.Now()
	for idx, round := range stats.Rounds {
		if round.ScoreDiff == 0 {
			round = stats.Rounds[idx].scoreDifferentialCalculation()
			stats.Rounds[idx] = round
		}
	}
	// handicap := Handicap {}


	return stats
}

func (stats Stats) dumpData() {
	content, err := json.Marshal(stats)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("golfdata.json", content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func readData() Stats {
	content, err := ioutil.ReadFile("golfdata.json")
	if err != nil {
		fmt.Println(err)
	}
	stats := Stats{}
	err = json.Unmarshal(content, &stats)
	if err != nil {
		fmt.Println(err)
	}
	return stats
}

func main() {
	stats := readData()
	// uid := uuid.New()
	// uuid := uid.String()
	stats.CalculateHandicap()
	// fmt.Println(stats.Rounds[0])
	fmt.Println(stats)
	stats.dumpData()
}