package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchHttpsURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/tunnels" {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("content-Type", "application/json")
				io.WriteString(w, `{"tunnels":[{"name":"command_line (http)","uri":"/api/tunnels/command_line+%28http%29","public_url":"http://a6c25881.ngrok.io","proto":"http","config":{"addr":"localhost:80","inspect":true},"metrics":{"conns":{"count":0,"gauge":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0},"http":{"count":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0}}},{"name":"command_line","uri":"/api/tunnels/command_line","public_url":"https://a6c25881.ngrok.io","proto":"https","config":{"addr":"localhost:80","inspect":true},"metrics":{"conns":{"count":0,"gauge":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0},"http":{"count":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0}}}],"uri":"/api/tunnels"}`)
				return
			}
			w.WriteHeader(http.StatusNotFound)
			return
		},
	))
	defer ts.Close()
	ngrok := &Ngrok{BaseURL: ts.URL}
	res, _ := ngrok.FetchHttpsURL()
	if res != "https://a6c25881.ngrok.io" {
		t.Errorf("Expected: https://a6c25881.ngrok.io, Actual: " + res)
	}
}

func TestFetchHttpURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/tunnels" {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("content-Type", "application/json")
				io.WriteString(w, `{"tunnels":[{"name":"command_line (http)","uri":"/api/tunnels/command_line+%28http%29","public_url":"http://a6c25881.ngrok.io","proto":"http","config":{"addr":"localhost:80","inspect":true},"metrics":{"conns":{"count":0,"gauge":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0},"http":{"count":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0}}},{"name":"command_line","uri":"/api/tunnels/command_line","public_url":"https://a6c25881.ngrok.io","proto":"https","config":{"addr":"localhost:80","inspect":true},"metrics":{"conns":{"count":0,"gauge":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0},"http":{"count":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0}}}],"uri":"/api/tunnels"}`)
				return
			}
			w.WriteHeader(http.StatusNotFound)
			return
		},
	))
	defer ts.Close()
	ngrok := &Ngrok{BaseURL: ts.URL}
	res, _ := ngrok.FetchHttpURL()
	if res != "http://a6c25881.ngrok.io" {
		t.Errorf("Expected: http://a6c25881.ngrok.io, Actual: " + res)
	}
}

func TestFetchURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/tunnels" {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("content-Type", "application/json")
				io.WriteString(w, `{"tunnels":[{"name":"command_line (http)","uri":"/api/tunnels/command_line+%28http%29","public_url":"http://a6c25881.ngrok.io","proto":"http","config":{"addr":"localhost:80","inspect":true},"metrics":{"conns":{"count":0,"gauge":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0},"http":{"count":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0}}},{"name":"command_line","uri":"/api/tunnels/command_line","public_url":"https://a6c25881.ngrok.io","proto":"https","config":{"addr":"localhost:80","inspect":true},"metrics":{"conns":{"count":0,"gauge":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0},"http":{"count":0,"rate1":0,"rate5":0,"rate15":0,"p50":0,"p90":0,"p95":0,"p99":0}}}],"uri":"/api/tunnels"}`)
				return
			}
			w.WriteHeader(http.StatusNotFound)
			return
		},
	))
	defer ts.Close()
	ngrok := &Ngrok{BaseURL: ts.URL}
	res, _ := ngrok.FetchURL("https")
	if res != "https://a6c25881.ngrok.io" {
		t.Errorf("Expected: https://a6c25881.ngrok.io, Actual: " + res)
	}
}
