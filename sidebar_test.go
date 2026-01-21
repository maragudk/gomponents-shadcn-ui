package ui_test

import (
	"strings"
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestSidebarProvider(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SidebarProvider(ui.SidebarProviderProps{}))
		is.True(t, strings.Contains(got, `data-slot="sidebar-wrapper"`))
		is.True(t, strings.Contains(got, `--sidebar-width`))
	})
}

func TestSidebar(t *testing.T) {
	t.Run("renders with default left side", func(t *testing.T) {
		got := render(t, ui.Sidebar(ui.SidebarProps{}))
		is.True(t, strings.Contains(got, `data-slot="sidebar"`))
		is.True(t, strings.Contains(got, `data-side="left"`))
		is.True(t, strings.Contains(got, `data-variant="sidebar"`))
	})

	t.Run("renders with right side", func(t *testing.T) {
		got := render(t, ui.Sidebar(ui.SidebarProps{Side: ui.SidebarSideRight}))
		is.True(t, strings.Contains(got, `data-side="right"`))
	})

	t.Run("renders with floating variant", func(t *testing.T) {
		got := render(t, ui.Sidebar(ui.SidebarProps{Variant: ui.SidebarVariantFloating}))
		is.True(t, strings.Contains(got, `data-variant="floating"`))
	})
}

func TestSidebarContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SidebarContent(ui.SidebarContentProps{}, Text("Content")))
		is.True(t, strings.Contains(got, `data-slot="sidebar-content"`))
		is.True(t, strings.Contains(got, `overflow-auto`))
	})
}

func TestSidebarHeader(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SidebarHeader(ui.SidebarHeaderProps{}, Text("Header")))
		is.True(t, strings.Contains(got, `data-slot="sidebar-header"`))
		is.True(t, strings.Contains(got, `flex flex-col gap-2 p-2`))
	})
}

func TestSidebarFooter(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SidebarFooter(ui.SidebarFooterProps{}, Text("Footer")))
		is.True(t, strings.Contains(got, `data-slot="sidebar-footer"`))
	})
}

func TestSidebarTrigger(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SidebarTrigger(ui.SidebarTriggerProps{}))
		is.True(t, strings.Contains(got, `data-slot="sidebar-trigger"`))
		is.True(t, strings.Contains(got, `type="button"`))
		is.True(t, strings.Contains(got, `Toggle Sidebar`))
	})
}

func TestSidebarGroup(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SidebarGroup(ui.SidebarGroupProps{}))
		is.True(t, strings.Contains(got, `data-slot="sidebar-group"`))
	})
}

func TestSidebarGroupLabel(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SidebarGroupLabel(ui.SidebarGroupLabelProps{}, Text("Label")))
		is.True(t, strings.Contains(got, `data-slot="sidebar-group-label"`))
	})
}

func TestSidebarGroupContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SidebarGroupContent(ui.SidebarGroupContentProps{}))
		is.True(t, strings.Contains(got, `data-slot="sidebar-group-content"`))
	})
}

func TestSidebarMenu(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SidebarMenu(ui.SidebarMenuProps{}))
		is.True(t, strings.Contains(got, `data-slot="sidebar-menu"`))
		is.True(t, strings.Contains(got, `<ul`))
	})
}

func TestSidebarMenuItem(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SidebarMenuItem(ui.SidebarMenuItemProps{}))
		is.True(t, strings.Contains(got, `data-slot="sidebar-menu-item"`))
		is.True(t, strings.Contains(got, `<li`))
	})
}

func TestSidebarMenuButton(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SidebarMenuButton(ui.SidebarMenuButtonProps{}, Text("Dashboard")))
		is.True(t, strings.Contains(got, `data-slot="sidebar-menu-button"`))
		is.True(t, strings.Contains(got, `data-size="default"`))
	})

	t.Run("renders with active state", func(t *testing.T) {
		got := render(t, ui.SidebarMenuButton(ui.SidebarMenuButtonProps{IsActive: true}, Text("Dashboard")))
		is.True(t, strings.Contains(got, `data-active="true"`))
	})

	t.Run("renders with small size", func(t *testing.T) {
		got := render(t, ui.SidebarMenuButton(ui.SidebarMenuButtonProps{Size: ui.SidebarMenuButtonSizeSm}, Text("Dashboard")))
		is.True(t, strings.Contains(got, `data-size="sm"`))
	})
}

func TestSidebarSeparator(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SidebarSeparator(ui.SidebarSeparatorProps{}))
		is.True(t, strings.Contains(got, `data-slot="sidebar-separator"`))
		is.True(t, strings.Contains(got, `role="separator"`))
	})
}

func TestSidebarInset(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.SidebarInset(ui.SidebarInsetProps{}, Text("Main content")))
		is.True(t, strings.Contains(got, `data-slot="sidebar-inset"`))
		is.True(t, strings.Contains(got, `<main`))
	})
}
