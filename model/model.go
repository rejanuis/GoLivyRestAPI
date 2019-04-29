package model

type ResultData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	// Data    map[string]interface{} `json:"data"`
	Data string `json:"data"`
}

type RespData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type Error struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   int    `json:"error"`
}

type Params struct {
	Criteria_id string `json:"criteria_id"`
	Start       string `json:"start"`
	End         string `json:"end"`
}
