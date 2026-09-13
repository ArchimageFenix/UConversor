# ARCHITECTURE.md

# UConversor — Conversor Automático de Unidades

## 1. Propósito del proyecto

Desarrollar una aplicación capaz de interpretar automáticamente una cantidad escrita por el usuario junto con su unidad y generar las conversiones equivalentes dentro de la misma magnitud física.

Ejemplo conceptual:

```text
Entrada:
24mph

Interpretación:
Valor: 24
Unidad: mph
Magnitud: velocidad

Salida:
24.0000 mph
38.6243 km/h
10.7290 m/s
20.8554 kn
35.2000 ft/s
```

El usuario no debe navegar por menús para seleccionar manualmente la unidad de origen y la unidad de destino.

La entrada debe contener suficiente información para que el sistema determine automáticamente qué magnitud se está utilizando y cuáles conversiones son compatibles.

---

## 2. Problema que resuelve

Los conversores tradicionales suelen exigir varios pasos:

1. seleccionar una categoría;
2. seleccionar una unidad de origen;
3. seleccionar una unidad de destino;
4. introducir el valor;
5. ejecutar la conversión.

Este proyecto pretende reducir ese proceso a una entrada directa:

```text
24mph
5km
2200Ω
12.5kg
```

El sistema será responsable de:

1. interpretar la entrada;
2. separar cantidad y unidad;
3. reconocer la unidad;
4. identificar la magnitud correspondiente;
5. identificar su familia;
6. seleccionar la relación matemática correcta;
7. convertir el valor;
8. generar automáticamente las conversiones compatibles;
9. presentar los resultados con cuatro decimales.

---

## 3. Principios principales

### 3.1 Entrada automática

El usuario expresará directamente:

```text
valor + unidad
```

Ejemplos:

```text
24mph
24 mph
5km
2200Ω
```

La aplicación deberá interpretar la entrada sin obligar al usuario a seleccionar previamente una categoría.

---

### 3.2 Familia y magnitud son conceptos diferentes

Una familia puede contener varias magnitudes.

Ejemplo:

```text
Electrónica
├── Voltaje
├── Corriente
├── Resistencia
└── Potencia
```

No debe permitirse convertir cantidades incompatibles únicamente porque pertenezcan a la misma familia.

Ejemplo inválido:

```text
12 V → mA
```

Voltaje y corriente pertenecen al ámbito electrónico, pero representan magnitudes físicas diferentes.

---

### 3.3 Conversión mediante unidad de referencia

Las conversiones no deben implementarse creando manualmente una fórmula independiente para cada pareja posible de unidades.

Cada magnitud deberá disponer de una unidad de referencia o de una estrategia matemática equivalente.

Ejemplo:

```text
Longitud
Unidad de referencia: metro

mi ──→ m ──→ km
          ├─→ cm
          ├─→ ft
          └─→ yd
```

Esto reduce duplicación y facilita añadir nuevas unidades.

---

### 3.4 Precisión de cálculo y presentación

Las operaciones internas deben conservar una precisión superior a la mostrada al usuario.

La salida visible utilizará cuatro decimales de precisión de presentación. Para mantener la legibilidad de cantidades extremadamente grandes o pequeñas, el formateador podrá seleccionar notación científica sin alterar el valor interno.

Política inicial de presentación:

```text
|valor| >= 1e9       → notación científica con 4 decimales
0 < |valor| < 1e-4   → notación científica con 4 decimales
resto                → notación decimal fija con 4 decimales
```

Regla:

```text
CALCULAR con precisión suficiente
        ↓
CONVERTIR
        ↓
FORMATEAR únicamente al presentar
```

La notación científica es exclusivamente una decisión de presentación. No reduce el uso de memoria, no modifica la precisión interna y no debe emplearse como mecanismo para evitar desbordamientos numéricos.

No se deben truncar prematuramente los valores intermedios.

---

## 4. Validez científica de las conversiones

La exactitud científica tiene prioridad sobre la comodidad de implementación.

Las relaciones, constantes, factores y fórmulas deberán proceder preferentemente de referencias normativas o científicas reconocidas.

