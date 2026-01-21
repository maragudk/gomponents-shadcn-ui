package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestCard(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Card(ui.CardProps{}))
		want := `<div class="bg-card text-card-foreground flex flex-col gap-6 rounded-xl border py-6 shadow-sm"></div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.Card(ui.CardProps{}, Class("w-96")))
		want := `<div class="bg-card text-card-foreground flex flex-col gap-6 rounded-xl border py-6 shadow-sm w-96"></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders full card structure", func(t *testing.T) {
		got := render(t, ui.Card(ui.CardProps{},
			ui.CardHeader(ui.CardHeaderProps{},
				ui.CardTitle(ui.CardTitleProps{}, Text("Title")),
				ui.CardDescription(ui.CardDescriptionProps{}, Text("Description")),
			),
			ui.CardContent(ui.CardContentProps{}, Text("Content")),
			ui.CardFooter(ui.CardFooterProps{}, Text("Footer")),
		))
		want := `<div class="bg-card text-card-foreground flex flex-col gap-6 rounded-xl border py-6 shadow-sm"><div class="@container/card-header grid auto-rows-min grid-rows-[auto_auto] items-start gap-2 px-6 has-data-[slot=card-action]:grid-cols-[1fr_auto] [.border-b]:pb-6"><div class="leading-none font-semibold">Title</div><div class="text-muted-foreground text-sm">Description</div></div><div class="px-6">Content</div><div class="flex items-center px-6 [.border-t]:pt-6">Footer</div></div>`
		is.Equal(t, want, got)
	})
}

func TestCardHeader(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.CardHeader(ui.CardHeaderProps{}, Text("Header")))
		want := `<div class="@container/card-header grid auto-rows-min grid-rows-[auto_auto] items-start gap-2 px-6 has-data-[slot=card-action]:grid-cols-[1fr_auto] [.border-b]:pb-6">Header</div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.CardHeader(ui.CardHeaderProps{}, Class("border-b")))
		want := `<div class="@container/card-header grid auto-rows-min grid-rows-[auto_auto] items-start gap-2 px-6 has-data-[slot=card-action]:grid-cols-[1fr_auto] [.border-b]:pb-6 border-b"></div>`
		is.Equal(t, want, got)
	})
}

func TestCardTitle(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.CardTitle(ui.CardTitleProps{}, Text("My Title")))
		want := `<div class="leading-none font-semibold">My Title</div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.CardTitle(ui.CardTitleProps{}, Class("text-xl")))
		want := `<div class="leading-none font-semibold text-xl"></div>`
		is.Equal(t, want, got)
	})
}

func TestCardDescription(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.CardDescription(ui.CardDescriptionProps{}, Text("My Description")))
		want := `<div class="text-muted-foreground text-sm">My Description</div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.CardDescription(ui.CardDescriptionProps{}, Class("italic")))
		want := `<div class="text-muted-foreground text-sm italic"></div>`
		is.Equal(t, want, got)
	})
}

func TestCardAction(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.CardAction(ui.CardActionProps{}, Text("Action")))
		want := `<div class="col-start-2 row-span-2 row-start-1 self-start justify-self-end">Action</div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.CardAction(ui.CardActionProps{}, Class("gap-2")))
		want := `<div class="col-start-2 row-span-2 row-start-1 self-start justify-self-end gap-2"></div>`
		is.Equal(t, want, got)
	})
}

func TestCardContent(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.CardContent(ui.CardContentProps{}, Text("Content")))
		want := `<div class="px-6">Content</div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.CardContent(ui.CardContentProps{}, Class("space-y-4")))
		want := `<div class="px-6 space-y-4"></div>`
		is.Equal(t, want, got)
	})
}

func TestCardFooter(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.CardFooter(ui.CardFooterProps{}, Text("Footer")))
		want := `<div class="flex items-center px-6 [.border-t]:pt-6">Footer</div>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.CardFooter(ui.CardFooterProps{}, Class("justify-between")))
		want := `<div class="flex items-center px-6 [.border-t]:pt-6 justify-between"></div>`
		is.Equal(t, want, got)
	})
}
