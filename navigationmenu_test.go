package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestNavigationMenu(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.NavigationMenu(ui.NavigationMenuProps{}))
		want := `<nav data-slot="navigation-menu" class="relative flex max-w-max flex-1 items-center justify-center"></nav>`
		is.Equal(t, want, got)
	})
}

func TestNavigationMenuList(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.NavigationMenuList(ui.NavigationMenuListProps{}))
		want := `<ul data-slot="navigation-menu-list" class="group flex flex-1 list-none items-center justify-center gap-1"></ul>`
		is.Equal(t, want, got)
	})
}

func TestNavigationMenuItem(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.NavigationMenuItem(ui.NavigationMenuItemProps{}))
		want := `<li data-slot="navigation-menu-item" class="relative"></li>`
		is.Equal(t, want, got)
	})
}

func TestNavigationMenuTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.NavigationMenuTrigger(ui.NavigationMenuTriggerProps{}, Text("Getting Started")))
		want := `<button type="button" data-slot="navigation-menu-trigger" class="group inline-flex h-9 w-max items-center justify-center rounded-md bg-background px-4 py-2 text-sm font-medium hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=open]:bg-accent/50 data-[state=open]:text-accent-foreground outline-none transition-colors cursor-pointer">Getting Started</button>`
		is.Equal(t, want, got)
	})
}

func TestNavigationMenuContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.NavigationMenuContent(ui.NavigationMenuContentProps{}))
		want := `<div data-slot="navigation-menu-content" class="bg-popover text-popover-foreground absolute top-full left-0 z-50 mt-1.5 w-auto overflow-hidden rounded-md border p-2 shadow"></div>`
		is.Equal(t, want, got)
	})
}

func TestNavigationMenuLink(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.NavigationMenuLink(ui.NavigationMenuLinkProps{}, Text("Introduction")))
		want := `<a data-slot="navigation-menu-link" class="data-[active=true]:bg-accent/50 data-[active=true]:text-accent-foreground hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground flex flex-col gap-1 rounded-sm p-2 text-sm transition-colors outline-none">Introduction</a>`
		is.Equal(t, want, got)
	})

	t.Run("renders with active state", func(t *testing.T) {
		got := render(t, ui.NavigationMenuLink(ui.NavigationMenuLinkProps{Active: true}, Text("Introduction")))
		want := `<a data-slot="navigation-menu-link" data-active="true" class="data-[active=true]:bg-accent/50 data-[active=true]:text-accent-foreground hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground flex flex-col gap-1 rounded-sm p-2 text-sm transition-colors outline-none">Introduction</a>`
		is.Equal(t, want, got)
	})
}

func TestNavigationMenuIndicator(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.NavigationMenuIndicator(ui.NavigationMenuIndicatorProps{}))
		want := `<div data-slot="navigation-menu-indicator" class="top-full z-[1] flex h-1.5 items-end justify-center overflow-hidden"><div class="bg-border relative top-[60%] h-2 w-2 rotate-45 rounded-tl-sm shadow-md"></div></div>`
		is.Equal(t, want, got)
	})
}

func TestNavigationMenuViewport(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.NavigationMenuViewport(ui.NavigationMenuViewportProps{}))
		want := `<div data-slot="navigation-menu-viewport" class="bg-popover text-popover-foreground absolute top-full left-0 z-50 mt-1.5 overflow-hidden rounded-md border shadow"></div>`
		is.Equal(t, want, got)
	})
}
