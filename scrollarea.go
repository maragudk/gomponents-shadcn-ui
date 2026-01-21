package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// ScrollAreaOrientation represents the scroll direction.
type ScrollAreaOrientation string

const (
	ScrollAreaOrientationVertical   ScrollAreaOrientation = "vertical"
	ScrollAreaOrientationHorizontal ScrollAreaOrientation = "horizontal"
)

// ScrollAreaProps for [ScrollArea].
type ScrollAreaProps struct {
	// Orientation controls the primary scroll direction.
	// Defaults to "vertical".
	Orientation ScrollAreaOrientation
}

// ScrollArea renders a scrollable area container with styled scrollbars.
// Use CSS classes to control height/width constraints on the scroll area.
func ScrollArea(props ScrollAreaProps, children ...Node) Node {
	orientation := props.Orientation
	if orientation == "" {
		orientation = ScrollAreaOrientationVertical
	}

	var scrollClass string
	if orientation == ScrollAreaOrientationHorizontal {
		scrollClass = scrollAreaHorizontalClass
	} else {
		scrollClass = scrollAreaVerticalClass
	}

	return h.Div(
		h.Data("slot", "scroll-area"),
		h.Data("orientation", string(orientation)),
		JoinAttrs("class",
			h.Class(scrollAreaBaseClass+" "+scrollClass),
			Group(children),
		),
	)
}

const scrollAreaBaseClass = "relative"
const scrollAreaVerticalClass = "overflow-y-auto overflow-x-hidden"
const scrollAreaHorizontalClass = "overflow-x-auto overflow-y-hidden"

// ScrollAreaViewportProps for [ScrollAreaViewport].
type ScrollAreaViewportProps struct{}

// ScrollAreaViewport renders the viewport container for scrollable content.
func ScrollAreaViewport(props ScrollAreaViewportProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "scroll-area-viewport"),
		JoinAttrs("class",
			h.Class(scrollAreaViewportBaseClass),
			Group(children),
		),
	)
}

const scrollAreaViewportBaseClass = "size-full rounded-[inherit]"

// ScrollBarProps for [ScrollBar].
type ScrollBarProps struct {
	// Orientation can be "vertical" or "horizontal".
	// Defaults to "vertical".
	Orientation ScrollAreaOrientation
}

// ScrollBar renders a styled scrollbar indicator.
// Note: This is a visual element; actual scrolling is handled by CSS overflow.
func ScrollBar(props ScrollBarProps, children ...Node) Node {
	orientation := props.Orientation
	if orientation == "" {
		orientation = ScrollAreaOrientationVertical
	}

	var orientationClass string
	if orientation == ScrollAreaOrientationHorizontal {
		orientationClass = "h-2.5 flex-col border-t border-t-transparent"
	} else {
		orientationClass = "h-full w-2.5 border-l border-l-transparent"
	}

	return h.Div(
		h.Data("slot", "scroll-area-scrollbar"),
		h.Data("orientation", string(orientation)),
		JoinAttrs("class",
			h.Class(scrollBarBaseClass+" "+orientationClass),
			Group(children),
		),
		h.Div(
			h.Data("slot", "scroll-area-thumb"),
			h.Class(scrollBarThumbBaseClass),
		),
	)
}

const scrollBarBaseClass = "flex touch-none p-px transition-colors select-none"
const scrollBarThumbBaseClass = "bg-border relative flex-1 rounded-full"
