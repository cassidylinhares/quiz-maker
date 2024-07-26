package models

type Question struct {
	Question string   `json:"question"`
	Answer   string   `json:"answer"`
	Options  []string `json:"options"`
}

type Quiz struct {
	Questions []Question `json:"questions"`
}
