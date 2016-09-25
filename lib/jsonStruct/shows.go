package jsonStruct

type Shows struct {
	Score float64 `json:"score"`
	Show  `json:"show"`
}

type Show struct {
	Links      `json:"_links"`
	Externals  `json:"externals"`
	Genres     []string `json:"genres"`
	ID         int      `json:"id"`
	Image      `json:"image"`
	Language   string `json:"language"`
	Name       string `json:"name"`
	Network    `json:"network"`
	Premiered  string `json:"premiered"`
	Rating     `json:"rating"`
	Runtime    int `json:"runtime"`
	Schedule   `json:"schedule"`
	Status     string      `json:"status"`
	Summary    string      `json:"summary"`
	Type       string      `json:"type"`
	Updated    int         `json:"updated"`
	URL        string      `json:"url"`
	WebChannel interface{} `json:"webChannel"`
	Weight     int         `json:"weight"`
	Episodes   []Episodes
}

type Links struct {
	Previousepisode `json:"previousepisode"`
	Self            `json:"self"`
}

type Externals struct {
	Imdb    string `json:"imdb"`
	Thetvdb int    `json:"thetvdb"`
	Tvrage  int    `json:"tvrage"`
}

type Image struct {
	Medium   string `json:"medium"`
	Original string `json:"original"`
}

type Network struct {
	Country `json:"country"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
}

type Rating struct {
	Average float64 `json:"average"`
}

type Schedule struct {
	Days []string `json:"days"`
	Time string   `json:"time"`
}

type Previousepisode struct {
	Href string `json:"href"`
}

type Self struct {
	Href string `json:"href"`
}

type Country struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Timezone string `json:"timezone"`
}
