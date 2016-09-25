package tvgo

import (
	"log"
	"net/http"
	"net/url"
	"strconv"

	"encoding/json"
	"fmt"
	c "github.com/blackdev1l/tvgo/lib/config"
	fs "github.com/blackdev1l/tvgo/lib/filesystem"
	str "github.com/blackdev1l/tvgo/lib/jsonStruct"
	"io/ioutil"
)

const API string = "http://api.tvmaze.com"

type Tvgo struct {
	Conf  c.Configuration
	shows []str.Show
}

func (tv *Tvgo) AddSeries(id int) {

	var episodes []str.Episodes

	log.Println("id is " + strconv.Itoa(id))

	var show str.Show
	body, err := tv.apiCall(API, "/shows/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &show)

	body, err = tv.apiCall(API, "/shows/"+strconv.Itoa(id)+"/episodes", nil)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &episodes)
	show.Episodes = episodes

	tv.saveShow(show)
}

func (tv *Tvgo) SearchSeries(seriesName string) {
	params := url.Values{}
	params.Add("q", seriesName)

	body, err := tv.apiCall(API, "/search/shows", params)
	if err != nil {
		log.Fatal(err)
	}

	var shows []str.Shows

	json.Unmarshal(body, &shows)

	if len(shows) == 0 {
		fmt.Println("Show not found")
		return
	}

	tv.renderSearch(shows)
}

func (t *Tvgo) RemoveSeries(id int) {
	path := t.Conf.Path + "/shows.json"
	if fs.Exists(path) {
		t.loadShows()
	}

	if len(t.shows) < 1 {
		log.Fatalln("Database is empty, exiting")
	}

	i := t.findPosition(id)
	if i != -1 {
		shows := append(t.shows[:i], t.shows[i+1:]...)

		data, err := json.Marshal(shows)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Saving...")
		ioutil.WriteFile(path, data, 0644)

	} else {
		log.Fatalln("id is incorrect")
	}
}

func (t *Tvgo) ListSeries() {

}

func (tv *Tvgo) apiCall(tmpUrl string, path string, params url.Values) ([]byte, error) {

	query, err := url.Parse(API)
	query.Path += path
	if err != nil {
		log.Fatal(err)

	}

	if params != nil {
		query.RawQuery = params.Encode()
	}

	log.Println("URL is " + query.String())
	resp, err := http.Get(query.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (t *Tvgo) saveShow(show str.Show) {
	path := t.Conf.Path + "/shows.json"

	if fs.Exists(path) {
		t.loadShows()
	}

	if !t.isAlreadyInDB(show) {
		t.shows = append(t.shows, show)
		data, err := json.Marshal(t.shows)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Saving...")
		ioutil.WriteFile(path, data, 0644)
	} else {
		log.Println("Entry already in DB, exiting")
	}
}

func (t *Tvgo) loadShows() {
	data, err := ioutil.ReadFile(t.Conf.Path + "/shows.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &t.shows)

}

func (t *Tvgo) isAlreadyInDB(show str.Show) bool {
	for _, v := range t.shows {
		if v.ID == show.ID {
			return true
		}
	}
	return false
}

func (t *Tvgo) findPosition(id int) int {
	for k, v := range t.shows {
		if v.ID == id {
			return k
		}
	}
	return -1
}
