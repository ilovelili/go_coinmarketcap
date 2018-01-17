package coinmarketcap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strconv"
	"time"
)

// Ticker ticker model
type Ticker struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Symbol               string `json:"symbol"`
	Rank                 string `json:"rank"`
	PriceUSD             string `json:"price_usd"`
	PriceBTC             string `json:"price_btc"`
	DailyVolumeUSD       string `json:"24h_volume_usd"`
	PercentChangeOneHour string `json:"percent_change_1h"`
	PercentChangeOneDay  string `json:"percent_change_24h"`
	PercentChangeOneWeek string `json:"percent_change_7d"`
	LastUpdated          string `json:"last_updated"`
}

// Tickers Alias of []Ticker
type Tickers []*Ticker

const (
	// BaseURL coinmarketcap API base URL
	BaseURL = "https://api.coinmarketcap.com/v1/"
)

// Client struct represents CoinMarketCap API client.
type Client struct {
	accessKey       string
	secretAccessKey string
	BaseURL         *url.URL
	HTTPClient      *http.Client
}

// Len Sort interface Len
func (t Tickers) Len() int {
	return len(t)
}

// Swap Sort interface Swap
func (t Tickers) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Less Sort interface Less => Sort by percent change in one hour
func (t Tickers) Less(i, j int) bool {
	percentchangeinanhourI, _ := strconv.ParseFloat(t[i].PercentChangeOneHour, 64)
	percentchangeinanhourJ, _ := strconv.ParseFloat(t[j].PercentChangeOneHour, 64)

	return percentchangeinanhourI > percentchangeinanhourJ
}

// MarshalJSON customize MarshalJSON http://choly.ca/post/go-json-marshalling/
func (t *Ticker) MarshalJSON() ([]byte, error) {
	type Alias Ticker
	return json.Marshal(&struct {
		LastUpdated time.Time `json:"last_updated"`
		*Alias
	}{
		LastUpdated: toTime(t.LastUpdated),
		Alias:       (*Alias)(t),
	})
}

// NewClient Init Client
func NewClient(key, secretKey string) (*Client, error) {
	baseurl, _ := url.Parse(BaseURL)
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	cli := &Client{accessKey: key, secretAccessKey: secretKey, BaseURL: baseurl, HTTPClient: client}
	return cli, nil
}

// GetTicker get ticker by id
func (cli *Client) GetTicker(ctx context.Context, id string) (*Ticker, error) {
	req, err := cli.newRequest(ctx, http.MethodGet, fmt.Sprintf("/ticker/%s", id), []byte(""))
	if err != nil {
		return nil, err
	}

	res, err := cli.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var tickers []Ticker
	err = decodeBody(res, &tickers)
	if err != nil {
		return nil, err
	}

	return &tickers[0], nil
}

// GetTickers get tickers
func (cli *Client) GetTickers(ctx context.Context) ([]*Ticker, error) {
	req, err := cli.newRequest(ctx, http.MethodGet, "/ticker", []byte(""))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	// yeah I like hard code
	q.Add("limit", "1000")
	req.URL.RawQuery = q.Encode()

	res, err := cli.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var tickers []*Ticker
	err = decodeBody(res, &tickers)
	if err != nil {
		return nil, err
	}

	// sort by change percent in one hour
	sort.Sort(Tickers(tickers))
	return tickers, nil
}

// newRequest create new request
func (cli *Client) newRequest(ctx context.Context, method, endpoint string, body []byte) (*http.Request, error) {
	u := *cli.BaseURL
	u.Path = path.Join(cli.BaseURL.Path, endpoint)
	fmt.Println(u.String())
	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	return req, nil
}

// decodeBody decode response body
func decodeBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(out)
}

// encodeBody encode body
func encodeBody(in interface{}) ([]byte, error) {
	return json.Marshal(in)
}

func toTime(unixtimestamp string) time.Time {
	i, _ := strconv.ParseInt(unixtimestamp, 10, 64)
	return time.Unix(i, 0)
}

