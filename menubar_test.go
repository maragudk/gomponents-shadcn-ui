package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestMenubar(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Menubar(ui.MenubarProps{}))
		want := `<div data-slot="menubar" class="bg-background flex h-9 items-center gap-1 rounded-md border p-1 shadow-xs"></div>`
		is.Equal(t, want, got)
	})
}

func TestMenubarMenu(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.MenubarMenu(ui.MenubarMenuProps{}))
		want := `<div data-slot="menubar-menu" class="relative"></div>`
		is.Equal(t, want, got)
	})
}

func TestMenubarTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.MenubarTrigger(ui.MenubarTriggerProps{}, Text("File")))
		want := `<button type="button" data-slot="menubar-trigger" class="focus:bg-accent focus:text-accent-foreground data-[state=open]:bg-accent data-[state=open]:text-accent-foreground flex items-center rounded-sm px-2 py-1 text-sm font-medium outline-hidden select-none cursor-pointer">File</button>`
		is.Equal(t, want, got)
	})
}

func TestMenubarContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.MenubarContent(ui.MenubarContentProps{}))
		want := `<div role="menu" data-slot="menubar-content" class="bg-popover text-popover-foreground absolute z-50 min-w-[12rem] overflow-hidden rounded-md border p-1 shadow-md top-full left-0 mt-1"></div>`
		is.Equal(t, want, got)
	})
}

func TestMenubarGroup(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.MenubarGroup(ui.MenubarGroupProps{}))
		want := `<div role="group" data-slot="menubar-group" class=""></div>`
		is.Equal(t, want, got)
	})
}

func TestMenubarItem(t *testing.T) {
	t.Run("renders default variant", func(t *testing.T) {
		got := render(t, ui.MenubarItem(ui.MenubarItemProps{}, Text("New File")))
		want := `<div role="menuitem" data-slot="menubar-item" class="focus:bg-accent focus:text-accent-foreground data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 relative flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 [&amp;_svg]:pointer-events-none [&amp;_svg]:shrink-0 [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4">New File</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders destructive variant", func(t *testing.T) {
		got := render(t, ui.MenubarItem(ui.MenubarItemProps{Variant: ui.MenubarItemVariantDestructive}, Text("Delete")))
		want := `<div role="menuitem" data-slot="menubar-item" data-variant="destructive" class="focus:bg-accent focus:text-accent-foreground data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 relative flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 [&amp;_svg]:pointer-events-none [&amp;_svg]:shrink-0 [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4">Delete</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with inset", func(t *testing.T) {
		got := render(t, ui.MenubarItem(ui.MenubarItemProps{Inset: true}, Text("Item")))
		want := `<div role="menuitem" data-slot="menubar-item" data-inset="true" class="focus:bg-accent focus:text-accent-foreground data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 relative flex cursor-pointer items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 [&amp;_svg]:pointer-events-none [&amp;_svg]:shrink-0 [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4">Item</div>`
		is.Equal(t, want, got)
	})
}

func TestMenubarCheckboxItem(t *testing.T) {
	t.Run("renders unchecked", func(t *testing.T) {
		got := render(t, ui.MenubarCheckboxItem(ui.MenubarCheckboxItemProps{}, Text("Show Toolbar")))
		want := `<div role="menuitemcheckbox" data-slot="menubar-checkbox-item" aria-checked="false" data-state="unchecked" class="focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50">Show Toolbar</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders checked", func(t *testing.T) {
		got := render(t, ui.MenubarCheckboxItem(ui.MenubarCheckboxItemProps{Checked: true}, Text("Show Toolbar")))
		want := `<div role="menuitemcheckbox" data-slot="menubar-checkbox-item" aria-checked="true" data-state="checked" class="focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50">Show Toolbar</div>`
		is.Equal(t, want, got)
	})
}

func TestMenubarRadioGroup(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.MenubarRadioGroup(ui.MenubarRadioGroupProps{}))
		want := `<div role="group" data-slot="menubar-radio-group" class=""></div>`
		is.Equal(t, want, got)
	})
}

func TestMenubarRadioItem(t *testing.T) {
	t.Run("renders unchecked", func(t *testing.T) {
		got := render(t, ui.MenubarRadioItem(ui.MenubarRadioItemProps{}, Text("Option A")))
		want := `<div role="menuitemradio" data-slot="menubar-radio-item" aria-checked="false" data-state="unchecked" class="focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50">Option A</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders checked", func(t *testing.T) {
		got := render(t, ui.MenubarRadioItem(ui.MenubarRadioItemProps{Checked: true}, Text("Option A")))
		want := `<div role="menuitemradio" data-slot="menubar-radio-item" aria-checked="true" data-state="checked" class="focus:bg-accent focus:text-accent-foreground relative flex cursor-pointer items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50">Option A</div>`
		is.Equal(t, want, got)
	})
}

func TestMenubarLabel(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.MenubarLabel(ui.MenubarLabelProps{}, Text("Label")))
		want := `<div data-slot="menubar-label" class="text-foreground px-2 py-1.5 text-sm font-medium data-[inset]:pl-8">Label</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with inset", func(t *testing.T) {
		got := render(t, ui.MenubarLabel(ui.MenubarLabelProps{Inset: true}, Text("Label")))
		want := `<div data-slot="menubar-label" data-inset="true" class="text-foreground px-2 py-1.5 text-sm font-medium data-[inset]:pl-8">Label</div>`
		is.Equal(t, want, got)
	})
}

func TestMenubarSeparator(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.MenubarSeparator(ui.MenubarSeparatorProps{}))
		want := `<div role="separator" data-slot="menubar-separator" class="bg-border -mx-1 my-1 h-px"></div>`
		is.Equal(t, want, got)
	})
}

func TestMenubarShortcut(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.MenubarShortcut(ui.MenubarShortcutProps{}, Text("⌘N")))
		want := `<span data-slot="menubar-shortcut" class="text-muted-foreground ml-auto text-xs tracking-widest">⌘N</span>`
		is.Equal(t, want, got)
	})
}

func TestMenubarSub(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.MenubarSub(ui.MenubarSubProps{}))
		want := `<div data-slot="menubar-sub" class="relative"></div>`
		is.Equal(t, want, got)
	})
}

func TestMenubarSubTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.MenubarSubTrigger(ui.MenubarSubTriggerProps{}, Text("More Options")))
		want := `<div data-slot="menubar-sub-trigger" class="focus:bg-accent focus:text-accent-foreground data-[state=open]:bg-accent data-[state=open]:text-accent-foreground flex cursor-pointer items-center rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[inset]:pl-8">More Options</div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with inset", func(t *testing.T) {
		got := render(t, ui.MenubarSubTrigger(ui.MenubarSubTriggerProps{Inset: true}, Text("More Options")))
		want := `<div data-slot="menubar-sub-trigger" data-inset="true" class="focus:bg-accent focus:text-accent-foreground data-[state=open]:bg-accent data-[state=open]:text-accent-foreground flex cursor-pointer items-center rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[inset]:pl-8">More Options</div>`
		is.Equal(t, want, got)
	})
}

func TestMenubarSubContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.MenubarSubContent(ui.MenubarSubContentProps{}))
		want := `<div role="menu" data-slot="menubar-sub-content" class="bg-popover text-popover-foreground absolute z-50 min-w-[8rem] overflow-hidden rounded-md border p-1 shadow-lg left-full top-0 ml-1"></div>`
		is.Equal(t, want, got)
	})
}
