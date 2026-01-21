package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// SidebarSide represents which side the sidebar appears on.
type SidebarSide string

const (
	SidebarSideLeft  SidebarSide = "left"
	SidebarSideRight SidebarSide = "right"
)

// SidebarVariant represents the visual variant of the sidebar.
type SidebarVariant string

const (
	SidebarVariantSidebar  SidebarVariant = "sidebar"
	SidebarVariantFloating SidebarVariant = "floating"
	SidebarVariantInset    SidebarVariant = "inset"
)

// SidebarProviderProps for [SidebarProvider].
type SidebarProviderProps struct{}

// SidebarProvider renders the outer wrapper for a sidebar layout.
// Use with Datastar for managing open/collapsed state.
func SidebarProvider(props SidebarProviderProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "sidebar-wrapper"),
		h.Style("--sidebar-width: 16rem; --sidebar-width-icon: 3rem"),
		JoinAttrs("class",
			h.Class(sidebarProviderBaseClass),
			Group(children),
		),
	)
}

const sidebarProviderBaseClass = "group/sidebar-wrapper has-data-[variant=inset]:bg-sidebar flex min-h-svh w-full"

// SidebarProps for [Sidebar].
type SidebarProps struct {
	// Side controls which side the sidebar appears on.
	// Defaults to "left".
	Side SidebarSide
	// Variant controls the visual style.
	// Defaults to "sidebar".
	Variant SidebarVariant
}

// Sidebar renders the main sidebar container.
// Use with Datastar's data-state attribute for expanded/collapsed states.
func Sidebar(props SidebarProps, children ...Node) Node {
	side := props.Side
	if side == "" {
		side = SidebarSideLeft
	}
	variant := props.Variant
	if variant == "" {
		variant = SidebarVariantSidebar
	}

	return h.Div(
		h.Data("slot", "sidebar"),
		h.Data("side", string(side)),
		h.Data("variant", string(variant)),
		JoinAttrs("class",
			h.Class(sidebarBaseClass),
			Group(children),
		),
	)
}

const sidebarBaseClass = "group peer text-sidebar-foreground hidden md:block"

// SidebarContentProps for [SidebarContent].
type SidebarContentProps struct{}

// SidebarContent renders the scrollable content area of the sidebar.
func SidebarContent(props SidebarContentProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "sidebar-content"),
		h.Data("sidebar", "content"),
		JoinAttrs("class",
			h.Class(sidebarContentBaseClass),
			Group(children),
		),
	)
}

const sidebarContentBaseClass = "flex min-h-0 flex-1 flex-col gap-2 overflow-auto group-data-[collapsible=icon]:overflow-hidden"

// SidebarHeaderProps for [SidebarHeader].
type SidebarHeaderProps struct{}

// SidebarHeader renders the header section of the sidebar.
func SidebarHeader(props SidebarHeaderProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "sidebar-header"),
		h.Data("sidebar", "header"),
		JoinAttrs("class",
			h.Class(sidebarHeaderBaseClass),
			Group(children),
		),
	)
}

const sidebarHeaderBaseClass = "flex flex-col gap-2 p-2"

// SidebarFooterProps for [SidebarFooter].
type SidebarFooterProps struct{}

// SidebarFooter renders the footer section of the sidebar.
func SidebarFooter(props SidebarFooterProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "sidebar-footer"),
		h.Data("sidebar", "footer"),
		JoinAttrs("class",
			h.Class(sidebarFooterBaseClass),
			Group(children),
		),
	)
}

const sidebarFooterBaseClass = "flex flex-col gap-2 p-2"

// SidebarTriggerProps for [SidebarTrigger].
type SidebarTriggerProps struct{}

// SidebarTrigger renders the button to toggle the sidebar.
func SidebarTrigger(props SidebarTriggerProps, children ...Node) Node {
	buttonClass := buttonClasses(ButtonProps{Variant: ButtonVariantGhost, Size: ButtonSizeIcon})

	return h.Button(
		h.Type("button"),
		h.Data("slot", "sidebar-trigger"),
		h.Data("sidebar", "trigger"),
		JoinAttrs("class",
			h.Class(buttonClass+" size-7"),
			Group(children),
		),
		Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 3v18"/></svg>`),
		h.Span(h.Class("sr-only"), Text("Toggle Sidebar")),
	)
}

// SidebarGroupProps for [SidebarGroup].
type SidebarGroupProps struct{}

// SidebarGroup renders a group container for sidebar items.
func SidebarGroup(props SidebarGroupProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "sidebar-group"),
		h.Data("sidebar", "group"),
		JoinAttrs("class",
			h.Class(sidebarGroupBaseClass),
			Group(children),
		),
	)
}

const sidebarGroupBaseClass = "relative flex w-full min-w-0 flex-col p-2"

// SidebarGroupLabelProps for [SidebarGroupLabel].
type SidebarGroupLabelProps struct{}

// SidebarGroupLabel renders a label for a sidebar group.
func SidebarGroupLabel(props SidebarGroupLabelProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "sidebar-group-label"),
		h.Data("sidebar", "group-label"),
		JoinAttrs("class",
			h.Class(sidebarGroupLabelBaseClass),
			Group(children),
		),
	)
}

const sidebarGroupLabelBaseClass = "text-sidebar-foreground/70 ring-sidebar-ring flex h-8 shrink-0 items-center rounded-md px-2 text-xs font-medium outline-hidden transition-[margin,opacity] duration-200 ease-linear focus-visible:ring-2 [&>svg]:size-4 [&>svg]:shrink-0 group-data-[collapsible=icon]:-mt-8 group-data-[collapsible=icon]:opacity-0"

