package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestToastViewport(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ToastViewport(ui.ToastViewportProps{}))
		want := `<div data-slot="toast-viewport" class="fixed top-0 z-[100] flex max-h-screen w-full flex-col-reverse p-4 sm:bottom-0 sm:right-0 sm:top-auto sm:flex-col md:max-w-[420px]"></div>`
		is.Equal(t, want, got)
	})
}

func TestToast(t *testing.T) {
	t.Run("renders default variant", func(t *testing.T) {
		got := render(t, ui.Toast(ui.ToastProps{}, Text("Toast message")))
		want := `<div role="status" aria-live="polite" data-slot="toast" class="group pointer-events-auto relative flex w-full items-center justify-between space-x-2 overflow-hidden rounded-md border p-4 pr-6 shadow-lg border bg-background text-foreground">Toast message</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders destructive variant", func(t *testing.T) {
		got := render(t, ui.Toast(ui.ToastProps{Variant: ui.ToastVariantDestructive}, Text("Error!")))
		want := `<div role="status" aria-live="polite" data-slot="toast" data-variant="destructive" class="group pointer-events-auto relative flex w-full items-center justify-between space-x-2 overflow-hidden rounded-md border p-4 pr-6 shadow-lg destructive border-destructive bg-destructive text-destructive-foreground">Error!</div>`
		is.Equal(t, want, got)
	})
}

func TestToastAction(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ToastAction(ui.ToastActionProps{}, Text("Undo")))
		want := `<button type="button" data-slot="toast-action" class="inline-flex h-8 shrink-0 items-center justify-center rounded-md border bg-transparent px-3 text-sm font-medium transition-colors hover:bg-secondary focus:outline-none focus:ring-1 focus:ring-ring disabled:pointer-events-none disabled:opacity-50 group-[.destructive]:border-muted/40 group-[.destructive]:hover:border-destructive/30 group-[.destructive]:hover:bg-destructive group-[.destructive]:hover:text-destructive-foreground group-[.destructive]:focus:ring-destructive cursor-pointer">Undo</button>`
		is.Equal(t, want, got)
	})
}

func TestToastClose(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ToastClose(ui.ToastCloseProps{}, Text("×")))
		want := `<button type="button" data-slot="toast-close" class="absolute right-1 top-1 rounded-md p-1 text-foreground/50 opacity-0 transition-opacity hover:text-foreground focus:opacity-100 focus:outline-none focus:ring-1 group-hover:opacity-100 group-[.destructive]:text-red-300 group-[.destructive]:hover:text-red-50 group-[.destructive]:focus:ring-red-400 group-[.destructive]:focus:ring-offset-red-600 cursor-pointer">×</button>`
		is.Equal(t, want, got)
	})
}

func TestToastTitle(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ToastTitle(ui.ToastTitleProps{}, Text("Title")))
		want := `<div data-slot="toast-title" class="text-sm font-semibold [&amp;+div]:text-xs">Title</div>`
		is.Equal(t, want, got)
	})
}

func TestToastDescription(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ToastDescription(ui.ToastDescriptionProps{}, Text("Description")))
		want := `<div data-slot="toast-description" class="text-sm opacity-90">Description</div>`
		is.Equal(t, want, got)
	})
}
