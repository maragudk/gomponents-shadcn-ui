package ui

import (
	"strconv"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// resizableIntToString converts an int to a string for attribute values.
func resizableIntToString(i int) string {
	return strconv.Itoa(i)
}

// ResizableDirection represents the direction of the resizable panel group.
type ResizableDirection string

const (
	ResizableDirectionHorizontal ResizableDirection = "horizontal"
	ResizableDirectionVertical   ResizableDirection = "vertical"
)

// ResizablePanelGroupProps for [ResizablePanelGroup].
type ResizablePanelGroupProps struct {
	// Direction controls the layout direction.
	// Defaults to "horizontal".
	Direction ResizableDirection
}

// ResizablePanelGroup renders a container for resizable panels.
// Use with JavaScript for actual resize functionality.
func ResizablePanelGroup(props ResizablePanelGroupProps, children ...Node) Node {
	direction := props.Direction
	if direction == "" {
		direction = ResizableDirectionHorizontal
	}

	return h.Div(
		h.Data("slot", "resizable-panel-group"),
		h.Data("panel-group-direction", string(direction)),
		JoinAttrs("class",
			h.Class(resizablePanelGroupBaseClass),
			Group(children),
		),
	)
}

const resizablePanelGroupBaseClass = "flex h-full w-full data-[panel-group-direction=vertical]:flex-col"

// ResizablePanelProps for [ResizablePanel].
type ResizablePanelProps struct {
	// DefaultSize sets the initial size as a percentage (0-100).
	DefaultSize int
}

// ResizablePanel renders an individual resizable panel.
func ResizablePanel(props ResizablePanelProps, children ...Node) Node {
	attrs := []Node{
		h.Data("slot", "resizable-panel"),
	}

	if props.DefaultSize > 0 {
		attrs = append(attrs, h.Data("default-size", resizableIntToString(props.DefaultSize)))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(resizablePanelBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const resizablePanelBaseClass = "flex-1"

// ResizableHandleProps for [ResizableHandle].
type ResizableHandleProps struct {
	// WithHandle shows a visible grip handle.
	WithHandle bool
}

// ResizableHandle renders a drag handle between panels.
// Use with JavaScript for actual resize functionality.
func ResizableHandle(props ResizableHandleProps, children ...Node) Node {
	var handleNode Node
	if props.WithHandle {
		handleNode = h.Div(
			h.Class("bg-border z-10 flex h-4 w-3 items-center justify-center rounded-xs border"),
			Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="9" cy="12" r="1"/><circle cx="9" cy="5" r="1"/><circle cx="9" cy="19" r="1"/><circle cx="15" cy="12" r="1"/><circle cx="15" cy="5" r="1"/><circle cx="15" cy="19" r="1"/></svg>`),
		)
	}

	return h.Div(
		h.Data("slot", "resizable-handle"),
		JoinAttrs("class",
			h.Class(resizableHandleBaseClass),
			Group(children),
		),
		handleNode,
	)
}

const resizableHandleBaseClass = "bg-border focus-visible:ring-ring relative flex w-px items-center justify-center after:absolute after:inset-y-0 after:left-1/2 after:w-1 after:-translate-x-1/2 focus-visible:ring-1 focus-visible:ring-offset-1 focus-visible:outline-hidden data-[panel-group-direction=vertical]:h-px data-[panel-group-direction=vertical]:w-full data-[panel-group-direction=vertical]:after:left-0 data-[panel-group-direction=vertical]:after:h-1 data-[panel-group-direction=vertical]:after:w-full data-[panel-group-direction=vertical]:after:translate-x-0 data-[panel-group-direction=vertical]:after:-translate-y-1/2 [&[data-panel-group-direction=vertical]>div]:rotate-90"
