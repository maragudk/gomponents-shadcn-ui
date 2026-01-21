package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestHoverCard(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.HoverCard(ui.HoverCardProps{}))
		want := `<div data-slot="hover-card" class="relative inline-flex group"></div>`
		is.Equal(t, want, got)
	})
}

func TestHoverCardTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.HoverCardTrigger(ui.HoverCardTriggerProps{}, Text("Hover me")))
		want := `<span data-slot="hover-card-trigger" class="">Hover me</span>`
		is.Equal(t, want, got)
	})
}

func TestHoverCardContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.HoverCardContent(ui.HoverCardContentProps{}, Text("Card content")))
		want := `<div data-slot="hover-card-content" class="bg-popover text-popover-foreground absolute z-50 hidden group-hover:block w-64 rounded-md border p-4 shadow-md top-full left-0 mt-2">Card content</div>`
		is.Equal(t, want, got)
	})
}
