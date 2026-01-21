package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestToggleGroup(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ToggleGroup(ui.ToggleGroupProps{}))
		want := `<div role="group" data-slot="toggle-group" data-type="single" class="inline-flex items-center rounded-md"></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with multiple type", func(t *testing.T) {
		got := render(t, ui.ToggleGroup(ui.ToggleGroupProps{Type: "multiple"}))
		want := `<div role="group" data-slot="toggle-group" data-type="multiple" class="inline-flex items-center rounded-md"></div>`
		is.Equal(t, want, got)
	})
}

func TestToggleGroupItem(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ToggleGroupItem(ui.ToggleGroupItemProps{}, Text("Item")))
		want := `<button type="button" data-slot="toggle-group-item" aria-pressed="false" data-state="off" class="inline-flex items-center justify-center gap-2 text-sm font-medium hover:bg-muted hover:text-muted-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground [&amp;_svg]:pointer-events-none [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 [&amp;_svg]:shrink-0 focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] outline-none transition-[color,box-shadow] whitespace-nowrap cursor-pointer focus:z-10 focus-visible:z-10 rounded-none first:rounded-l-md last:rounded-r-md border-l-0 first:border-l bg-transparent h-9 px-3 min-w-9">Item</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders pressed state", func(t *testing.T) {
		got := render(t, ui.ToggleGroupItem(ui.ToggleGroupItemProps{Pressed: true}, Text("Item")))
		want := `<button type="button" data-slot="toggle-group-item" aria-pressed="true" data-state="on" class="inline-flex items-center justify-center gap-2 text-sm font-medium hover:bg-muted hover:text-muted-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground [&amp;_svg]:pointer-events-none [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 [&amp;_svg]:shrink-0 focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] outline-none transition-[color,box-shadow] whitespace-nowrap cursor-pointer focus:z-10 focus-visible:z-10 rounded-none first:rounded-l-md last:rounded-r-md border-l-0 first:border-l bg-transparent h-9 px-3 min-w-9">Item</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders outline variant", func(t *testing.T) {
		got := render(t, ui.ToggleGroupItem(ui.ToggleGroupItemProps{Variant: ui.ToggleVariantOutline}, Text("Item")))
		want := `<button type="button" data-slot="toggle-group-item" aria-pressed="false" data-state="off" class="inline-flex items-center justify-center gap-2 text-sm font-medium hover:bg-muted hover:text-muted-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground [&amp;_svg]:pointer-events-none [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 [&amp;_svg]:shrink-0 focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] outline-none transition-[color,box-shadow] whitespace-nowrap cursor-pointer focus:z-10 focus-visible:z-10 rounded-none first:rounded-l-md last:rounded-r-md border-l-0 first:border-l border border-input bg-transparent shadow-xs hover:bg-accent hover:text-accent-foreground h-9 px-3 min-w-9">Item</button>`
		is.Equal(t, want, got)
	})
}
