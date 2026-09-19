// File: tutorial_viewmodel.go
//
// Responsibility:
//   - Define the presentation state required by the thematic
//     tutorial pages in the UConversor Learning area.
//   - Identify Learning as the active Web section.
//   - Identify which static tutorial must be rendered.
//
// Receives:
//   - A controlled tutorial identifier resolved by the Web layer.
//
// Produces:
//   - TutorialViewModel consumed by tutorial.html.
//
// Previous logical stage:
//   - Tutorial routing and handlers.go.
//
// Next logical stage:
//   - render.go and tutorial.html.
//
// Important restrictions:
//   - Must not perform conversions.
//   - Must not call app.App or the conversion Engine.
//   - Must not inspect Registry or Catalog.
//   - Must not contain scientific conversion factors.
//   - Must not contain tutorial educational content.
//   - Tutorial identifiers must originate from a controlled
//     Web-side mapping and never directly from arbitrary user input.
package web

// TutorialViewModel contains only the presentation state required
// to render one static Learning tutorial.
type TutorialViewModel struct {
	ActivePage string
	Tutorial   string
}

// NewTutorialViewModel creates the presentation model for a
// previously validated tutorial identifier.
func NewTutorialViewModel(tutorial string) TutorialViewModel {
	return TutorialViewModel{
		ActivePage: "learning",
		Tutorial:   tutorial,
	}
}
