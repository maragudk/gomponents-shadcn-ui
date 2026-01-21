package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// AvatarProps for [Avatar].
type AvatarProps struct{}

// Avatar renders a span element as the avatar container with shadcn/ui styling.
// Pass additional attributes (like [h.Class]) and children (typically [AvatarImage] or [AvatarFallback]) as needed.
func Avatar(props AvatarProps, children ...Node) Node {
	return h.Span(
		JoinAttrs("class",
			h.Class(avatarBaseClass),
			Group(children),
		),
	)
}

const avatarBaseClass = "relative flex size-8 shrink-0 overflow-hidden rounded-full"

// AvatarImageProps for [AvatarImage].
type AvatarImageProps struct{}

// AvatarImage renders an img element with shadcn/ui styling for use inside [Avatar].
// Pass additional attributes (like [h.Src], [h.Alt], [h.Class]) as needed.
func AvatarImage(props AvatarImageProps, children ...Node) Node {
	return h.Img(
		JoinAttrs("class",
			h.Class(avatarImageBaseClass),
			Group(children),
		),
	)
}

const avatarImageBaseClass = "aspect-square size-full"

// AvatarFallbackProps for [AvatarFallback].
type AvatarFallbackProps struct{}

// AvatarFallback renders a span element with shadcn/ui styling for use inside [Avatar] when no image is available.
// Pass additional attributes (like [h.Class]) and children (typically initials) as needed.
func AvatarFallback(props AvatarFallbackProps, children ...Node) Node {
	return h.Span(
		JoinAttrs("class",
			h.Class(avatarFallbackBaseClass),
			Group(children),
		),
	)
}

const avatarFallbackBaseClass = "bg-muted flex size-full items-center justify-center rounded-full"
