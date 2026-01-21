package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// PopoverProps for [Popover].
type PopoverProps struct{}

// Popover renders a container for popover trigger and content.
// Use with Datastar for open/close state management.
func Popover(props PopoverProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "popover"),
		JoinAttrs("class",
			h.Class(popoverBaseClass),
			Group(children),
		),
	)
}

const popoverBaseClass = "relative inline-block"

// PopoverTriggerProps for [PopoverTrigger].
type PopoverTriggerProps struct{}

// PopoverTrigger renders the element that triggers the popover.
// Typically contains a button that toggles the popover visibility.
func PopoverTrigger(props PopoverTriggerProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "popover-trigger"),
		JoinAttrs("class",
			h.Class(popoverTriggerBaseClass),
			Group(children),
		),
	)
}

const popoverTriggerBaseClass = ""

// PopoverContentProps for [PopoverContent].
type PopoverContentProps struct{}

// PopoverContent renders the popover content that appears when triggered.
// Use Datastar's data-show directive to control visibility.
func PopoverContent(props PopoverContentProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "popover-content"),
		JoinAttrs("class",
			h.Class(popoverContentBaseClass),
			Group(children),
		),
	)
}

const popoverContentBaseClass = "bg-popover text-popover-foreground absolute z-50 w-72 rounded-md border p-4 shadow-md top-full left-0 mt-1"
