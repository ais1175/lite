package system

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Release struct {
	TagName string `json:"tag_name"`
}

type VersionStatus struct {
	Current         string    `json:"current"`
	Latest          string    `json:"latest"`
	UpdateAvailable bool      `json:"update_available"`
	LastChecked     time.Time `json:"last_checked"`
}

type Service struct {
	currentVersion string
	latestVersion  string
	lastChecked    time.Time
	mu             sync.RWMutex
}

func NewService(currentVersion string) *Service {
	return &Service{
		currentVersion: currentVersion,
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
	resp, err := http.Get("https://api.github.com/repos/fivemanage/lite/releases")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("github api returned status: %s", resp.Status)
	}

	var releases []Release
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return "", err
	}

	if len(releases) == 0 {
		return "", fmt.Errorf("no releases found")
	}

	return releases[0].TagName, nil
}
