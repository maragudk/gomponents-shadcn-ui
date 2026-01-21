package ui_test

import (
	"strings"
	"testing"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestButton(t *testing.T) {
	t.Run("renders with default variant and size", func(t *testing.T) {
		got := render(t, ui.Button(ui.ButtonProps{}, Text("Click me")))
		want := `<button class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg]:size-4 [&amp;_svg]:shrink-0 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-9 px-4 py-2">Click me</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders destructive variant", func(t *testing.T) {
		got := render(t, ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantDestructive}, Text("Delete")))
		want := `<button class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg]:size-4 [&amp;_svg]:shrink-0 bg-destructive text-destructive-foreground shadow-sm hover:bg-destructive/90 h-9 px-4 py-2">Delete</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders outline variant", func(t *testing.T) {
		got := render(t, ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}, Text("Outline")))
		want := `<button class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg]:size-4 [&amp;_svg]:shrink-0 border border-input bg-background shadow-sm hover:bg-accent hover:text-accent-foreground h-9 px-4 py-2">Outline</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders secondary variant", func(t *testing.T) {
		got := render(t, ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantSecondary}, Text("Secondary")))
		want := `<button class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg]:size-4 [&amp;_svg]:shrink-0 bg-secondary text-secondary-foreground shadow-sm hover:bg-secondary/80 h-9 px-4 py-2">Secondary</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders ghost variant", func(t *testing.T) {
		got := render(t, ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantGhost}, Text("Ghost")))
		want := `<button class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg]:size-4 [&amp;_svg]:shrink-0 hover:bg-accent hover:text-accent-foreground h-9 px-4 py-2">Ghost</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders link variant", func(t *testing.T) {
		got := render(t, ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantLink}, Text("Link")))
		want := `<button class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg]:size-4 [&amp;_svg]:shrink-0 text-primary underline-offset-4 hover:underline h-9 px-4 py-2">Link</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders small size", func(t *testing.T) {
		got := render(t, ui.Button(ui.ButtonProps{Size: ui.ButtonSizeSm}, Text("Small")))
		want := `<button class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg]:size-4 [&amp;_svg]:shrink-0 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-8 rounded-md px-3 text-xs">Small</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders large size", func(t *testing.T) {
		got := render(t, ui.Button(ui.ButtonProps{Size: ui.ButtonSizeLg}, Text("Large")))
		want := `<button class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg]:size-4 [&amp;_svg]:shrink-0 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-10 rounded-md px-8">Large</button>`
		is.Equal(t, want, got)
	})

	t.Run("renders icon size", func(t *testing.T) {
		got := render(t, ui.Button(ui.ButtonProps{Size: ui.ButtonSizeIcon}, Text("+")))
		want := `<button class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg]:size-4 [&amp;_svg]:shrink-0 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-9 w-9">+</button>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.Button(ui.ButtonProps{}, Class("my-custom-class"), Text("Custom")))
		want := `<button class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg]:size-4 [&amp;_svg]:shrink-0 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-9 px-4 py-2 my-custom-class">Custom</button>`
		is.Equal(t, want, got)
	})

	t.Run("accepts additional attributes", func(t *testing.T) {
		got := render(t, ui.Button(ui.ButtonProps{}, Type("submit"), Disabled(), Text("Submit")))
		want := `<button class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg]:size-4 [&amp;_svg]:shrink-0 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-9 px-4 py-2" type="submit" disabled>Submit</button>`
		is.Equal(t, want, got)
	})
}

func TestButtonA(t *testing.T) {
	t.Run("renders an anchor styled as button", func(t *testing.T) {
		got := render(t, ui.ButtonA(ui.ButtonProps{}, Href("/home"), Text("Go Home")))
		want := `<a class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg]:size-4 [&amp;_svg]:shrink-0 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-9 px-4 py-2" href="/home">Go Home</a>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.ButtonA(ui.ButtonProps{}, Class("my-custom-class"), Href("/"), Text("Home")))
		want := `<a class="inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&amp;_svg]:pointer-events-none [&amp;_svg]:size-4 [&amp;_svg]:shrink-0 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-9 px-4 py-2 my-custom-class" href="/">Home</a>`
		is.Equal(t, want, got)
	})
}

func render(t *testing.T, n Node) string {
	t.Helper()
	var b strings.Builder
	err := n.Render(&b)
	is.NotError(t, err)
	return b.String()
}
