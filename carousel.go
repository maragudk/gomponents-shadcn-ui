package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// CarouselOrientation represents the orientation of the carousel.
type CarouselOrientation string

const (
	CarouselOrientationHorizontal CarouselOrientation = "horizontal"
	CarouselOrientationVertical   CarouselOrientation = "vertical"
)

// CarouselProps for [Carousel].
type CarouselProps struct {
	// Orientation controls the carousel direction.
	// Defaults to "horizontal".
	Orientation CarouselOrientation
}

// Carousel renders a container for carousel slides.
// Use with Datastar or JavaScript for slide navigation.
func Carousel(props CarouselProps, children ...Node) Node {
	orientation := props.Orientation
	if orientation == "" {
		orientation = CarouselOrientationHorizontal
	}

	return h.Div(
		h.Role("region"),
		h.Aria("roledescription", "carousel"),
		h.Data("slot", "carousel"),
		h.Data("orientation", string(orientation)),
		JoinAttrs("class",
			h.Class(carouselBaseClass),
			Group(children),
		),
	)
}

const carouselBaseClass = "relative"

// CarouselContentProps for [CarouselContent].
type CarouselContentProps struct {
	// Orientation should match the parent carousel orientation.
	// Defaults to "horizontal".
	Orientation CarouselOrientation
}

// CarouselContent renders the container for carousel items.
func CarouselContent(props CarouselContentProps, children ...Node) Node {
	orientation := props.Orientation
	if orientation == "" {
		orientation = CarouselOrientationHorizontal
	}

	var orientationClass string
	if orientation == CarouselOrientationVertical {
		orientationClass = "-mt-4 flex-col"
	} else {
		orientationClass = "-ml-4"
	}

	return h.Div(
		h.Data("slot", "carousel-content"),
		h.Class("overflow-hidden"),
		h.Div(
			JoinAttrs("class",
				h.Class("flex "+orientationClass),
				Group(children),
			),
		),
	)
}

// CarouselItemProps for [CarouselItem].
type CarouselItemProps struct {
	// Orientation should match the parent carousel orientation.
	// Defaults to "horizontal".
	Orientation CarouselOrientation
}

// CarouselItem renders an individual carousel slide.
func CarouselItem(props CarouselItemProps, children ...Node) Node {
	orientation := props.Orientation
	if orientation == "" {
		orientation = CarouselOrientationHorizontal
	}

	var orientationClass string
	if orientation == CarouselOrientationVertical {
		orientationClass = "pt-4"
	} else {
		orientationClass = "pl-4"
	}

	return h.Div(
		h.Role("group"),
		h.Aria("roledescription", "slide"),
		h.Data("slot", "carousel-item"),
		JoinAttrs("class",
			h.Class(carouselItemBaseClass+" "+orientationClass),
			Group(children),
		),
	)
}

const carouselItemBaseClass = "min-w-0 shrink-0 grow-0 basis-full"

// CarouselPreviousProps for [CarouselPrevious].
type CarouselPreviousProps struct {
	// Orientation should match the parent carousel orientation.
	// Defaults to "horizontal".
	Orientation CarouselOrientation
}

// CarouselPrevious renders the previous slide button.
func CarouselPrevious(props CarouselPreviousProps, children ...Node) Node {
	orientation := props.Orientation
	if orientation == "" {
		orientation = CarouselOrientationHorizontal
	}

	var orientationClass string
	if orientation == CarouselOrientationVertical {
		orientationClass = "-top-12 left-1/2 -translate-x-1/2 rotate-90"
	} else {
		orientationClass = "top-1/2 -left-12 -translate-y-1/2"
	}

	// Use the button classes from the button component
	buttonClass := buttonClasses(ButtonProps{Variant: ButtonVariantOutline, Size: ButtonSizeIcon})

	return h.Button(
		h.Type("button"),
		h.Data("slot", "carousel-previous"),
		JoinAttrs("class",
			h.Class(buttonClass+" "+carouselPreviousBaseClass+" "+orientationClass),
			Group(children),
		),
		Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m12 19-7-7 7-7"/><path d="M19 12H5"/></svg>`),
		h.Span(h.Class("sr-only"), Text("Previous slide")),
	)
}

const carouselPreviousBaseClass = "absolute size-8 rounded-full"

// CarouselNextProps for [CarouselNext].
type CarouselNextProps struct {
	// Orientation should match the parent carousel orientation.
	// Defaults to "horizontal".
	Orientation CarouselOrientation
}

// CarouselNext renders the next slide button.
func CarouselNext(props CarouselNextProps, children ...Node) Node {
	orientation := props.Orientation
	if orientation == "" {
		orientation = CarouselOrientationHorizontal
	}

	var orientationClass string
	if orientation == CarouselOrientationVertical {
		orientationClass = "-bottom-12 left-1/2 -translate-x-1/2 rotate-90"
	} else {
		orientationClass = "top-1/2 -right-12 -translate-y-1/2"
	}

	// Use the button classes from the button component
	buttonClass := buttonClasses(ButtonProps{Variant: ButtonVariantOutline, Size: ButtonSizeIcon})

	return h.Button(
		h.Type("button"),
		h.Data("slot", "carousel-next"),
		JoinAttrs("class",
			h.Class(buttonClass+" "+carouselNextBaseClass+" "+orientationClass),
			Group(children),
		),
		Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>`),
		h.Span(h.Class("sr-only"), Text("Next slide")),
	)
}

const carouselNextBaseClass = "absolute size-8 rounded-full"
