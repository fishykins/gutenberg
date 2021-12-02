package webbers

import (
	"fmt"
	"strings"

	"github.com/fishykins/gutenberg/pkg/formatting"
	"github.com/fishykins/gutenberg/pkg/lang"
)

func BuildWord(r *Region, id int) (*lang.Word, error) {
	// Ensure there are at least two lines so we can get title and type.
	if len(r.Lines) < 3 {
		return nil, fmt.Errorf("region has too few lines (%d)", len(r.Lines))
	}

	fmt.Printf("Building region %d (lines %d -> %d)...\n", id, r.Start, r.End)

	title := r.Lines[0]
	subTitle := r.Lines[1]

	// We should handle this properly some time, but for now just refuse the word.
	if strings.Contains(title, ";") {
		//fmt.Println("bad title")
		return nil, fmt.Errorf("%s: title contains ';'", title)
	}

	wordType, err := getWordType(subTitle)
	if err != nil {
		//fmt.Println("bad type")
		return nil, err
	}

	definitions := BuildWordDefinitions(r, id)

	word := lang.Word{
		ID:          int32(id),
		Inner:       title,
		Type:        wordType,
		Definitions: definitions,
	}

	fmt.Printf("R%d: \"%s\" has %d definitons\n", id, word.Header(formatting.FormatType_Plain), len(definitions))
	for _, d := range word.Definitions {
		fmt.Println(d.Text + "\n")
	}

	return &word, nil
}

func getWordType(line string) (lang.WordType, error) {
	if strings.Contains(line, " adv.") {
		return lang.WordTypeAdverb, nil
	}
	if strings.Contains(line, " prep.") {
		return lang.WordTypePreposition, nil
	}
	if strings.Contains(line, " conj.") {
		return lang.WordTypeConjunction, nil
	}
	if strings.Contains(line, " n.") {
		return lang.WordTypeNoun, nil
	}
	if strings.Contains(line, " a.") {
		return lang.WordTypeAdjective, nil
	}
	if strings.Contains(line, " v.") {
		return lang.WordTypeVerb, nil
	}
	return lang.WordTypeNone, fmt.Errorf("cannot determine type from line: %s", line)
}
