package clg_json

import "encoding/json"

type Challenge struct {
	Header      string `json:"header"`
	Description string `json:"description"`
}

func LoadChallenges() ([]Challenge, error) {
	var challenges []Challenge
	err := json.Unmarshal(challenges_json, &challenges)
	if err != nil {
		return nil, err
	}
	return challenges, nil
}
