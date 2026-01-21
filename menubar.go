package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// MenubarProps for [Menubar].
type MenubarProps struct{}

// Menubar renders a horizontal menu bar container.
// Use with Datastar for open/close state management of individual menus.
func Menubar(props MenubarProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "menubar"),
		JoinAttrs("class",
			h.Class(menubarBaseClass),
			Group(children),
		),
	)
}

const menubarBaseClass = "bg-background flex h-9 items-center gap-1 rounded-md border p-1 shadow-xs"

// MenubarMenuProps for [MenubarMenu].
type MenubarMenuProps struct{}

// MenubarMenu renders a container for a single menu in the menubar.
func MenubarMenu(props MenubarMenuProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "menubar-menu"),
		JoinAttrs("class",
			h.Class(menubarMenuBaseClass),
			Group(children),
		),
	)
}

const menubarMenuBaseClass = "relative"

// MenubarTriggerProps for [MenubarTrigger].
type MenubarTriggerProps struct{}

// MenubarTrigger renders a button that triggers a menu.
func MenubarTrigger(props MenubarTriggerProps, children ...Node) Node {
	return h.Button(
		h.Type("button"),
		h.Data("slot", "menubar-trigger"),
		JoinAttrs("class",
			h.Class(menubarTriggerBaseClass),
			Group(children),
		),
	)
}

const menubarTriggerBaseClass = "focus:bg-accent focus:text-accent-foreground data-[state=open]:bg-accent data-[state=open]:text-accent-foreground flex items-center rounded-sm px-2 py-1 text-sm font-medium outline-hidden select-none cursor-pointer"

// MenubarContentProps for [MenubarContent].
type MenubarContentProps struct{}

// MenubarContent renders the dropdown content of a menu.
// Use with Datastar's data-show directive to control visibility.
func MenubarContent(props MenubarContentProps, children ...Node) Node {
	return h.Div(
		h.Role("menu"),
		h.Data("slot", "menubar-content"),
		JoinAttrs("class",
			h.Class(menubarContentBaseClass),
			Group(children),
		),
	)
}

const menubarContentBaseClass = "bg-popover text-popover-foreground absolute z-50 min-w-[12rem] overflow-hidden rounded-md border p-1 shadow-md top-full left-0 mt-1"

// MenubarGroupProps for [MenubarGroup].
type MenubarGroupProps struct{}

// MenubarGroup renders a group of menu items.
func MenubarGroup(props MenubarGroupProps, children ...Node) Node {
	return h.Div(
		h.Role("group"),
		h.Data("slot", "menubar-group"),
		JoinAttrs("class",
			h.Class(menubarGroupBaseClass),
			Group(children),
		),
	)
}

const menubarGroupBaseClass = ""

// MenubarItemVariant represents the variant of a menubar item.
type MenubarItemVariant string

const (
	MenubarItemVariantDefault     MenubarItemVariant = "default"
	MenubarItemVariantDestructive MenubarItemVariant = "destructive"
)

// MenubarItemProps for [MenubarItem].
type MenubarItemProps struct {
	// Variant can be "default" or "destructive".
	Variant MenubarItemVariant
	// Inset adds left padding for alignment with checkbox/radio items.
	Inset bool
}

