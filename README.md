# go-vite-js

Slim adapter for GO templates and VITE.JS

---

## Example:

### Initialization in application

Create vite.js adapter & define adapter like payload for template, for example in GIN handler:

```go
func main() {
    // creating new adapter
    govite := govite.Default()

    // define payload for route
    payload := &gin.H{
        "vite": vitejs,
    }

    // create http router
    router := gin.Default()
    
    // define templates
    router.LoadHTMLGlob("templates/*")
    
    // define route
    router.GET("/", func(context *gin.Context) {
        context.HTML(200, "website", c.payload)
    })

    // run http server
    router.Run(":8080")
}
```

### Using in GO templates

Create addition template for inject it to footer:

```html
{{ define "vite.footer" }}
    {{ if ne .vite.IsProduction true }}
        <script type="module" src="{{ .vite.Client }}"></script>
    {{ end }}
{{ end }}
```

Define template for some page

```html
{{ define "website" }}
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>go-vite-js</title>
        <link rel="stylesheet" href="{{ .vite.Asset `sass/website.sass` }}">
    </head>
    <body id="app">
    <h1>Go with Vite.js</h1>
    {{ template "vite.footer" . }}
    <script src="{{ .vite.Asset `ts/website.ts` }}"></script>
    </body>
    </html>
{{ end }}
```
