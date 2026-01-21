package ui_test

import (
	"testing"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestAvatar(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Avatar(ui.AvatarProps{}))
		want := `<span class="relative flex size-8 shrink-0 overflow-hidden rounded-full"></span>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.Avatar(ui.AvatarProps{}, Class("size-12")))
		want := `<span class="relative flex size-8 shrink-0 overflow-hidden rounded-full size-12"></span>`
		is.Equal(t, want, got)
	})

	t.Run("renders with image child", func(t *testing.T) {
		got := render(t, ui.Avatar(ui.AvatarProps{},
			ui.AvatarImage(ui.AvatarImageProps{}, Src("/avatar.png"), Alt("User")),
		))
		want := `<span class="relative flex size-8 shrink-0 overflow-hidden rounded-full"><img class="aspect-square size-full" src="/avatar.png" alt="User"></span>`
		is.Equal(t, want, got)
	})

	t.Run("renders with fallback child", func(t *testing.T) {
		got := render(t, ui.Avatar(ui.AvatarProps{},
			ui.AvatarFallback(ui.AvatarFallbackProps{}, Text("JD")),
		))
		want := `<span class="relative flex size-8 shrink-0 overflow-hidden rounded-full"><span class="bg-muted flex size-full items-center justify-center rounded-full">JD</span></span>`
		is.Equal(t, want, got)
	})
}

func TestAvatarImage(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AvatarImage(ui.AvatarImageProps{}, Src("/test.png"), Alt("Test")))
		want := `<img class="aspect-square size-full" src="/test.png" alt="Test">`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.AvatarImage(ui.AvatarImageProps{}, Class("object-cover"), Src("/test.png")))
		want := `<img class="aspect-square size-full object-cover" src="/test.png">`
		is.Equal(t, want, got)
	})
}

func TestAvatarFallback(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.AvatarFallback(ui.AvatarFallbackProps{}, Text("AB")))
		want := `<span class="bg-muted flex size-full items-center justify-center rounded-full">AB</span>`
		is.Equal(t, want, got)
	})

	t.Run("merges custom class into single attribute", func(t *testing.T) {
		got := render(t, ui.AvatarFallback(ui.AvatarFallbackProps{}, Class("text-lg"), Text("CD")))
		want := `<span class="bg-muted flex size-full items-center justify-center rounded-full text-lg">CD</span>`
		is.Equal(t, want, got)
	})
}
