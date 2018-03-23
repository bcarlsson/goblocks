package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davidscholberg/go-i3barjson"
)

// Bitstamp represents the configuration for the bitstamp display block.
type Bitstamp struct {
	BlockConfigBase `yaml:",inline"`
	Currency        string `yaml:"currency"`
}

// Response struct is used to parse JSON response from Bitstamp.
type Response struct {
	High      string `json:"high"`
	Last      string `json:"last"`
	Timestamp string `json:"timestamp"`
	Bid       string `json:"bid"`
	Vwap      string `json:"vwap"`
	Volume    string `json:"volume"`
	Low       string `json:"low"`
	Ask       string `json:"ask"`
	Open      string `json:"open"`
}

// UpdateBlock updates the bitstamp block.
func (c Bitstamp) UpdateBlock(b *i3barjson.Block) {
	b.Color = c.Color
	fullTextFmt := fmt.Sprintf("%s%%s", c.Label)

	bs := &Response{}
	url := "https://www.bitstamp.net/api/v2/ticker/" + c.Currency + "usd"
	resp, err := http.Get(url)

	if err != nil {
		b.Urgent = true
		b.FullText = fmt.Sprintf(fullTextFmt, err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		b.Urgent = true
		b.FullText = fmt.Sprintf(fullTextFmt, err.Error())
		return
	}

	err = json.Unmarshal(body, bs)
	if err != nil {
		b.Urgent = true
		b.FullText = fmt.Sprintf(fullTextFmt, err.Error())
		return
	}

	b.FullText = fmt.Sprintf(fullTextFmt, "$"+bs.Last)
}
