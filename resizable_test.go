package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestResizablePanelGroup(t *testing.T) {
	t.Run("renders with default horizontal direction", func(t *testing.T) {
		got := render(t, ui.ResizablePanelGroup(ui.ResizablePanelGroupProps{}))
		want := `<div data-slot="resizable-panel-group" data-panel-group-direction="horizontal" class="flex h-full w-full data-[panel-group-direction=vertical]:flex-col"></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with vertical direction", func(t *testing.T) {
		got := render(t, ui.ResizablePanelGroup(ui.ResizablePanelGroupProps{Direction: ui.ResizableDirectionVertical}))
		want := `<div data-slot="resizable-panel-group" data-panel-group-direction="vertical" class="flex h-full w-full data-[panel-group-direction=vertical]:flex-col"></div>`
		is.Equal(t, want, got)
	})
}

func TestResizablePanel(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ResizablePanel(ui.ResizablePanelProps{}, Text("Content")))
		want := `<div data-slot="resizable-panel" class="flex-1">Content</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with default size", func(t *testing.T) {
		got := render(t, ui.ResizablePanel(ui.ResizablePanelProps{DefaultSize: 50}, Text("Content")))
		want := `<div data-slot="resizable-panel" data-default-size="50" class="flex-1">Content</div>`
		is.Equal(t, want, got)
	})
}

func TestResizableHandle(t *testing.T) {
	t.Run("renders without handle", func(t *testing.T) {
		got := render(t, ui.ResizableHandle(ui.ResizableHandleProps{}))
		want := `<div data-slot="resizable-handle" class="bg-border focus-visible:ring-ring relative flex w-px items-center justify-center after:absolute after:inset-y-0 after:left-1/2 after:w-1 after:-translate-x-1/2 focus-visible:ring-1 focus-visible:ring-offset-1 focus-visible:outline-hidden data-[panel-group-direction=vertical]:h-px data-[panel-group-direction=vertical]:w-full data-[panel-group-direction=vertical]:after:left-0 data-[panel-group-direction=vertical]:after:h-1 data-[panel-group-direction=vertical]:after:w-full data-[panel-group-direction=vertical]:after:translate-x-0 data-[panel-group-direction=vertical]:after:-translate-y-1/2 [&amp;[data-panel-group-direction=vertical]&gt;div]:rotate-90"></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with handle", func(t *testing.T) {
		got := render(t, ui.ResizableHandle(ui.ResizableHandleProps{WithHandle: true}))
		want := `<div data-slot="resizable-handle" class="bg-border focus-visible:ring-ring relative flex w-px items-center justify-center after:absolute after:inset-y-0 after:left-1/2 after:w-1 after:-translate-x-1/2 focus-visible:ring-1 focus-visible:ring-offset-1 focus-visible:outline-hidden data-[panel-group-direction=vertical]:h-px data-[panel-group-direction=vertical]:w-full data-[panel-group-direction=vertical]:after:left-0 data-[panel-group-direction=vertical]:after:h-1 data-[panel-group-direction=vertical]:after:w-full data-[panel-group-direction=vertical]:after:translate-x-0 data-[panel-group-direction=vertical]:after:-translate-y-1/2 [&amp;[data-panel-group-direction=vertical]&gt;div]:rotate-90"><div class="bg-border z-10 flex h-4 w-3 items-center justify-center rounded-xs border"><svg xmlns="http://www.w3.org/2000/svg" width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="9" cy="12" r="1"/><circle cx="9" cy="5" r="1"/><circle cx="9" cy="19" r="1"/><circle cx="15" cy="12" r="1"/><circle cx="15" cy="5" r="1"/><circle cx="15" cy="19" r="1"/></svg></div></div>`
		is.Equal(t, want, got)
	})
}
