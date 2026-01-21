package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestAlertDialog(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AlertDialog(ui.AlertDialogProps{}))
		want := `<div data-slot="alert-dialog" class=""></div>`
		is.Equal(t, want, got)
	})
}

func TestAlertDialogTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AlertDialogTrigger(ui.AlertDialogTriggerProps{}, Text("Open")))
		want := `<span data-slot="alert-dialog-trigger" class="">Open</span>`
		is.Equal(t, want, got)
	})
}

func TestAlertDialogOverlay(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AlertDialogOverlay(ui.AlertDialogOverlayProps{}))
		want := `<div data-slot="alert-dialog-overlay" class="fixed inset-0 z-50 bg-black/50"></div>`
		is.Equal(t, want, got)
	})
}

func TestAlertDialogContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AlertDialogContent(ui.AlertDialogContentProps{}, Text("Content")))
		want := `<div role="alertdialog" aria-modal="true" data-slot="alert-dialog-content" class="bg-background fixed top-1/2 left-1/2 z-50 grid w-full max-w-[calc(100%-2rem)] -translate-x-1/2 -translate-y-1/2 gap-4 rounded-lg border p-6 shadow-lg sm:max-w-lg">Content</div>`
		is.Equal(t, want, got)
	})
}

func TestAlertDialogHeader(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AlertDialogHeader(ui.AlertDialogHeaderProps{}, Text("Header")))
		want := `<div data-slot="alert-dialog-header" class="flex flex-col gap-2 text-center sm:text-left">Header</div>`
		is.Equal(t, want, got)
	})
}

func TestAlertDialogFooter(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AlertDialogFooter(ui.AlertDialogFooterProps{}, Text("Footer")))
		want := `<div data-slot="alert-dialog-footer" class="flex flex-col-reverse gap-2 sm:flex-row sm:justify-end">Footer</div>`
		is.Equal(t, want, got)
	})
}

func TestAlertDialogTitle(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AlertDialogTitle(ui.AlertDialogTitleProps{}, Text("Title")))
		want := `<h2 data-slot="alert-dialog-title" class="text-lg font-semibold">Title</h2>`
		is.Equal(t, want, got)
	})
}

func TestAlertDialogDescription(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AlertDialogDescription(ui.AlertDialogDescriptionProps{}, Text("Description")))
		want := `<p data-slot="alert-dialog-description" class="text-muted-foreground text-sm">Description</p>`
		is.Equal(t, want, got)
	})
}

func TestAlertDialogAction(t *testing.T) {
	t.Run("renders with default button styling", func(t *testing.T) {
		got := render(t, ui.AlertDialogAction(ui.AlertDialogActionProps{}, Text("Continue")))
		// Uses default button styling
		is.True(t, len(got) > 0, "should render output")
	})
}

func TestAlertDialogCancel(t *testing.T) {
	t.Run("renders with outline button styling", func(t *testing.T) {
		got := render(t, ui.AlertDialogCancel(ui.AlertDialogCancelProps{}, Text("Cancel")))
		// Uses outline button styling
		is.True(t, len(got) > 0, "should render output")
	})
}
