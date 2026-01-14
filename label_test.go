package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestLabel(t *testing.T) {
	t.Run("renders a label with default styling", func(t *testing.T) {
		got := render(t, ui.Label(ui.LabelProps{}, Text("Email")))
		want := `<label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">Email</label>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.Label(ui.LabelProps{}, Class("my-custom-class"), Text("Password")))
		want := `<label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70 my-custom-class">Password</label>`
		is.Equal(t, want, got)
	})

	t.Run("accepts additional attributes", func(t *testing.T) {
		got := render(t, ui.Label(ui.LabelProps{}, For("email-input"), ID("my-label"), Text("Email")))
		want := `<label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70" for="email-input" id="my-label">Email</label>`
		is.Equal(t, want, got)
	})
}
