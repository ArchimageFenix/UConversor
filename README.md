<div align="center">

<img src="docs/assets/uconversor-logo.png"
     alt="UConversor - Universal Conversor"
     width="220">

# UConversor

### Universal Conversor

**Conversión automática de unidades con enfoque en simplicidad, rendimiento y trazabilidad científica.**

</div>

**Conversión automática de unidades con enfoque en simplicidad,
rendimiento y trazabilidad científica.**

> Escribe una cantidad con su unidad. UConversor identifica qué
> representa y genera automáticamente las conversiones compatibles.
:::

------------------------------------------------------------------------

## ✨ ¿Qué es UConversor?

**UConversor (Universal Conversor)** es una aplicación de conversión
automática de unidades desarrollada en **Go**.

Su idea principal es eliminar buena parte de los pasos habituales de un
conversor tradicional. En lugar de elegir primero una categoría, después
una unidad de origen y finalmente una unidad de destino, el usuario
escribe directamente una expresión como:

``` text
24mph
5km
2200Ω
300Mbps
1sol
```

UConversor analiza la expresión, reconoce la unidad, determina su
**familia** y su **magnitud**, valida que exista una relación científica
admitida y genera las conversiones compatibles disponibles en su
catálogo.

El proyecto busca combinar:

**facilidad de uso + rigor científico + software ligero y mantenible.**

------------------------------------------------------------------------

## 🎯 ¿Qué problema intenta resolver?

Un conversor tradicional suele exigir:

``` text
Elegir categoría
      ↓
Elegir unidad de origen
      ↓
Elegir unidad de destino
      ↓
Introducir valor
      ↓
Convertir
```

UConversor propone:

``` text
24mph
  ↓
Reconocer automáticamente "mph"
  ↓
Identificar: Velocidad
  ↓
Validar la conversión
  ↓
Generar equivalencias compatibles
```

El usuario proporciona la información esencial ---**valor + unidad**---
y el programa interpreta el resto.

------------------------------------------------------------------------

## ⌨️ ¿Cómo se introducen los datos?

La sintaxis principal es:

``` text
valor + unidad
```

Ejemplos:

``` text
24mph
24 mph
5km
2.5kg
2200Ω
```

No es necesario indicar previamente que `mph` pertenece a velocidad o
que `Ω` pertenece a resistencia eléctrica.

UConversor mantiene **símbolos científicos canónicos** y puede admitir
aliases explícitos para símbolos difíciles de escribir.

> ⚠️ Las mayúsculas y minúsculas pueden tener significado científico.
> UConversor no trata indiscriminadamente `m`, `M`, `mV`, `MV`, `B` y
> `b` como equivalentes.

------------------------------------------------------------------------

## 🧠 ¿Cómo funciona?

``` text
👤 Entrada del usuario
        ↓
🧹 Normalización
        ↓
🔎 Separación de valor y unidad
        ↓
🏷️ Reconocimiento de la unidad
        ↓
🧭 Identificación de familia y magnitud
        ↓
🛡️ Comprobación de compatibilidad
        ↓
📚 Selección de la definición científica
        ↓
🧮 Conversión mediante el motor
        ↓
🔄 Generación de equivalencias compatibles
        ↓
🖥️ Presentación del resultado
```

Una característica importante es que **familia y magnitud no significan
lo mismo**.

``` text
Electrónica
├── Voltaje
├── Corriente
├── Resistencia
└── Potencia
```

Aunque voltaje y corriente pertenezcan a la misma familia, no
representan la misma magnitud. Por eso una expresión como:

``` text
12 V → mA
```

no es una simple conversión de unidades válida.

------------------------------------------------------------------------

## ⚡ Rendimiento: hacer únicamente el trabajo necesario

UConversor sigue una regla importante: **una entrada que ya sabemos que
no está soportada no debe recorrer innecesariamente el resto del
sistema**.

``` text
24xyz
   ↓
Unidad no reconocida
   ↓
⛔ FIN
```

Una entrada rechazada tempranamente no debe continuar hacia el validador
científico, selección de fórmulas, motor de conversión ni generación de
resultados.

