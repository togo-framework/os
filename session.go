package os

import (
	"context"

	"github.com/togo-framework/settings"
)

// IconPos is a desktop icon's saved (x, y) position.
type IconPos struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// DesktopPrefs is one user's desktop personalization state, persisted through
// the settings plugin (scope "os:desktop", key = user id).
type DesktopPrefs struct {
	Theme         string             `json:"theme"`
	Accent        string             `json:"accent"`
	Wallpaper     string             `json:"wallpaper"`
	LockWallpaper string             `json:"lock_wallpaper"`
	DockPinned    []string           `json:"dock_pinned"`
	DesktopHidden []string           `json:"desktop_hidden"`
	IconPositions map[string]IconPos `json:"icon_positions"`
}

const desktopScope = "os:desktop"

func defaultPrefs() DesktopPrefs {
	return DesktopPrefs{
		Theme:         "dark",
		Accent:        "purple",
		Wallpaper:     "aurora",
		LockWallpaper: "monterey",
		DockPinned:    []string{},
		DesktopHidden: []string{},
		IconPositions: map[string]IconPos{},
	}
}

// Prefs returns the saved desktop preferences for userID, or sensible
// defaults if none have been saved yet (or the settings plugin isn't
// installed).
func (s *Service) Prefs(ctx context.Context, userID string) DesktopPrefs {
	if s.settings == nil || userID == "" {
		return defaultPrefs()
	}
	if v, ok, err := settings.Get[DesktopPrefs](ctx, s.settings, desktopScope, userID); err == nil && ok {
		return v
	}
	return defaultPrefs()
}

// SavePrefs persists userID's desktop preferences.
func (s *Service) SavePrefs(ctx context.Context, userID string, p DesktopPrefs) error {
	if s.settings == nil || userID == "" {
		return nil
	}
	return settings.Set(ctx, s.settings, desktopScope, userID, p)
}
