package ui_test

import (
	"testing"

	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestSeparator(t *testing.T) {
	t.Run("renders horizontal by default", func(t *testing.T) {
		got := render(t, ui.Separator(ui.SeparatorProps{}))
		want := `<div role="separator" class="bg-border shrink-0 h-px w-full"></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders horizontal orientation", func(t *testing.T) {
		got := render(t, ui.Separator(ui.SeparatorProps{Orientation: ui.SeparatorOrientationHorizontal}))
		want := `<div role="separator" class="bg-border shrink-0 h-px w-full"></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders vertical orientation", func(t *testing.T) {
		got := render(t, ui.Separator(ui.SeparatorProps{Orientation: ui.SeparatorOrientationVertical}))
		want := `<div role="separator" class="bg-border shrink-0 h-full w-px"></div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.Separator(ui.SeparatorProps{}, Class("my-4")))
		want := `<div role="separator" class="bg-border shrink-0 h-px w-full my-4"></div>`
		is.Equal(t, want, got)
	})
}
