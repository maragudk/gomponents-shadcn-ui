package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// SelectProps for [Select].
type SelectProps struct{}

// Select renders a container for select components.
// Use with Datastar for open/close state management.
func Select(props SelectProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "select"),
		JoinAttrs("class",
			h.Class(selectBaseClass),
			Group(children),
		),
	)
}

const selectBaseClass = "relative inline-block"

// SelectTriggerSize represents the size of a select trigger.
type SelectTriggerSize string

const (
	SelectTriggerSizeDefault SelectTriggerSize = "default"
	SelectTriggerSizeSm      SelectTriggerSize = "sm"
)

// SelectTriggerProps for [SelectTrigger].
type SelectTriggerProps struct {
	// Size can be "default" or "sm".
	Size SelectTriggerSize
}

// SelectTrigger renders the button that opens the select dropdown.
func SelectTrigger(props SelectTriggerProps, children ...Node) Node {
	size := props.Size
	if size == "" {
		size = SelectTriggerSizeDefault
	}

	return h.Button(
		h.Type("button"),
		h.Data("slot", "select-trigger"),
		h.Data("size", string(size)),
		JoinAttrs("class",
			h.Class(selectTriggerBaseClass),
			Group(children),
		),
	)
}

const selectTriggerBaseClass = "border-input data-[placeholder]:text-muted-foreground [&_svg:not([class*='text-'])]:text-muted-foreground focus-visible:border-ring focus-visible:ring-ring/50 dark:bg-input/30 dark:hover:bg-input/50 flex w-full items-center justify-between gap-2 rounded-md border bg-transparent px-3 py-2 text-sm whitespace-nowrap shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50 data-[size=default]:h-9 data-[size=sm]:h-8 cursor-pointer"

// SelectValueProps for [SelectValue].
type SelectValueProps struct {
	// Placeholder is the text shown when no value is selected.
	Placeholder string
}

// SelectValue renders the selected value display.
func SelectValue(props SelectValueProps, children ...Node) Node {
	attrs := []Node{
		h.Data("slot", "select-value"),
	}

	if props.Placeholder != "" && len(children) == 0 {
		attrs = append(attrs, h.Data("placeholder", "true"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(selectValueBaseClass),
			Group(children),
		),
	)

	// If no children and placeholder is set, show placeholder
	if len(children) == 0 && props.Placeholder != "" {
		return h.Span(append(attrs, Text(props.Placeholder))...)
	}

	return h.Span(attrs...)
}

const selectValueBaseClass = "line-clamp-1 flex items-center gap-2"

// SelectContentProps for [SelectContent].
type SelectContentProps struct{}

// SelectContent renders the dropdown content panel.
// Use with Datastar's data-show directive to control visibility.
func SelectContent(props SelectContentProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "select-content"),
		JoinAttrs("class",
			h.Class(selectContentBaseClass),
			Group(children),
		),
	)
}

const selectContentBaseClass = "bg-popover text-popover-foreground absolute z-50 min-w-[8rem] overflow-hidden rounded-md border p-1 shadow-md top-full left-0 mt-1 w-full"

// SelectGroupProps for [SelectGroup].
type SelectGroupProps struct{}

// SelectGroup renders a group of select items.
func SelectGroup(props SelectGroupProps, children ...Node) Node {
	return h.Div(
		h.Role("group"),
		h.Data("slot", "select-group"),
		JoinAttrs("class",
			h.Class(selectGroupBaseClass),
			Group(children),
		),
	)
}

const selectGroupBaseClass = ""

// SelectItemProps for [SelectItem].
type SelectItemProps struct {
	// Selected indicates if this item is currently selected.
	Selected bool
}

// SelectItem renders a single select option.
func SelectItem(props SelectItemProps, children ...Node) Node {
	attrs := []Node{
		h.Role("option"),
		h.Data("slot", "select-item"),
	}

	if props.Selected {
		attrs = append(attrs, h.Aria("selected", "true"), h.Data("state", "checked"))
	} else {
		attrs = append(attrs, h.Aria("selected", "false"), h.Data("state", "unchecked"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(selectItemBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const selectItemBaseClass = "focus:bg-accent focus:text-accent-foreground relative flex w-full cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-8 pl-2 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 hover:bg-accent hover:text-accent-foreground"

// SelectLabelProps for [SelectLabel].
type SelectLabelProps struct{}

// SelectLabel renders a label in the select dropdown.
func SelectLabel(props SelectLabelProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "select-label"),
		JoinAttrs("class",
			h.Class(selectLabelBaseClass),
			Group(children),
		),
	)
}

const selectLabelBaseClass = "text-muted-foreground px-2 py-1.5 text-xs"

// SelectSeparatorProps for [SelectSeparator].
type SelectSeparatorProps struct{}

// SelectSeparator renders a separator line in the select dropdown.
func SelectSeparator(props SelectSeparatorProps, children ...Node) Node {
	return h.Div(
		h.Role("separator"),
		h.Data("slot", "select-separator"),
		JoinAttrs("class",
			h.Class(selectSeparatorBaseClass),
			Group(children),
		),
	)
}

const selectSeparatorBaseClass = "bg-border -mx-1 my-1 h-px"
