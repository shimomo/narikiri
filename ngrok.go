package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Ngrok struct {
	Tunnels []struct {
		PublicURL string `json:"public_url"`
		Proto     string `json:"proto"`
	} `json:"tunnels"`
}

func (n *Ngrok) FetchHttpsURL() (string, error) {
	return n.FetchURL("https")
}

func (n *Ngrok) FetchHttpURL() (string, error) {
	return n.FetchURL("http")
}

func (n *Ngrok) FetchURL(proto string) (string, error) {
	var url string
	res, err := http.Get("http://localhost:4040/api/tunnels")
	if err != nil {
		return url, err
	}
	defer res.Body.Close()
	con, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(con, &n)
	for _, t := range n.Tunnels {
		if proto == t.Proto {
			url = t.PublicURL
		}
	}
	return url, err
}
