package templates

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rakyll/statik/fs"
	"github.com/shurcooL/httpfs/html/vfstemplate"
)

var templates map[string]*template.Template

// Template will never be used.
type Template struct {
	templates *template.Template
}

// Render implements the echo#Render function for template rendering
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := templates[name]
	if !ok {
		err := errors.New("template: name not found: " + name)
		return err
	}
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	components := strings.Split(name, ":")
	return tmpl.ExecuteTemplate(w, components[0], data) // layout -> defined in each layout template
}

func match(pattern string, name string) bool {
	matched, _ := path.Match(pattern, name)
	return matched
}

// LoadTemplates loads templates on program initialisation
func LoadTemplates(hfs http.FileSystem, base string) {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	var layouts []string
	var pages []string
	var globalShared []string

	fs.Walk(hfs, "/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if match(base+"/layouts/*.html", path) {
			layouts = append(layouts, path)
			return nil
		}

		if match(base+"/pages/*.html", path) {
			pages = append(pages, path)
			return nil
		}

		if match(base+"/shared/*.html", path) {
			globalShared = append(globalShared, path)
			return nil
		}

		return nil
	})

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		for _, page := range pages {
			files := append(globalShared, layout, page)
			// todo - crawl_folder func call if include is folder
			layoutBase := filepath.Base(layout)
			layoutShort := layoutBase[0:strings.LastIndex(layoutBase, ".")]
			pageBase := filepath.Base(page)
			pageShort := pageBase[0:strings.LastIndex(pageBase, ".")]
			tmp := template.New(pageShort).Delims("{{", "}}")
			templates[layoutShort+":"+pageShort] = template.Must(vfstemplate.ParseFiles(hfs, tmp, files...))
		}
	}
}