Esto forma parte de la arquitectura, no es solamente una decisión
visual.

Cuando corresponde, una magnitud utiliza una **unidad o representación
de referencia**, evitando implementar una fórmula independiente para
cada pareja posible:

``` text
mi ──→ m ──→ km
          ├─→ cm
          ├─→ ft
          └─→ yd
```

Esta estrategia reduce duplicación, facilita añadir unidades y mantiene
sencillo el camino lógico de una conversión.

La CLI y la Web reutilizan el mismo núcleo científico. No existe un
segundo motor de conversión exclusivo para el navegador.

> 🚀 El objetivo de rendimiento no es obtener velocidad sacrificando
> exactitud, sino **evitar trabajo innecesario conservando la precisión
> del cálculo**.

------------------------------------------------------------------------

## 🎯 Precisión: calcular primero, formatear después

``` text
CALCULAR con precisión suficiente
        ↓
CONVERTIR
        ↓
FORMATEAR al presentar
```

Los valores intermedios no deben redondearse prematuramente. La forma
visible del resultado es responsabilidad de la capa de presentación.

La notación científica es igualmente una **decisión de presentación** y
no un mecanismo para reducir la precisión interna.

------------------------------------------------------------------------

## 🌍 Ámbitos de aplicación

La versión actual contiene catálogos relacionados con distintas áreas:

  -----------------------------------------------------------------------
  Área                                Ejemplos
  ----------------------------------- -----------------------------------
  📏 **Longitud y distancia**         mm, cm, m, km, pulgadas, pies,
                                      millas y otras

  ⚖️ **Masa**                         mg, g, kg, oz, lb

  🚗 **Velocidad**                    m/s, km/h, mph, nudos, ft/s

  ⏱️ **Tiempo**                       ns, µs, ms, s, min, h, días y otras
                                      escalas

  🔌 **Electrónica**                  voltaje, corriente, resistencia y
                                      potencia

  🧪 **Física**                       magnitudes físicas adicionales
                                      según el catálogo

  💾 **Datos digitales**              bits, bytes y múltiplos

  🌐 **Tasa de transferencia**        bit/s, byte/s y múltiplos

  📐 **Ángulos**                      conversiones angulares

  💪 **Fuerza**                       unidades de fuerza

  🪐 **Escalas científicas            por ejemplo, año luz y sol marciano
  especiales**                        
  -----------------------------------------------------------------------

El catálogo está pensado para crecer sin convertir el núcleo en una
colección de casos especiales.

------------------------------------------------------------------------

## 🪐 De la Tierra a Marte

UConversor no está limitado a unidades cotidianas.

Un ejemplo incorporado al proyecto es el **sol marciano**, tratado como
una unidad de duración dentro de Tiempo:

``` text
1sol
```

La idea representa bien la filosofía del proyecto: si una nueva unidad
pertenece a una magnitud ya conocida, su incorporación debe concentrarse
principalmente en el catálogo y sus pruebas, no en reconstruir el motor.

------------------------------------------------------------------------

## 📚 Fuentes internacionales y trazabilidad científica

UConversor no pretende ser una colección de factores numéricos copiados
sin contexto.

Cada definición puede conservar:

``` text
Unidad
├── nombre
├── símbolo
├── aliases
├── familia
├── magnitud
├── factor o fórmula
├── tipo de relación
├── fuente científica/normativa
└── observaciones
```

Cuando existe una relación oficial o científicamente establecida, tiene
prioridad sobre una fórmula propia creada por conveniencia.

Jerarquía conceptual:

``` text
1. Definición normativa u oficial
        ↓
2. Definición SI / organismo metrológico reconocido
        ↓
3. Relación científica establecida
        ↓
4. Relación derivada de definiciones validadas
        ↓
5. Fórmula propia documentada y comprobable
```

Entre las referencias internacionales que pueden respaldar distintas
definiciones, **según cada magnitud concreta**, se encuentran:

