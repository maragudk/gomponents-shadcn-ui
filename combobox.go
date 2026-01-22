package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// ComboboxProps for [Combobox].
type ComboboxProps struct{}

// Combobox renders a container for combobox components.
// Use with Datastar for open/close state and selection management.
func Combobox(props ComboboxProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "combobox"),
		JoinAttrs("class",
			h.Class(comboboxBaseClass),
			Group(children),
		),
	)
}

const comboboxBaseClass = "relative inline-block"

// ComboboxTriggerProps for [ComboboxTrigger].
type ComboboxTriggerProps struct{}

// ComboboxTrigger renders the button that opens the combobox dropdown.
func ComboboxTrigger(props ComboboxTriggerProps, children ...Node) Node {
	return h.Button(
		h.Type("button"),
		h.Data("slot", "combobox-trigger"),
		h.Role("combobox"),
		h.Aria("autocomplete", "list"),
		JoinAttrs("class",
			h.Class(comboboxTriggerBaseClass),
			Group(children),
		),
	)
}

const comboboxTriggerBaseClass = "border-input data-[placeholder]:text-muted-foreground [&_svg:not([class*='text-'])]:text-muted-foreground focus-visible:border-ring focus-visible:ring-ring/50 dark:bg-input/30 dark:hover:bg-input/50 flex w-full items-center justify-between gap-2 rounded-md border bg-transparent px-3 py-2 text-sm whitespace-nowrap shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 h-9 cursor-pointer"

// ComboboxValueProps for [ComboboxValue].
type ComboboxValueProps struct {
	// Placeholder is the text shown when no value is selected.
	Placeholder string
}

// ComboboxValue renders the selected value display.
func ComboboxValue(props ComboboxValueProps, children ...Node) Node {
	attrs := []Node{
		h.Data("slot", "combobox-value"),
	}

	if props.Placeholder != "" && len(children) == 0 {
		attrs = append(attrs, h.Data("placeholder", "true"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(comboboxValueBaseClass),
			Group(children),
		),
	)

	// If no children and placeholder is set, show placeholder
	if len(children) == 0 && props.Placeholder != "" {
		return h.Span(append(attrs, Text(props.Placeholder))...)
	}

	return h.Span(attrs...)
}

const comboboxValueBaseClass = "line-clamp-1 flex items-center gap-2"

// ComboboxContentProps for [ComboboxContent].
type ComboboxContentProps struct{}

// ComboboxContent renders the dropdown panel with search input and items.
// Use with Datastar's data-show directive to control visibility.
func ComboboxContent(props ComboboxContentProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "combobox-content"),
		JoinAttrs("class",
			h.Class(comboboxContentBaseClass),
			Group(children),
		),
	)
}

const comboboxContentBaseClass = "bg-popover text-popover-foreground absolute z-50 min-w-[200px] overflow-hidden rounded-md border p-0 shadow-md top-full left-0 mt-1 w-full"

// ComboboxInputProps for [ComboboxInput].
type ComboboxInputProps struct {
	Placeholder string
}

// ComboboxInput renders the search input inside the combobox dropdown.
// Use data-bind to bind the input value to a signal for filtering.
func ComboboxInput(props ComboboxInputProps, children ...Node) Node {
	placeholder := props.Placeholder
	if placeholder == "" {
		placeholder = "Search..."
	}

	return h.Div(
		h.Data("slot", "combobox-input-wrapper"),
		h.Class("flex items-center border-b px-3"),
		// Search icon
		Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-2 size-4 shrink-0 opacity-50"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>`),
		h.Input(
			h.Type("text"),
			h.Data("slot", "combobox-input"),
			h.Placeholder(placeholder),
			JoinAttrs("class",
				h.Class(comboboxInputBaseClass),
				Group(children),
			),
		),
	)
}

const comboboxInputBaseClass = "placeholder:text-muted-foreground flex h-10 w-full rounded-md bg-transparent py-3 text-sm outline-hidden disabled:cursor-not-allowed disabled:opacity-50"

// ComboboxListProps for [ComboboxList].
type ComboboxListProps struct{}

