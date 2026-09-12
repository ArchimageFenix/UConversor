// File: magnitude.go
// Responsibility: Define physical magnitudes that are convertible only within themselves.
// Receives: Nothing at runtime.
// Produces: Stable magnitude identifiers.
// Previous logical stage: units/family.go.
// Next logical stage: units/unit.go and catalog files.
// Important restrictions: Different magnitudes must never be directly converted into each other.
package units

type Magnitude string

const (
	MagnitudeLength      Magnitude = "Longitud"
	MagnitudeMass        Magnitude = "Masa"
	MagnitudeSpeed       Magnitude = "Velocidad"
	MagnitudeVoltage     Magnitude = "Voltaje"
	MagnitudeCurrent     Magnitude = "Corriente"
	MagnitudeResistance  Magnitude = "Resistencia"
	MagnitudePower       Magnitude = "Potencia"
	MagnitudeTemperature Magnitude = "Temperatura"
	MagnitudePressure    Magnitude = "Presión"
	MagnitudeEnergy      Magnitude = "Energía"
	MagnitudeForce       Magnitude = "Fuerza"
	MagnitudeData        Magnitude = "Almacenamiento de datos"
	MagnitudePlaneAngle  Magnitude = "Ángulo plano"
	MagnitudeTime        Magnitude = "Tiempo"
	MagnitudeDataRate    Magnitude = "Tasa de transferencia de datos"
)
