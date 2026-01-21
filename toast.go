package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// ToastVariant represents the visual variant of a toast.
type ToastVariant string

const (
	ToastVariantDefault     ToastVariant = "default"
	ToastVariantDestructive ToastVariant = "destructive"
)

// ToastViewportProps for [ToastViewport].
type ToastViewportProps struct{}

// ToastViewport renders a container for positioning toasts.
func ToastViewport(props ToastViewportProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "toast-viewport"),
		JoinAttrs("class",
			h.Class(toastViewportBaseClass),
			Group(children),
		),
	)
}

const toastViewportBaseClass = "fixed top-0 z-[100] flex max-h-screen w-full flex-col-reverse p-4 sm:bottom-0 sm:right-0 sm:top-auto sm:flex-col md:max-w-[420px]"

// ToastProps for [Toast].
type ToastProps struct {
	// Variant can be "default" or "destructive".
	Variant ToastVariant
}

// Toast renders a toast notification container.
// Use with Datastar for open/close state management.
func Toast(props ToastProps, children ...Node) Node {
	var variantClass string
	switch props.Variant {
	case ToastVariantDestructive:
		variantClass = toastDestructiveClass
	default:
		variantClass = toastDefaultClass
	}

	attrs := []Node{
		h.Role("status"),
		h.Aria("live", "polite"),
		h.Data("slot", "toast"),
	}

	if props.Variant == ToastVariantDestructive {
		attrs = append(attrs, h.Data("variant", "destructive"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(toastBaseClass+" "+variantClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const toastBaseClass = "group pointer-events-auto relative flex w-full items-center justify-between space-x-2 overflow-hidden rounded-md border p-4 pr-6 shadow-lg"
const toastDefaultClass = "border bg-background text-foreground"
const toastDestructiveClass = "destructive border-destructive bg-destructive text-destructive-foreground"

// ToastActionProps for [ToastAction].
type ToastActionProps struct{}

// ToastAction renders an action button within a toast.
func ToastAction(props ToastActionProps, children ...Node) Node {
	return h.Button(
		h.Type("button"),
		h.Data("slot", "toast-action"),
		JoinAttrs("class",
			h.Class(toastActionBaseClass),
			Group(children),
		),
	)
}

const toastActionBaseClass = "inline-flex h-8 shrink-0 items-center justify-center rounded-md border bg-transparent px-3 text-sm font-medium transition-colors hover:bg-secondary focus:outline-none focus:ring-1 focus:ring-ring disabled:pointer-events-none disabled:opacity-50 group-[.destructive]:border-muted/40 group-[.destructive]:hover:border-destructive/30 group-[.destructive]:hover:bg-destructive group-[.destructive]:hover:text-destructive-foreground group-[.destructive]:focus:ring-destructive cursor-pointer"

// ToastCloseProps for [ToastClose].
type ToastCloseProps struct{}

// ToastClose renders a close button for the toast.
func ToastClose(props ToastCloseProps, children ...Node) Node {
	return h.Button(
		h.Type("button"),
		h.Data("slot", "toast-close"),
		JoinAttrs("class",
			h.Class(toastCloseBaseClass),
			Group(children),
		),
	)
}

const toastCloseBaseClass = "absolute right-1 top-1 rounded-md p-1 text-foreground/50 opacity-0 transition-opacity hover:text-foreground focus:opacity-100 focus:outline-none focus:ring-1 group-hover:opacity-100 group-[.destructive]:text-red-300 group-[.destructive]:hover:text-red-50 group-[.destructive]:focus:ring-red-400 group-[.destructive]:focus:ring-offset-red-600 cursor-pointer"

// ToastTitleProps for [ToastTitle].
type ToastTitleProps struct{}

// ToastTitle renders the title text of a toast.
func ToastTitle(props ToastTitleProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "toast-title"),
		JoinAttrs("class",
			h.Class(toastTitleBaseClass),
			Group(children),
		),
	)
}

const toastTitleBaseClass = "text-sm font-semibold [&+div]:text-xs"

// ToastDescriptionProps for [ToastDescription].
type ToastDescriptionProps struct{}

// ToastDescription renders the description text of a toast.
func ToastDescription(props ToastDescriptionProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "toast-description"),
		JoinAttrs("class",
			h.Class(toastDescriptionBaseClass),
			Group(children),
		),
	)
}

const toastDescriptionBaseClass = "text-sm opacity-90"
