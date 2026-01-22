package ui_test

import (
	"strings"
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestCombobox(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCombobox(t, ui.Combobox(ui.ComboboxProps{}))
		want := `<div data-slot="combobox" class="relative inline-block"></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with children", func(t *testing.T) {
		got := renderCombobox(t, ui.Combobox(ui.ComboboxProps{}, Text("Content")))
		is.True(t, strings.Contains(got, "Content"))
	})
}

func TestComboboxTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxTrigger(ui.ComboboxTriggerProps{}))
		is.True(t, strings.Contains(got, `data-slot="combobox-trigger"`))
		is.True(t, strings.Contains(got, `role="combobox"`))
		is.True(t, strings.Contains(got, `type="button"`))
	})
}

func TestComboboxValue(t *testing.T) {
	t.Run("renders with placeholder", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxValue(ui.ComboboxValueProps{Placeholder: "Select a framework..."}))
		is.True(t, strings.Contains(got, "Select a framework..."))
		is.True(t, strings.Contains(got, `data-placeholder="true"`))
	})

	t.Run("renders with children instead of placeholder", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxValue(ui.ComboboxValueProps{Placeholder: "Select..."}, Text("React")))
		is.True(t, strings.Contains(got, "React"))
		is.True(t, !strings.Contains(got, `data-placeholder="true"`))
	})
}

func TestComboboxContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxContent(ui.ComboboxContentProps{}))
		is.True(t, strings.Contains(got, `data-slot="combobox-content"`))
		is.True(t, strings.Contains(got, "bg-popover"))
	})
}

func TestComboboxInput(t *testing.T) {
	t.Run("renders with default placeholder", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxInput(ui.ComboboxInputProps{}))
		is.True(t, strings.Contains(got, `placeholder="Search..."`))
		is.True(t, strings.Contains(got, `data-slot="combobox-input"`))
	})

	t.Run("renders with custom placeholder", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxInput(ui.ComboboxInputProps{Placeholder: "Search frameworks..."}))
		is.True(t, strings.Contains(got, `placeholder="Search frameworks..."`))
	})

	t.Run("renders search icon", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxInput(ui.ComboboxInputProps{}))
		is.True(t, strings.Contains(got, "<svg"))
	})
}

func TestComboboxList(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxList(ui.ComboboxListProps{}))
		want := `<div data-slot="combobox-list" role="listbox" class="max-h-[300px] overflow-y-auto overflow-x-hidden"></div>`
		is.Equal(t, want, got)
	})
}

func TestComboboxEmpty(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxEmpty(ui.ComboboxEmptyProps{}, Text("No framework found.")))
		want := `<div data-slot="combobox-empty" class="py-6 text-center text-sm">No framework found.</div>`
		is.Equal(t, want, got)
	})
}

func TestComboboxGroup(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxGroup(ui.ComboboxGroupProps{}))
		is.True(t, strings.Contains(got, `data-slot="combobox-group"`))
		is.True(t, strings.Contains(got, `role="group"`))
	})
}

func TestComboboxLabel(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxLabel(ui.ComboboxLabelProps{}, Text("Frameworks")))
		want := `<div data-slot="combobox-label" class="text-muted-foreground px-2 py-1.5 text-xs font-medium">Frameworks</div>`
		is.Equal(t, want, got)
	})
}

func TestComboboxItem(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxItem(ui.ComboboxItemProps{}, Text("React")))
		is.True(t, strings.Contains(got, `data-slot="combobox-item"`))
		is.True(t, strings.Contains(got, `role="option"`))
		is.True(t, strings.Contains(got, "React"))
	})

	t.Run("renders with value", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxItem(ui.ComboboxItemProps{Value: "react"}, Text("React")))
		is.True(t, strings.Contains(got, `data-value="react"`))
	})

	t.Run("renders with selected state", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxItem(ui.ComboboxItemProps{Selected: true}, Text("React")))
		is.True(t, strings.Contains(got, `aria-selected="true"`))
		is.True(t, strings.Contains(got, `data-selected="true"`))
	})

	t.Run("renders with disabled state", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxItem(ui.ComboboxItemProps{Disabled: true}, Text("React")))
		is.True(t, strings.Contains(got, `data-disabled="true"`))
	})
}

func TestComboboxItemIndicator(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxItemIndicator(ui.ComboboxItemIndicatorProps{}, Text("✓")))
		want := `<span data-slot="combobox-item-indicator" class="absolute right-2 flex size-4 items-center justify-center">✓</span>`
		is.Equal(t, want, got)
	})
}

func TestComboboxSeparator(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := renderCombobox(t, ui.ComboboxSeparator(ui.ComboboxSeparatorProps{}))
		want := `<div data-slot="combobox-separator" role="separator" class="bg-border -mx-1 my-1 h-px"></div>`
		is.Equal(t, want, got)
	})
}

func renderCombobox(t *testing.T, n Node) string {
	t.Helper()
	var b strings.Builder
	err := n.Render(&b)
	is.NotError(t, err)
	return b.String()
}
