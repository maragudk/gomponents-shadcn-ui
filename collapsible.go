package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// CollapsibleProps for [Collapsible].
type CollapsibleProps struct{}

// Collapsible renders a container for collapsible content.
// Use with [CollapsibleTrigger] and [CollapsibleContent].
// For interactivity, use Datastar to manage open/closed state.
func Collapsible(props CollapsibleProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "collapsible"),
		JoinAttrs("class",
			h.Class(collapsibleBaseClass),
			Group(children),
		),
	)
}

const collapsibleBaseClass = ""

// CollapsibleTriggerProps for [CollapsibleTrigger].
type CollapsibleTriggerProps struct{}

// CollapsibleTrigger renders a button that toggles the collapsible content.
// Use Datastar attributes to handle the toggle action.
func CollapsibleTrigger(props CollapsibleTriggerProps, children ...Node) Node {
	return h.Button(
		h.Type("button"),
		h.Data("slot", "collapsible-trigger"),
		JoinAttrs("class",
			h.Class(collapsibleTriggerBaseClass),
			Group(children),
		),
	)
}

const collapsibleTriggerBaseClass = "cursor-pointer"

// CollapsibleContentProps for [CollapsibleContent].
type CollapsibleContentProps struct{}

// CollapsibleContent renders the content that can be expanded or collapsed.
// Use Datastar's data-show directive to toggle visibility.
func CollapsibleContent(props CollapsibleContentProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "collapsible-content"),
		JoinAttrs("class",
			h.Class(collapsibleContentBaseClass),
			Group(children),
		),
	)
}

const collapsibleContentBaseClass = "overflow-hidden"
