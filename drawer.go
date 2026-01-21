package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// DrawerDirection represents the direction from which the drawer slides in.
type DrawerDirection string

const (
	DrawerDirectionBottom DrawerDirection = "bottom"
	DrawerDirectionTop    DrawerDirection = "top"
	DrawerDirectionLeft   DrawerDirection = "left"
	DrawerDirectionRight  DrawerDirection = "right"
)

// DrawerProps for [Drawer].
type DrawerProps struct{}

// Drawer renders a container for drawer components.
// Use with Datastar for open/close state management.
func Drawer(props DrawerProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "drawer"),
		JoinAttrs("class",
			h.Class(drawerBaseClass),
			Group(children),
		),
	)
}

const drawerBaseClass = ""

// DrawerTriggerProps for [DrawerTrigger].
type DrawerTriggerProps struct{}

// DrawerTrigger renders an element that triggers the drawer.
func DrawerTrigger(props DrawerTriggerProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "drawer-trigger"),
		JoinAttrs("class",
			h.Class(drawerTriggerBaseClass),
			Group(children),
		),
	)
}

const drawerTriggerBaseClass = ""

// DrawerOverlayProps for [DrawerOverlay].
type DrawerOverlayProps struct{}

// DrawerOverlay renders the backdrop overlay behind the drawer.
func DrawerOverlay(props DrawerOverlayProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "drawer-overlay"),
		JoinAttrs("class",
			h.Class(drawerOverlayBaseClass),
			Group(children),
		),
	)
}

const drawerOverlayBaseClass = "fixed inset-0 z-50 bg-black/50"

// DrawerContentProps for [DrawerContent].
type DrawerContentProps struct {
	// Direction specifies from which edge the drawer slides in.
	// Defaults to "bottom".
	Direction DrawerDirection
}

// DrawerContent renders the main drawer content panel.
// Use with Datastar's data-show directive to control visibility.
func DrawerContent(props DrawerContentProps, children ...Node) Node {
	direction := props.Direction
	if direction == "" {
		direction = DrawerDirectionBottom
	}

	var directionClass string
	switch direction {
	case DrawerDirectionTop:
		directionClass = "inset-x-0 top-0 max-h-[80vh] rounded-b-lg border-b"
	case DrawerDirectionBottom:
		directionClass = "inset-x-0 bottom-0 max-h-[80vh] rounded-t-lg border-t"
	case DrawerDirectionLeft:
		directionClass = "inset-y-0 left-0 w-3/4 sm:max-w-sm border-r"
	case DrawerDirectionRight:
		directionClass = "inset-y-0 right-0 w-3/4 sm:max-w-sm border-l"
	}

	return h.Div(
		h.Role("dialog"),
		h.Aria("modal", "true"),
		h.Data("slot", "drawer-content"),
		h.Data("direction", string(direction)),
		JoinAttrs("class",
			h.Class(drawerContentBaseClass+" "+directionClass),
			Group(children),
		),
	)
}

const drawerContentBaseClass = "bg-background fixed z-50 flex h-auto flex-col"

// DrawerHeaderProps for [DrawerHeader].
type DrawerHeaderProps struct{}

// DrawerHeader renders the header section of a drawer.
func DrawerHeader(props DrawerHeaderProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "drawer-header"),
		JoinAttrs("class",
			h.Class(drawerHeaderBaseClass),
			Group(children),
		),
	)
}

const drawerHeaderBaseClass = "flex flex-col gap-1.5 p-4 text-center sm:text-left"

// DrawerFooterProps for [DrawerFooter].
type DrawerFooterProps struct{}

// DrawerFooter renders the footer section of a drawer.
func DrawerFooter(props DrawerFooterProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "drawer-footer"),
		JoinAttrs("class",
			h.Class(drawerFooterBaseClass),
			Group(children),
		),
	)
}

const drawerFooterBaseClass = "mt-auto flex flex-col gap-2 p-4"

// DrawerTitleProps for [DrawerTitle].
type DrawerTitleProps struct{}

// DrawerTitle renders the title of a drawer.
func DrawerTitle(props DrawerTitleProps, children ...Node) Node {
	return h.H2(
		h.Data("slot", "drawer-title"),
		JoinAttrs("class",
			h.Class(drawerTitleBaseClass),
			Group(children),
		),
	)
}

const drawerTitleBaseClass = "text-foreground font-semibold"

// DrawerDescriptionProps for [DrawerDescription].
type DrawerDescriptionProps struct{}

// DrawerDescription renders the description text of a drawer.
func DrawerDescription(props DrawerDescriptionProps, children ...Node) Node {
	return h.P(
		h.Data("slot", "drawer-description"),
		JoinAttrs("class",
			h.Class(drawerDescriptionBaseClass),
			Group(children),
		),
	)
}

const drawerDescriptionBaseClass = "text-muted-foreground text-sm"

// DrawerCloseProps for [DrawerClose].
type DrawerCloseProps struct{}

// DrawerClose renders a close button for the drawer.
func DrawerClose(props DrawerCloseProps, children ...Node) Node {
	return h.Button(
		h.Type("button"),
		h.Data("slot", "drawer-close"),
		JoinAttrs("class",
			h.Class(drawerCloseBaseClass),
			Group(children),
		),
	)
}

const drawerCloseBaseClass = "ring-offset-background focus:ring-ring absolute top-4 right-4 rounded-sm opacity-70 transition-opacity hover:opacity-100 focus:ring-2 focus:ring-offset-2 focus:outline-hidden disabled:pointer-events-none cursor-pointer"

// DrawerHandleProps for [DrawerHandle].
type DrawerHandleProps struct{}

// DrawerHandle renders the drag handle indicator for a bottom drawer.
func DrawerHandle(props DrawerHandleProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "drawer-handle"),
		JoinAttrs("class",
			h.Class(drawerHandleBaseClass),
			Group(children),
		),
	)
}

const drawerHandleBaseClass = "bg-muted mx-auto mt-4 h-2 w-[100px] shrink-0 rounded-full"