Jerarquía conceptual:

```text
1. Definición normativa u oficial
        ↓
2. Definición del SI / organismo metrológico reconocido
        ↓
3. Relación científica establecida
        ↓
4. Fórmula derivada a partir de relaciones validadas
        ↓
5. Fórmula propia documentada y comprobable
```

Referencias prioritarias cuando correspondan:

- ISO / IEC, especialmente normas relativas a cantidades y unidades.
- Sistema Internacional de Unidades (SI).
- BIPM.
- Organismos metrológicos o científicos reconocidos.
- Otras fuentes técnicas oficiales pertinentes a una magnitud concreta.

### Regla fundamental

Una fórmula propia no debe sustituir una relación oficial, normativa o científicamente establecida cuando esta exista.

---

## 5. Trazabilidad de fórmulas

Las conversiones no deben almacenarse como números sin contexto.

Cada definición de unidad o fórmula deberá poder asociarse conceptualmente con metadatos como:

```text
Unidad
├── nombre
├── símbolo
├── aliases
├── familia
├── magnitud
├── unidad de referencia
├── factor o fórmula
├── tipo de relación
│   ├── oficial
│   ├── exacta
│   ├── derivada
│   └── propia
├── fuente científica/normativa
└── observaciones
```

La fuente no necesariamente deberá mostrarse siempre en la interfaz, pero deberá conservarse en la definición interna para permitir auditoría y mantenimiento.

---

## 6. Alcance inicial

La primera versión incluirá como mínimo cinco familias principales.

### 6.1 Longitud / distancia

Magnitud principal:

```text
Longitud
```

Unidades candidatas:

```text
mm
cm
m
km
in
ft
yd
mi
ly
```

El año luz (`ly`) se tratará como una unidad de longitud. Su relación deberá documentarse mediante definiciones científicas reconocidas y no mediante una constante arbitraria. Cuando el valor utilizado sea derivado de definiciones establecidas —por ejemplo, velocidad de la luz y año juliano— deberá identificarse como relación derivada y conservar la trazabilidad de las fuentes correspondientes. Las cifras divulgativas redondeadas pueden utilizarse como referencia explicativa, pero no deberán sustituir una definición científica más precisa cuando esta esté disponible.

---

### 6.2 Masa / peso de uso común

Magnitud principal:

```text
Masa
```

Unidades candidatas:

```text
mg
g
kg
oz
lb
```

Nota:

Aunque coloquialmente se utilice la palabra "peso", internamente deberá distinguirse correctamente entre masa y fuerza cuando sea científicamente necesario.

---

### 6.3 Velocidad

Magnitud principal:

```text
Velocidad
```

Unidades candidatas:

```text
m/s
km/h
mph
kn
ft/s
```

---

### 6.4 Electrónica

Familia:

```text
Electrónica
```

Magnitudes iniciales candidatas:

```text
Voltaje
Corriente
Resistencia
Potencia
```

Ejemplos de unidades:

```text
Voltaje:
mV
V
kV

Corriente:
µA
mA
A

Resistencia:
Ω
kΩ
MΩ

Potencia:
mW
W
kW
```

Cada magnitud deberá convertirse únicamente dentro de sí misma.

---

### 6.5 Física

Familia general para magnitudes físicas comunes adicionales.

Magnitudes candidatas:

```text
Temperatura
Presión
Energía
```

Ejemplos:

```text
Temperatura:
°C
°F
K

Presión:
Pa
kPa
bar
psi

Energía:
J
kJ
cal
kcal
Wh
kWh
```

La lista definitiva de unidades deberá cerrarse durante la especificación detallada de cada magnitud.

---

## 7. Flujo conceptual general

