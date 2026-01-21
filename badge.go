package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// BadgeVariant defines the visual style of a [Badge].
type BadgeVariant string

const (
	BadgeVariantDefault     BadgeVariant = "default"
	BadgeVariantSecondary   BadgeVariant = "secondary"
	BadgeVariantDestructive BadgeVariant = "destructive"
	BadgeVariantOutline     BadgeVariant = "outline"
)

// BadgeProps for [Badge].
type BadgeProps struct {
	Variant BadgeVariant
}

// Badge renders a span element with shadcn/ui badge styling.
// Pass additional attributes (like [h.Class]) and children as needed.
func Badge(props BadgeProps, children ...Node) Node {
	return h.Span(
		JoinAttrs("class",
			h.Class(badgeClasses(props)),
			Group(children),
		),
	)
}

func badgeClasses(props BadgeProps) string {
	variant := props.Variant
	if variant == "" {
		variant = BadgeVariantDefault
	}

	variantClass, ok := badgeVariantClasses[variant]
	if !ok {
		panic("ui: invalid BadgeVariant: " + string(variant))
	}

	return badgeBaseClass + " " + variantClass
}

const badgeBaseClass = "inline-flex items-center justify-center rounded-full border px-2 py-0.5 text-xs font-medium w-fit whitespace-nowrap shrink-0 [&>svg]:size-3 gap-1 [&>svg]:pointer-events-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive transition-[color,box-shadow] overflow-hidden"

var badgeVariantClasses = map[BadgeVariant]string{
	BadgeVariantDefault:     "border-transparent bg-primary text-primary-foreground [a&]:hover:bg-primary/90",
	BadgeVariantSecondary:   "border-transparent bg-secondary text-secondary-foreground [a&]:hover:bg-secondary/90",
	BadgeVariantDestructive: "border-transparent bg-destructive text-white [a&]:hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40 dark:bg-destructive/60",
	BadgeVariantOutline:     "text-foreground [a&]:hover:bg-accent [a&]:hover:text-accent-foreground",
}
