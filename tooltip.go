package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// TooltipProps for [Tooltip].
type TooltipProps struct{}

// Tooltip renders a container for tooltip trigger and content.
// The tooltip appears on hover using CSS group hover.
func Tooltip(props TooltipProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "tooltip"),
		JoinAttrs("class",
			h.Class(tooltipBaseClass),
			Group(children),
		),
	)
}

const tooltipBaseClass = "relative inline-flex group"

// TooltipTriggerProps for [TooltipTrigger].
type TooltipTriggerProps struct{}

// TooltipTrigger renders the element that triggers the tooltip on hover.
func TooltipTrigger(props TooltipTriggerProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "tooltip-trigger"),
		JoinAttrs("class",
			h.Class(tooltipTriggerBaseClass),
			Group(children),
		),
	)
}

const tooltipTriggerBaseClass = ""

// TooltipContentProps for [TooltipContent].
type TooltipContentProps struct{}

// TooltipContent renders the tooltip content that appears on hover.
// Uses CSS group-hover to show/hide.
func TooltipContent(props TooltipContentProps, children ...Node) Node {
	return h.Div(
		h.Role("tooltip"),
		h.Data("slot", "tooltip-content"),
		JoinAttrs("class",
			h.Class(tooltipContentBaseClass),
			Group(children),
		),
	)
}

const tooltipContentBaseClass = "bg-foreground text-background absolute z-50 hidden group-hover:block w-max rounded-md px-3 py-1.5 text-xs bottom-full left-1/2 -translate-x-1/2 mb-2"
