package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestDialog(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Dialog(ui.DialogProps{}))
		want := `<div data-slot="dialog" class=""></div>`
		is.Equal(t, want, got)
	})
}

func TestDialogTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DialogTrigger(ui.DialogTriggerProps{}, Text("Open")))
		want := `<span data-slot="dialog-trigger" class="">Open</span>`
		is.Equal(t, want, got)
	})
}

func TestDialogOverlay(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DialogOverlay(ui.DialogOverlayProps{}))
		want := `<div data-slot="dialog-overlay" class="fixed inset-0 z-50 bg-black/50"></div>`
		is.Equal(t, want, got)
	})
}

func TestDialogContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DialogContent(ui.DialogContentProps{}, Text("Content")))
		want := `<div role="dialog" aria-modal="true" data-slot="dialog-content" class="bg-background fixed top-1/2 left-1/2 z-50 grid w-full max-w-[calc(100%-2rem)] -translate-x-1/2 -translate-y-1/2 gap-4 rounded-lg border p-6 shadow-lg sm:max-w-lg">Content</div>`
		is.Equal(t, want, got)
	})
}

func TestDialogHeader(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DialogHeader(ui.DialogHeaderProps{}, Text("Header")))
		want := `<div data-slot="dialog-header" class="flex flex-col gap-2 text-center sm:text-left">Header</div>`
		is.Equal(t, want, got)
	})
}

func TestDialogFooter(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DialogFooter(ui.DialogFooterProps{}, Text("Footer")))
		want := `<div data-slot="dialog-footer" class="flex flex-col-reverse gap-2 sm:flex-row sm:justify-end">Footer</div>`
		is.Equal(t, want, got)
	})
}

func TestDialogTitle(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DialogTitle(ui.DialogTitleProps{}, Text("Title")))
		want := `<h2 data-slot="dialog-title" class="text-lg leading-none font-semibold">Title</h2>`
		is.Equal(t, want, got)
	})
}

func TestDialogDescription(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DialogDescription(ui.DialogDescriptionProps{}, Text("Description")))
		want := `<p data-slot="dialog-description" class="text-muted-foreground text-sm">Description</p>`
		is.Equal(t, want, got)
	})
}

func TestDialogClose(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DialogClose(ui.DialogCloseProps{}, Text("×")))
		want := `<button type="button" data-slot="dialog-close" class="ring-offset-background focus:ring-ring absolute top-4 right-4 rounded-sm opacity-70 transition-opacity hover:opacity-100 focus:ring-2 focus:ring-offset-2 focus:outline-hidden disabled:pointer-events-none cursor-pointer">×</button>`
		is.Equal(t, want, got)
	})
}
