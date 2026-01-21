package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// DialogProps for [Dialog].
type DialogProps struct{}

// Dialog renders a container for dialog components.
// Use with Datastar for open/close state management.
func Dialog(props DialogProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "dialog"),
		JoinAttrs("class",
			h.Class(dialogBaseClass),
			Group(children),
		),
	)
}

const dialogBaseClass = ""

// DialogTriggerProps for [DialogTrigger].
type DialogTriggerProps struct{}

// DialogTrigger renders an element that triggers the dialog.
func DialogTrigger(props DialogTriggerProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "dialog-trigger"),
		JoinAttrs("class",
			h.Class(dialogTriggerBaseClass),
			Group(children),
		),
	)
}

const dialogTriggerBaseClass = ""

// DialogOverlayProps for [DialogOverlay].
type DialogOverlayProps struct{}

// DialogOverlay renders the backdrop overlay behind the dialog.
func DialogOverlay(props DialogOverlayProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "dialog-overlay"),
		JoinAttrs("class",
			h.Class(dialogOverlayBaseClass),
			Group(children),
		),
	)
}

const dialogOverlayBaseClass = "fixed inset-0 z-50 bg-black/50"

// DialogContentProps for [DialogContent].
type DialogContentProps struct{}

// DialogContent renders the main dialog content panel.
// Use with Datastar's data-show directive to control visibility.
func DialogContent(props DialogContentProps, children ...Node) Node {
	return h.Div(
		h.Role("dialog"),
		h.Aria("modal", "true"),
		h.Data("slot", "dialog-content"),
		JoinAttrs("class",
			h.Class(dialogContentBaseClass),
			Group(children),
		),
	)
}

const dialogContentBaseClass = "bg-background fixed top-1/2 left-1/2 z-50 grid w-full max-w-[calc(100%-2rem)] -translate-x-1/2 -translate-y-1/2 gap-4 rounded-lg border p-6 shadow-lg sm:max-w-lg"

// DialogHeaderProps for [DialogHeader].
type DialogHeaderProps struct{}

// DialogHeader renders the header section of a dialog.
func DialogHeader(props DialogHeaderProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "dialog-header"),
		JoinAttrs("class",
			h.Class(dialogHeaderBaseClass),
			Group(children),
		),
	)
}

const dialogHeaderBaseClass = "flex flex-col gap-2 text-center sm:text-left"

// DialogFooterProps for [DialogFooter].
type DialogFooterProps struct{}

// DialogFooter renders the footer section of a dialog.
func DialogFooter(props DialogFooterProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "dialog-footer"),
		JoinAttrs("class",
			h.Class(dialogFooterBaseClass),
			Group(children),
		),
	)
}

const dialogFooterBaseClass = "flex flex-col-reverse gap-2 sm:flex-row sm:justify-end"

// DialogTitleProps for [DialogTitle].
type DialogTitleProps struct{}

// DialogTitle renders the title of a dialog.
func DialogTitle(props DialogTitleProps, children ...Node) Node {
	return h.H2(
		h.Data("slot", "dialog-title"),
		JoinAttrs("class",
			h.Class(dialogTitleBaseClass),
			Group(children),
		),
	)
}

const dialogTitleBaseClass = "text-lg leading-none font-semibold"

// DialogDescriptionProps for [DialogDescription].
type DialogDescriptionProps struct{}

// DialogDescription renders the description text of a dialog.
func DialogDescription(props DialogDescriptionProps, children ...Node) Node {
	return h.P(
		h.Data("slot", "dialog-description"),
		JoinAttrs("class",
			h.Class(dialogDescriptionBaseClass),
			Group(children),
		),
	)
}

const dialogDescriptionBaseClass = "text-muted-foreground text-sm"

// DialogCloseProps for [DialogClose].
type DialogCloseProps struct{}

// DialogClose renders a close button for the dialog.
func DialogClose(props DialogCloseProps, children ...Node) Node {
	return h.Button(
		h.Type("button"),
		h.Data("slot", "dialog-close"),
		JoinAttrs("class",
			h.Class(dialogCloseBaseClass),
			Group(children),
		),
	)
}

const dialogCloseBaseClass = "ring-offset-background focus:ring-ring absolute top-4 right-4 rounded-sm opacity-70 transition-opacity hover:opacity-100 focus:ring-2 focus:ring-offset-2 focus:outline-hidden disabled:pointer-events-none cursor-pointer"
