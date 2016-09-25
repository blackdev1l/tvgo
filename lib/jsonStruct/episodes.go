package jsonStruct

type Episodes struct {
	Links    `json:"_links"`
	Airdate  string `json:"airdate"`
	Airstamp string `json:"airstamp"`
	Airtime  string `json:"airtime"`
	ID       int    `json:"id"`
	Image    `json:"image"`
	Name     string `json:"name"`
	Number   int    `json:"number"`
	Runtime  int    `json:"runtime"`
	Season   int    `json:"season"`
	Summary  string `json:"summary"`
	URL      string `json:"url"`
	Seen     bool
}
