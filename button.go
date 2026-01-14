package ui

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ButtonVariant defines the visual style of a button.
type ButtonVariant string

const (
	ButtonVariantDefault     ButtonVariant = "default"
	ButtonVariantDestructive ButtonVariant = "destructive"
	ButtonVariantOutline     ButtonVariant = "outline"
	ButtonVariantSecondary   ButtonVariant = "secondary"
	ButtonVariantGhost       ButtonVariant = "ghost"
	ButtonVariantLink        ButtonVariant = "link"
)

// ButtonSize defines the size of a button.
type ButtonSize string

const (
	ButtonSizeDefault ButtonSize = "default"
	ButtonSizeSm      ButtonSize = "sm"
	ButtonSizeLg      ButtonSize = "lg"
	ButtonSizeIcon    ButtonSize = "icon"
)

// ButtonProps configures a Button component.
type ButtonProps struct {
	Variant ButtonVariant
	Size    ButtonSize
	Class   string
}

// Button renders a button element with shadcn/ui styling.
// Pass additional attributes and children as needed.
func Button(props ButtonProps, children ...g.Node) g.Node {
	return h.Button(
		h.Class(buttonClasses(props)),
		g.Group(children),
	)
}

// ButtonA renders an anchor element styled as a button.
func ButtonA(props ButtonProps, children ...g.Node) g.Node {
	return h.A(
		h.Class(buttonClasses(props)),
		g.Group(children),
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

	base := "inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0"

	variantClass := buttonVariantClasses[variant]
	sizeClass := buttonSizeClasses[size]

	classes := base + " " + variantClass + " " + sizeClass

	if props.Class != "" {
		classes += " " + props.Class
	}

	return classes
}

var buttonVariantClasses = map[ButtonVariant]string{
	ButtonVariantDefault:     "bg-primary text-primary-foreground shadow hover:bg-primary/90",
	ButtonVariantDestructive: "bg-destructive text-destructive-foreground shadow-sm hover:bg-destructive/90",
	ButtonVariantOutline:     "border border-input bg-background shadow-sm hover:bg-accent hover:text-accent-foreground",
	ButtonVariantSecondary:   "bg-secondary text-secondary-foreground shadow-sm hover:bg-secondary/80",
	ButtonVariantGhost:       "hover:bg-accent hover:text-accent-foreground",
	ButtonVariantLink:        "text-primary underline-offset-4 hover:underline",
}

var buttonSizeClasses = map[ButtonSize]string{
	ButtonSizeDefault: "h-9 px-4 py-2",
	ButtonSizeSm:      "h-8 rounded-md px-3 text-xs",
	ButtonSizeLg:      "h-10 rounded-md px-8",
	ButtonSizeIcon:    "h-9 w-9",
}