// ComboboxList renders the scrollable list container for combobox items.
func ComboboxList(props ComboboxListProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "combobox-list"),
		h.Role("listbox"),
		JoinAttrs("class",
			h.Class(comboboxListBaseClass),
			Group(children),
		),
	)
}

const comboboxListBaseClass = "max-h-[300px] overflow-y-auto overflow-x-hidden"

// ComboboxEmptyProps for [ComboboxEmpty].
type ComboboxEmptyProps struct{}

// ComboboxEmpty renders the empty state when no results are found.
// Use with data-show to display only when the filtered list is empty.
func ComboboxEmpty(props ComboboxEmptyProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "combobox-empty"),
		JoinAttrs("class",
			h.Class(comboboxEmptyBaseClass),
			Group(children),
		),
	)
}

const comboboxEmptyBaseClass = "py-6 text-center text-sm"

// ComboboxGroupProps for [ComboboxGroup].
type ComboboxGroupProps struct{}

// ComboboxGroup renders a group of combobox items.
func ComboboxGroup(props ComboboxGroupProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "combobox-group"),
		h.Role("group"),
		JoinAttrs("class",
			h.Class(comboboxGroupBaseClass),
			Group(children),
		),
	)
}

const comboboxGroupBaseClass = "overflow-hidden p-1 text-foreground"

// ComboboxLabelProps for [ComboboxLabel].
type ComboboxLabelProps struct{}

// ComboboxLabel renders a label/heading for a group of items.
func ComboboxLabel(props ComboboxLabelProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "combobox-label"),
		JoinAttrs("class",
			h.Class(comboboxLabelBaseClass),
			Group(children),
		),
	)
}

const comboboxLabelBaseClass = "text-muted-foreground px-2 py-1.5 text-xs font-medium"

// ComboboxItemProps for [ComboboxItem].
type ComboboxItemProps struct {
	// Value is the searchable value for this item (used for filtering)
	Value string
	// Selected indicates if this item is currently selected
	Selected bool
	// Disabled makes the item non-interactive
	Disabled bool
}

// ComboboxItem renders a selectable combobox item.
// Use data-show with a filter expression to show/hide based on search.
// Use data-on:click for selection handling.
func ComboboxItem(props ComboboxItemProps, children ...Node) Node {
	attrs := []Node{
		h.Data("slot", "combobox-item"),
		h.Role("option"),
	}

	if props.Value != "" {
		attrs = append(attrs, h.Data("value", props.Value))
	}
	if props.Selected {
		attrs = append(attrs, h.Aria("selected", "true"), h.Data("selected", "true"))
	} else {
		attrs = append(attrs, h.Aria("selected", "false"))
	}
	if props.Disabled {
		attrs = append(attrs, h.Data("disabled", "true"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(comboboxItemBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const comboboxItemBaseClass = "relative flex w-full cursor-default items-center gap-2 rounded-sm py-1.5 pr-8 pl-2 text-sm outline-hidden select-none data-[disabled=true]:pointer-events-none data-[disabled=true]:opacity-50 cursor-pointer hover:bg-accent hover:text-accent-foreground data-[selected=true]:bg-accent data-[selected=true]:text-accent-foreground"

// ComboboxItemIndicatorProps for [ComboboxItemIndicator].
type ComboboxItemIndicatorProps struct{}

// ComboboxItemIndicator renders the check icon for selected items.
// Only visible when the parent ComboboxItem has Selected=true.
func ComboboxItemIndicator(props ComboboxItemIndicatorProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "combobox-item-indicator"),
		JoinAttrs("class",
			h.Class(comboboxItemIndicatorBaseClass),
			Group(children),
		),
	)
}

const comboboxItemIndicatorBaseClass = "absolute right-2 flex size-4 items-center justify-center"

// ComboboxSeparatorProps for [ComboboxSeparator].
type ComboboxSeparatorProps struct{}

// ComboboxSeparator renders a visual separator between items or groups.
func ComboboxSeparator(props ComboboxSeparatorProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "combobox-separator"),
		h.Role("separator"),
		JoinAttrs("class",
			h.Class(comboboxSeparatorBaseClass),
			Group(children),
		),
	)
}

const comboboxSeparatorBaseClass = "bg-border -mx-1 my-1 h-px"
