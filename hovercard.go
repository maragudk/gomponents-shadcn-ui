package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// HoverCardProps for [HoverCard].
type HoverCardProps struct{}

// HoverCard renders a container for hover card trigger and content.
// The card appears on hover using CSS group hover.
func HoverCard(props HoverCardProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "hover-card"),
		JoinAttrs("class",
			h.Class(hoverCardBaseClass),
			Group(children),
		),
	)
}

const hoverCardBaseClass = "relative inline-flex group"

// HoverCardTriggerProps for [HoverCardTrigger].
type HoverCardTriggerProps struct{}

// HoverCardTrigger renders the element that triggers the hover card on hover.
func HoverCardTrigger(props HoverCardTriggerProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "hover-card-trigger"),
		JoinAttrs("class",
			h.Class(hoverCardTriggerBaseClass),
			Group(children),
		),
	)
}

const hoverCardTriggerBaseClass = ""

// HoverCardContentProps for [HoverCardContent].
type HoverCardContentProps struct{}

// HoverCardContent renders the hover card content that appears on hover.
// Uses CSS group-hover to show/hide.
func HoverCardContent(props HoverCardContentProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "hover-card-content"),
		JoinAttrs("class",
			h.Class(hoverCardContentBaseClass),
			Group(children),
		),
	)
}

const hoverCardContentBaseClass = "bg-popover text-popover-foreground absolute z-50 hidden group-hover:block w-64 rounded-md border p-4 shadow-md top-full left-0 mt-2"
