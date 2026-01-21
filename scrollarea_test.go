package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestScrollArea(t *testing.T) {
	t.Run("renders with default vertical orientation", func(t *testing.T) {
		got := render(t, ui.ScrollArea(ui.ScrollAreaProps{}, Text("Content")))
		want := `<div data-slot="scroll-area" data-orientation="vertical" class="relative overflow-y-auto overflow-x-hidden">Content</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with horizontal orientation", func(t *testing.T) {
		got := render(t, ui.ScrollArea(ui.ScrollAreaProps{Orientation: ui.ScrollAreaOrientationHorizontal}, Text("Content")))
		want := `<div data-slot="scroll-area" data-orientation="horizontal" class="relative overflow-x-auto overflow-y-hidden">Content</div>`
		is.Equal(t, want, got)
	})
}

func TestScrollAreaViewport(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ScrollAreaViewport(ui.ScrollAreaViewportProps{}, Text("Content")))
		want := `<div data-slot="scroll-area-viewport" class="size-full rounded-[inherit]">Content</div>`
		is.Equal(t, want, got)
	})
}

func TestScrollBar(t *testing.T) {
	t.Run("renders with default vertical orientation", func(t *testing.T) {
		got := render(t, ui.ScrollBar(ui.ScrollBarProps{}))
		want := `<div data-slot="scroll-area-scrollbar" data-orientation="vertical" class="flex touch-none p-px transition-colors select-none h-full w-2.5 border-l border-l-transparent"><div data-slot="scroll-area-thumb" class="bg-border relative flex-1 rounded-full"></div></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with horizontal orientation", func(t *testing.T) {
		got := render(t, ui.ScrollBar(ui.ScrollBarProps{Orientation: ui.ScrollAreaOrientationHorizontal}))
		want := `<div data-slot="scroll-area-scrollbar" data-orientation="horizontal" class="flex touch-none p-px transition-colors select-none h-2.5 flex-col border-t border-t-transparent"><div data-slot="scroll-area-thumb" class="bg-border relative flex-1 rounded-full"></div></div>`
		is.Equal(t, want, got)
	})
}
