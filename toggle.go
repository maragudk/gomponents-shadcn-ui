package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// ToggleVariant for [Toggle].
type ToggleVariant string

const (
	ToggleVariantDefault ToggleVariant = "default"
	ToggleVariantOutline ToggleVariant = "outline"
)

// ToggleSize for [Toggle].
type ToggleSize string

const (
	ToggleSizeDefault ToggleSize = "default"
	ToggleSizeSm      ToggleSize = "sm"
	ToggleSizeLg      ToggleSize = "lg"
)

// ToggleProps for [Toggle].
type ToggleProps struct {
	Variant ToggleVariant
	Size    ToggleSize
	Pressed bool
}

// Toggle renders a button element that can be toggled on/off with shadcn/ui styling.
// Set Pressed to true for the "on" state. Use Datastar for interactive toggling.
// Pass additional attributes and children as needed.
func Toggle(props ToggleProps, children ...Node) Node {
	variant := props.Variant
	if variant == "" {
		variant = ToggleVariantDefault
	}
	size := props.Size
	if size == "" {
		size = ToggleSizeDefault
	}

	attrs := []Node{
		h.Type("button"),
		h.Data("slot", "toggle"),
	}

	if props.Pressed {
		attrs = append(attrs, h.Aria("pressed", "true"), h.Data("state", "on"))
	} else {
		attrs = append(attrs, h.Aria("pressed", "false"), h.Data("state", "off"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(toggleBaseClass+" "+toggleVariantClasses[variant]+" "+toggleSizeClasses[size]),
			Group(children),
		),
	)

	return h.Button(attrs...)
}

const toggleBaseClass = "inline-flex items-center justify-center gap-2 rounded-md text-sm font-medium hover:bg-muted hover:text-muted-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 [&_svg]:shrink-0 focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] outline-none transition-[color,box-shadow] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive whitespace-nowrap cursor-pointer"

var toggleVariantClasses = map[ToggleVariant]string{
	ToggleVariantDefault: "bg-transparent",
	ToggleVariantOutline: "border border-input bg-transparent shadow-xs hover:bg-accent hover:text-accent-foreground",
}

var toggleSizeClasses = map[ToggleSize]string{
	ToggleSizeDefault: "h-9 px-2 min-w-9",
	ToggleSizeSm:      "h-8 px-1.5 min-w-8",
	ToggleSizeLg:      "h-10 px-2.5 min-w-10",
}
