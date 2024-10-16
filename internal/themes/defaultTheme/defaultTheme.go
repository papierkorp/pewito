package main

import (
	"github.com/a-h/templ"
	"pewito/internal/themes"
	"pewito/internal/themes/defaultTheme/templates"
)

type DefaultTheme struct{}

func (t *DefaultTheme) Home() (templ.Component, error) {
	return templates.Home(), nil
}

func (t *DefaultTheme) Help() (templ.Component, error) {
	return templates.Help(), nil
}

func (t *DefaultTheme) Settings() (templ.Component, error) {
	return templates.Settings(), nil
}

func (t *DefaultTheme) Search() (templ.Component, error) {
	return templates.Search(), nil
}

func (t *DefaultTheme) DocsRoot() (templ.Component, error) {
	return templates.DocsRoot(), nil
}

func (t *DefaultTheme) Docs(content string) (templ.Component, error) {
	return templates.Docs(content), nil
}

func (t *DefaultTheme) Playground() (templ.Component, error) {
	return templates.Playground(), nil
}

func (t *DefaultTheme) Plugins() (templ.Component, error) {
    return templates.Plugins(), nil
}

var Theme DefaultTheme

func main() {}

func init() {
	themes.RegisterTheme("defaultTheme", &DefaultTheme{})
}