package structs

import "time"

type ClimaCellPayload struct {
	Data struct {
		Timelines []struct {
			Timestep  string    `json:"timestep"`
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
			Intervals []struct {
				StartTime time.Time `json:"startTime"`
				Values    struct {
					Temperature float64 `json:"temperature"`
				} `json:"values"`
			} `json:"intervals"`
		} `json:"timelines"`
	} `json:"data"`
}
