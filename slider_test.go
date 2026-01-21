package ui_test

import (
	"testing"

	. "maragu.dev/gomponents/html"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestSlider(t *testing.T) {
	t.Run("renders with default styling", func(t *testing.T) {
		got := render(t, ui.Slider(ui.SliderProps{}))
		want := `<input type="range" data-slot="slider" class="w-full h-1.5 cursor-pointer appearance-none rounded-full bg-muted outline-none disabled:opacity-50 disabled:cursor-not-allowed [&amp;::-webkit-slider-thumb]:appearance-none [&amp;::-webkit-slider-thumb]:size-4 [&amp;::-webkit-slider-thumb]:rounded-full [&amp;::-webkit-slider-thumb]:border [&amp;::-webkit-slider-thumb]:border-primary [&amp;::-webkit-slider-thumb]:bg-white [&amp;::-webkit-slider-thumb]:shadow-sm [&amp;::-webkit-slider-thumb]:transition-shadow [&amp;::-webkit-slider-thumb]:hover:ring-4 [&amp;::-webkit-slider-thumb]:hover:ring-ring/50 [&amp;::-webkit-slider-thumb]:focus:ring-4 [&amp;::-webkit-slider-thumb]:focus:ring-ring/50 [&amp;::-moz-range-thumb]:appearance-none [&amp;::-moz-range-thumb]:size-4 [&amp;::-moz-range-thumb]:rounded-full [&amp;::-moz-range-thumb]:border [&amp;::-moz-range-thumb]:border-primary [&amp;::-moz-range-thumb]:bg-white [&amp;::-moz-range-thumb]:shadow-sm [&amp;::-moz-range-thumb]:transition-shadow [&amp;::-moz-range-thumb]:hover:ring-4 [&amp;::-moz-range-thumb]:hover:ring-ring/50 [&amp;::-moz-range-thumb]:focus:ring-4 [&amp;::-moz-range-thumb]:focus:ring-ring/50">`
		is.Equal(t, want, got)
	})

	t.Run("renders with min, max, and value", func(t *testing.T) {
		got := render(t, ui.Slider(ui.SliderProps{}, Min("0"), Max("100"), Value("50")))
		want := `<input type="range" data-slot="slider" class="w-full h-1.5 cursor-pointer appearance-none rounded-full bg-muted outline-none disabled:opacity-50 disabled:cursor-not-allowed [&amp;::-webkit-slider-thumb]:appearance-none [&amp;::-webkit-slider-thumb]:size-4 [&amp;::-webkit-slider-thumb]:rounded-full [&amp;::-webkit-slider-thumb]:border [&amp;::-webkit-slider-thumb]:border-primary [&amp;::-webkit-slider-thumb]:bg-white [&amp;::-webkit-slider-thumb]:shadow-sm [&amp;::-webkit-slider-thumb]:transition-shadow [&amp;::-webkit-slider-thumb]:hover:ring-4 [&amp;::-webkit-slider-thumb]:hover:ring-ring/50 [&amp;::-webkit-slider-thumb]:focus:ring-4 [&amp;::-webkit-slider-thumb]:focus:ring-ring/50 [&amp;::-moz-range-thumb]:appearance-none [&amp;::-moz-range-thumb]:size-4 [&amp;::-moz-range-thumb]:rounded-full [&amp;::-moz-range-thumb]:border [&amp;::-moz-range-thumb]:border-primary [&amp;::-moz-range-thumb]:bg-white [&amp;::-moz-range-thumb]:shadow-sm [&amp;::-moz-range-thumb]:transition-shadow [&amp;::-moz-range-thumb]:hover:ring-4 [&amp;::-moz-range-thumb]:hover:ring-ring/50 [&amp;::-moz-range-thumb]:focus:ring-4 [&amp;::-moz-range-thumb]:focus:ring-ring/50" min="0" max="100" value="50">`
		is.Equal(t, want, got)
	})

	t.Run("renders disabled state", func(t *testing.T) {
		got := render(t, ui.Slider(ui.SliderProps{}, Disabled()))
		want := `<input type="range" data-slot="slider" class="w-full h-1.5 cursor-pointer appearance-none rounded-full bg-muted outline-none disabled:opacity-50 disabled:cursor-not-allowed [&amp;::-webkit-slider-thumb]:appearance-none [&amp;::-webkit-slider-thumb]:size-4 [&amp;::-webkit-slider-thumb]:rounded-full [&amp;::-webkit-slider-thumb]:border [&amp;::-webkit-slider-thumb]:border-primary [&amp;::-webkit-slider-thumb]:bg-white [&amp;::-webkit-slider-thumb]:shadow-sm [&amp;::-webkit-slider-thumb]:transition-shadow [&amp;::-webkit-slider-thumb]:hover:ring-4 [&amp;::-webkit-slider-thumb]:hover:ring-ring/50 [&amp;::-webkit-slider-thumb]:focus:ring-4 [&amp;::-webkit-slider-thumb]:focus:ring-ring/50 [&amp;::-moz-range-thumb]:appearance-none [&amp;::-moz-range-thumb]:size-4 [&amp;::-moz-range-thumb]:rounded-full [&amp;::-moz-range-thumb]:border [&amp;::-moz-range-thumb]:border-primary [&amp;::-moz-range-thumb]:bg-white [&amp;::-moz-range-thumb]:shadow-sm [&amp;::-moz-range-thumb]:transition-shadow [&amp;::-moz-range-thumb]:hover:ring-4 [&amp;::-moz-range-thumb]:hover:ring-ring/50 [&amp;::-moz-range-thumb]:focus:ring-4 [&amp;::-moz-range-thumb]:focus:ring-ring/50" disabled>`
		is.Equal(t, want, got)
	})
}
