package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// LabelProps for [Label].
type LabelProps struct{}

// Label renders a label element with shadcn/ui styling.
// Pass additional attributes (like [h.For], [h.Class]) and children as needed.
func Label(props LabelProps, children ...Node) Node {
	return h.Label(
		JoinAttrs("class",
			h.Class(labelBaseClass),
			Group(children),
		),
	)
}

const labelBaseClass = "text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
