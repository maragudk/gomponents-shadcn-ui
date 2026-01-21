package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// SeparatorOrientation defines the orientation of a [Separator].
type SeparatorOrientation string

const (
	SeparatorOrientationHorizontal SeparatorOrientation = "horizontal"
	SeparatorOrientationVertical   SeparatorOrientation = "vertical"
)

// SeparatorProps for [Separator].
type SeparatorProps struct {
	Orientation SeparatorOrientation
}

// Separator renders a div element as a visual separator with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) as needed.
func Separator(props SeparatorProps, children ...Node) Node {
	return h.Div(
		h.Role("separator"),
		JoinAttrs("class",
			h.Class(separatorClasses(props)),
			Group(children),
		),
	)
}

func separatorClasses(props SeparatorProps) string {
	orientation := props.Orientation
	if orientation == "" {
		orientation = SeparatorOrientationHorizontal
	}

	orientationClass, ok := separatorOrientationClasses[orientation]
	if !ok {
		panic("ui: invalid SeparatorOrientation: " + string(orientation))
	}

	return separatorBaseClass + " " + orientationClass
}

const separatorBaseClass = "bg-border shrink-0"

var separatorOrientationClasses = map[SeparatorOrientation]string{
	SeparatorOrientationHorizontal: "h-px w-full",
	SeparatorOrientationVertical:   "h-full w-px",
}
