package system

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type GHCRToken struct {
	Token string `json:"token"`
}

type GHCRTags struct {
	Tags []string `json:"tags"`
}

type VersionStatus struct {
	Current         string    `json:"current"`
	Latest          string    `json:"latest"`
	UpdateAvailable bool      `json:"update_available"`
	LastChecked     time.Time `json:"last_checked"`
}

type SystemConfig struct {
	BucketDomain string `json:"bucket_domain"`
}

type Service struct {
	currentVersion string
	latestVersion  string
	lastChecked    time.Time
	bucketDomain   string
	mu             sync.RWMutex
}

func NewService(currentVersion string, bucketDomain string) *Service {
	return &Service{
		currentVersion: currentVersion,
		bucketDomain:   bucketDomain,
	}
}

func (s *Service) GetConfig() *SystemConfig {
	return &SystemConfig{
		BucketDomain: s.bucketDomain,
	}
}

func (s *Service) GetVersionStatus() (*VersionStatus, error) {
	s.mu.RLock()
	if time.Since(s.lastChecked) < 12*time.Hour && s.latestVersion != "" {
		defer s.mu.RUnlock()
		return &VersionStatus{
			Current:         s.currentVersion,
			Latest:          s.latestVersion,
			UpdateAvailable: s.currentVersion != s.latestVersion && s.currentVersion != "dev",
			LastChecked:     s.lastChecked,
		}, nil
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()

	if time.Since(s.lastChecked) < 12*time.Hour && s.latestVersion != "" {
		return &VersionStatus{
			Current:         s.currentVersion,
			Latest:          s.latestVersion,
			UpdateAvailable: s.currentVersion != s.latestVersion && s.currentVersion != "dev",
			LastChecked:     s.lastChecked,
		}, nil
	}

	latest, err := s.fetchLatestVersion()
	if err != nil {
		return nil, err
	}

	s.latestVersion = latest
	s.lastChecked = time.Now()

	return &VersionStatus{
		Current:         s.currentVersion,
		Latest:          s.latestVersion,
		UpdateAvailable: s.currentVersion != s.latestVersion && s.currentVersion != "dev",
		LastChecked:     s.lastChecked,
	}, nil
}

func (s *Service) fetchLatestVersion() (string, error) {
	tokenResp, err := http.Get("https://ghcr.io/token?service=ghcr.io&scope=repository:fivemanage/lite:pull")
	if err != nil {
		return "", err
	}
	defer tokenResp.Body.Close()

	var token GHCRToken
	if err := json.NewDecoder(tokenResp.Body).Decode(&token); err != nil {
		return "", err
	}

	req, _ := http.NewRequest("GET", "https://ghcr.io/v2/fivemanage/lite/tags/list", nil)
	req.Header.Set("Authorization", "Bearer "+token.Token)

	client := &http.Client{}
	tagsResp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer tagsResp.Body.Close()

	if tagsResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ghcr api returned status: %s", tagsResp.Status)
	}

	var tags GHCRTags
	if err := json.NewDecoder(tagsResp.Body).Decode(&tags); err != nil {
		return "", err
	}

	if len(tags.Tags) == 0 {
		return "", fmt.Errorf("no tags found")
	}

	for i := len(tags.Tags) - 1; i >= 0; i-- {
		tag := tags.Tags[i]
		if tag != "latest" {
			return tag, nil
		}
	}

	return "latest", nil
}