```text
ENTRADA DEL USUARIO
        ↓
NORMALIZACIÓN
        ↓
SEPARACIÓN DE VALOR Y UNIDAD
        ↓
RECONOCIMIENTO DE UNIDAD / FAMILIA / MAGNITUD
        ↓
¿ESTÁ SOPORTADA?
   ┌────┴────┐
   │         │
  NO        SÍ
   │         │
   ↓         ↓
MOSTRAR      VALIDACIÓN DEL VALOR
"NO          ↓
DISPONIBLE" SELECCIÓN DE DEFINICIÓN CIENTÍFICA
   │         ↓
   ↓        CONVERSIÓN A REPRESENTACIÓN DE REFERENCIA
  FIN        ↓
            GENERACIÓN DE CONVERSIONES COMPATIBLES
            ↓
            REDONDEO PARA PRESENTACIÓN
            ↓
            SALIDA
```

### 7.1 Regla de rechazo temprano

Si la unidad, familia o magnitud indicada por la entrada no existe en el catálogo soportado por la aplicación, el proceso debe detenerse inmediatamente.

El programa deberá informar al usuario que la conversión solicitada no está disponible.

Regla obligatoria:

```text
ENTRADA NO SOPORTADA
        ↓
MENSAJE DE NO DISPONIBILIDAD
        ↓
FIN DEL PROCESO
```

Una entrada no soportada **no debe**:

- entrar al motor de conversión;
- entrar al validador científico;
- seleccionar fórmulas;
- generar resultados parciales;
- incorporarse a una cola de procesamiento posterior;
- ejecutar trabajo innecesario después de conocerse que no está soportada.

El rechazo debe producirse en el primer punto en que el sistema pueda determinar de forma fiable que la unidad, familia o magnitud no pertenece al catálogo disponible.

---

## 8. Casos que debe distinguir el sistema

Ejemplos válidos:

```text
24mph
24 mph
5km
2.5kg
2200Ω
```

Casos que requieren tratamiento de error o reglas específicas:

```text
abc
24xyz
mph
24
```

En particular, cuando la entrada contiene una unidad, familia o magnitud que no forma parte del catálogo disponible, el comportamiento ya queda definido:

```text
24xyz
  ↓
No disponible
  ↓
FIN
```

La solicitud no debe continuar hacia las etapas de validación científica o conversión.

Casos cuya política definitiva aún debe especificarse:

```text
-15km
0km
2,5km
24 MPH
24 mi/h
```

Estos casos deberán resolverse antes de considerar cerradas las especificaciones funcionales.

---

## 9. Responsabilidades arquitectónicas

El sistema debe separar claramente las siguientes responsabilidades.

### 9.1 Entrada

Recibe la expresión escrita por el usuario.

No realiza conversiones.

---

### 9.2 Parser

Separa:

```text
valor
unidad
```

Ejemplo:

```text
"24mph"
    ↓
24
mph
```

No decide fórmulas científicas.

---

### 9.3 Normalización

Normaliza las variantes permitidas de escritura sin destruir diferencias científicamente significativas entre símbolos.

La normalización no deberá convertir globalmente todos los símbolos a mayúsculas o minúsculas, porque la capitalización puede tener significado científico.

Ejemplos:

```text
m   ≠ M
mV  ≠ MV
mW  ≠ MW
Pa
kPa
K
A
V
W
```

Los símbolos canónicos conservarán la escritura científica definida en el catálogo. La tolerancia de entrada se resolverá mediante aliases explícitos y controlados, no destruyendo la capitalización original de todos los símbolos.

Ejemplo de política de aliases:

```text
2200ohm → símbolo canónico Ω
4.7kohm → símbolo canónico kΩ
2Mohm   → símbolo canónico MΩ
```

La salida deberá mostrar el símbolo canónico aunque el usuario haya empleado un alias más fácil de escribir.

---

### 9.4 Registro de unidades

Mantiene las definiciones conocidas de:

- familias;
- magnitudes;
- unidades;
- símbolos;
- aliases;
- factores;
- fórmulas;
- fuentes.

También actúa como referencia autorizada para determinar si una entrada pertenece al conjunto de conversiones soportadas.

Si una unidad, familia o magnitud no existe en este registro, la solicitud debe marcarse como no disponible y detener su avance hacia cualquier etapa de conversión.

No debe contener lógica de interfaz.

---

### 9.5 Clasificador

Determina a qué magnitud pertenece la unidad reconocida.

Ejemplo:

```text
mph
 ↓
Velocidad
```

