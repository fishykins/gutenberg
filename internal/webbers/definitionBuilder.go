package webbers

import (
	"strconv"
	"strings"

	"github.com/fishykins/gutenberg/pkg/lang"
)

const defnTag string = "Defn:"

func BuildWordDefinitions(r *Region, id int) []lang.Definition {

	// Find definitions
	subRegions := make([]Region, 0)
	subRegion := NewEmptyRegion()

	// Go through each line looking for numbered definitions.
	for i := 2; i < len(r.Lines); i++ {
		line := r.Lines[i]

		split := strings.SplitN(line, ".", 2)
		firstSentance := split[0]

		// If the line start parses into an int, we know this is a definition.
		if _, err := strconv.ParseInt(firstSentance, 10, 32); err == nil {
			followOn := split[1]
			if subRegion.IsOpen() {
				subRegion.Close(i - 1)
				subRegions = append(subRegions, subRegion)
			}
			subRegion = NewOpenRegion(i)
			subRegion.AddLine(followOn)
		}
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

	for _, line := range r.Lines {
		if strings.HasPrefix(line, defnTag) && !open {
			open = true
			description = strings.TrimPrefix(line, defnTag)
		}
		if open {
			if line != "" {
				description = description + " " + line
			} else {
				break
			}
		}
	}

	defn := lang.Definition{Theme: -1, Obs: false, Text: description}
	return []lang.Definition{defn}
}

// builds multiple defns based on the given regions
func definitions(regions []Region, id int) []lang.Definition {
	defns := make([]lang.Definition, 0)
	for _, r := range regions {
		text := ""
		for _, line := range r.Lines {
			if line != "" {
				text = text + line + " "
			}
		}
		defn := lang.Definition{Theme: -1, Obs: false, Text: text}
		defns = append(defns, defn)
	}
	return defns
}
