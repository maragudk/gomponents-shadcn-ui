package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// AlertDialogProps for [AlertDialog].
type AlertDialogProps struct{}

// AlertDialog renders a container for alert dialog components.
// Use with Datastar for open/close state management.
func AlertDialog(props AlertDialogProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "alert-dialog"),
		JoinAttrs("class",
			h.Class(alertDialogBaseClass),
			Group(children),
		),
	)
}

const alertDialogBaseClass = ""

// AlertDialogTriggerProps for [AlertDialogTrigger].
type AlertDialogTriggerProps struct{}

// AlertDialogTrigger renders an element that triggers the alert dialog.
func AlertDialogTrigger(props AlertDialogTriggerProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "alert-dialog-trigger"),
		JoinAttrs("class",
			h.Class(alertDialogTriggerBaseClass),
			Group(children),
		),
	)
}

const alertDialogTriggerBaseClass = ""

// AlertDialogOverlayProps for [AlertDialogOverlay].
type AlertDialogOverlayProps struct{}

// AlertDialogOverlay renders the backdrop overlay behind the alert dialog.
func AlertDialogOverlay(props AlertDialogOverlayProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "alert-dialog-overlay"),
		JoinAttrs("class",
			h.Class(alertDialogOverlayBaseClass),
			Group(children),
		),
	)
}

const alertDialogOverlayBaseClass = "fixed inset-0 z-50 bg-black/50"

// AlertDialogContentProps for [AlertDialogContent].
type AlertDialogContentProps struct{}

// AlertDialogContent renders the main alert dialog content panel.
// Use with Datastar's data-show directive to control visibility.
func AlertDialogContent(props AlertDialogContentProps, children ...Node) Node {
	return h.Div(
		h.Role("alertdialog"),
		h.Aria("modal", "true"),
		h.Data("slot", "alert-dialog-content"),
		JoinAttrs("class",
			h.Class(alertDialogContentBaseClass),
			Group(children),
		),
	)
}

const alertDialogContentBaseClass = "bg-background fixed top-1/2 left-1/2 z-50 grid w-full max-w-[calc(100%-2rem)] -translate-x-1/2 -translate-y-1/2 gap-4 rounded-lg border p-6 shadow-lg sm:max-w-lg"

// AlertDialogHeaderProps for [AlertDialogHeader].
type AlertDialogHeaderProps struct{}

// AlertDialogHeader renders the header section of an alert dialog.
func AlertDialogHeader(props AlertDialogHeaderProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "alert-dialog-header"),
		JoinAttrs("class",
			h.Class(alertDialogHeaderBaseClass),
			Group(children),
		),
	)
}

const alertDialogHeaderBaseClass = "flex flex-col gap-2 text-center sm:text-left"

// AlertDialogFooterProps for [AlertDialogFooter].
type AlertDialogFooterProps struct{}

// AlertDialogFooter renders the footer section of an alert dialog.
func AlertDialogFooter(props AlertDialogFooterProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "alert-dialog-footer"),
		JoinAttrs("class",
			h.Class(alertDialogFooterBaseClass),
			Group(children),
		),
	)
}

const alertDialogFooterBaseClass = "flex flex-col-reverse gap-2 sm:flex-row sm:justify-end"

// AlertDialogTitleProps for [AlertDialogTitle].
type AlertDialogTitleProps struct{}

// AlertDialogTitle renders the title of an alert dialog.
func AlertDialogTitle(props AlertDialogTitleProps, children ...Node) Node {
	return h.H2(
		h.Data("slot", "alert-dialog-title"),
		JoinAttrs("class",
			h.Class(alertDialogTitleBaseClass),
			Group(children),
		),
	)
}

const alertDialogTitleBaseClass = "text-lg font-semibold"

// AlertDialogDescriptionProps for [AlertDialogDescription].
type AlertDialogDescriptionProps struct{}

// AlertDialogDescription renders the description text of an alert dialog.
func AlertDialogDescription(props AlertDialogDescriptionProps, children ...Node) Node {
	return h.P(
		h.Data("slot", "alert-dialog-description"),
		JoinAttrs("class",
			h.Class(alertDialogDescriptionBaseClass),
			Group(children),
		),
	)
}

const alertDialogDescriptionBaseClass = "text-muted-foreground text-sm"

// AlertDialogActionProps for [AlertDialogAction].
type AlertDialogActionProps struct{}

// AlertDialogAction renders the primary action button of an alert dialog.
// Styled as a default button.
func AlertDialogAction(props AlertDialogActionProps, children ...Node) Node {
	return h.Button(
		h.Type("button"),
		h.Data("slot", "alert-dialog-action"),
		JoinAttrs("class",
			h.Class(buttonClasses(ButtonProps{})),
			Group(children),
		),
	)
}

// AlertDialogCancelProps for [AlertDialogCancel].
type AlertDialogCancelProps struct{}

// AlertDialogCancel renders the cancel button of an alert dialog.
// Styled as an outline button.
func AlertDialogCancel(props AlertDialogCancelProps, children ...Node) Node {
	return h.Button(
		h.Type("button"),
		h.Data("slot", "alert-dialog-cancel"),
		JoinAttrs("class",
			h.Class(buttonClasses(ButtonProps{Variant: ButtonVariantOutline})),
			Group(children),
		),
	)
}
