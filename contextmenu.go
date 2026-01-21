package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// ContextMenuProps for [ContextMenu].
type ContextMenuProps struct{}

// ContextMenu renders a container for context menu components.
// Use with Datastar for open/close state management on right-click.
func ContextMenu(props ContextMenuProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "context-menu"),
		JoinAttrs("class",
			h.Class(contextMenuBaseClass),
			Group(children),
		),
	)
}

const contextMenuBaseClass = "relative"

// ContextMenuTriggerProps for [ContextMenuTrigger].
type ContextMenuTriggerProps struct{}

// ContextMenuTrigger renders the element that triggers the context menu on right-click.
func ContextMenuTrigger(props ContextMenuTriggerProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "context-menu-trigger"),
		JoinAttrs("class",
			h.Class(contextMenuTriggerBaseClass),
			Group(children),
		),
	)
}

const contextMenuTriggerBaseClass = ""

// ContextMenuContentProps for [ContextMenuContent].
type ContextMenuContentProps struct{}

// ContextMenuContent renders the context menu content panel.
// Use with Datastar's data-show directive to control visibility.
func ContextMenuContent(props ContextMenuContentProps, children ...Node) Node {
	return h.Div(
		h.Role("menu"),
		h.Data("slot", "context-menu-content"),
		JoinAttrs("class",
			h.Class(contextMenuContentBaseClass),
			Group(children),
		),
	)
}

const contextMenuContentBaseClass = "bg-popover text-popover-foreground absolute z-50 min-w-[8rem] overflow-hidden rounded-md border p-1 shadow-md"

// ContextMenuGroupProps for [ContextMenuGroup].
type ContextMenuGroupProps struct{}

// ContextMenuGroup renders a group of menu items.
func ContextMenuGroup(props ContextMenuGroupProps, children ...Node) Node {
	return h.Div(
		h.Role("group"),
		h.Data("slot", "context-menu-group"),
		JoinAttrs("class",
			h.Class(contextMenuGroupBaseClass),
			Group(children),
		),
	)
}

const contextMenuGroupBaseClass = ""

// ContextMenuItemVariant represents the variant of a context menu item.
type ContextMenuItemVariant string

const (
	ContextMenuItemVariantDefault     ContextMenuItemVariant = "default"
	ContextMenuItemVariantDestructive ContextMenuItemVariant = "destructive"
)

// ContextMenuItemProps for [ContextMenuItem].
type ContextMenuItemProps struct {
	// Variant can be "default" or "destructive".
	Variant ContextMenuItemVariant
	// Inset adds left padding for alignment with checkbox/radio items.
	Inset bool
}

// ContextMenuItem renders a single menu item.
func ContextMenuItem(props ContextMenuItemProps, children ...Node) Node {
	attrs := []Node{
		h.Role("menuitem"),
		h.Data("slot", "context-menu-item"),
	}

	if props.Variant == ContextMenuItemVariantDestructive {
		attrs = append(attrs, h.Data("variant", "destructive"))
	}

	if props.Inset {
		attrs = append(attrs, h.Data("inset", "true"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(contextMenuItemBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const contextMenuItemBaseClass = "focus:bg-accent focus:text-accent-foreground data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 relative flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4"

// ContextMenuCheckboxItemProps for [ContextMenuCheckboxItem].
type ContextMenuCheckboxItemProps struct {
	// Checked indicates if the item is checked.
	Checked bool
}

// ContextMenuCheckboxItem renders a checkbox menu item.
func ContextMenuCheckboxItem(props ContextMenuCheckboxItemProps, children ...Node) Node {
	attrs := []Node{
		h.Role("menuitemcheckbox"),
		h.Data("slot", "context-menu-checkbox-item"),
	}

	if props.Checked {
		attrs = append(attrs, h.Aria("checked", "true"), h.Data("state", "checked"))
	} else {
		attrs = append(attrs, h.Aria("checked", "false"), h.Data("state", "unchecked"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(contextMenuCheckboxItemBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const contextMenuCheckboxItemBaseClass = "focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50"

// ContextMenuRadioGroupProps for [ContextMenuRadioGroup].
type ContextMenuRadioGroupProps struct{}

// ContextMenuRadioGroup renders a group of radio menu items.
func ContextMenuRadioGroup(props ContextMenuRadioGroupProps, children ...Node) Node {
	return h.Div(
		h.Role("group"),
		h.Data("slot", "context-menu-radio-group"),
		JoinAttrs("class",
			h.Class(contextMenuRadioGroupBaseClass),
			Group(children),
		),
	)
}

const contextMenuRadioGroupBaseClass = ""

// ContextMenuRadioItemProps for [ContextMenuRadioItem].
type ContextMenuRadioItemProps struct {
	// Checked indicates if the item is selected.
	Checked bool
}

// ContextMenuRadioItem renders a radio menu item.
func ContextMenuRadioItem(props ContextMenuRadioItemProps, children ...Node) Node {
	attrs := []Node{
		h.Role("menuitemradio"),
		h.Data("slot", "context-menu-radio-item"),
	}

	if props.Checked {
		attrs = append(attrs, h.Aria("checked", "true"), h.Data("state", "checked"))
	} else {
		attrs = append(attrs, h.Aria("checked", "false"), h.Data("state", "unchecked"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(contextMenuRadioItemBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const contextMenuRadioItemBaseClass = "focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50"

// ContextMenuLabelProps for [ContextMenuLabel].
type ContextMenuLabelProps struct {
	// Inset adds left padding for alignment with checkbox/radio items.
	Inset bool
}

// ContextMenuLabel renders a label in the context menu.
func ContextMenuLabel(props ContextMenuLabelProps, children ...Node) Node {
	attrs := []Node{
		h.Data("slot", "context-menu-label"),
	}

	if props.Inset {
		attrs = append(attrs, h.Data("inset", "true"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(contextMenuLabelBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const contextMenuLabelBaseClass = "text-foreground px-2 py-1.5 text-sm font-medium data-[inset]:pl-8"

// ContextMenuSeparatorProps for [ContextMenuSeparator].
type ContextMenuSeparatorProps struct{}

// ContextMenuSeparator renders a separator line in the context menu.
func ContextMenuSeparator(props ContextMenuSeparatorProps, children ...Node) Node {
	return h.Div(
		h.Role("separator"),
		h.Data("slot", "context-menu-separator"),
		JoinAttrs("class",
			h.Class(contextMenuSeparatorBaseClass),
			Group(children),
		),
	)
}

const contextMenuSeparatorBaseClass = "bg-border -mx-1 my-1 h-px"

// ContextMenuShortcutProps for [ContextMenuShortcut].
type ContextMenuShortcutProps struct{}

// ContextMenuShortcut renders a keyboard shortcut hint in a menu item.
func ContextMenuShortcut(props ContextMenuShortcutProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "context-menu-shortcut"),
		JoinAttrs("class",
			h.Class(contextMenuShortcutBaseClass),
			Group(children),
		),
	)
}

const contextMenuShortcutBaseClass = "text-muted-foreground ml-auto text-xs tracking-widest"
