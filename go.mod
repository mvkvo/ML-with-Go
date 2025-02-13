module web

go 1.23.5

require github.com/labstack/echo/v4 v4.13.3

require (
	dario.cat/mergo v1.0.1 // indirect
	git.sr.ht/~sbinet/gg v0.6.0 // indirect
	github.com/BurntSushi/freetype-go v0.0.0-20160129220410-b763ddbfe298 // indirect
	github.com/BurntSushi/graphics-go v0.0.0-20160129215708-b43f31a4a966 // indirect
	github.com/BurntSushi/xgb v0.0.0-20160522181843-27f122750802 // indirect
	github.com/BurntSushi/xgbutil v0.0.0-20190907113008-ad855c713046 // indirect
	github.com/ByteArena/poly2tri-go v0.0.0-20170716161910-d102ad91854f // indirect
	github.com/Kagami/go-avif v0.1.0 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.3.0 // indirect
	github.com/ajstarks/svgo v0.0.0-20211024235047-1546f124cd8b // indirect
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/benoitkugler/textlayout v0.3.0 // indirect
	github.com/benoitkugler/textprocessing v0.0.3 // indirect
	github.com/campoy/embedmd v1.0.0 // indirect
	github.com/go-fonts/latin-modern v0.3.3 // indirect
	github.com/go-fonts/liberation v0.3.3 // indirect
	github.com/go-latex/latex v0.0.0-20240709081214-31cef3c7570e // indirect
	github.com/go-pdf/fpdf v0.9.0 // indirect
	github.com/go-text/typesetting v0.2.1 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/huandu/xstrings v1.5.0 // indirect
	github.com/kolesa-team/go-webp v1.0.4 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/spf13/cast v1.7.0 // indirect
	github.com/srwiley/rasterx v0.0.0-20220730225603-2ab79fcdd4ef // indirect
	github.com/srwiley/scanx v0.0.0-20190309010443-e94503791388 // indirect
	github.com/tdewolff/font v0.0.0-20250206205927-2dd4de7757d6 // indirect
	github.com/tdewolff/minify/v2 v2.21.1 // indirect
	github.com/tdewolff/parse/v2 v2.7.21-0.20250206205826-9029f397cf8a // indirect
	github.com/wcharczuk/go-chart/v2 v2.1.2 // indirect
	golang.org/x/image v0.23.0 // indirect
	star-tex.org/x/tex v0.5.0 // indirect
)

require (
	github.com/Masterminds/sprig/v3 v3.3.0
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/tdewolff/canvas v0.0.0-20250209140343-015076d8ff76
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gonum.org/v1/plot v0.15.0
)

replace web/handler => ../handler

replace web/calc => ../calc

replace web/csv => ../csv

replace web/utils => ../utils
