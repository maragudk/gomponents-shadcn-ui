package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// DropdownMenuProps for [DropdownMenu].
type DropdownMenuProps struct{}

// DropdownMenu renders a container for dropdown menu components.
// Use with Datastar for open/close state management.
func DropdownMenu(props DropdownMenuProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "dropdown-menu"),
		JoinAttrs("class",
			h.Class(dropdownMenuBaseClass),
			Group(children),
		),
	)
}

const dropdownMenuBaseClass = "relative inline-block"

// DropdownMenuTriggerProps for [DropdownMenuTrigger].
type DropdownMenuTriggerProps struct{}

// DropdownMenuTrigger renders the element that triggers the dropdown menu.
func DropdownMenuTrigger(props DropdownMenuTriggerProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "dropdown-menu-trigger"),
		JoinAttrs("class",
			h.Class(dropdownMenuTriggerBaseClass),
			Group(children),
		),
	)
}

const dropdownMenuTriggerBaseClass = ""

// DropdownMenuContentProps for [DropdownMenuContent].
type DropdownMenuContentProps struct{}

// DropdownMenuContent renders the dropdown menu content panel.
// Use with Datastar's data-show directive to control visibility.
func DropdownMenuContent(props DropdownMenuContentProps, children ...Node) Node {
	return h.Div(
		h.Role("menu"),
		h.Data("slot", "dropdown-menu-content"),
		JoinAttrs("class",
			h.Class(dropdownMenuContentBaseClass),
			Group(children),
		),
	)
}

const dropdownMenuContentBaseClass = "bg-popover text-popover-foreground absolute z-50 min-w-[8rem] overflow-hidden rounded-md border p-1 shadow-md top-full left-0 mt-1"

// DropdownMenuGroupProps for [DropdownMenuGroup].
type DropdownMenuGroupProps struct{}

// DropdownMenuGroup renders a group of menu items.
func DropdownMenuGroup(props DropdownMenuGroupProps, children ...Node) Node {
	return h.Div(
		h.Role("group"),
		h.Data("slot", "dropdown-menu-group"),
		JoinAttrs("class",
			h.Class(dropdownMenuGroupBaseClass),
			Group(children),
		),
	)
}

const dropdownMenuGroupBaseClass = ""

// DropdownMenuItemVariant represents the variant of a dropdown menu item.
type DropdownMenuItemVariant string

const (
	DropdownMenuItemVariantDefault     DropdownMenuItemVariant = "default"
	DropdownMenuItemVariantDestructive DropdownMenuItemVariant = "destructive"
)

// DropdownMenuItemProps for [DropdownMenuItem].
type DropdownMenuItemProps struct {
	// Variant can be "default" or "destructive".
	Variant DropdownMenuItemVariant
	// Inset adds left padding for alignment with checkbox/radio items.
	Inset bool
}

// DropdownMenuItem renders a single menu item.
func DropdownMenuItem(props DropdownMenuItemProps, children ...Node) Node {
	attrs := []Node{
		h.Role("menuitem"),
		h.Data("slot", "dropdown-menu-item"),
	}

	if props.Variant == DropdownMenuItemVariantDestructive {
		attrs = append(attrs, h.Data("variant", "destructive"))
	}

	if props.Inset {
		attrs = append(attrs, h.Data("inset", "true"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(dropdownMenuItemBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const dropdownMenuItemBaseClass = "focus:bg-accent focus:text-accent-foreground data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 relative flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4"

// DropdownMenuCheckboxItemProps for [DropdownMenuCheckboxItem].
type DropdownMenuCheckboxItemProps struct {
	// Checked indicates if the item is checked.
	Checked bool
}

// DropdownMenuCheckboxItem renders a checkbox menu item.
func DropdownMenuCheckboxItem(props DropdownMenuCheckboxItemProps, children ...Node) Node {
	attrs := []Node{
		h.Role("menuitemcheckbox"),
		h.Data("slot", "dropdown-menu-checkbox-item"),
	}

	if props.Checked {
		attrs = append(attrs, h.Aria("checked", "true"), h.Data("state", "checked"))
	} else {
		attrs = append(attrs, h.Aria("checked", "false"), h.Data("state", "unchecked"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(dropdownMenuCheckboxItemBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const dropdownMenuCheckboxItemBaseClass = "focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50"

// DropdownMenuRadioGroupProps for [DropdownMenuRadioGroup].
type DropdownMenuRadioGroupProps struct{}

// DropdownMenuRadioGroup renders a group of radio menu items.
func DropdownMenuRadioGroup(props DropdownMenuRadioGroupProps, children ...Node) Node {
	return h.Div(
		h.Role("group"),
		h.Data("slot", "dropdown-menu-radio-group"),
		JoinAttrs("class",
			h.Class(dropdownMenuRadioGroupBaseClass),
			Group(children),
		),
	)
}

const dropdownMenuRadioGroupBaseClass = ""

// DropdownMenuRadioItemProps for [DropdownMenuRadioItem].
type DropdownMenuRadioItemProps struct {
	// Checked indicates if the item is selected.
	Checked bool
}

// DropdownMenuRadioItem renders a radio menu item.
func DropdownMenuRadioItem(props DropdownMenuRadioItemProps, children ...Node) Node {
	attrs := []Node{
		h.Role("menuitemradio"),
		h.Data("slot", "dropdown-menu-radio-item"),
	}

	if props.Checked {
		attrs = append(attrs, h.Aria("checked", "true"), h.Data("state", "checked"))
	} else {
		attrs = append(attrs, h.Aria("checked", "false"), h.Data("state", "unchecked"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(dropdownMenuRadioItemBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const dropdownMenuRadioItemBaseClass = "focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50"

// DropdownMenuLabelProps for [DropdownMenuLabel].
type DropdownMenuLabelProps struct {
	// Inset adds left padding for alignment with checkbox/radio items.
	Inset bool
}

// DropdownMenuLabel renders a label in the dropdown menu.
func DropdownMenuLabel(props DropdownMenuLabelProps, children ...Node) Node {
	attrs := []Node{
		h.Data("slot", "dropdown-menu-label"),
	}

	if props.Inset {
		attrs = append(attrs, h.Data("inset", "true"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(dropdownMenuLabelBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const dropdownMenuLabelBaseClass = "px-2 py-1.5 text-sm font-medium data-[inset]:pl-8"

// DropdownMenuSeparatorProps for [DropdownMenuSeparator].
type DropdownMenuSeparatorProps struct{}

// DropdownMenuSeparator renders a separator line in the dropdown menu.
func DropdownMenuSeparator(props DropdownMenuSeparatorProps, children ...Node) Node {
	return h.Div(
		h.Role("separator"),
		h.Data("slot", "dropdown-menu-separator"),
		JoinAttrs("class",
			h.Class(dropdownMenuSeparatorBaseClass),
			Group(children),
		),
	)
}

const dropdownMenuSeparatorBaseClass = "bg-border -mx-1 my-1 h-px"

// DropdownMenuShortcutProps for [DropdownMenuShortcut].
type DropdownMenuShortcutProps struct{}

// DropdownMenuShortcut renders a keyboard shortcut hint in a menu item.
func DropdownMenuShortcut(props DropdownMenuShortcutProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "dropdown-menu-shortcut"),
		JoinAttrs("class",
			h.Class(dropdownMenuShortcutBaseClass),
			Group(children),
		),
	)
}

const dropdownMenuShortcutBaseClass = "text-muted-foreground ml-auto text-xs tracking-widest"
