package help

// File: content.go
// Responsibility: Define the user-facing help documentation for UConversor.
// Receives: No runtime input.
// Produces: The help manual as an ordered collection of text lines.
// Previous logical stage: CLI command detection for the "h" command.
// Next logical stage: internal/output/help/viewer.go.
// Important restrictions:
//   - This file contains documentation only.
//   - It must not read keyboard input.
//   - It must not perform conversions.
//   - It must not validate scientific formulas.
//   - Canonical unit symbols and documented aliases must remain consistent
//     with the unit catalog.

import "strings"

const manual = `UCONVERSOR - MANUAL DE AYUDA

1. PROPÓSITO

UConversor es una herramienta de conversión automática de unidades.

El usuario introduce directamente un valor seguido de su unidad y el programa
identifica automáticamente la familia y la magnitud correspondiente para
generar las conversiones compatibles disponibles.

Ejemplos:

    24mph
    5km
    40in
    2200ohm
    100N
    1TB
    1GiB
    180deg

No es necesario seleccionar previamente una categoría, indicar una unidad
destino ni navegar por menús.

La filosofía de UConversor es:

    La entrada debe expresar qué quiere convertir;
    el usuario no debe tener que explicarle al programa cómo convertirlo.


2. FORMA GENERAL DE ENTRADA

La sintaxis básica es:

    <valor><unidad>

Ejemplos:

    24mph
    40in
    500mA
    20PB
    90deg

También pueden aceptarse espacios cuando la entrada lo permita:

    24 mph
    5 km

El valor y la unidad permiten a UConversor determinar automáticamente qué
conversión debe realizarse.


3. COMANDOS DE LA CONSOLA

    h       Abrir este manual de ayuda.

Dentro del visor de ayuda están disponibles los controles de navegación
correspondientes.

Los controles exactos se muestran permanentemente en la parte inferior del
visor.


4. CONVENCIONES DE UNIDADES

UConversor conserva los símbolos científicos y técnicos definidos para las
unidades.

Las mayúsculas y minúsculas pueden tener significados diferentes y no deben
considerarse equivalentes indiscriminadamente.

Ejemplos:

    m       metro

    mV      milivoltio
    V       voltio
    kV      kilovoltio

    µA      microamperio
    mA      miliamperio
    A       amperio

    mW      milivatio
    W       vatio
    kW      kilovatio

    K       kelvin
    Pa      pascal
    kPa     kilopascal

En información digital esta distinción es especialmente importante:

    MB      megabyte
    Mb      megabit

    GB      gigabyte
    Gb      gigabit

    TB      terabyte
    Tb      terabit

    PB      petabyte
    Pb      petabit

Por esta razón UConversor no debe convertir indiscriminadamente todos los
símbolos a mayúsculas o minúsculas.


5. SÍMBOLOS CANÓNICOS Y ALIASES

UConversor puede aceptar determinadas formas alternativas de escritura para
facilitar la entrada desde un teclado convencional.

Los resultados muestran siempre el símbolo canónico de la unidad.

Ejemplo:

    Entrada:

        2200ohm

    Representación utilizada por UConversor:

        2200 Ω

Aliases admitidos para resistencia:

    ohm     -> Ω
    kohm    -> kΩ
    Mohm    -> MΩ

Para corriente también puede utilizarse:

    uA      -> µA

Para grados puede utilizarse:

    deg     -> °

Ejemplo:

    180deg

se representa internamente como:

    180°

En información digital pueden utilizarse formas compactas para bits:

    kb      -> kbit
    Mb      -> Mbit
    Gb      -> Gbit
    Tb      -> Tbit
    Pb      -> Pbit

Los aliases son únicamente mecanismos de entrada. No reemplazan los símbolos
canónicos empleados por el catálogo.


6. FAMILIAS Y MAGNITUDES DISPONIBLES


LONGITUD / DISTANCIA

Magnitud:

    Longitud

Unidades:

    mm      milímetro
    cm      centímetro
    m       metro
    km      kilómetro
    in      pulgada
    ft      pie
    yd      yarda
    mi      milla
    ly      año luz

Ejemplos:

    5km
    40in
    150cm
    3mi
    2ly


MASA

Magnitud:

    Masa

Unidades:

    mg      miligramo
    g       gramo
    kg      kilogramo
    oz      onza
    lb      libra

Ejemplos:

    500g
    25kg
    10lb
    32oz


VELOCIDAD

Magnitud:

    Velocidad

Unidades:

    m/s     metro por segundo
    km/h    kilómetro por hora
    mph     milla por hora
    kn      nudo
    ft/s    pie por segundo

Ejemplos:

    24mph
    100km/h
    10m/s
    15kn

Una entrada como:

    24mph

representa velocidad y no distancia.


ELECTRÓNICA

La familia Electrónica contiene varias magnitudes independientes.

Voltaje:

    mV      milivoltio
    V       voltio
    kV      kilovoltio

Ejemplos:

    500mV
    12V
    2kV

Corriente:

    µA      microamperio
    mA      miliamperio
    A       amperio

Alias:

    uA      -> µA

Ejemplos:

    100uA
    500mA
    2A

Resistencia:

    Ω       ohmio
    kΩ      kiloohmio
    MΩ      megaohmio

Aliases:

    ohm     -> Ω
    kohm    -> kΩ
    Mohm    -> MΩ

Ejemplo:

    2200ohm

Potencia:

    mW      milivatio
    W       vatio
    kW      kilovatio

Ejemplos:

    500mW
    100W
    2kW


FÍSICA

Temperatura:

    K       kelvin
    °C      grado Celsius
    °F      grado Fahrenheit

Las conversiones de temperatura utilizan transformaciones de escala y
desplazamiento.

Presión:

    Pa      pascal
    kPa     kilopascal
    bar     bar
    psi     libra por pulgada cuadrada

Ejemplos:

    100kPa
    1bar
    32psi

Energía:

    J       julio
    kJ      kilojulio
    cal     caloría
    kcal    kilocaloría
    Wh      vatio-hora
    kWh     kilovatio-hora

Ejemplos:

    500J
    10kJ
    250kcal
    2kWh


FUERZA

Magnitud:

    Fuerza

Unidades:

    N       newton
    kN      kilonewton
    lbf     libra-fuerza

Ejemplos:

    100N
    5kN
    10lbf

Importante:

    kN      kilonewton
    kn      nudo

Son unidades diferentes y pertenecen a magnitudes diferentes.


DATOS / INFORMACIÓN DIGITAL

Magnitud:

    Almacenamiento de datos

La familia Datos permite convertir bits, bytes, unidades decimales y unidades
binarias.


Bits:

    bit     bit
    kbit    kilobit
    Mbit    megabit
    Gbit    gigabit
    Tbit    terabit
    Pbit    petabit

Formas compactas disponibles:

    kb
    Mb
    Gb
    Tb
    Pb


Bytes decimales:

    B       byte
    kB      kilobyte
    MB      megabyte
    GB      gigabyte
    TB      terabyte
    PB      petabyte

Los prefijos decimales representan potencias de 1000:

    1 kB = 1000 B
    1 MB = 1000 kB
    1 GB = 1000 MB
    1 TB = 1000 GB
    1 PB = 1000 TB


Bytes binarios:

    KiB     kibibyte
    MiB     mebibyte
    GiB     gibibyte
    TiB     tebibyte
    PiB     pebibyte

Los prefijos binarios representan potencias de 1024:

    1 KiB = 1024 B
    1 MiB = 1024 KiB
    1 GiB = 1024 MiB
    1 TiB = 1024 GiB
    1 PiB = 1024 TiB


Relación entre bit y byte:

    1 B = 8 bit

Por ejemplo:

    1 Tbit = 125 GB

y:

    1 TB = 8 Tbit


Distinciones importantes:

    MB != Mb
    GB != Gb
    TB != Tb
    PB != Pb

También:

    GB  != GiB
    TB  != TiB
    PB  != PiB

Ejemplos válidos:

    1TB
    1Tb
    1GiB
    8Mbit
    20PB


ÁNGULOS

Magnitud:

    Ángulo plano

Unidades:

    rad     radián
    °       grado
    gon     grado centesimal
    rev     revolución

Como el símbolo ° puede ser incómodo de introducir desde el teclado, se puede
utilizar:

    deg     -> °

Ejemplos:

    90deg
    180deg
    400deg
    3.1416rad
    200gon
    1rev

Relaciones fundamentales:

    360° = 2π rad = 400 gon = 1 rev

    180° = π rad = 200 gon = 0.5 rev

    90°  = π/2 rad = 100 gon = 0.25 rev

Ejemplo:

    180deg

equivale aproximadamente a:

    3.1416 rad
    180.0000 °
    200.0000 gon
    0.5000 rev

Los ángulos no están limitados al intervalo entre 0° y 360°.

Por ejemplo:

    400deg

es un valor válido y representa una revolución completa más 40 grados.

Convertir unidades angulares y normalizar un ángulo son operaciones
matemáticas diferentes.


7. COMPATIBILIDAD ENTRE UNIDADES

UConversor solo convierte unidades pertenecientes a la misma magnitud.

Ejemplos válidos:

    mph -> km/h
    mph -> m/s

    Ω -> kΩ
    Ω -> MΩ

    N -> kN
    N -> lbf

    TB -> GiB
    Tbit -> GB

    deg -> rad
    gon -> rev

Ejemplos que no representan conversiones directas:

    mph -> km
    V -> A
    Ω -> W
    N -> kg

Aunque varias magnitudes pueden pertenecer a una misma familia, no significa
que sean convertibles entre sí.

Voltaje, corriente, resistencia y potencia, por ejemplo, pertenecen a
Electrónica pero son magnitudes diferentes.


8. AÑO LUZ

El símbolo:

    ly

representa el año luz.

El año luz es una unidad de longitud y no una unidad de tiempo.

UConversor lo integra dentro de la magnitud Longitud y permite conversiones
compatibles como:

    ly -> km
    ly -> m
    ly -> mi

y también en sentido contrario.

La relación utilizada por el catálogo debe mantenerse respaldada por
definiciones y referencias científicas documentadas por el proyecto.


9. CÓMO FUNCIONAN LAS CONVERSIONES

UConversor utiliza una unidad de referencia dentro de cada magnitud.

En una conversión lineal, conceptualmente realiza:

    entrada
       |
       v
    unidad de referencia
       |
       v
    unidades compatibles

La relación general es:

    referencia = valor * Scale

y para obtener otra unidad:

    resultado = referencia / Scale de la unidad destino


Ejemplo de longitud:

    pulgadas -> metros -> demás unidades de longitud


Ejemplo de ángulos:

    grados -> radianes -> demás unidades angulares


Ejemplo de información digital:

    PB -> bit -> demás unidades de información digital


Gracias a este diseño no es necesario programar una fórmula independiente para
cada posible pareja de unidades.

Las conversiones de temperatura utilizan además un desplazamiento:

    referencia = valor * Scale + Offset


10. PRECISIÓN Y PRESENTACIÓN

UConversor utiliza valores float64 internamente para efectuar los cálculos.

Los valores internos no se reducen anticipadamente a cuatro decimales.

El redondeo pertenece únicamente a la presentación.

Para valores ordinarios se muestran cuatro decimales:

    2.2000
    24.0000
    140.0000

Para valores extremadamente grandes o pequeños se utiliza notación científica.

Ejemplos:

    1.6000e+17
    1.0739e-16

La expresión:

    1.6000e+17

representa:

    1.6 x 10^17

La notación científica modifica únicamente la representación visual del
resultado y no la precisión interna empleada por el motor.


11. TRAZABILIDAD CIENTÍFICA

Las conversiones de UConversor deben basarse prioritariamente en relaciones
científicas reconocidas y documentadas.

El proyecto utiliza la siguiente jerarquía:

    1. Definición oficial o normativa.
    2. SI, BIPM u organismo metrológico o científico competente.
    3. Relación científica establecida.
    4. Relación derivada de definiciones validadas.
    5. Fórmula propia únicamente cuando sea necesaria.

Una fórmula propia no debe reemplazar arbitrariamente una definición oficial
o una relación científica establecida.

Las fuentes correspondientes a las unidades forman parte de la información
interna del catálogo para conservar trazabilidad.

Ejemplos de organismos y referencias utilizadas por el proyecto incluyen:

    BIPM
    NIST
    IEC

Las relaciones basadas en constantes matemáticas, como las conversiones
angulares con π, deben conservarse mediante expresiones matemáticas apropiadas
cuando sea posible.


12. UNIDADES NO DISPONIBLES

Si UConversor no reconoce una unidad, la solicitud se rechaza antes de entrar
al motor científico de conversión.

Ejemplo:

    24xyz

produce un error indicando que la unidad o familia no está disponible.

Una unidad desconocida no pasa al validador ni al motor de conversión.

Cuando esto ocurra puede utilizar:

    h

para consultar este manual.


13. EJEMPLOS RÁPIDOS

Longitud:

    5km
    40in

Velocidad:

    24mph

Masa:

    10lb

Año luz:

    3ly

Resistencia:

    2200ohm

Voltaje:

    12V

Corriente:

    500mA

Energía:

    2kWh

Fuerza:

    100N
    10lbf

Datos:

    1TB
    1Tb
    1GiB
    20PB

Ángulos:

    90deg
    180deg
    200gon
    1rev


14. REFERENCIA RÁPIDA DE ALIASES

    ohm     -> Ω
    kohm    -> kΩ
    Mohm    -> MΩ

    uA      -> µA

    deg     -> °

    kb      -> kbit
    Mb      -> Mbit
    Gb      -> Gbit
    Tb      -> Tbit
    Pb      -> Pbit


15. RECORDATORIOS IMPORTANTES

    kN != kn

    MB != Mb
    GB != Gb
    TB != Tb
    PB != Pb

    GB != GiB
    TB != TiB
    PB != PiB

    1 B = 8 bit

    1 KiB = 1024 B
    1 kB  = 1000 B

    360° = 2π rad
    360° = 400 gon
    360° = 1 rev


16. PRINCIPIO DE USO

UConversor detecta la unidad introducida, determina su familia y magnitud y
genera automáticamente las equivalencias compatibles disponibles.

El usuario proporciona la medida.

El catálogo define las unidades y sus relaciones.

El motor aplica el mecanismo general de conversión.

La presentación muestra los resultados de forma legible.

Esta separación permite ampliar UConversor con nuevas unidades y familias sin
modificar innecesariamente el motor de conversión.
`

// Lines returns the help manual separated into individual display lines.
//
// The viewer receives these lines and decides which portion is currently
// visible. Keeping pagination outside this file prevents documentation
// content from becoming coupled to terminal behavior.
func Lines() []string {
	return strings.Split(manual, "\n")
}
