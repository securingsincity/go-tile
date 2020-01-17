package main

import (
	"io"
	"net/http"
	"github.com/google/uuid"
)

const API_URL_SCAFFOLD = "https://production.tile-api.com/api/v1"
const DEFAULT_APP_ID = "ios-tile-production"
const DEFAULT_APP_VERSION = "2.55.1.3707"
const DEFAULT_LOCALE = "en-US"

func main() {
	http.Get("")
	tile := CreateTileClient("","")
	tile.Login()
	// tile.All()
}

func CreateTileClient(email string, password string) tileClient {
	return tileClient{
		email:    email,
		password: password,
		clientUUID: uuid.
	}
}

type tileClient struct {
	email      string
	password   string
	clientUUID string
}

func (t tileClient) Login() string {
	return ""
}

func (t tileClient) request(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Tile_app_id", DEFAULT_APP_ID)
	req.Header.Add("Tile_app_version", DEFAULT_APP_VERSION)
	req.Header.Add("Tile_client_uuid", t.clientUUID)
	return http.DefaultClient.Do(req)
}

// func request()