Si no puede asociar la entrada con una unidad, familia o magnitud soportada, debe producir un resultado de no disponibilidad que corte el flujo inmediatamente.

No debe enviar entradas desconocidas al motor de conversión para que este descubra el error posteriormente.

---

### 9.6 Motor de conversión

Recibe:

```text
valor
unidad origen
magnitud
definición validada
```

Produce las equivalencias compatibles.

El motor solo puede recibir solicitudes previamente reconocidas, clasificadas y admitidas como soportadas.

No interpreta entrada de usuario, no decide si una familia existe y no imprime resultados.

---

### 9.7 Validador científico

Garantiza que la definición utilizada posea una procedencia aceptada por las reglas del proyecto.

Debe permitir distinguir entre:

```text
relación oficial
relación exacta
relación derivada
fórmula propia
```

---

### 9.8 Formateador

Convierte los resultados internos en representación visible.

Responsabilidades principales:

- representar los resultados con cuatro decimales de presentación;
- seleccionar notación decimal fija o científica según la política definida en 3.4;
- mantener alineada y legible la salida tabular.

No debe modificar los valores usados internamente por el motor. El cambio entre notación fija y científica pertenece exclusivamente a la presentación.

---

### 9.9 Presentación

Muestra al usuario las conversiones generadas.

La implementación CLI actual puede utilizar tabla, colores ANSI, mensajes de estado y una pantalla de bienvenida de UConversor. La presentación debe permanecer separada del núcleo de conversión.

`NO_COLOR=1` podrá desactivar los colores ANSI sin modificar los resultados.

`console.go` concentra la presentación normal de la terminal. `main.go` únicamente coordina su invocación y no debe contener arte ASCII, tablas, colores ni reglas de formato.

No realiza cálculos.

---


## 9.10 Aplicación / orquestación

`internal/app/app.go` coordina el flujo principal de la aplicación.

Responsabilidades:

- recibir la entrada desde la interfaz activa;
- invocar parser y normalizador;
- consultar el registro de unidades;
- solicitar validación científica;
- ejecutar el motor de conversión;
- entregar resultados al formateador y presentador.

No debe contener fórmulas de conversión ni lógica científica específica.

---

## 9.11 Interfaces CLI y Web actuales

UConversor dispone de interfaces desacopladas que reutilizan la misma capa de aplicación y el mismo núcleo científico. La CLI continúa siendo válida y la interfaz web ya forma parte de la implementación actual.

```text
CLI ─┐
     ├→ APP → NÚCLEO DE CONVERSIÓN
WEB ─┘
```

El núcleo no debe depender de HTTP, HTML, plantillas, CSS ni decisiones propias de una interfaz concreta. La Web actúa como adaptador de presentación: recibe HTTP, delega operaciones de dominio en `app.App`, crea modelos de vista y renderiza plantillas.

### 9.11.1 Vistas web independientes

La interfaz web crecerá lateralmente mediante rutas y vistas independientes, preservando Inicio. Rutas de esta etapa:

```text
/                  → Inicio
/examples          → Ejemplos / unidades soportadas
/#resultPanel      → Resultado dentro de Inicio
/#about            → Acerca de dentro de Inicio
```

Los ejemplos rápidos del convertidor permanecen en Inicio. `/examples` será una vista documental y de descubrimiento alimentada por las unidades realmente registradas, nunca por una segunda lista manual en HTML.

```text
catalog/* → catalog.All() → units.Registry → app.App → ExamplesViewModel → /examples
```

El `Registry` continúa siendo la autoridad sobre las unidades soportadas. Podrá exponer una consulta de solo lectura que devuelva una copia de las unidades registradas; `App` expondrá esa información a las interfaces, evitando que Web dependa directamente de los detalles internos del registro.

Regla de dependencia: `Web → App → Registry → Catalog`. Se prohíbe duplicar el catálogo en HTML o hacer que Web lea directamente los archivos científicos del catálogo.

### 9.11.2 Plantillas web

Las plantillas se cargan conjuntamente mediante `template.ParseFS`; por ello, vistas independientes no deben competir redefiniendo un mismo template global `content`. Las páginas usarán templates raíz con nombres propios y reutilizarán parciales comunes como `header` y `footer`.

