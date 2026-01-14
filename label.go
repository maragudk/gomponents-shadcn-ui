package ui

import (
	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// LabelProps configures a Label component.
type LabelProps struct {
	Class string
}

// Label renders a label element with shadcn/ui styling.
// Pass additional attributes (like h.For) and children as needed.
func Label(props LabelProps, children ...g.Node) g.Node {
	return c.JoinAttrs("class",
		h.Label(
			h.Class(labelBaseClass),
			g.If(props.Class != "", h.Class(props.Class)),
			g.Group(children),
		),
	)
}

const labelBaseClass = "text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
