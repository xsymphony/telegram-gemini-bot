package gin_extra

import (
	"embed"
	"html/template"
	"io/fs"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func LoadHTMLFromEmbedFS(engine *gin.Engine, embedFS embed.FS, pattern string) {
	root := template.New("")
	tmpl := template.Must(root, loadAndAddToRoot(engine.FuncMap, root, embedFS, pattern))
	engine.SetHTMLTemplate(tmpl)
}

func loadAndAddToRoot(funcMap template.FuncMap, rootTemplate *template.Template, embedFS embed.FS, pattern string) error {
	pattern = strings.ReplaceAll(pattern, ".", "\\.")
	pattern = strings.ReplaceAll(pattern, "*", ".*")

	err := fs.WalkDir(embedFS, ".", func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if matched, _ := regexp.MatchString(pattern, path); !d.IsDir() && matched {
			data, readErr := embedFS.ReadFile(path)
			if readErr != nil {
				return readErr
			}
			t := rootTemplate.New(path).Funcs(funcMap)
			if _, parseErr := t.Parse(string(data)); parseErr != nil {
				return parseErr
			}
		}
		return nil
	})
	return err
}
