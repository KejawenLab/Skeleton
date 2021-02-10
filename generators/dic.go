package generators

import (
	"fmt"
	"os"
	engine "text/template"

	configs "github.com/crowdeco/skeleton/configs"
)

type Dic struct {
}

func (g *Dic) Generate(template *configs.Template, modulePath string, workDir string, templatePath string) {
	dicTemplate, _ := engine.ParseFiles(fmt.Sprintf("%s/%s/dic.tpl", workDir, templatePath))
	dicFile, err := os.Create(fmt.Sprintf("%s/%s.go", modulePath, template.ModuleLowercase))
	if err != nil {
		panic(err)
	}

	dicTemplate.Execute(dicFile, template)
}
