

# Escáner OCR de Texto

Herramienta de software de escaneo OCR de texto. Está diseñada para extraer texto de imágenes y archivos PDF.

Este software está __EN DESARROLLO INTENSO__. Por ahora, utilice [este software de OCR](https://abanoubhanna.com/posts/ocr).

## Comandos

```sh
# obtener todas las dependencias del proyecto
go mod tidy

# compilar: genera un archivo ejecutable/app de CLI llamado "ocr"
go build -o ocr .

# ejecutar la app para identificar una imagen de ejemplo
./ocr --lang=eng --img=img/default.png

# dependencias, compilar, ejecutar
go mod tidy && go build -o ocr . && ./ocr --lang=eng --img=img/default.png

# detectar condiciones de carrera y fugas de memoria
go run -race .
```

## Hoja de ruta y tareas

Hecho y POR HACER de características, objetivos y valores. Pequeños pasos hacia la meta.

- v0.1.0
  - La biblioteca OCR utilizada es [tesseract-OCR](https://github.com/tesseract-ocr/tesseract) con el envoltorio (wrapper) [gosseract](https://github.com/otiai10/gosseract) v2.
  - imagen a texto
  - soporte para OCR en inglés
  - soporte para OCR en árabe
  - agregar pruebas
  - usar algunas imágenes de prueba de [renard314/textfairy](https://github.com/renard314/textfairy)
- Próximos pasos
  - PDF a texto
  - PDF a docx
  - PDF a PDF con texto seleccionable
  - __escalable__ : aprovechar todos los núcleos de la CPU para completar el trabajo más rápido
  - __en lote__ / __procesamiento por bloques__ : corutinas y paralelismo para tareas / trabajos
  - app de CLI componible para scripts y automatización
  - UX / fácil de usar / amigable para el usuario
  - dividir la imagen en pedazos/segmentos y aplicar OCR de forma concurrente. (rendimiento)
  - disponible en Debian y distribuciones basadas en Debian
  - disponible como snap
  - disponible como flatpak
  - disponible en Elementary OS
  - disponible en Mac OS (probablemente vía HomeBrew)
  - disponible en Windows
  - historial de escaneos : { imagen original, imagen procesada, texto extraído }
  - migrar desde opencv a bibliotecas de Go

### Referencias

- [3b1b : ¿qué es una convolución?](https://youtu.be/KuXjwB4LzSA)
- [entrenar / refinar el OCR de tesseract](https://github.com/abanoubha/train-tesseract-ocr)
- [tesseractfonts : ajuste fino de tesseract para nuevas fuentes](https://github.com/dhivehi/tesseractfonts)
- [Biblioteca Pre-Recognize](https://github.com/leha-bot/PRLib) - biblioteca con algoritmos para mejorar la calidad del OCR.

## Código fuente

El código fuente del proyecto OCR se puede encontrar en:

- GitHub: <https://github.com/abanoubha/ocr.git>
- GitLab: <https://gitlab.com/abanoubha/ocr.git>
- CodeBerg: <https://codeberg.org/abanoubha/ocr.git>
