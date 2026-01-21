package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestCollapsible(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Collapsible(ui.CollapsibleProps{}))
		want := `<div data-slot="collapsible" class=""></div>`
		is.Equal(t, want, got)
	})
}

func TestCollapsibleTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.CollapsibleTrigger(ui.CollapsibleTriggerProps{}, Text("Toggle")))
		want := `<button type="button" data-slot="collapsible-trigger" class="cursor-pointer">Toggle</button>`
		is.Equal(t, want, got)
	})
}

func TestCollapsibleContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.CollapsibleContent(ui.CollapsibleContentProps{}, Text("Content")))
		want := `<div data-slot="collapsible-content" class="overflow-hidden">Content</div>`
		is.Equal(t, want, got)
	})
}