`PageViewModel` permanece orientado a Inicio/conversión. Ejemplos utilizará un modelo separado (`ExamplesViewModel`). No se introducirá una SPA, router JavaScript ni flags globales de página cuando una vista independiente sea suficiente.

Navegación común:

```text
Inicio      → /
Ejemplos    → /examples
Resultado   → /#resultPanel
Acerca de   → /#about
```

---

### 9.12 Sistema de ayuda interactiva de la CLI

La CLI incorporará un sistema de ayuda navegable de solo lectura para mantener la consola principal limpia y concentrar en un espacio separado la documentación de uso de UConversor.

El comando inicial será:

```text
h
```

`h` debe reconocerse en la capa CLI antes de enviar la entrada al núcleo de conversión.

```text
Entrada del usuario
        ↓
Detección de comando CLI
     ┌──┴─────────────┐
     │                │
    h            expresión
     │                │
     ↓                ↓
Help Viewer       APP / conversión
```

El sistema de ayuda tendrá dos responsabilidades separadas:

```text
help/content.go  → contenido documental
help/viewer.go   → navegación y presentación interactiva
```

`content.go` documentará como mínimo:

- propósito de UConversor;
- sintaxis de entrada;
- comandos disponibles;
- familias, magnitudes, unidades y símbolos admitidos;
- aliases de entrada;
- convenciones de mayúsculas/minúsculas;
- símbolos científicos difíciles de escribir y sus alternativas admitidas;
- política de notación científica;
- ejemplos de uso;
- principios generales de trazabilidad científica.

`viewer.go` será un visor de solo lectura inspirado en herramientas de terminal Unix/Linux. Permitirá navegar el contenido sin convertir el sistema de ayuda en un editor. La implementación inicial podrá utilizar navegación por líneas o páginas y un comando de salida claramente indicado.

El sistema de ayuda pertenece exclusivamente a la interfaz CLI y no debe realizar conversiones, validar relaciones científicas, modificar el registro ni formar parte del motor.

Cuando una unidad o familia no esté disponible, el rechazo temprano se conserva y la presentación añadirá únicamente una sugerencia:

```text
ERROR: unidad o familia no disponible: xyz
Escriba 'h' para consultar la ayuda.
```

La sugerencia no permite que la entrada desconocida avance hacia el validador científico ni hacia el motor.

---
## 10. Estructura actual de carpetas y archivos

UConversor está implementado en Go y dispone actualmente de dos interfaces funcionales: CLI y Web. Ambas reutilizan la misma capa `App` y el mismo núcleo de conversión.

La estructura relevante actual es:

```text
UConversor/
│
├── ARCHITECTURE.md
├── go.mod
│
├── cmd/
│   ├── converter/
│   │   └── main.go              # punto de entrada CLI
│   └── web/
│       └── main.go              # punto de entrada Web
│
├── internal/
│   ├── app/
│   │   └── app.go               # orquestación compartida
│   ├── input/
│   │   ├── parser.go
│   │   └── normalizer.go
│   ├── units/
│   │   ├── registry.go
│   │   ├── unit.go
│   │   ├── magnitude.go
│   │   └── family.go
│   ├── conversion/
│   │   ├── engine.go
│   │   ├── formula.go
│   │   └── validator.go
│   ├── catalog/
│   │   └── ...                  # catálogos por magnitud/familia
│   ├── output/
│   │   ├── formatter.go
│   │   ├── console.go
│   │   └── help/
│   │       ├── content.go
│   │       └── viewer.go
│   ├── model/
│   │   ├── request.go
│   │   ├── result.go
│   │   └── source.go
│   └── web/
│       ├── handlers.go
│       ├── server.go
│       ├── render.go
│       ├── templates.go
│       ├── viewmodel.go
│       ├── examples_viewmodel.go
│       └── ...                  # recursos/vistas Web asociados
│
├── tests/
└── docs/
```

La estructura sigue convenciones idiomáticas de Go y mantiene puntos de entrada mínimos. `cmd/converter` inicia la CLI y `cmd/web` inicia el servidor Web. Ninguno de los dos debe duplicar lógica científica.

