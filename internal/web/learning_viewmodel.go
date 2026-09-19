// File: learning_viewmodel.go
//
// Responsibility:
//   - Define the presentation state required by the
//     UConversor Learning page.
//   - Identify the Learning page as the active Web section.
//
// Receives:
//   - No domain or scientific data.
//
// Produces:
//   - LearningViewModel consumed by the Learning HTML template.
//
// Previous logical stage:
//   - handlers.go.
//
// Next logical stage:
//   - render.go and the Learning HTML template.
//
// Important restrictions:
//   - Must not perform conversions.
//   - Must not call app.App or the conversion Engine.
//   - Must not inspect the unit Registry or catalog.
//   - Must not define scientific constants or formulas.
//   - Must not duplicate scientific or catalog information.
//   - Must remain a Web presentation model only.
package web

// LearningViewModel represents the presentation state
// required to render the UConversor Learning page.
//
// The educational content is currently static, so this model
// contains only the navigation state required by the shared header.
type LearningViewModel struct {
	ActivePage string
}

// NewLearningViewModel creates the presentation model
// for the Learning page.
//
// Produces:
//   - LearningViewModel with "learning" as the active page.
//
// Important restrictions:
//   - Does not obtain data from the application layer.
//   - Does not perform calculations.
//   - Does not access the scientific core.
func NewLearningViewModel() LearningViewModel {
	return LearningViewModel{
		ActivePage: "learning",
	}
}