// MenubarItem renders a single menu item.
func MenubarItem(props MenubarItemProps, children ...Node) Node {
	attrs := []Node{
		h.Role("menuitem"),
		h.Data("slot", "menubar-item"),
	}

	if props.Variant == MenubarItemVariantDestructive {
		attrs = append(attrs, h.Data("variant", "destructive"))
	}

	if props.Inset {
		attrs = append(attrs, h.Data("inset", "true"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(menubarItemBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const menubarItemBaseClass = "focus:bg-accent focus:text-accent-foreground data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 relative flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4"

// MenubarCheckboxItemProps for [MenubarCheckboxItem].
type MenubarCheckboxItemProps struct {
	// Checked indicates if the item is checked.
	Checked bool
}

// MenubarCheckboxItem renders a checkbox menu item.
func MenubarCheckboxItem(props MenubarCheckboxItemProps, children ...Node) Node {
	attrs := []Node{
		h.Role("menuitemcheckbox"),
		h.Data("slot", "menubar-checkbox-item"),
	}

	if props.Checked {
		attrs = append(attrs, h.Aria("checked", "true"), h.Data("state", "checked"))
	} else {
		attrs = append(attrs, h.Aria("checked", "false"), h.Data("state", "unchecked"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(menubarCheckboxItemBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const menubarCheckboxItemBaseClass = "focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50"

// MenubarRadioGroupProps for [MenubarRadioGroup].
type MenubarRadioGroupProps struct{}

// MenubarRadioGroup renders a group of radio menu items.
func MenubarRadioGroup(props MenubarRadioGroupProps, children ...Node) Node {
	return h.Div(
		h.Role("group"),
		h.Data("slot", "menubar-radio-group"),
		JoinAttrs("class",
			h.Class(menubarRadioGroupBaseClass),
			Group(children),
		),
	)
}

const menubarRadioGroupBaseClass = ""

// MenubarRadioItemProps for [MenubarRadioItem].
type MenubarRadioItemProps struct {
	// Checked indicates if the item is selected.
	Checked bool
}

// MenubarRadioItem renders a radio menu item.
func MenubarRadioItem(props MenubarRadioItemProps, children ...Node) Node {
	attrs := []Node{
		h.Role("menuitemradio"),
		h.Data("slot", "menubar-radio-item"),
	}

	if props.Checked {
		attrs = append(attrs, h.Aria("checked", "true"), h.Data("state", "checked"))
	} else {
		attrs = append(attrs, h.Aria("checked", "false"), h.Data("state", "unchecked"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(menubarRadioItemBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const menubarRadioItemBaseClass = "focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50"

// MenubarLabelProps for [MenubarLabel].
type MenubarLabelProps struct {
	// Inset adds left padding for alignment with checkbox/radio items.
	Inset bool
}

// MenubarLabel renders a label in the menubar.
func MenubarLabel(props MenubarLabelProps, children ...Node) Node {
	attrs := []Node{
		h.Data("slot", "menubar-label"),
	}

	if props.Inset {
		attrs = append(attrs, h.Data("inset", "true"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(menubarLabelBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const menubarLabelBaseClass = "text-foreground px-2 py-1.5 text-sm font-medium data-[inset]:pl-8"

// MenubarSeparatorProps for [MenubarSeparator].
type MenubarSeparatorProps struct{}

// MenubarSeparator renders a separator line in the menubar.
func MenubarSeparator(props MenubarSeparatorProps, children ...Node) Node {
	return h.Div(
		h.Role("separator"),
		h.Data("slot", "menubar-separator"),
		JoinAttrs("class",
			h.Class(menubarSeparatorBaseClass),
			Group(children),
		),
	)
}

const menubarSeparatorBaseClass = "bg-border -mx-1 my-1 h-px"

// MenubarShortcutProps for [MenubarShortcut].
type MenubarShortcutProps struct{}

// MenubarShortcut renders a keyboard shortcut hint in a menu item.
func MenubarShortcut(props MenubarShortcutProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "menubar-shortcut"),
		JoinAttrs("class",
			h.Class(menubarShortcutBaseClass),
			Group(children),
		),
	)
}

const menubarShortcutBaseClass = "text-muted-foreground ml-auto text-xs tracking-widest"

// MenubarSubProps for [MenubarSub].
type MenubarSubProps struct{}

// MenubarSub renders a container for a submenu.
func MenubarSub(props MenubarSubProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "menubar-sub"),
		JoinAttrs("class",
			h.Class(menubarSubBaseClass),
			Group(children),
		),
	)
}

const menubarSubBaseClass = "relative"

// MenubarSubTriggerProps for [MenubarSubTrigger].
type MenubarSubTriggerProps struct {
	// Inset adds left padding for alignment with checkbox/radio items.
	Inset bool
}

// MenubarSubTrigger renders a trigger for a submenu.
func MenubarSubTrigger(props MenubarSubTriggerProps, children ...Node) Node {
	attrs := []Node{
		h.Data("slot", "menubar-sub-trigger"),
	}

	if props.Inset {
		attrs = append(attrs, h.Data("inset", "true"))
	}

	attrs = append(attrs,
		JoinAttrs("class",
			h.Class(menubarSubTriggerBaseClass),
			Group(children),
		),
	)

	return h.Div(attrs...)
}

const menubarSubTriggerBaseClass = "focus:bg-accent focus:text-accent-foreground data-[state=open]:bg-accent data-[state=open]:text-accent-foreground flex cursor-pointer items-center rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[inset]:pl-8"

// MenubarSubContentProps for [MenubarSubContent].
type MenubarSubContentProps struct{}

// MenubarSubContent renders the content of a submenu.
// Use with Datastar's data-show directive to control visibility.
func MenubarSubContent(props MenubarSubContentProps, children ...Node) Node {
	return h.Div(
		h.Role("menu"),
		h.Data("slot", "menubar-sub-content"),
		JoinAttrs("class",
			h.Class(menubarSubContentBaseClass),
			Group(children),
		),
	)
}

const menubarSubContentBaseClass = "bg-popover text-popover-foreground absolute z-50 min-w-[8rem] overflow-hidden rounded-md border p-1 shadow-lg left-full top-0 ml-1"