La Web fue contemplada originalmente como una evolución posterior de la CLI. Esa evolución ya fue implementada: hoy ambas interfaces son operativas y reutilizan `internal/app` y el mismo núcleo de conversión. Esta nota conserva la historia de la decisión sin describir la Web como una capacidad futura.

---

## 11. Secuencia lógica entre módulos

### 11.1 Flujo CLI

```text
cmd/converter/main.go
        ↓
detección de comando CLI
   ┌────┴──────────────┐
   │                   │
   h               expresión
   │                   │
   ▼                   ▼
output/help         app/app.go
                       ↓
                 input/parser
                       ↓
                 input/normalizer
                       ↓
                 units/registry
                       ↓
                 magnitud/familia
                       ↓
              conversion/validator
                       ↓
               conversion/engine
                       ↓
               output/formatter
                       ↓
               output/console
```

Los catálogos proporcionan datos al registro y al motor:

```text
catalog/*
    ↓
units/registry
    ↓
conversion/engine
```

El camino `h → output/help` es exclusivamente de interfaz y nunca entra al pipeline científico. Las expresiones de conversión continúan utilizando el flujo normal de la aplicación.

### 11.2 Flujo Web

```text
cmd/web/main.go
      ↓
internal/web/server.go
      ↓
HTTP / rutas
   ┌──┴─────────────────────────┐
   │                            │
Conversión                  /examples
   │                            │
   ↓                            ↓
handler                    handler
   │                            │
   ↓                            ↓
App.Convert()             App.SupportedUnits()
   │                            │
   ↓                            ↓
Registry / Validator /    Registry
Engine                         │
   │                            ↓
   ↓                      ExamplesViewModel
ResultViewModel                 │
   │                            ↓
   └──────────────→ HTML / navegador
```

La interfaz Web es una capa de presentación operativa. No contiene fórmulas científicas ni mantiene un catálogo paralelo. Las rutas Web dependen de `App`; `App` conserva la frontera con el núcleo y el `Registry`.
---

## 12. Regla obligatoria para cada archivo de código

Cada archivo relevante deberá comenzar con comentarios que documenten como mínimo:

```text
File
Responsibility
Receives
Produces
Previous logical stage
Next logical stage
Important restrictions
```

Ejemplo conceptual:

```text
File: engine.<ext>

Responsibility:
    Ejecutar conversiones entre unidades compatibles.

Receives:
    Valor validado, unidad reconocida y definición de magnitud.

Produces:
    Resultados numéricos de conversión.

Previous logical stage:
    conversion/validator.<ext>

Next logical stage:
    output/formatter.<ext>

Important restrictions:
    No debe interpretar texto de entrada.
    No debe imprimir resultados.
    No debe redondear valores internos prematuramente.
```

---

## 13. Regla de modularidad

Un archivo debe tener una responsabilidad principal claramente identificable.

Ejemplo:

```text
parser
```

debe interpretar la estructura de entrada.

No debe:

- calcular conversiones;
- presentar resultados;
- definir factores científicos;
- decidir políticas globales.

Del mismo modo:

```text
formatter
```

debe encargarse de la representación final y no del cálculo.

---

## 14. Extensibilidad

Añadir una nueva unidad dentro de una magnitud existente deberá requerir modificar únicamente el catálogo o definición correspondiente y, cuando sea necesario, sus pruebas.

Ejemplo conceptual:

```text
Agregar una nueva unidad de longitud
        ↓
catalog/length
        ↓
pruebas asociadas
```

El parser, la presentación y el motor general no deberían necesitar reescribirse.

Añadir una nueva magnitud podrá requerir:

```text
nueva definición de magnitud
nuevo catálogo
registro
pruebas
```

pero no una modificación general de toda la aplicación.

---

## 15. Fórmulas propias

El sistema podrá admitir fórmulas propias cuando una necesidad legítima no esté cubierta por una definición oficial directa.

Toda fórmula propia deberá:

