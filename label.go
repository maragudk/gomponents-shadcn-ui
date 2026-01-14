package ui

import (
	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// LabelProps configures a Label component.
type LabelProps struct{}

// Label renders a label element with shadcn/ui styling.
// Pass additional attributes (like h.For, h.Class) and children as needed.
func Label(props LabelProps, children ...g.Node) g.Node {
	return h.Label(
		c.JoinAttrs("class",
			h.Class(labelBaseClass),
			g.Group(children),
		),
	)
}

const labelBaseClass = "text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
