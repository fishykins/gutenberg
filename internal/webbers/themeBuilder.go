package webbers

import "github.com/fishykins/gutenberg/pkg/lang"

var Themes = map[string]lang.Theme{}
var id int32 = 0

func GetTheme(name string) *lang.Theme {
	if theme, ok := Themes[name]; ok {
		return &theme
	}
	id++
	newTheme := &lang.Theme{
		Name: name,
		ID:   id,
	}
	Themes[name] = *newTheme
	return newTheme
}

func GetThemeFromID(id int32) *lang.Theme {
	for _, theme := range Themes {
		if theme.ID == id {
			return &theme
		}
	}
	return nil
}