1. estar documentada;
2. indicar qué magnitud representa;
3. indicar su dominio de validez;
4. especificar de qué relaciones científicas se deriva;
5. incluir pruebas;
6. identificarse internamente como fórmula derivada o propia;
7. no sustituir una definición normativa existente sin justificación.

---

## 16. Estrategias de conversión

No todas las magnitudes deben asumir necesariamente una simple multiplicación por factor.

El motor deberá permitir al menos conceptualmente:

```text
Conversión lineal por factor
Conversión afín
Fórmula especializada
```

Ejemplo importante:

```text
Temperatura
```

no puede tratarse exactamente igual que una distancia porque conversiones como Celsius/Fahrenheit requieren desplazamiento además de escala.

Por ello el diseño del motor no debe limitarse a:

```text
resultado = valor * factor
```

como única estrategia posible.

---

## 17. Pruebas científicas

Las pruebas deberán validar tanto el software como las relaciones matemáticas.

Tipos mínimos:

### Parser

Verificar entradas válidas e inválidas.

### Conversión

Verificar resultados conocidos contra valores de referencia.

### Precisión

Confirmar que los cálculos internos no se limiten prematuramente a cuatro decimales.

### Formato

Confirmar presentación con cuatro decimales y el cambio controlado a notación científica para valores extremadamente grandes o pequeños.

### Fuente

Verificar que toda unidad o fórmula registrada tenga la metadata científica requerida.

### Compatibilidad

Impedir conversiones entre magnitudes distintas.

Ejemplo:

```text
V → mA
```

debe rechazarse como conversión directa de unidades.

---

## 18. Principios de mantenimiento

La arquitectura debe permitir que una persona pueda retomar el proyecto después de semanas o meses y localizar rápidamente:

- dónde se interpreta la entrada;
- dónde se define una unidad;
- dónde se almacena una fórmula;
- dónde se documenta su fuente;
- dónde se ejecuta una conversión;
- dónde se redondea la salida;
- dónde se agrega una nueva familia;
- dónde se encuentran las pruebas.

---

## 19. Regla de sincronización documental

Si una modificación cambia:

- la estructura de carpetas;
- las responsabilidades;
- el flujo;
- una dependencia;
- una estrategia de conversión;
- una política científica;
- una familia o magnitud fundamental;

`ARCHITECTURE.md` deberá actualizarse junto con el cambio.

El documento no debe quedar desfasado respecto a la implementación.

---

## 20. Principio central del proyecto

```text
INTERPRETAR correctamente
        ↓
IDENTIFICAR la magnitud
        ↓
VALIDAR la relación científica
        ↓
CALCULAR con precisión
        ↓
CONVERTIR solo entre magnitudes compatibles
        ↓
PRESENTAR con 4 decimales y notación legible cuando corresponda
```

La facilidad de uso no debe conseguirse sacrificando rigor científico.

---

# Estado actual del diseño

**Estado del documento: ARQUITECTURA ACTUAL — CLI + WEB FUNCIONALES.**

Cualquier cambio posterior que altere responsabilidades, flujo, catálogo, estructura o reglas de procesamiento deberá reflejarse mediante una actualización explícita de este documento.

## Definido