-   🌐 **BIPM** --- Bureau International des Poids et Mesures;
-   📐 **Sistema Internacional de Unidades (SI)**;
-   🧪 **NIST** --- National Institute of Standards and Technology;
-   ⚡ **IEC** --- International Electrotechnical Commission;
-   📘 **ISO** --- International Organization for Standardization;
-   🚀 **NASA / JPL** y otras instituciones científicas oficiales cuando
    una magnitud especializada lo requiere.

Esto no significa que cada unidad proceda de todas estas instituciones.
La intención es conservar la **fuente específica aplicable** a cada
relación.

Así, un factor puede ser auditado, mantenido y contrastado en lugar de
permanecer como un número sin procedencia.

------------------------------------------------------------------------

## 🧮 Estrategias matemáticas

UConversor no presupone que todas las conversiones sean una simple
multiplicación.

### 🔹 Conversión lineal

Adecuada para relaciones por escala.

### 🔸 Conversión afín

Necesaria cuando además de escala existe desplazamiento, como en
determinadas conversiones de temperatura.

### 🧩 Fórmula especializada

Disponible conceptualmente cuando una magnitud legítimamente lo
requiera, siempre con documentación, dominio de validez y pruebas.

------------------------------------------------------------------------

## 🖥️ Dos formas de utilizar UConversor

### 💻 CLI --- Línea de comandos

Desde el código fuente:

``` bash
go run ./cmd/converter
```

Después pueden introducirse expresiones como:

``` text
24mph
5km
2200Ω
1sol
```

La CLI dispone de ayuda mediante:

``` text
h
```

La presentación de terminal está separada del motor científico.

### 🌐 Interfaz Web

El servidor Web utiliza la misma capa de aplicación y el mismo núcleo:

``` bash
go run ./cmd/web
```

Después abre en el navegador la dirección indicada por el servidor.

La Web incluye página principal de conversión, resultados visuales y una
vista `/examples` para descubrir ejemplos y unidades soportadas.

Las unidades mostradas en `/examples` proceden del registro real:

``` text
Catálogo
   ↓
Registry
   ↓
App
   ↓
Web
   ↓
/examples
```

No se mantiene una segunda lista científica duplicada manualmente en
HTML.

------------------------------------------------------------------------

## 🛠️ Tecnologías

  Tecnología                                Función
  ----------------------------------------- ---------------------------------------
  🟦 **Go**                                 aplicación, dominio, motor y servidor
  🌐 **HTTP / biblioteca estándar de Go**   servidor Web
  🧩 **HTML Templates**                     generación de vistas
  🎨 **HTML + CSS**                         interfaz Web
  🧪 **Testing de Go**                      pruebas automatizadas
  📄 **Markdown**                           arquitectura y documentación

La separación principal es:

``` text
CLI ─────┐
         ├──→ APP ──→ NÚCLEO DE CONVERSIÓN
WEB ─────┘
```

------------------------------------------------------------------------

## 💻 Sistemas operativos

Al estar escrito en Go, UConversor puede compilarse para las principales
plataformas soportadas por el toolchain, sujeto a validación de la
versión concreta.

### 🪟 Windows

Es el entorno principal de desarrollo actual.

``` powershell
go build -o converter.exe ./cmd/converter
go build -o web.exe ./cmd/web
```

### 🐧 Linux

``` bash
go build -o converter ./cmd/converter
go build -o web ./cmd/web
```

### 🍎 macOS

La arquitectura no depende conceptualmente de una interfaz exclusiva de
Windows y puede compilarse con Go para macOS siempre que la versión
concreta y sus pruebas sean compatibles.

> ℹ️ **Portabilidad no equivale a certificación automática.** Cada
> plataforma debe validarse ejecutando las pruebas correspondientes.

------------------------------------------------------------------------

## 📦 Compilar desde el código fuente

Requisitos:

-   Go;
-   una terminal;
-   Git para clonar el repositorio;
-   un navegador moderno para la interfaz Web.

``` bash
git clone https://github.com/ArchimageFenix/UConversor.git
cd UConversor
go test ./...
```

Ejecutar CLI:

``` bash
go run ./cmd/converter
```

