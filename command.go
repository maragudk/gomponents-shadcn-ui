package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// CommandProps for [Command].
type CommandProps struct{}

// Command renders a command palette container.
// Use with Datastar for search filtering and keyboard navigation.
func Command(props CommandProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "command"),
		JoinAttrs("class",
			h.Class(commandBaseClass),
			Group(children),
		),
	)
}

const commandBaseClass = "bg-popover text-popover-foreground flex h-full w-full flex-col overflow-hidden rounded-md"

// CommandInputProps for [CommandInput].
type CommandInputProps struct {
	Placeholder string
}

// CommandInput renders the search input for a command palette.
// Use data-bind to bind the input value to a signal for filtering.
func CommandInput(props CommandInputProps, children ...Node) Node {
	placeholder := props.Placeholder
	if placeholder == "" {
		placeholder = "Type a command or search..."
	}

	return h.Div(
		h.Data("slot", "command-input-wrapper"),
		h.Class("flex h-9 items-center gap-2 border-b px-3"),
		// Search icon
		Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4 shrink-0 opacity-50"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>`),
		h.Input(
			h.Type("text"),
			h.Data("slot", "command-input"),
			h.Placeholder(placeholder),
			JoinAttrs("class",
				h.Class(commandInputBaseClass),
				Group(children),
			),
		),
	)
}

const commandInputBaseClass = "placeholder:text-muted-foreground flex h-10 w-full rounded-md bg-transparent py-3 text-sm outline-hidden disabled:cursor-not-allowed disabled:opacity-50"

// CommandListProps for [CommandList].
type CommandListProps struct{}

// CommandList renders the scrollable list container for command items.
func CommandList(props CommandListProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "command-list"),
		h.Role("listbox"),
		JoinAttrs("class",
			h.Class(commandListBaseClass),
			Group(children),
		),
	)
}

const commandListBaseClass = "max-h-[300px] scroll-py-1 overflow-x-hidden overflow-y-auto"

// CommandEmptyProps for [CommandEmpty].
type CommandEmptyProps struct{}

// CommandEmpty renders the empty state when no results are found.
// Use with data-show to display only when the filtered list is empty.
func CommandEmpty(props CommandEmptyProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "command-empty"),
		JoinAttrs("class",
			h.Class(commandEmptyBaseClass),
			Group(children),
		),
	)
}

const commandEmptyBaseClass = "py-6 text-center text-sm"

// CommandGroupProps for [CommandGroup].
type CommandGroupProps struct {
	Heading string
}

// CommandGroup renders a group of command items with an optional heading.
func CommandGroup(props CommandGroupProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "command-group"),
		h.Role("group"),
		JoinAttrs("class",
			h.Class(commandGroupBaseClass),
			If(props.Heading != "", h.Div(
				h.Data("slot", "command-group-heading"),
				h.Class("text-muted-foreground px-2 py-1.5 text-xs font-medium"),
				Text(props.Heading),
			)),
			Group(children),
		),
	)
}

const commandGroupBaseClass = "text-foreground overflow-hidden p-1"

// CommandItemProps for [CommandItem].
type CommandItemProps struct {
	// Value is the searchable value for this item (used for filtering)
	Value string
	// Disabled makes the item non-interactive
	Disabled bool
}

// CommandItem renders a selectable command item.
// Use data-show with a filter expression to show/hide based on search.
// Use data-on:click for selection handling.
func CommandItem(props CommandItemProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "command-item"),
		h.Role("option"),
		If(props.Value != "", h.Data("value", props.Value)),
		If(props.Disabled, h.Data("disabled", "true")),
		JoinAttrs("class",
			h.Class(commandItemBaseClass),
			Group(children),
		),
	)
}

const commandItemBaseClass = "data-[selected=true]:bg-accent data-[selected=true]:text-accent-foreground [&_svg:not([class*='text-'])]:text-muted-foreground relative flex cursor-default items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled=true]:pointer-events-none data-[disabled=true]:opacity-50 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4 cursor-pointer hover:bg-accent hover:text-accent-foreground"

// CommandShortcutProps for [CommandShortcut].
type CommandShortcutProps struct{}

// CommandShortcut renders a keyboard shortcut hint.
func CommandShortcut(props CommandShortcutProps, children ...Node) Node {
	return h.Span(
		h.Data("slot", "command-shortcut"),
		JoinAttrs("class",
			h.Class(commandShortcutBaseClass),
			Group(children),
		),
	)
}

const commandShortcutBaseClass = "text-muted-foreground ml-auto text-xs tracking-widest"

// CommandSeparatorProps for [CommandSeparator].
type CommandSeparatorProps struct{}

// CommandSeparator renders a visual separator between items or groups.
func CommandSeparator(props CommandSeparatorProps, children ...Node) Node {
	return h.Div(
		h.Data("slot", "command-separator"),
		JoinAttrs("class",
			h.Class(commandSeparatorBaseClass),
			Group(children),
		),
	)
}

const commandSeparatorBaseClass = "bg-border -mx-1 h-px"
