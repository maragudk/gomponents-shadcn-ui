package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestTooltip(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Tooltip(ui.TooltipProps{}))
		want := `<div data-slot="tooltip" class="relative inline-flex group"></div>`
		is.Equal(t, want, got)
	})
}

func TestTooltipTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.TooltipTrigger(ui.TooltipTriggerProps{}, Text("Hover me")))
		want := `<span data-slot="tooltip-trigger" class="">Hover me</span>`
		is.Equal(t, want, got)
	})
}

func TestTooltipContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.TooltipContent(ui.TooltipContentProps{}, Text("Tooltip text")))
		want := `<div role="tooltip" data-slot="tooltip-content" class="bg-foreground text-background absolute z-50 hidden group-hover:block w-max rounded-md px-3 py-1.5 text-xs bottom-full left-1/2 -translate-x-1/2 mb-2">Tooltip text</div>`
		is.Equal(t, want, got)
	})
}
