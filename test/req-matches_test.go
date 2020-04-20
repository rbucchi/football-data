package test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	apiclient "github.com/rbucchi/football-data/pkg/api-client"
	"github.com/rbucchi/football-data/pkg/data/business"
	"github.com/rbucchi/football-data/pkg/data/request"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	client := new(apiclient.ApiClient)
	client.HttpClient = MockClient{}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New(
			"Error from web server",
		)
	}
	var matches business.MatchList
	err := client.Get(
		&matches,
		request.MatchesRequest{
			Filter:      request.Filter{Matchday: 1, Plan: ""},
			Competition: 2019,
		},
	)
	assert.NotNil(t, err)
	assert.EqualValues(t, 0, matches.Count)
}

func TestMatchesSuccess(t *testing.T) {
	client := new(apiclient.ApiClient)
	client.HttpClient = MockClient{}

	json := `{"count":10,"filters":{"matchday":"1"},"competition":{"id":2019,"area":{"id":2114,"name":"Italy"},"name":"Serie A","code":"SA","plan":"TIER_ONE","lastUpdated":"2020-04-18T00:00:07Z"},"matches":[{"id":279230,"season":{"id":530,"startDate":"2019-08-24","endDate":"2020-05-24","currentMatchday":32},"utcDate":"2019-08-24T16:00:00Z","status":"FINISHED","matchday":1,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2019-09-26T15:34:12Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"AWAY_TEAM","duration":"REGULAR","fullTime":{"homeTeam":0,"awayTeam":1},"halfTime":{"homeTeam":0,"awayTeam":1},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":112,"name":"Parma Calcio 1913"},"awayTeam":{"id":109,"name":"Juventus FC"},"referees":[{"id":10977,"name":"Fabio Maresca","nationality":null},{"id":10987,"name":"Filippo Valeriani","nationality":null},{"id":11067,"name":"Pasquale De Meo","nationality":null},{"id":11039,"name":"Marco Piccinini","nationality":null},{"id":11057,"name":"Paolo Mazzoleni","nationality":null}]},{"id":279228,"season":{"id":530,"startDate":"2019-08-24","endDate":"2020-05-24","currentMatchday":32},"utcDate":"2019-08-24T18:45:00Z","status":"FINISHED","matchday":1,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2019-09-26T15:34:12Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"AWAY_TEAM","duration":"REGULAR","fullTime":{"homeTeam":3,"awayTeam":4},"halfTime":{"homeTeam":1,"awayTeam":2},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":99,"name":"ACF Fiorentina"},"awayTeam":{"id":113,"name":"SSC Napoli"},"referees":[{"id":11029,"name":"Davide Massa","nationality":null},{"id":11048,"name":"Alberto Tegoni","nationality":null},{"id":10992,"name":"Stefano Alassio","nationality":null},{"id":11065,"name":"Daniele Doveri","nationality":null},{"id":11075,"name":"Paolo Valeri","nationality":null}]},{"id":279231,"season":{"id":530,"startDate":"2019-08-24","endDate":"2020-05-24","currentMatchday":32},"utcDate":"2019-08-25T16:00:00Z","status":"FINISHED","matchday":1,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2019-09-26T15:34:12Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"HOME_TEAM","duration":"REGULAR","fullTime":{"homeTeam":1,"awayTeam":0},"halfTime":{"homeTeam":0,"awayTeam":0},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":115,"name":"Udinese Calcio"},"awayTeam":{"id":98,"name":"AC Milan"},"referees":[{"id":11021,"name":"Fabrizio Pasqua","nationality":null},{"id":11001,"name":"Mauro Vivenzi","nationality":null},{"id":11080,"name":"Sergio Ranghetti","nationality":null},{"id":57829,"name":"Lorenzo Maggioni","nationality":null},{"id":10990,"name":"Piero Giacomelli","nationality":null}]},{"id":279229,"season":{"id":530,"startDate":"2019-08-24","endDate":"2020-05-24","currentMatchday":32},"utcDate":"2019-08-25T18:45:00Z","status":"FINISHED","matchday":1,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2019-09-26T15:34:12Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"DRAW","duration":"REGULAR","fullTime":{"homeTeam":3,"awayTeam":3},"halfTime":{"homeTeam":2,"awayTeam":2},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":100,"name":"AS Roma"},"awayTeam":{"id":107,"name":"Genoa CFC"},"referees":[{"id":11050,"name":"Giampaolo Calvarese","nationality":null},{"id":11070,"name":"Stefano Liberti","nationality":null},{"id":11118,"name":"Dario Cecconi","nationality":null},{"id":11086,"name":"Riccardo Ros","nationality":null},{"id":10985,"name":"Michael Fabbri","nationality":null}]},{"id":279232,"season":{"id":530,"startDate":"2019-08-24","endDate":"2020-05-24","currentMatchday":32},"utcDate":"2019-08-25T18:45:00Z","status":"FINISHED","matchday":1,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2019-09-26T15:34:12Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"AWAY_TEAM","duration":"REGULAR","fullTime":{"homeTeam":0,"awayTeam":3},"halfTime":{"homeTeam":0,"awayTeam":1},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":584,"name":"UC Sampdoria"},"awayTeam":{"id":110,"name":"SS Lazio"},"referees":[{"id":10988,"name":"Gianluca Rocchi","nationality":null},{"id":11055,"name":"Filippo Meli","nationality":null},{"id":10991,"name":"Giorgio Peretti","nationality":null},{"id":11071,"name":"Marco Serra","nationality":null},{"id":11047,"name":"Luca Banti","nationality":null}]},{"id":279233,"season":{"id":530,"startDate":"2019-08-24","endDate":"2020-05-24","currentMatchday":32},"utcDate":"2019-08-25T18:45:00Z","status":"FINISHED","matchday":1,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2019-09-26T15:34:12Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"AWAY_TEAM","duration":"REGULAR","fullTime":{"homeTeam":0,"awayTeam":1},"halfTime":{"homeTeam":0,"awayTeam":0},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":104,"name":"Cagliari Calcio"},"awayTeam":{"id":449,"name":"Brescia Calcio"},"referees":[{"id":11115,"name":"Eugenio Abbattista","nationality":null},{"id":11044,"name":"Giacomo Paganessi","nationality":null},{"id":11182,"name":"Marco Bresmes","nationality":null},{"id":11095,"name":"Lorenzo Illuzzi","nationality":null},{"id":11002,"name":"Luca Pairetto","nationality":null}]},{"id":279234,"season":{"id":530,"startDate":"2019-08-24","endDate":"2020-05-24","currentMatchday":32},"utcDate":"2019-08-25T18:45:00Z","status":"FINISHED","matchday":1,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2019-09-26T15:34:12Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"HOME_TEAM","duration":"REGULAR","fullTime":{"homeTeam":2,"awayTeam":1},"halfTime":{"homeTeam":1,"awayTeam":0},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":586,"name":"Torino FC"},"awayTeam":{"id":471,"name":"US Sassuolo Calcio"},"referees":[{"id":11043,"name":"Maurizio Mariani","nationality":null},{"id":11087,"name":"Alberto Santoro","nationality":null},{"id":11089,"name":"Giovanni Baccini","nationality":null},{"id":11084,"name":"Davide Ghersini","nationality":null},{"id":11006,"name":"Rosario Abisso","nationality":null}]},{"id":279235,"season":{"id":530,"startDate":"2019-08-24","endDate":"2020-05-24","currentMatchday":32},"utcDate":"2019-08-25T18:45:00Z","status":"FINISHED","matchday":1,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2019-09-26T15:34:12Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"DRAW","duration":"REGULAR","fullTime":{"homeTeam":1,"awayTeam":1},"halfTime":{"homeTeam":1,"awayTeam":1},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":450,"name":"Hellas Verona FC"},"awayTeam":{"id":103,"name":"Bologna FC 1909"},"referees":[{"id":11280,"name":"Antonio Giua","nationality":null},{"id":11045,"name":"Giorgio Schenone","nationality":null},{"id":57247,"name":"Davide Imperiale","nationality":null},{"id":11082,"name":"Antonio Di Martino","nationality":null},{"id":11083,"name":"Luigi Nasca","nationality":null}]},{"id":279236,"season":{"id":530,"startDate":"2019-08-24","endDate":"2020-05-24","currentMatchday":32},"utcDate":"2019-08-25T18:45:00Z","status":"FINISHED","matchday":1,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2019-09-26T15:34:12Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"AWAY_TEAM","duration":"REGULAR","fullTime":{"homeTeam":2,"awayTeam":3},"halfTime":{"homeTeam":2,"awayTeam":1},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":1107,"name":"SPAL 2013"},"awayTeam":{"id":102,"name":"Atalanta BC"},"referees":[{"id":11072,"name":"Gianluca Manganiello","nationality":null},{"id":55938,"name":"Domenico Rocca","nationality":null},{"id":57236,"name":"Niccolò Pagliardini","nationality":null},{"id":11032,"name":"Niccolò Baroni","nationality":null},{"id":11053,"name":"Daniele Chiffi","nationality":null}]},{"id":279227,"season":{"id":530,"startDate":"2019-08-24","endDate":"2020-05-24","currentMatchday":32},"utcDate":"2019-08-26T18:45:00Z","status":"FINISHED","matchday":1,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2019-09-26T15:34:12Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"HOME_TEAM","duration":"REGULAR","fullTime":{"homeTeam":4,"awayTeam":0},"halfTime":{"homeTeam":2,"awayTeam":0},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":108,"name":"FC Internazionale Milano"},"awayTeam":{"id":5890,"name":"US Lecce"},"referees":[{"id":10980,"name":"Federico La Penna","nationality":null},{"id":11076,"name":"Alessandro Giallatini","nationality":null},{"id":11023,"name":"Emanuele Prenna","nationality":null},{"id":57830,"name":"Manuel Volpi","nationality":null},{"id":11068,"name":"Gianluca Aureliano","nationality":null}]}]}`
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	var matches business.MatchesResponse
	err := client.Get(
		&matches,
		request.MatchesRequest{
			Filter:      request.Filter{Matchday: 1},
			Competition: 2019,
		},
	)
	assert.Nil(t, err)
	assert.EqualValues(t, 10, matches.Count)
	assert.EqualValues(t, "1", matches.Filters.Matchday)
	assert.EqualValues(t, 2019, matches.Competition.ID)
	assert.EqualValues(t, 2114, matches.Competition.Area.ID)
	timeResult, _ := time.Parse(time.RFC3339, "2020-04-18T00:00:07Z")
	assert.EqualValues(t, timeResult, matches.Competition.LastUpdated)

	assert.EqualValues(t, "Juventus FC", matches.Matches[0].AwayTeam.Name)
	assert.EqualValues(t, "Fabio Maresca", matches.Matches[0].Referees[0].Name)
	assert.EqualValues(t, "AWAY_TEAM", matches.Matches[0].Score.Winner)
	assert.EqualValues(t, 1, matches.Matches[0].Score.FullTime.AwayTeam)
	assert.EqualValues(t, 0, matches.Matches[0].Score.FullTime.HomeTeam)
	assert.EqualValues(t, 32, matches.Matches[0].Season.CurrentMatchday)
}