// SidebarGroupContentProps for [SidebarGroupContent].
type SidebarGroupContentProps struct{}

// SidebarGroupContent renders the content area of a sidebar group.
func SidebarGroupContent(props SidebarGroupContentProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "sidebar-group-content"),
		h.Data("sidebar", "group-content"),
		JoinAttrs("class",
			h.Class(sidebarGroupContentBaseClass),
			Group(children),
		),
	)
}

const sidebarGroupContentBaseClass = "w-full text-sm"

// SidebarMenuProps for [SidebarMenu].
type SidebarMenuProps struct{}

// SidebarMenu renders a menu container for sidebar items.
func SidebarMenu(props SidebarMenuProps, children ...Node) Node {
	return h.Ul(
		h.Data("slot", "sidebar-menu"),
		h.Data("sidebar", "menu"),
		JoinAttrs("class",
			h.Class(sidebarMenuBaseClass),
			Group(children),
		),
	)
}

const sidebarMenuBaseClass = "flex w-full min-w-0 flex-col gap-1"

// SidebarMenuItemProps for [SidebarMenuItem].
type SidebarMenuItemProps struct{}

// SidebarMenuItem renders a single menu item container.
func SidebarMenuItem(props SidebarMenuItemProps, children ...Node) Node {
	return h.Li(
		h.Data("slot", "sidebar-menu-item"),
		h.Data("sidebar", "menu-item"),
		JoinAttrs("class",
			h.Class(sidebarMenuItemBaseClass),
			Group(children),
		),
	)
}

const sidebarMenuItemBaseClass = "group/menu-item relative"

// SidebarMenuButtonSize represents the size of a sidebar menu button.
type SidebarMenuButtonSize string

const (
	SidebarMenuButtonSizeDefault SidebarMenuButtonSize = "default"
	SidebarMenuButtonSizeSm      SidebarMenuButtonSize = "sm"
	SidebarMenuButtonSizeLg      SidebarMenuButtonSize = "lg"
)

// SidebarMenuButtonProps for [SidebarMenuButton].
type SidebarMenuButtonProps struct {
	// IsActive indicates if this menu item is currently active.
	IsActive bool
	// Size controls the button size.
	// Defaults to "default".
	Size SidebarMenuButtonSize
}

// SidebarMenuButton renders a button within a sidebar menu item.
func SidebarMenuButton(props SidebarMenuButtonProps, children ...Node) Node {
	size := props.Size
	if size == "" {
		size = SidebarMenuButtonSizeDefault
	}

	var sizeClass string
	switch size {
	case SidebarMenuButtonSizeSm:
		sizeClass = "h-7 text-xs"
	case SidebarMenuButtonSizeLg:
		sizeClass = "h-12 text-sm group-data-[collapsible=icon]:p-0!"
	default:
		sizeClass = "h-8 text-sm"
	}

	attrs := []Node{
		h.Data("slot", "sidebar-menu-button"),
		h.Data("sidebar", "menu-button"),
		h.Data("size", string(size)),
	}

	if props.IsActive {
		attrs = append(attrs, h.Data("active", "true"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(sidebarMenuButtonBaseClass+" "+sizeClass),
			Group(children),
		),
	)

	return h.Button(attrs...)
}

const sidebarMenuButtonBaseClass = "peer/menu-button flex w-full items-center gap-2 overflow-hidden rounded-md p-2 text-left outline-hidden ring-sidebar-ring transition-[width,height,padding] hover:bg-sidebar-accent hover:text-sidebar-accent-foreground focus-visible:ring-2 active:bg-sidebar-accent active:text-sidebar-accent-foreground disabled:pointer-events-none disabled:opacity-50 data-[active=true]:bg-sidebar-accent data-[active=true]:font-medium data-[active=true]:text-sidebar-accent-foreground group-data-[collapsible=icon]:size-8! group-data-[collapsible=icon]:p-2! [&>span:last-child]:truncate [&>svg]:size-4 [&>svg]:shrink-0 cursor-pointer"

// SidebarSeparatorProps for [SidebarSeparator].
type SidebarSeparatorProps struct{}

// SidebarSeparator renders a visual separator in the sidebar.
func SidebarSeparator(props SidebarSeparatorProps, children ...Node) Node {
	return h.Div(
		h.Role("separator"),
		h.Data("slot", "sidebar-separator"),
		h.Data("sidebar", "separator"),
		JoinAttrs("class",
			h.Class(sidebarSeparatorBaseClass),
			Group(children),
		),
	)
}

const sidebarSeparatorBaseClass = "bg-sidebar-border mx-2 my-2 h-px w-auto"

// SidebarInsetProps for [SidebarInset].
type SidebarInsetProps struct{}

// SidebarInset renders the main content area next to the sidebar.
func SidebarInset(props SidebarInsetProps, children ...Node) Node {
	return h.Main(
		h.Data("slot", "sidebar-inset"),
		JoinAttrs("class",
			h.Class(sidebarInsetBaseClass),
			Group(children),
		),
	)
}

const sidebarInsetBaseClass = "bg-background relative flex w-full flex-1 flex-col md:peer-data-[variant=inset]:m-2 md:peer-data-[variant=inset]:ml-0 md:peer-data-[variant=inset]:rounded-xl md:peer-data-[variant=inset]:shadow-sm md:peer-data-[variant=inset]:peer-data-[state=collapsed]:ml-2"
