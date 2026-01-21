package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestSelect(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Select(ui.SelectProps{}))
		want := `<div data-slot="select" class="relative inline-block"></div>`
		is.Equal(t, want, got)
	})
}

func TestSelectTrigger(t *testing.T) {
	t.Run("renders with default size", func(t *testing.T) {
		got := render(t, ui.SelectTrigger(ui.SelectTriggerProps{}, Text("Select")))
		want := `<button type="button" data-slot="select-trigger" data-size="default" class="border-input data-[placeholder]:text-muted-foreground [&amp;_svg:not([class*=&#39;text-&#39;])]:text-muted-foreground focus-visible:border-ring focus-visible:ring-ring/50 dark:bg-input/30 dark:hover:bg-input/50 flex w-full items-center justify-between gap-2 rounded-md border bg-transparent px-3 py-2 text-sm whitespace-nowrap shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 data-[size=default]:h-9 data-[size=sm]:h-8 cursor-pointer">Select</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders with sm size", func(t *testing.T) {
		got := render(t, ui.SelectTrigger(ui.SelectTriggerProps{Size: ui.SelectTriggerSizeSm}, Text("Select")))
		want := `<button type="button" data-slot="select-trigger" data-size="sm" class="border-input data-[placeholder]:text-muted-foreground [&amp;_svg:not([class*=&#39;text-&#39;])]:text-muted-foreground focus-visible:border-ring focus-visible:ring-ring/50 dark:bg-input/30 dark:hover:bg-input/50 flex w-full items-center justify-between gap-2 rounded-md border bg-transparent px-3 py-2 text-sm whitespace-nowrap shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 data-[size=default]:h-9 data-[size=sm]:h-8 cursor-pointer">Select</button>`
		is.Equal(t, want, got)
	})
}

func TestSelectValue(t *testing.T) {
	t.Run("renders with value", func(t *testing.T) {
		got := render(t, ui.SelectValue(ui.SelectValueProps{}, Text("Apple")))
		want := `<span data-slot="select-value" class="line-clamp-1 flex items-center gap-2">Apple</span>`
		is.Equal(t, want, got)
	})

	t.Run("renders with placeholder", func(t *testing.T) {
		got := render(t, ui.SelectValue(ui.SelectValueProps{Placeholder: "Select a fruit"}))
		want := `<span data-slot="select-value" data-placeholder="true" class="line-clamp-1 flex items-center gap-2">Select a fruit</span>`
		is.Equal(t, want, got)
	})
}

func TestSelectContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SelectContent(ui.SelectContentProps{}))
		want := `<div data-slot="select-content" class="bg-popover text-popover-foreground absolute z-50 min-w-[8rem] overflow-hidden rounded-md border p-1 shadow-md top-full left-0 mt-1 w-full"></div>`
		is.Equal(t, want, got)
	})
}

func TestSelectGroup(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SelectGroup(ui.SelectGroupProps{}))
		want := `<div role="group" data-slot="select-group" class=""></div>`
		is.Equal(t, want, got)
	})
}

func TestSelectItem(t *testing.T) {
	t.Run("renders unchecked", func(t *testing.T) {
		got := render(t, ui.SelectItem(ui.SelectItemProps{}, Text("Apple")))
		want := `<div role="option" data-slot="select-item" aria-selected="false" data-state="unchecked" class="focus:bg-accent focus:text-accent-foreground relative flex w-full cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-8 pl-2 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 hover:bg-accent hover:text-accent-foreground">Apple</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders checked", func(t *testing.T) {
		got := render(t, ui.SelectItem(ui.SelectItemProps{Selected: true}, Text("Apple")))
		want := `<div role="option" data-slot="select-item" aria-selected="true" data-state="checked" class="focus:bg-accent focus:text-accent-foreground relative flex w-full cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-8 pl-2 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 hover:bg-accent hover:text-accent-foreground">Apple</div>`
		is.Equal(t, want, got)
	})
}

func TestSelectLabel(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SelectLabel(ui.SelectLabelProps{}, Text("Fruits")))
		want := `<div data-slot="select-label" class="text-muted-foreground px-2 py-1.5 text-xs">Fruits</div>`
		is.Equal(t, want, got)
	})
}

func TestSelectSeparator(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SelectSeparator(ui.SelectSeparatorProps{}))
		want := `<div role="separator" data-slot="select-separator" class="bg-border -mx-1 my-1 h-px"></div>`
		is.Equal(t, want, got)
	})
}
