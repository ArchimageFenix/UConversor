// File: family.go
// Responsibility: Define supported high-level families.
// Receives: Nothing at runtime.
// Produces: Stable family identifiers used by catalogs and results.
// Previous logical stage: architectural domain definition.
// Next logical stage: units/magnitude.go and catalog files.
// Important restrictions: A family groups magnitudes; it does not imply convertibility between them.
package units

type Family string

const (
	FamilyLength      Family = "Longitud / distancia"
	FamilyMass        Family = "Masa"
	FamilySpeed       Family = "Velocidad"
	FamilyElectronics Family = "Electrónica"
	FamilyPhysics     Family = "Física"
	FamilyForce       Family = "Fuerza"
	FamilyData        Family = "Datos"
	FamilyAngle       Family = "Ángulos"
	FamilyTime        Family = "Tiempo"
)
