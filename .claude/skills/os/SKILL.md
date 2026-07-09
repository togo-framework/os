---
name: os
description: Work with the os desktop-shell plugin and package togo plugins as installable OS apps.
---

# os plugin

Turns the app's frontend into a desktop shell (login screen, wallpaper, desktop
icons, dock, top bar) built on `@togo-framework/ui`'s `DesktopShell`/`Window`
components (`web/` in this repo, injected into the consuming app's frontend by
`togo install`).

## Backend

- `os.FromKernel(k)` → `*os.Service`.
- `GET/PUT /api/os/session` — the current user's `DesktopPrefs` (theme, accent,
  wallpaper, dock pins, icon positions), stored via the `settings` plugin.
- `GET /api/os/apps`, `GET /api/os/apps/{slug}` — registered OS apps
  (`os.AppMeta` joined with `plugin-host`'s enabled state, if installed).

## Packaging a plugin as an OS app

Scaffold one with `togo make:os-app <name> --icon <icon> --color "#rrggbb"
--category <category>`. It generates a normal togo plugin (provider + web +
MCP + `.claude/agents`) whose `init()` additionally calls:

```go
os.RegisterApp(os.AppMeta{
    Slug: "<name>", Name: "<Name>", Icon: "<icon>", Color: "#rrggbb",
    Category: "<category>", Window: os.WindowSpec{Width: 800, Height: 600, Resizable: true},
})
```

That's the only thing that makes a plugin "OS-aware" — everything else
(backend routes, MCP tools, `.claude/agents`+`.claude/skills`) is a normal
plugin, installed and injected the usual way (`togo install <owner>/<repo>`).
