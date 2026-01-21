package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestDropdownMenu(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DropdownMenu(ui.DropdownMenuProps{}))
		want := `<div data-slot="dropdown-menu" class="relative inline-block"></div>`
		is.Equal(t, want, got)
	})
}

func TestDropdownMenuTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DropdownMenuTrigger(ui.DropdownMenuTriggerProps{}, Text("Menu")))
		want := `<span data-slot="dropdown-menu-trigger" class="">Menu</span>`
		is.Equal(t, want, got)
	})
}

func TestDropdownMenuContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DropdownMenuContent(ui.DropdownMenuContentProps{}))
		want := `<div role="menu" data-slot="dropdown-menu-content" class="bg-popover text-popover-foreground absolute z-50 min-w-[8rem] overflow-hidden rounded-md border p-1 shadow-md top-full left-0 mt-1"></div>`
		is.Equal(t, want, got)
	})
}

func TestDropdownMenuGroup(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DropdownMenuGroup(ui.DropdownMenuGroupProps{}))
		want := `<div role="group" data-slot="dropdown-menu-group" class=""></div>`
		is.Equal(t, want, got)
	})
}

func TestDropdownMenuItem(t *testing.T) {
	t.Run("renders default variant", func(t *testing.T) {
		got := render(t, ui.DropdownMenuItem(ui.DropdownMenuItemProps{}, Text("Item")))
		want := `<div role="menuitem" data-slot="dropdown-menu-item" class="focus:bg-accent focus:text-accent-foreground data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 relative flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 [&amp;_svg]:pointer-events-none [&amp;_svg]:shrink-0 [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4">Item</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders destructive variant", func(t *testing.T) {
		got := render(t, ui.DropdownMenuItem(ui.DropdownMenuItemProps{Variant: ui.DropdownMenuItemVariantDestructive}, Text("Delete")))
		want := `<div role="menuitem" data-slot="dropdown-menu-item" data-variant="destructive" class="focus:bg-accent focus:text-accent-foreground data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 relative flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 [&amp;_svg]:pointer-events-none [&amp;_svg]:shrink-0 [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4">Delete</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with inset", func(t *testing.T) {
		got := render(t, ui.DropdownMenuItem(ui.DropdownMenuItemProps{Inset: true}, Text("Item")))
		want := `<div role="menuitem" data-slot="dropdown-menu-item" data-inset="true" class="focus:bg-accent focus:text-accent-foreground data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 relative flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 [&amp;_svg]:pointer-events-none [&amp;_svg]:shrink-0 [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4">Item</div>`
		is.Equal(t, want, got)
	})
}

func TestDropdownMenuCheckboxItem(t *testing.T) {
	t.Run("renders unchecked", func(t *testing.T) {
		got := render(t, ui.DropdownMenuCheckboxItem(ui.DropdownMenuCheckboxItemProps{}, Text("Option")))
		want := `<div role="menuitemcheckbox" data-slot="dropdown-menu-checkbox-item" aria-checked="false" data-state="unchecked" class="focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50">Option</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders checked", func(t *testing.T) {
		got := render(t, ui.DropdownMenuCheckboxItem(ui.DropdownMenuCheckboxItemProps{Checked: true}, Text("Option")))
		want := `<div role="menuitemcheckbox" data-slot="dropdown-menu-checkbox-item" aria-checked="true" data-state="checked" class="focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50">Option</div>`
		is.Equal(t, want, got)
	})
}

func TestDropdownMenuRadioGroup(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DropdownMenuRadioGroup(ui.DropdownMenuRadioGroupProps{}))
		want := `<div role="group" data-slot="dropdown-menu-radio-group" class=""></div>`
		is.Equal(t, want, got)
	})
}

func TestDropdownMenuRadioItem(t *testing.T) {
	t.Run("renders unchecked", func(t *testing.T) {
		got := render(t, ui.DropdownMenuRadioItem(ui.DropdownMenuRadioItemProps{}, Text("Option")))
		want := `<div role="menuitemradio" data-slot="dropdown-menu-radio-item" aria-checked="false" data-state="unchecked" class="focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50">Option</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders checked", func(t *testing.T) {
		got := render(t, ui.DropdownMenuRadioItem(ui.DropdownMenuRadioItemProps{Checked: true}, Text("Option")))
		want := `<div role="menuitemradio" data-slot="dropdown-menu-radio-item" aria-checked="true" data-state="checked" class="focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50">Option</div>`
		is.Equal(t, want, got)
	})
}

func TestDropdownMenuLabel(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DropdownMenuLabel(ui.DropdownMenuLabelProps{}, Text("Label")))
		want := `<div data-slot="dropdown-menu-label" class="px-2 py-1.5 text-sm font-medium data-[inset]:pl-8">Label</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with inset", func(t *testing.T) {
		got := render(t, ui.DropdownMenuLabel(ui.DropdownMenuLabelProps{Inset: true}, Text("Label")))
		want := `<div data-slot="dropdown-menu-label" data-inset="true" class="px-2 py-1.5 text-sm font-medium data-[inset]:pl-8">Label</div>`
		is.Equal(t, want, got)
	})
}

func TestDropdownMenuSeparator(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DropdownMenuSeparator(ui.DropdownMenuSeparatorProps{}))
		want := `<div role="separator" data-slot="dropdown-menu-separator" class="bg-border -mx-1 my-1 h-px"></div>`
		is.Equal(t, want, got)
	})
}

func TestDropdownMenuShortcut(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.DropdownMenuShortcut(ui.DropdownMenuShortcutProps{}, Text("⌘K")))
		want := `<span data-slot="dropdown-menu-shortcut" class="text-muted-foreground ml-auto text-xs tracking-widest">⌘K</span>`
		is.Equal(t, want, got)
	})
}
