package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// ToggleGroupProps for [ToggleGroup].
type ToggleGroupProps struct {
	// Type can be "single" (only one item active) or "multiple" (multiple items can be active).
	Type string
}

// ToggleGroup renders a container for toggle items with shadcn/ui styling.
// Pass [ToggleGroupItem] children as needed.
func ToggleGroup(props ToggleGroupProps, children ...Node) Node {
	groupType := props.Type
	if groupType == "" {
		groupType = "single"
	}
	return h.Div(
		h.Role("group"),
		h.Data("slot", "toggle-group"),
		h.Data("type", groupType),
		JoinAttrs("class",
			h.Class(toggleGroupBaseClass),
			Group(children),
		),
	)
}

const toggleGroupBaseClass = "inline-flex items-center rounded-md"

// ToggleGroupItemProps for [ToggleGroupItem].
type ToggleGroupItemProps struct {
	Variant ToggleVariant
	Size    ToggleSize
	Pressed bool
}

// ToggleGroupItem renders a toggle button within a [ToggleGroup] with shadcn/ui styling.
// Pass additional attributes and children as needed.
func ToggleGroupItem(props ToggleGroupItemProps, children ...Node) Node {
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
		h.Data("slot", "toggle-group-item"),
	}

	if props.Pressed {
		attrs = append(attrs, h.Aria("pressed", "true"), h.Data("state", "on"))
	} else {
		attrs = append(attrs, h.Aria("pressed", "false"), h.Data("state", "off"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(toggleGroupItemBaseClass+" "+toggleVariantClasses[variant]+" "+toggleGroupItemSizeClasses[size]),
			Group(children),
		),
	)

	return h.Button(attrs...)
}

const toggleGroupItemBaseClass = "inline-flex items-center justify-center gap-2 text-sm font-medium hover:bg-muted hover:text-muted-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 [&_svg]:shrink-0 focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] outline-none transition-[color,box-shadow] whitespace-nowrap cursor-pointer focus:z-10 focus-visible:z-10 rounded-none first:rounded-l-md last:rounded-r-md border-l-0 first:border-l"

var toggleGroupItemSizeClasses = map[ToggleSize]string{
	ToggleSizeDefault: "h-9 px-3 min-w-9",
	ToggleSizeSm:      "h-8 px-2 min-w-8",
	ToggleSizeLg:      "h-10 px-3 min-w-10",
}
