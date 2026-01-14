package ui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// LabelProps configures a Label component.
type LabelProps struct {
	Class string
}

// Label renders a label element with shadcn/ui styling.
// Pass additional attributes (like h.For) and children as needed.
func Label(props LabelProps, children ...g.Node) g.Node {
	return h.Label(
		h.Class(labelClasses(props)),
		g.Group(children),
	)
}

func labelClasses(props LabelProps) string {
	base := "text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"

	if props.Class != "" {
		return base + " " + props.Class
	}

	return base
}