- Nombre de la aplicación: **UConversor**.
- Entrada automática basada en valor + unidad.
- Reconocimiento de magnitudes.
- Separación entre familia y magnitud.
- Al menos cinco familias iniciales.
- Conversión mediante representación de referencia cuando sea apropiado.
- Precisión interna superior a la precisión visible.
- Presentación ordinaria con cuatro decimales.
- Notación científica automática para `|valor| >= 1e9` o `0 < |valor| < 1e-4`.
- La notación científica es una decisión de presentación y no altera `float64` ni la precisión interna.
- Salida CLI estructurada mediante tabla, colores ANSI y mensajes de estado.
- Posibilidad de desactivar colores mediante `NO_COLOR=1`.
- Pantalla de bienvenida de UConversor gestionada por la capa de presentación.
- Prioridad de fórmulas y relaciones científicamente validadas.
- Posibilidad controlada de fórmulas propias.
- Trazabilidad de las fuentes y distinción entre relaciones oficiales, exactas, derivadas y propias.
- Año luz (`ly`) incorporado conceptualmente como unidad de longitud y sujeto a trazabilidad científica.
- Símbolos canónicos científicos conservados en la salida.
- Aliases controlados para facilitar la entrada sin sustituir el símbolo canónico, incluyendo `ohm → Ω`, `kohm → kΩ` y `Mohm → MΩ`.
- No se aplicará conversión global a mayúsculas o minúsculas cuando pueda destruir significado científico.
- Arquitectura modular.
- Comentarios obligatorios de responsabilidad y secuencia en archivos.
- Estructura de carpetas y archivos documentada en este archivo.
- Rechazo temprano obligatorio para unidades, familias o magnitudes no soportadas.
- Las entradas no soportadas no entran al motor, al validador científico ni a ninguna cola o etapa posterior de procesamiento.
- Sistema de ayuda interactiva accesible mediante `h`.
- Separación entre contenido de ayuda (`content.go`) y navegación (`viewer.go`).
- Los comandos CLI se interceptan antes del núcleo de conversión.
- Las entradas desconocidas sugieren `h` sin alterar el rechazo temprano.
- La documentación detallada de uso y convenciones se mantendrá fuera de la consola principal.

- Servidor web integrado sobre la misma capa `App` y núcleo científico.
- Página principal `/` preservada como vista de conversión.
- Vista independiente `/examples` definida para descubrir unidades soportadas.
- `/examples` obtiene datos mediante `App → Registry`, sin listas HTML duplicadas.
- Navegación común: `/`, `/examples`, `/#resultPanel`, `/#about`.
- Vistas web independientes reutilizan parciales y mantienen modelos de vista separados.

## Pendiente / decisiones aún abiertas

- Ampliación futura de la lista de unidades y magnitudes.
- Uso de coma decimal.
- Política general para valores negativos según magnitud.
- Política general para cero según magnitud.
- Completar y auditar las referencias normativas concretas de cada unidad/familia.
- Completar pruebas automáticas específicas para nuevas unidades, aliases, notación científica y sistema de ayuda.
- Definir los controles exactos de navegación del Help Viewer durante su implementación.

---

## 21. Web operativa: página Ejemplos

La incorporación de `/examples` fue una ampliación lateral y no una modificación del motor. Actualmente forma parte de la interfaz Web operativa y hace visible el conocimiento ya registrado por UConversor.

Responsabilidades actuales:

- `units/registry.go`: consulta segura de todas las unidades registradas, sin exponer el slice interno.
- `app/app.go`: operación de aplicación para consultar unidades soportadas.
- `internal/web/examples_viewmodel.go`: adaptación de `units.Unit` a estructuras de presentación.
- `internal/web/handlers.go`: handler GET del recurso `/examples`.
- `internal/web/server.go`: registro de la ruta.
- `internal/web/render.go` y `templates.go`: soporte de vistas independientes sin colisión de nombres de template.
- `templates/examples.html`: contenido propio de Ejemplos.
- `templates/partials/header.html`: navegación común actualizada.

El catálogo científico, parser, validador, motor y flujo `App.Convert()` no deben modificarse para implementar esta vista.

Flujos Web:

```text
Conversión:
Browser → handler → App.Convert() → Registry/Validator/Engine → ResultViewModel → HTML

Ejemplos:
Browser → /examples → handler → App.SupportedUnits() → Registry → ExamplesViewModel → HTML
```

Regla de extensibilidad: una nueva unidad correctamente añadida al catálogo y registrada debe poder aparecer en `/examples` sin modificar manualmente el HTML.

---

# Flujo de desarrollo acordado

```text
IDEA / PROBLEMA
        ↓
ANÁLISIS DEL PROBLEMA
        ↓
ESPECIFICACIONES
        ↓
DISEÑO CONCEPTUAL
        ↓
ARCHITECTURE.md
        ↓
IMPLEMENTACIÓN MODULAR
        ↓
PRUEBAS
        ↓
EVOLUCIÓN / MANTENIMIENTO
```

Ninguna implementación deberá adelantarse a las decisiones arquitectónicas necesarias.
