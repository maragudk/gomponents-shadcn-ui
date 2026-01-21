package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestContextMenu(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ContextMenu(ui.ContextMenuProps{}))
		want := `<div data-slot="context-menu" class="relative"></div>`
		is.Equal(t, want, got)
	})
}

func TestContextMenuTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ContextMenuTrigger(ui.ContextMenuTriggerProps{}, Text("Right click me")))
		want := `<div data-slot="context-menu-trigger" class="">Right click me</div>`
		is.Equal(t, want, got)
	})
}

func TestContextMenuContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ContextMenuContent(ui.ContextMenuContentProps{}))
		want := `<div role="menu" data-slot="context-menu-content" class="bg-popover text-popover-foreground absolute z-50 min-w-[8rem] overflow-hidden rounded-md border p-1 shadow-md"></div>`
		is.Equal(t, want, got)
	})
}

func TestContextMenuGroup(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ContextMenuGroup(ui.ContextMenuGroupProps{}))
		want := `<div role="group" data-slot="context-menu-group" class=""></div>`
		is.Equal(t, want, got)
	})
}

func TestContextMenuItem(t *testing.T) {
	t.Run("renders default variant", func(t *testing.T) {
		got := render(t, ui.ContextMenuItem(ui.ContextMenuItemProps{}, Text("Item")))
		want := `<div role="menuitem" data-slot="context-menu-item" class="focus:bg-accent focus:text-accent-foreground data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 relative flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 [&amp;_svg]:pointer-events-none [&amp;_svg]:shrink-0 [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4">Item</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders destructive variant", func(t *testing.T) {
		got := render(t, ui.ContextMenuItem(ui.ContextMenuItemProps{Variant: ui.ContextMenuItemVariantDestructive}, Text("Delete")))
		want := `<div role="menuitem" data-slot="context-menu-item" data-variant="destructive" class="focus:bg-accent focus:text-accent-foreground data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 relative flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 [&amp;_svg]:pointer-events-none [&amp;_svg]:shrink-0 [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4">Delete</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with inset", func(t *testing.T) {
		got := render(t, ui.ContextMenuItem(ui.ContextMenuItemProps{Inset: true}, Text("Item")))
		want := `<div role="menuitem" data-slot="context-menu-item" data-inset="true" class="focus:bg-accent focus:text-accent-foreground data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 relative flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 [&amp;_svg]:pointer-events-none [&amp;_svg]:shrink-0 [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4">Item</div>`
		is.Equal(t, want, got)
	})
}

func TestContextMenuCheckboxItem(t *testing.T) {
	t.Run("renders unchecked", func(t *testing.T) {
		got := render(t, ui.ContextMenuCheckboxItem(ui.ContextMenuCheckboxItemProps{}, Text("Option")))
		want := `<div role="menuitemcheckbox" data-slot="context-menu-checkbox-item" aria-checked="false" data-state="unchecked" class="focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50">Option</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders checked", func(t *testing.T) {
		got := render(t, ui.ContextMenuCheckboxItem(ui.ContextMenuCheckboxItemProps{Checked: true}, Text("Option")))
		want := `<div role="menuitemcheckbox" data-slot="context-menu-checkbox-item" aria-checked="true" data-state="checked" class="focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50">Option</div>`
		is.Equal(t, want, got)
	})
}

func TestContextMenuRadioGroup(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ContextMenuRadioGroup(ui.ContextMenuRadioGroupProps{}))
		want := `<div role="group" data-slot="context-menu-radio-group" class=""></div>`
		is.Equal(t, want, got)
	})
}

func TestContextMenuRadioItem(t *testing.T) {
	t.Run("renders unchecked", func(t *testing.T) {
		got := render(t, ui.ContextMenuRadioItem(ui.ContextMenuRadioItemProps{}, Text("Option")))
		want := `<div role="menuitemradio" data-slot="context-menu-radio-item" aria-checked="false" data-state="unchecked" class="focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50">Option</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders checked", func(t *testing.T) {
		got := render(t, ui.ContextMenuRadioItem(ui.ContextMenuRadioItemProps{Checked: true}, Text("Option")))
		want := `<div role="menuitemradio" data-slot="context-menu-radio-item" aria-checked="true" data-state="checked" class="focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50">Option</div>`
		is.Equal(t, want, got)
	})
}

func TestContextMenuLabel(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ContextMenuLabel(ui.ContextMenuLabelProps{}, Text("Label")))
		want := `<div data-slot="context-menu-label" class="text-foreground px-2 py-1.5 text-sm font-medium data-[inset]:pl-8">Label</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with inset", func(t *testing.T) {
		got := render(t, ui.ContextMenuLabel(ui.ContextMenuLabelProps{Inset: true}, Text("Label")))
		want := `<div data-slot="context-menu-label" data-inset="true" class="text-foreground px-2 py-1.5 text-sm font-medium data-[inset]:pl-8">Label</div>`
		is.Equal(t, want, got)
	})
}

func TestContextMenuSeparator(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ContextMenuSeparator(ui.ContextMenuSeparatorProps{}))
		want := `<div role="separator" data-slot="context-menu-separator" class="bg-border -mx-1 my-1 h-px"></div>`
		is.Equal(t, want, got)
	})
}

func TestContextMenuShortcut(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.ContextMenuShortcut(ui.ContextMenuShortcutProps{}, Text("⌘K")))
		want := `<span data-slot="context-menu-shortcut" class="text-muted-foreground ml-auto text-xs tracking-widest">⌘K</span>`
		is.Equal(t, want, got)
	})
}
