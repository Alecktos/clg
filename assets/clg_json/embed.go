package clg_json

import _ "embed"

var (
	//go:embed challenges.json
	challenges_json []byte

	//go:embed endgame.json
	endgame_json []byte
)
