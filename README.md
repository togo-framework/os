<!-- togo-header -->
<div align="center">
  <picture><source media="(prefers-color-scheme: dark)" srcset=".github/assets/togo-mark-dark.svg" /><img src=".github/assets/togo-mark.svg" alt="ToGO" height="64" /></picture>
  <h1>togo-framework/os</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1F8A99" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/os"><img src="https://pkg.go.dev/badge/github.com/togo-framework/os.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/os
```

<!-- /togo-header -->

# os — an OS-style desktop shell for togo apps

Turns a togo app's web frontend into a **macOS-style desktop environment**: a
login/lock screen, a wallpapered desktop with draggable icons, a dock, draggable
/ resizable / snappable windows, a Launchpad app directory, Spotlight search, a
notification center, and a Settings app with live theming. Any togo plugin can
register itself as a launchable **OS app** (icon, color, category, window spec).

The shell is fully responsive — on phones windows go full-screen and the UI
drills in iOS-style.

## Register an app

Each OS-app plugin calls `os.RegisterApp` from its `init()`, alongside its normal
provider registration:

```go
func init() {
	os.RegisterApp(os.AppMeta{
		Slug:     "notes",
		Name:     "Notes",
		Icon:     "sticky-note",           // lucide icon (kebab-case)
		Color:    "#f59e0b",
		Category: "productivity",
		Window:   os.WindowSpec{Width: 720, Height: 520, Resizable: true},
	})
}
```

Scaffold a full app bundle (UI + backend + MCP + Claude agent + this metadata)
with `togo make:os-app <name>`, or start a whole project on the desktop shell
with `togo new <app> --frontend os`.

## REST API

| Method | Path | Purpose |
|---|---|---|
| `GET` | `/api/os/session` | Current user's `DesktopPrefs` (theme, wallpaper, dock pins, icon positions, …) |
| `PUT` | `/api/os/session` | Persist `DesktopPrefs` |
| `GET` | `/api/os/apps` | Registered apps (icon/color/category/window + enabled state) |
| `GET` | `/api/os/apps/{slug}` | One app's metadata |

Desktop preferences are stored through the [settings](https://github.com/togo-framework/settings)
plugin (scope `os:desktop`, key = user id); enabled state is joined from
[plugin-host](https://github.com/togo-framework/plugin-host) when present. The
plugin registers at `PriorityLate+15` and reads the current user generically, so
it works with any auth driver.

MIT

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
