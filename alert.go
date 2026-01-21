package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// AlertVariant defines the visual style of an [Alert].
type AlertVariant string

const (
	AlertVariantDefault     AlertVariant = "default"
	AlertVariantDestructive AlertVariant = "destructive"
)

// AlertProps for [Alert].
type AlertProps struct {
	Variant AlertVariant
}

// Alert renders a div element as an alert with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) and children as needed.
func Alert(props AlertProps, children ...Node) Node {
	return h.Div(
		h.Role("alert"),
		JoinAttrs("class",
			h.Class(alertClasses(props)),
			Group(children),
		),
	)
}

func alertClasses(props AlertProps) string {
	variant := props.Variant
	if variant == "" {
		variant = AlertVariantDefault
	}

	variantClass, ok := alertVariantClasses[variant]
	if !ok {
		panic("ui: invalid AlertVariant: " + string(variant))
	}

	return alertBaseClass + " " + variantClass
}

const alertBaseClass = "relative w-full rounded-lg border px-4 py-3 text-sm grid has-[>svg]:grid-cols-[calc(var(--spacing)*4)_1fr] grid-cols-[0_1fr] has-[>svg]:gap-x-3 gap-y-0.5 items-start [&>svg]:size-4 [&>svg]:translate-y-0.5 [&>svg]:text-current"

var alertVariantClasses = map[AlertVariant]string{
	AlertVariantDefault:     "bg-card text-card-foreground",
	AlertVariantDestructive: "text-destructive bg-card [&>svg]:text-current *:data-[slot=alert-description]:text-destructive/90",
}

// AlertTitleProps for [AlertTitle].
type AlertTitleProps struct{}

// AlertTitle renders a div element as the alert title with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) and children as needed.
func AlertTitle(props AlertTitleProps, children ...Node) Node {
	return h.Div(
		JoinAttrs("class",
			h.Class(alertTitleBaseClass),
			Group(children),
		),
	)
}

const alertTitleBaseClass = "col-start-2 line-clamp-1 min-h-4 font-medium tracking-tight"

// AlertDescriptionProps for [AlertDescription].
type AlertDescriptionProps struct{}

// AlertDescription renders a div element as the alert description with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) and children as needed.
func AlertDescription(props AlertDescriptionProps, children ...Node) Node {
	return h.Div(
		JoinAttrs("class",
			h.Class(alertDescriptionBaseClass),
			Group(children),
		),
	)
}

const alertDescriptionBaseClass = "text-muted-foreground col-start-2 grid justify-items-start gap-1 text-sm [&_p]:leading-relaxed"
