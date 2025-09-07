package trovo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	URL               = "https://open.trovo.live/page/login.html"
	URL_TOKEN_REFRESH = "https://open-api.trovo.live/openplatform/refreshtoken"
	URL_CHANNEL_INFO  = "https://open-api.trovo.live/openplatform/channels/id"
)

var (
	emptyScopeError          = errors.New("empty scope")
	serializeError           = errors.New("serialize error")
	writeAccessError         = errors.New("write access error")
	writeRefreshError        = errors.New("write refresh error")
	jsonSerializationError   = errors.New("json serialization error")
	creationRequestError     = errors.New("creation request error")
	jsonDeserializationError = errors.New("json deserialization error")
)

func NewTrovoClient() *ClientTrovo {
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
		DisableKeepAlives:   false,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &ClientTrovo{
		Client: &http.Client{
			Timeout:   10 * time.Second,
			Transport: transport,
		},
	}
}

func (clientTrovo *ClientTrovo) ConfigureAuthorizationURL(ClientID, ResponseType string, Scope []string, RedirectURL string) (formattedURL string, err error) {
	var fullScope string

	if len(Scope) == 0 {
		return "", emptyScopeError
	}

	fullScope = strings.Join(Scope, "+")

	formattedURL = fmt.Sprintf("%v?client_id=%s&response_type=%s&scope=%s&redirect_uri=%s", URL, ClientID, ResponseType, fullScope, RedirectURL)
	return formattedURL, nil
}

func (clientTrovo *ClientTrovo) RefreshAccessToken(accessToken, refreshToken string) (newAccess, nerRefresh string, err error) {
	data := RefreshTokenRequest{
		ClientSecret: os.Getenv("TROVO_CLIENT_SECRET"),
		GrantType:    "refresh_token",
		RefreshToken: os.Getenv("TROVO_REFRESH_TOKEN"),
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return "", "", serializeError
	}

	request, err := http.NewRequest("POST", URL_TOKEN_REFRESH, bytes.NewBuffer(dataJSON))
	if err != nil {
		return "", "", err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("client-id", os.Getenv("TROVO_CLIENT_ID"))

	response, err := clientTrovo.Client.Do(request)
	if err != nil {
		fmt.Println(response.StatusCode)
		return "", "", err
	}

	defer response.Body.Close()

	var responseJSON ResponseTokenRequest

	if err := json.NewDecoder(response.Body).Decode(&responseJSON); err != nil {
		return "", "", err
	}

	if responseJSON.AccessToken != accessToken {
		return "", "", writeAccessError
	}

	if responseJSON.RefreshToken != refreshToken {
		return "", "", writeRefreshError
	}

	return responseJSON.AccessToken, responseJSON.RefreshToken, nil
}

func (clientTrovo *ClientTrovo) ChannelByID(channelID string) (*ChannelInfo, error) {
	body := ChannelInfoByID{channelID}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	request, err := http.NewRequest("POST", URL_CHANNEL_INFO, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Client-ID", os.Getenv("TROVO_CLIENT_ID"))

	response, err := clientTrovo.Client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	if response == nil {
		return nil, errors.New("response is nil")
	}

	defer response.Body.Close()

	var responseJSON ChannelInfo

	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	if err != nil {
		return nil, jsonDeserializationError
	}

	return &responseJSON, nil
}

func (clientTrovo *ClientTrovo) ChannelByUsername(username string) (*ChannelInfo, error) {
	body := ChannelInfoByUsername{username}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	request, err := http.NewRequest("POST", URL_CHANNEL_INFO, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Client-ID", os.Getenv("TROVO_CLIENT_ID"))

	response, err := clientTrovo.Client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	if response == nil {
		return nil, errors.New("response is nil")
	}

	defer response.Body.Close()

	var responseJSON ChannelInfo

	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	if err != nil {
		fmt.Println(err)
		return nil, jsonDeserializationError
	}

	return &responseJSON, nil
}
