package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestDrawer(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Drawer(ui.DrawerProps{}))
		want := `<div data-slot="drawer" class=""></div>`
		is.Equal(t, want, got)
	})
}

func TestDrawerTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DrawerTrigger(ui.DrawerTriggerProps{}, Text("Open")))
		want := `<span data-slot="drawer-trigger" class="">Open</span>`
		is.Equal(t, want, got)
	})
}

func TestDrawerOverlay(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DrawerOverlay(ui.DrawerOverlayProps{}))
		want := `<div data-slot="drawer-overlay" class="fixed inset-0 z-50 bg-black/50"></div>`
		is.Equal(t, want, got)
	})
}

func TestDrawerContent(t *testing.T) {
	t.Run("renders with default bottom direction", func(t *testing.T) {
		got := render(t, ui.DrawerContent(ui.DrawerContentProps{}, Text("Content")))
		want := `<div role="dialog" aria-modal="true" data-slot="drawer-content" data-direction="bottom" class="bg-background fixed z-50 flex h-auto flex-col inset-x-0 bottom-0 max-h-[80vh] rounded-t-lg border-t">Content</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with top direction", func(t *testing.T) {
		got := render(t, ui.DrawerContent(ui.DrawerContentProps{Direction: ui.DrawerDirectionTop}, Text("Content")))
		want := `<div role="dialog" aria-modal="true" data-slot="drawer-content" data-direction="top" class="bg-background fixed z-50 flex h-auto flex-col inset-x-0 top-0 max-h-[80vh] rounded-b-lg border-b">Content</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with left direction", func(t *testing.T) {
		got := render(t, ui.DrawerContent(ui.DrawerContentProps{Direction: ui.DrawerDirectionLeft}, Text("Content")))
		want := `<div role="dialog" aria-modal="true" data-slot="drawer-content" data-direction="left" class="bg-background fixed z-50 flex h-auto flex-col inset-y-0 left-0 w-3/4 sm:max-w-sm border-r">Content</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with right direction", func(t *testing.T) {
		got := render(t, ui.DrawerContent(ui.DrawerContentProps{Direction: ui.DrawerDirectionRight}, Text("Content")))
		want := `<div role="dialog" aria-modal="true" data-slot="drawer-content" data-direction="right" class="bg-background fixed z-50 flex h-auto flex-col inset-y-0 right-0 w-3/4 sm:max-w-sm border-l">Content</div>`
		is.Equal(t, want, got)
	})
}

func TestDrawerHeader(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DrawerHeader(ui.DrawerHeaderProps{}, Text("Header")))
		want := `<div data-slot="drawer-header" class="flex flex-col gap-1.5 p-4 text-center sm:text-left">Header</div>`
		is.Equal(t, want, got)
	})
}

func TestDrawerFooter(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DrawerFooter(ui.DrawerFooterProps{}, Text("Footer")))
		want := `<div data-slot="drawer-footer" class="mt-auto flex flex-col gap-2 p-4">Footer</div>`
		is.Equal(t, want, got)
	})
}

func TestDrawerTitle(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DrawerTitle(ui.DrawerTitleProps{}, Text("Title")))
		want := `<h2 data-slot="drawer-title" class="text-foreground font-semibold">Title</h2>`
		is.Equal(t, want, got)
	})
}

func TestDrawerDescription(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DrawerDescription(ui.DrawerDescriptionProps{}, Text("Description")))
		want := `<p data-slot="drawer-description" class="text-muted-foreground text-sm">Description</p>`
		is.Equal(t, want, got)
	})
}

func TestDrawerClose(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DrawerClose(ui.DrawerCloseProps{}, Text("×")))
		want := `<button type="button" data-slot="drawer-close" class="ring-offset-background focus:ring-ring absolute top-4 right-4 rounded-sm opacity-70 transition-opacity hover:opacity-100 focus:ring-2 focus:ring-offset-2 focus:outline-hidden disabled:pointer-events-none cursor-pointer">×</button>`
		is.Equal(t, want, got)
	})
}

func TestDrawerHandle(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DrawerHandle(ui.DrawerHandleProps{}))
		want := `<div data-slot="drawer-handle" class="bg-muted mx-auto mt-4 h-2 w-[100px] shrink-0 rounded-full"></div>`
		is.Equal(t, want, got)
	})
}
