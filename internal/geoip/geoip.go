package geoip

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/movableink/sre-assignment-golang/internal/config"
)

// Response represents the GeoIP API response
type Response struct {
	IP          string  `json:"ip"`
	Location    string  `json:"location"`
	PostalCode  string  `json:"postal_code"`
	NetworkName string  `json:"network_name"`
	Domain      string  `json:"domain"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

// Service handles GeoIP lookups
type Service struct {
	client  *http.Client
	config  *config.Config
}

// New creates a new GeoIP service
func New(cfg *config.Config) *Service {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Service{
		client: client,
		config: cfg,
	}
}

// LookupIP looks up geographical information for an IP address
func (s *Service) LookupIP(ip string) (*Response, error) {
	url := fmt.Sprintf("%s/geoip/%s", s.config.APIURL, ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	if s.config.APIToken != "" {
		req.Header.Set("Authorization", "Bearer " + s.config.APIToken)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed: %s", string(body))
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return &result, nil
}
