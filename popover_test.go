package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestPopover(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Popover(ui.PopoverProps{}))
		want := `<div data-slot="popover" class="relative inline-block"></div>`
		is.Equal(t, want, got)
	})
}

func TestPopoverTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.PopoverTrigger(ui.PopoverTriggerProps{}, Text("Click me")))
		want := `<span data-slot="popover-trigger" class="">Click me</span>`
		is.Equal(t, want, got)
	})
}

func TestPopoverContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.PopoverContent(ui.PopoverContentProps{}, Text("Content")))
		want := `<div data-slot="popover-content" class="bg-popover text-popover-foreground absolute z-50 w-72 rounded-md border p-4 shadow-md top-full left-0 mt-1">Content</div>`
		is.Equal(t, want, got)
	})
}
