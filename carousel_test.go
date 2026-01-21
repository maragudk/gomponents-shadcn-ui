package ui_test

import (
	"strings"
	"testing"

	. "maragu.dev/gomponents"
	"maragu.dev/is"

	ui "maragu.dev/gomponents-shadcn-ui"
)

func TestCarousel(t *testing.T) {
	t.Run("renders with default horizontal orientation", func(t *testing.T) {
		got := render(t, ui.Carousel(ui.CarouselProps{}))
		want := `<div role="region" aria-roledescription="carousel" data-slot="carousel" data-orientation="horizontal" class="relative"></div>`
		is.Equal(t, want, got)
	})

	t.Run("renders with vertical orientation", func(t *testing.T) {
		got := render(t, ui.Carousel(ui.CarouselProps{Orientation: ui.CarouselOrientationVertical}))
		want := `<div role="region" aria-roledescription="carousel" data-slot="carousel" data-orientation="vertical" class="relative"></div>`
		is.Equal(t, want, got)
	})
}

func TestCarouselContent(t *testing.T) {
	t.Run("renders with horizontal orientation", func(t *testing.T) {
		got := render(t, ui.CarouselContent(ui.CarouselContentProps{}))
		is.True(t, strings.Contains(got, `data-slot="carousel-content"`))
		is.True(t, strings.Contains(got, `overflow-hidden`))
		is.True(t, strings.Contains(got, `-ml-4`))
	})

	t.Run("renders with vertical orientation", func(t *testing.T) {
		got := render(t, ui.CarouselContent(ui.CarouselContentProps{Orientation: ui.CarouselOrientationVertical}))
		is.True(t, strings.Contains(got, `data-slot="carousel-content"`))
		is.True(t, strings.Contains(got, `-mt-4 flex-col`))
	})
}

func TestCarouselItem(t *testing.T) {
	t.Run("renders with horizontal orientation", func(t *testing.T) {
		got := render(t, ui.CarouselItem(ui.CarouselItemProps{}, Text("Slide 1")))
		is.True(t, strings.Contains(got, `data-slot="carousel-item"`))
		is.True(t, strings.Contains(got, `role="group"`))
		is.True(t, strings.Contains(got, `aria-roledescription="slide"`))
		is.True(t, strings.Contains(got, `pl-4`))
	})

	t.Run("renders with vertical orientation", func(t *testing.T) {
		got := render(t, ui.CarouselItem(ui.CarouselItemProps{Orientation: ui.CarouselOrientationVertical}, Text("Slide 1")))
		is.True(t, strings.Contains(got, `pt-4`))
	})
}

func TestCarouselPrevious(t *testing.T) {
	t.Run("renders with horizontal orientation", func(t *testing.T) {
		got := render(t, ui.CarouselPrevious(ui.CarouselPreviousProps{}))
		is.True(t, strings.Contains(got, `data-slot="carousel-previous"`))
		is.True(t, strings.Contains(got, `type="button"`))
		is.True(t, strings.Contains(got, `-left-12`))
		is.True(t, strings.Contains(got, `Previous slide`))
	})

	t.Run("renders with vertical orientation", func(t *testing.T) {
		got := render(t, ui.CarouselPrevious(ui.CarouselPreviousProps{Orientation: ui.CarouselOrientationVertical}))
		is.True(t, strings.Contains(got, `-top-12`))
		is.True(t, strings.Contains(got, `rotate-90`))
	})
}

func TestCarouselNext(t *testing.T) {
	t.Run("renders with horizontal orientation", func(t *testing.T) {
		got := render(t, ui.CarouselNext(ui.CarouselNextProps{}))
		is.True(t, strings.Contains(got, `data-slot="carousel-next"`))
		is.True(t, strings.Contains(got, `type="button"`))
		is.True(t, strings.Contains(got, `-right-12`))
		is.True(t, strings.Contains(got, `Next slide`))
	})

	t.Run("renders with vertical orientation", func(t *testing.T) {
		got := render(t, ui.CarouselNext(ui.CarouselNextProps{Orientation: ui.CarouselOrientationVertical}))
		is.True(t, strings.Contains(got, `-bottom-12`))
		is.True(t, strings.Contains(got, `rotate-90`))
	})
}
