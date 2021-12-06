package webbers

import (
	"strconv"
	"strings"

	"github.com/fishykins/gutenberg/pkg/lang"
)

const defnTag string = "Defn:"
const obsoleteTag string = "[Obs.]"

func BuildWordDefinitions(r *Region, id int) []lang.Definition {

	// Find definitions
	subRegions := make([]Region, 0)
	subRegion := NewEmptyRegion()

	// Go through each line looking for numbered definitions.
	for i := 2; i < len(r.Lines); i++ {
		line := r.Lines[i]

		split := strings.SplitN(line, ". ", 2)

		// If the line start parses into an int, we know this is a definition.
		if _, err := strconv.ParseInt(split[0], 10, 32); err == nil {
			if subRegion.IsOpen() {
				if _, err := strconv.ParseInt(split[0], 10, 32); err == nil {
					subRegion.Close(i - 1)
					subRegions = append(subRegions, subRegion)
				}
			}
			followOn := split[1]
			subRegion = NewOpenRegion(i)
			subRegion.AddLine(followOn)
		} else {
			if subRegion.IsOpen() && line != "" {
				subRegion.AddLine(line)
			}
		}
	}

	if subRegion.IsOpen() {
		subRegion.Close(len(r.Lines) - 1)
		subRegions = append(subRegions, subRegion)
	}

	if len(subRegions) == 0 {
		// No definitions found, use the simple "defn:" tag instead
		return simpleDefinition(r, id)
	} else {
		// Handle each of the definitions
		return definitions(subRegions, id)
	}
}

// builds a single defn by looking for the "defn:" tag.
func simpleDefinition(r *Region, id int) []lang.Definition {
	open := false
	description := ""
	obsolete := false

	for _, line := range r.Lines {
		if strings.Contains(line, obsoleteTag) {
			obsolete = true
		}
		if strings.HasPrefix(line, defnTag) && !open {
			open = true
			description = strings.TrimPrefix(line, defnTag)
		} else if open {
			if line != "" {
				description = description + " " + line
			} else {
				break
			}
		}
	}
	defn := lang.Definition{Theme: -1, Obs: obsolete, Text: description}
	return []lang.Definition{defn}
}

// builds multiple defns based on the given regions
func definitions(regions []Region, id int) []lang.Definition {
	defns := make([]lang.Definition, 0)
	for _, r := range regions {
		defn := complexDefinition(&r)
		defns = append(defns, defn)
	}
	return defns
}

func complexDefinition(region *Region) lang.Definition {
	text := ""
	prelude := ""
	obsolete := false
	for _, line := range region.Lines {
		if line != "" {
			if strings.HasPrefix(line, defnTag+" ") {
				// Explicit definition- everything before this was just a prelude.
				prelude = text
				text = strings.TrimPrefix(line, defnTag+" ")
			} else {
				// Slap it on the end, we don't have anything else to go on.
				text = text + " " + line
				if strings.Contains(line, obsoleteTag) {
					obsolete = true
				}
			}
		}
	}
	defn := lang.Definition{Theme: -1, Obs: obsolete, Text: text}

	if prelude != "" {
		handlePrelude(prelude, &defn)
	}

	return defn
}

func handlePrelude(p string, d *lang.Definition) {
	tags := lang.GetTags(p, '(', ')')
	if len(tags) == 1 {
		if tags[0] != obsoleteTag {
			// Only one tag, so it's probably safe to use it as a theme title.
			themeID := GetTheme(tags[0]).ID
			d.Theme = themeID
		} else {
			d.Obs = true
		}
	} else {
		for _, tag := range tags {
			if tag == obsoleteTag {
				d.Obs = true
				continue
			}
			themeID := GetTheme(tag).ID
			d.Theme = themeID
			break
		}
	}
}