Ejecutar Web:

``` bash
go run ./cmd/web
```

------------------------------------------------------------------------

## 🧪 Pruebas

UConversor busca validar tanto el software como las relaciones
matemáticas.

-   🔎 **Parser:** entradas válidas e inválidas.
-   🧮 **Conversión:** resultados conocidos.
-   🎯 **Precisión:** evitar redondeos internos prematuros.
-   🖥️ **Formato:** representación final.
-   📚 **Fuente:** metadata científica requerida.
-   🧭 **Compatibilidad:** impedir conversiones entre magnitudes
    diferentes.
-   ⛔ **Rechazo temprano:** impedir trabajo posterior para entradas
    desconocidas.

Ejecutar todas:

``` bash
go test ./...
```

------------------------------------------------------------------------

## 🧱 Diseño modular

``` text
Entrada
  ↓
Parser
  ↓
Normalización
  ↓
Registro
  ↓
Validación
  ↓
Motor
  ↓
Formato
  ↓
Presentación
```

Añadir una unidad a una magnitud existente debería concentrarse
principalmente en:

``` text
Nueva unidad
    ↓
Catálogo correspondiente
    ↓
Pruebas
```

y no exigir reescribir parser, motor, CLI y Web.

------------------------------------------------------------------------

## 🧭 Principio central

``` text
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
PRESENTAR de forma clara
```

**La facilidad de uso no debe conseguirse sacrificando rigor
científico.**

------------------------------------------------------------------------

## 👨‍🎓 ¿Para quién puede ser útil?

UConversor puede resultar útil en:

-   🎓 educación y aprendizaje;
-   🧑‍🏫 docencia;
-   🔬 ciencias;
-   ⚙️ ingeniería;
-   🔌 electrónica;
-   💻 informática y redes;
-   💾 almacenamiento y transferencia de datos;
-   📐 matemáticas y física;
-   🪐 astronomía y exploración espacial;
-   🏠 conversiones cotidianas;
-   🧑‍💻 herramientas técnicas y automatización.

------------------------------------------------------------------------

## 📖 ¿Quieres profundizar?

Este README está pensado para explicar **qué es UConversor, qué puede
hacer y cómo utilizarlo**.

Para conocer las decisiones de ingeniería, responsabilidades internas,
trazabilidad científica, reglas de extensión, rechazo temprano,
estrategias matemáticas y evolución del sistema, consulta:

### 👉 [`ARCHITECTURE.md`](ARCHITECTURE.md)

Ese documento es la referencia técnica del diseño del proyecto.

------------------------------------------------------------------------

## 🌱 Estado del proyecto

UConversor continúa en evolución.

La arquitectura está preparada para incorporar nuevas unidades y
magnitudes de forma controlada, manteniendo separadas la definición
científica, el reconocimiento de entrada, el cálculo, la presentación y
las interfaces.

El objetivo no es acumular unidades sin criterio, sino construir un
catálogo cada vez más amplio que continúe siendo **trazable, comprobable
y mantenible**.

------------------------------------------------------------------------

## 🤝 Contribuciones

Las contribuciones son bienvenidas especialmente cuando:

-   añaden unidades con referencias fiables;
-   incorporan pruebas;
-   mejoran documentación;
-   corrigen relaciones con evidencia verificable;
-   mejoran portabilidad;
-   preservan la separación de responsabilidades.

Antes de modificar aspectos estructurales o científicos, consulta
[`ARCHITECTURE.md`](ARCHITECTURE.md).

------------------------------------------------------------------------

## 📜 Licencia

Actualmente debe aplicarse únicamente la licencia que se publique
explícitamente junto al repositorio.

> Si todavía no existe un archivo `LICENSE`, no debe asumirse
> automáticamente una licencia de código abierto concreta.

------------------------------------------------------------------------

::: {align="center"}
### 🌐 UConversor --- Universal Conversor

**Una entrada. Una magnitud. Todas las conversiones compatibles.**

🧮 Precisión · 📚 Trazabilidad · ⚡ Rendimiento · 🧱 Modularidad
:::
