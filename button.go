package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// ButtonVariant defines the visual style of a [Button].
type ButtonVariant string

const (
	ButtonVariantDefault     ButtonVariant = "default"
	ButtonVariantDestructive ButtonVariant = "destructive"
	ButtonVariantOutline     ButtonVariant = "outline"
	ButtonVariantSecondary   ButtonVariant = "secondary"
	ButtonVariantGhost       ButtonVariant = "ghost"
	ButtonVariantLink        ButtonVariant = "link"
)

// ButtonSize defines the size of a [Button].
type ButtonSize string

const (
	ButtonSizeDefault ButtonSize = "default"
	ButtonSizeSm      ButtonSize = "sm"
	ButtonSizeLg      ButtonSize = "lg"
	ButtonSizeIcon    ButtonSize = "icon"
	ButtonSizeIconSm  ButtonSize = "icon-sm"
	ButtonSizeIconLg  ButtonSize = "icon-lg"
)

// ButtonProps for [Button] and [ButtonA].
type ButtonProps struct {
	Variant ButtonVariant
	Size    ButtonSize
}

// Button renders a button element with shadcn/ui styling.
// Pass additional attributes (like [h.Type], [h.Disabled], [h.Class]) and children as needed.
func Button(props ButtonProps, children ...Node) Node {
	return h.Button(
		JoinAttrs("class",
			h.Class(buttonClasses(props)),
			Group(children),
		),
	)
}

// ButtonA renders an anchor element styled as a button.
// Pass additional attributes (like [h.Href], [h.Class]) and children as needed.
//
// Note: Anchor elements don't support the disabled attribute. To create a disabled anchor button,
// add [h.Aria]("disabled", "true") for screen readers, [h.TabIndex]("-1") to prevent keyboard focus,
// [h.Class]("pointer-events-none opacity-50") for visual styling, and handle click prevention in your code.
func ButtonA(props ButtonProps, children ...Node) Node {
	return h.A(
		JoinAttrs("class",
			h.Class(buttonClasses(props)),
			Group(children),
		),
	)
}

func buttonClasses(props ButtonProps) string {
	variant := props.Variant
	if variant == "" {
		variant = ButtonVariantDefault
	}

	size := props.Size
	if size == "" {
		size = ButtonSizeDefault
	}

	variantClass, ok := buttonVariantClasses[variant]
	if !ok {
		panic("ui: invalid ButtonVariant: " + string(variant))
	}

	sizeClass, ok := buttonSizeClasses[size]
	if !ok {
		panic("ui: invalid ButtonSize: " + string(size))
	}

	return buttonBaseClass + " " + variantClass + " " + sizeClass
}

const buttonBaseClass = "inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-all disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 shrink-0 [&_svg]:shrink-0 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive"

var buttonVariantClasses = map[ButtonVariant]string{
	ButtonVariantDefault:     "bg-primary text-primary-foreground hover:bg-primary/90",
	ButtonVariantDestructive: "bg-destructive text-white hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40 dark:bg-destructive/60",
	ButtonVariantOutline:     "border bg-background shadow-xs hover:bg-accent hover:text-accent-foreground dark:bg-input/30 dark:border-input dark:hover:bg-input/50",
	ButtonVariantSecondary:   "bg-secondary text-secondary-foreground hover:bg-secondary/80",
	ButtonVariantGhost:       "hover:bg-accent hover:text-accent-foreground dark:hover:bg-accent/50",
	ButtonVariantLink:        "text-primary underline-offset-4 hover:underline",
}

var buttonSizeClasses = map[ButtonSize]string{
	ButtonSizeDefault: "h-9 px-4 py-2 has-[>svg]:px-3",
	ButtonSizeSm:      "h-8 rounded-md gap-1.5 px-3 has-[>svg]:px-2.5",
	ButtonSizeLg:      "h-10 rounded-md px-6 has-[>svg]:px-4",
	ButtonSizeIcon:    "size-9",
	ButtonSizeIconSm:  "size-8",
	ButtonSizeIconLg:  "size-10",
}
