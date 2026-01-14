package ui_test

import (
	"strings"
	"testing"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestButton(t *testing.T) {
	t.Run("renders a button with default variant and size", func(t *testing.T) {
		b := ui.Button(ui.ButtonProps{}, g.Text("Click me"))
		html := render(t, b)

		is.True(t, strings.Contains(html, "<button"))
		is.True(t, strings.Contains(html, "Click me"))
		is.True(t, strings.Contains(html, "bg-primary"))
		is.True(t, strings.Contains(html, "h-9 px-4 py-2"))
	})

	t.Run("renders destructive variant", func(t *testing.T) {
		b := ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantDestructive}, g.Text("Delete"))
		html := render(t, b)

		is.True(t, strings.Contains(html, "bg-destructive"))
	})

	t.Run("renders outline variant", func(t *testing.T) {
		b := ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}, g.Text("Outline"))
		html := render(t, b)

		is.True(t, strings.Contains(html, "border border-input"))
	})

	t.Run("renders secondary variant", func(t *testing.T) {
		b := ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantSecondary}, g.Text("Secondary"))
		html := render(t, b)

		is.True(t, strings.Contains(html, "bg-secondary"))
	})

	t.Run("renders ghost variant", func(t *testing.T) {
		b := ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantGhost}, g.Text("Ghost"))
		html := render(t, b)

		is.True(t, strings.Contains(html, "hover:bg-accent"))
		is.True(t, !strings.Contains(html, "bg-primary"))
	})

	t.Run("renders link variant", func(t *testing.T) {
		b := ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantLink}, g.Text("Link"))
		html := render(t, b)

		is.True(t, strings.Contains(html, "underline-offset-4"))
	})

	t.Run("renders small size", func(t *testing.T) {
		b := ui.Button(ui.ButtonProps{Size: ui.ButtonSizeSm}, g.Text("Small"))
		html := render(t, b)

		is.True(t, strings.Contains(html, "h-8"))
		is.True(t, strings.Contains(html, "text-xs"))
	})

	t.Run("renders large size", func(t *testing.T) {
		b := ui.Button(ui.ButtonProps{Size: ui.ButtonSizeLg}, g.Text("Large"))
		html := render(t, b)

		is.True(t, strings.Contains(html, "h-10"))
		is.True(t, strings.Contains(html, "px-8"))
	})

	t.Run("renders icon size", func(t *testing.T) {
		b := ui.Button(ui.ButtonProps{Size: ui.ButtonSizeIcon}, g.Text("+"))
		html := render(t, b)

		is.True(t, strings.Contains(html, "h-9 w-9"))
	})

	t.Run("adds custom class", func(t *testing.T) {
		b := ui.Button(ui.ButtonProps{Class: "my-custom-class"}, g.Text("Custom"))
		html := render(t, b)

		is.True(t, strings.Contains(html, "my-custom-class"))
	})

	t.Run("accepts additional attributes", func(t *testing.T) {
		b := ui.Button(ui.ButtonProps{}, h.Type("submit"), h.Disabled(), g.Text("Submit"))
		html := render(t, b)

		is.True(t, strings.Contains(html, `type="submit"`))
		is.True(t, strings.Contains(html, "disabled"))
	})
}

func TestButtonA(t *testing.T) {
	t.Run("renders an anchor styled as button", func(t *testing.T) {
		b := ui.ButtonA(ui.ButtonProps{}, h.Href("/home"), g.Text("Go Home"))
		html := render(t, b)

		is.True(t, strings.Contains(html, "<a"))
		is.True(t, strings.Contains(html, `href="/home"`))
		is.True(t, strings.Contains(html, "Go Home"))
		is.True(t, strings.Contains(html, "bg-primary"))
	})
}

func render(t *testing.T, n g.Node) string {
	t.Helper()
	var b strings.Builder
	err := n.Render(&b)
	is.NotError(t, err)
	return b.String()
}
