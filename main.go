package main

type (
	Stats struct {
		Round []struct {
			Hole []struct {
				Par int `json:"par"`
				Score int `json:"score"`
				Putts int `json:putts`
			}
			Date int
		}
		NumRounds int `json:"num_rounds"`
		CurHdcp int `json:"cur_hdcp"`
	}
)