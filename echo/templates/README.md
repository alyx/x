# x/echo/templates

The `templates` package implements a template loader built to support handling templates built into the binary (via [Statik](https://github.com/rakyll/statik)). This package also supports shared components (e.g. a layout system) to extend Golang's standard `html/template` library.

## Usage

Before compiling an application using this component, you must first generate the inclusion source from Statik, via `static -src=./your/template/folder`. You may then initialize this library with:

```go
hfs, err := fs.New()
var t = &templates.Template{}
e.Renderer = t
templates.LoadTemplates(hfs, "/templates")
```

Where `LoadTemplates()` takes an http.FileSystem object and a string specifying the base directory for your templates.

This package will parse your templates folder looking for the following:

* A folder named `layouts` in the base directory, containing layouts.
* A folder named `pages` containing individual page content.
* A folder named `shared` containing any other shared components.

Once the application is initialized, you may render a template via `c.Render(http.StatusOK, "base:index", nil)` where `base:index` is the layout:page representation.