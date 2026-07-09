package os

import (
	"sort"
	"sync"

	pluginhost "github.com/togo-framework/plugin-host"
	"github.com/togo-framework/togo"
)

// WindowSpec is an OS app's default window geometry.
type WindowSpec struct {
	Width     int  `json:"width"`
	Height    int  `json:"height"`
	Resizable bool `json:"resizable"`
}

// AppMeta is the desktop-facing metadata an "OS app" plugin advertises in
// addition to its normal togo.RegisterProviderFunc registration. Call
// RegisterApp from the app plugin's own init(), alongside its provider
// registration — the same self-registration idiom togo.Register uses.
type AppMeta struct {
	Slug     string     `json:"slug"`
	Name     string     `json:"name"`
	Icon     string     `json:"icon"`
	Color    string     `json:"color"`
	Category string     `json:"category"`
	Window   WindowSpec `json:"window"`
}

// App is an AppMeta joined with its runtime enabled state.
type App struct {
	AppMeta
	Enabled bool `json:"enabled"`
}

var (
	appsMu sync.Mutex
	apps   []AppMeta
)

// RegisterApp adds an OS app's desktop metadata to the registry.
func RegisterApp(m AppMeta) {
	appsMu.Lock()
	defer appsMu.Unlock()
	apps = append(apps, m)
}

// Apps returns every registered OS app, sorted by slug, joined with
// plugin-host's enabled state when plugin-host is installed on k (an app
// with no matching plugin-host record is treated as enabled by default).
func Apps(k *togo.Kernel) []App {
	appsMu.Lock()
	metas := make([]AppMeta, len(apps))
	copy(metas, apps)
	appsMu.Unlock()

	sort.Slice(metas, func(i, j int) bool { return metas[i].Slug < metas[j].Slug })

	host, hosted := pluginhost.FromKernel(k)
	out := make([]App, 0, len(metas))
	for _, m := range metas {
		enabled := true
		if hosted {
			if _, ok := host.Get(m.Slug); ok {
				enabled = host.IsEnabled(m.Slug)
			}
		}
		out = append(out, App{AppMeta: m, Enabled: enabled})
	}
	return out
}

// AppBySlug returns one registered app by slug.
func AppBySlug(k *togo.Kernel, slug string) (App, bool) {
	for _, a := range Apps(k) {
		if a.Slug == slug {
			return a, true
		}
	}
	return App{}, false
}
