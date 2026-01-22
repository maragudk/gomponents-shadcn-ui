package ui_test

import (
	"strings"
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestCommand(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCommand(t, ui.Command(ui.CommandProps{}))
		want := `<div data-slot="command" class="bg-popover text-popover-foreground flex h-full w-full flex-col overflow-hidden rounded-md"></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with children", func(t *testing.T) {
		got := renderCommand(t, ui.Command(ui.CommandProps{}, Text("Content")))
		is.True(t, strings.Contains(got, "Content"))
	})
}

func TestCommandInput(t *testing.T) {
	t.Run("renders with default placeholder", func(t *testing.T) {
		got := renderCommand(t, ui.CommandInput(ui.CommandInputProps{}))
		is.True(t, strings.Contains(got, `placeholder="Type a command or search..."`))
		is.True(t, strings.Contains(got, `data-slot="command-input"`))
	})

	t.Run("renders with custom placeholder", func(t *testing.T) {
		got := renderCommand(t, ui.CommandInput(ui.CommandInputProps{Placeholder: "Search files..."}))
		is.True(t, strings.Contains(got, `placeholder="Search files..."`))
	})

	t.Run("renders search icon", func(t *testing.T) {
		got := renderCommand(t, ui.CommandInput(ui.CommandInputProps{}))
		is.True(t, strings.Contains(got, "<svg"))
		is.True(t, strings.Contains(got, "circle"))
	})
}

func TestCommandList(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCommand(t, ui.CommandList(ui.CommandListProps{}))
		want := `<div data-slot="command-list" role="listbox" class="max-h-[300px] scroll-py-1 overflow-x-hidden overflow-y-auto"></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with children", func(t *testing.T) {
		got := renderCommand(t, ui.CommandList(ui.CommandListProps{}, Text("Items")))
		is.True(t, strings.Contains(got, "Items"))
	})
}

func TestCommandEmpty(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCommand(t, ui.CommandEmpty(ui.CommandEmptyProps{}, Text("No results found.")))
		want := `<div data-slot="command-empty" class="py-6 text-center text-sm">No results found.</div>`
		is.Equal(t, want, got)
	})
}

func TestCommandGroup(t *testing.T) {
	t.Run("renders without heading", func(t *testing.T) {
		got := renderCommand(t, ui.CommandGroup(ui.CommandGroupProps{}, Text("Items")))
		want := `<div data-slot="command-group" role="group" class="text-foreground overflow-hidden p-1">Items</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with heading", func(t *testing.T) {
		got := renderCommand(t, ui.CommandGroup(ui.CommandGroupProps{Heading: "Suggestions"}, Text("Items")))
		is.True(t, strings.Contains(got, `data-slot="command-group-heading"`))
		is.True(t, strings.Contains(got, "Suggestions"))
	})
}

func TestCommandItem(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCommand(t, ui.CommandItem(ui.CommandItemProps{}, Text("Item")))
		is.True(t, strings.Contains(got, `data-slot="command-item"`))
		is.True(t, strings.Contains(got, `role="option"`))
		is.True(t, strings.Contains(got, "Item"))
	})

	t.Run("renders with value", func(t *testing.T) {
		got := renderCommand(t, ui.CommandItem(ui.CommandItemProps{Value: "settings"}, Text("Settings")))
		is.True(t, strings.Contains(got, `data-value="settings"`))
	})

	t.Run("renders with disabled state", func(t *testing.T) {
		got := renderCommand(t, ui.CommandItem(ui.CommandItemProps{Disabled: true}, Text("Disabled")))
		is.True(t, strings.Contains(got, `data-disabled="true"`))
	})
}

func TestCommandShortcut(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCommand(t, ui.CommandShortcut(ui.CommandShortcutProps{}, Text("⌘K")))
		want := `<span data-slot="command-shortcut" class="text-muted-foreground ml-auto text-xs tracking-widest">⌘K</span>`
		is.Equal(t, want, got)
	})
}

func TestCommandSeparator(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCommand(t, ui.CommandSeparator(ui.CommandSeparatorProps{}))
		want := `<div data-slot="command-separator" class="bg-border -mx-1 h-px"></div>`
		is.Equal(t, want, got)
	})
}

func renderCommand(t *testing.T, n Node) string {
	t.Helper()
	var b strings.Builder
	err := n.Render(&b)
	is.NotError(t, err)
	return b.String()
}
