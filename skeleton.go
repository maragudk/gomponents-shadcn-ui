package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

// SkeletonProps for [Skeleton].
type SkeletonProps struct{}

// Skeleton renders a div element as a loading placeholder with shadcn/ui styling.
// Pass additional attributes (like [h.Class] for sizing) as needed.
func Skeleton(props SkeletonProps, children ...Node) Node {
	return h.Div(
		JoinAttrs("class",
			h.Class(skeletonBaseClass),
			Group(children),
		),
	)
}

const skeletonBaseClass = "bg-accent animate-pulse rounded-md"
