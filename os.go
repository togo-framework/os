// Package os turns a togo app's frontend into a desktop-shell "operating
// system": a login screen, a desktop (wallpaper, icons, dock, top bar), and a
// convention for packaging other togo plugins as installable "OS apps" (icon,
// color, category, default window size) via RegisterApp.
//
// It composes existing togo plugins rather than reinventing them: settings
// stores per-user desktop preferences, plugin-host (if installed) supplies
// enable/disable state for hosted apps, and auth/notification-center are left
// to the consuming app to install — os only reads the acting user off the
// request the same way notification-center does (X-User-Id header or a
// context claim), so it works with any auth driver.
//
//	os, _ := os.FromKernel(k)
//	prefs, _ := os.Prefs(ctx, "user-1")
package os

import (
	"context"
	"net/http"

	"github.com/togo-framework/settings"
	"github.com/togo-framework/togo"
)

// Service is the os plugin runtime stored on the kernel (k.Get("os")).
type Service struct {
	k        *togo.Kernel
	settings *settings.Store
}

func init() {
	togo.RegisterProviderFunc("os", togo.PriorityLate+15, func(k *togo.Kernel) error {
		s := &Service{k: k, settings: settings.FromKernel(k)}
		k.Set("os", s)
		if k.Router != nil {
			s.mountRoutes(k.Router)
		}
		return nil
	})
}

// FromKernel returns the os Service, or nil if the plugin isn't installed.
func FromKernel(k *togo.Kernel) *Service {
	if v, ok := k.Get("os"); ok {
		if s, ok := v.(*Service); ok {
			return s
		}
	}
	return nil
}

// currentUser resolves the acting user from the request — an X-User-Id
// header, or a context claim set by WithUser (same convention as
// notification-center, so both plugins agree on "who is logged in").
func currentUser(req *http.Request) string {
	if v := req.Header.Get("X-User-Id"); v != "" {
		return v
	}
	if v, ok := req.Context().Value(userKey).(string); ok {
		return v
	}
	return ""
}

type ctxKey string

const userKey ctxKey = "os.user"

// WithUser stores the acting user on a context (for non-HTTP callers/tests).
func WithUser(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userKey, userID)
}
