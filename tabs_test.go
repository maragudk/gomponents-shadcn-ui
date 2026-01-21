package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestTabs(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Tabs(ui.TabsProps{}))
		want := `<div data-slot="tabs" class="flex flex-col gap-2"></div>`
		is.Equal(t, want, got)
	})
}

func TestTabsList(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.TabsList(ui.TabsListProps{}))
		want := `<div role="tablist" data-slot="tabs-list" class="bg-muted text-muted-foreground inline-flex h-9 w-fit items-center justify-center rounded-lg p-[3px]"></div>`
		is.Equal(t, want, got)
	})
}

func TestTabsTrigger(t *testing.T) {
	t.Run("renders inactive tab", func(t *testing.T) {
		got := render(t, ui.TabsTrigger(ui.TabsTriggerProps{}, Text("Tab 1")))
		want := `<button type="button" role="tab" data-slot="tabs-trigger" aria-selected="false" data-state="inactive" class="data-[state=active]:bg-background dark:data-[state=active]:text-foreground focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:outline-ring dark:data-[state=active]:border-input dark:data-[state=active]:bg-input/30 text-foreground dark:text-muted-foreground inline-flex h-[calc(100%-1px)] flex-1 items-center justify-center gap-1.5 rounded-md border border-transparent px-2 py-1 text-sm font-medium whitespace-nowrap transition-[color,box-shadow] focus-visible:ring-[3px] focus-visible:outline-1 disabled:pointer-events-none disabled:opacity-50 data-[state=active]:shadow-sm [&amp;_svg]:pointer-events-none [&amp;_svg]:shrink-0 [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 cursor-pointer">Tab 1</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders active tab", func(t *testing.T) {
		got := render(t, ui.TabsTrigger(ui.TabsTriggerProps{Active: true}, Text("Tab 1")))
		want := `<button type="button" role="tab" data-slot="tabs-trigger" aria-selected="true" data-state="active" class="data-[state=active]:bg-background dark:data-[state=active]:text-foreground focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:outline-ring dark:data-[state=active]:border-input dark:data-[state=active]:bg-input/30 text-foreground dark:text-muted-foreground inline-flex h-[calc(100%-1px)] flex-1 items-center justify-center gap-1.5 rounded-md border border-transparent px-2 py-1 text-sm font-medium whitespace-nowrap transition-[color,box-shadow] focus-visible:ring-[3px] focus-visible:outline-1 disabled:pointer-events-none disabled:opacity-50 data-[state=active]:shadow-sm [&amp;_svg]:pointer-events-none [&amp;_svg]:shrink-0 [&amp;_svg:not([class*=&#39;size-&#39;])]:size-4 cursor-pointer">Tab 1</button>`
		is.Equal(t, want, got)
	})
}

func TestTabsContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.TabsContent(ui.TabsContentProps{}, Text("Content")))
		want := `<div role="tabpanel" data-slot="tabs-content" class="flex-1 outline-none">Content</div>`
		is.Equal(t, want, got)
	})
}
