package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"encoding/json"
	"io/ioutil"

	"github.com/google/uuid"
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

var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}


func (hole *Hole) calculateAdjScore() {
	adj := hole.Score + hole.HandicapStrokes
	maxScore := hole.Par + hole.HandicapStrokes + 2
	hole.AdjustedScore = maxScore
	if adj < maxScore {
		hole.AdjustedScore = adj
	}
}

func (round *Round) scoreDifferentialCalculation() {
	if round.AdjGrossScore == 0 || round.Score == 0 {
		adjScore := 0
		grossScore := 0
		for _, hole := range round.Holes {
			grossScore += hole.Score
			hole.calculateAdjScore()
			adjScore += hole.AdjustedScore
		}
		round.Score, round.AdjGrossScore = grossScore, adjScore
	}
	if len(round.Holes) < 18 {
		fmt.Println("Unable to calculate handicap for incomplete rounds")
	} else {
		round.ScoreDiff = (float32(round.AdjGrossScore) - round.Rating) * (float32(113) / float32(round.Slope))
	}
}

func (stats *Stats) CalculateHandicap() {
	if stats.NumRounds < 3 {
		fmt.Println("Not enough rounds. Need", 3 - stats.NumRounds, "more to calculate a handicap")
	}
	// dt := time.Now()
	for idx, round := range stats.Rounds {
		if round.ScoreDiff == 0 {
			round.scoreDifferentialCalculation()
			stats.Rounds[idx] = round
		}
		fmt.Println(round.ScoreDiff)
	}
	// handicap := Handicap {}
	fmt.Println(stats.Rounds[0].ScoreDiff)
}

func getUserInput(printString string) string{
	fmt.Println(printString)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured when reading input. Please try again.")
		return ""
	}
	input = strings.TrimSuffix(input, "\n")
	return input
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func stringToFloat(s string) float32 {
	i, err := strconv.ParseFloat(s, 32)
	if err != nil {
		fmt.Println(err)
	}
	return float32(i)
}

func stringToBool(s string) bool {
	i, err := strconv.ParseBool(s)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func addNewHole(holeNum int, addingPutts string, hdcpStrokes string) Hole {
	hole := Hole{}
	hole.Number = holeNum
	text := "Hole " + strconv.Itoa(holeNum) + " "
	hole.Score = stringToInt(getUserInput(text + "score?"))
	hole.Par = stringToInt(getUserInput(text + "par?"))
	if strings.Contains(addingPutts, "y") {
		hole.Putts = stringToInt(getUserInput(text + "putts?"))
	}
	if strings.Contains(hdcpStrokes, "y") {
		hole.HandicapStrokes = stringToInt(getUserInput(text + "handicap strokes?"))
	}
	CallClear()
	return hole
}

func addNewRound() Round {
	round := Round{}
	round.UUID = uuid.New().String()
	round.Course = getUserInput("What is the course name?")
	round.Slope = stringToInt(getUserInput("What is the course slope?"))
	round.Rating = stringToFloat(getUserInput("What is the course rating?"))
	holes := stringToInt(getUserInput("How many holes do you want to add?"))
	addingPutts := getUserInput("Are you adding putts? (y/n)")
	hdcpStrokes := getUserInput("Do you have handicap strokes? (y/n)")
	CallClear()
	i := 1
	for {
		hole := addNewHole(i, addingPutts, hdcpStrokes)
		round.Holes = append(round.Holes, hole)

		if i == holes {
			break
		}
		i++
	}
	return round
}

func deleteRound(uid string) {
	
}

func (stats *Stats) dumpData() {
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
	// test pointers to see if that's what I wanted with being able to change
	// values without having to reassign
	stats := readData()
	newRound := getUserInput("Would you like to enter a new round?")
	if strings.Contains(newRound, "y") {
		stats.Rounds := append(stats.Rounds, addNewRound())
		fmt.Println(round)
	}
	// uid := uuid.New()
	// uuid := uid.String()
	stats.CalculateHandicap()
	// fmt.Println(stats.Rounds[0])
	fmt.Println(stats.Rounds[0].ScoreDiff)
	fmt.Println(stats)
	stats.dumpData()
}