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
		want := `<label class="flex items-center gap-2 text-sm leading-none font-medium select-none group-data-[disabled=true]:pointer-events-none group-data-[disabled=true]:opacity-50 peer-disabled:cursor-not-allowed peer-disabled:opacity-50">Email</label>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.Label(ui.LabelProps{}, Class("my-custom-class"), Text("Password")))
		want := `<label class="flex items-center gap-2 text-sm leading-none font-medium select-none group-data-[disabled=true]:pointer-events-none group-data-[disabled=true]:opacity-50 peer-disabled:cursor-not-allowed peer-disabled:opacity-50 my-custom-class">Password</label>`
		is.Equal(t, want, got)
	})

	t.Run("accepts additional attributes", func(t *testing.T) {
		got := render(t, ui.Label(ui.LabelProps{}, For("email-input"), ID("my-label"), Text("Email")))
		want := `<label class="flex items-center gap-2 text-sm leading-none font-medium select-none group-data-[disabled=true]:pointer-events-none group-data-[disabled=true]:opacity-50 peer-disabled:cursor-not-allowed peer-disabled:opacity-50" for="email-input" id="my-label">Email</label>`
		is.Equal(t, want, got)
	})
}
