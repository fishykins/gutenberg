package lang

import (
	"fmt"
	"strings"

	"github.com/fishykins/gutenberg/pkg/formatting"
	"go.mongodb.org/mongo-driver/bson"
)

type Word struct {
	ID          int32        `json:"ID" bson:"_id,omitempty"`
	Inner       string       `json:"word" bson:"inner,omitempty"`
	Type        WordType     `json:"type" bson:"type,omitempty"`
	Language    string       `json:"language" bson:"language,omitempty"`
	Source      string       `json:"source" bson:"source,omitempty"`
	Definitions []Definition `json:"definitions" bson:"definitions,omitempty"`
}

type WordLink struct {
	ID       int64   `json:"ID" bson:"_id,omitempty"`
	Word     int64   `json:"A" bson:"word,omitempty"`
	Strength float32 `json:"Strength" bson:"strength,omitempty"` // How strong the link is. 0 => no link, 1 => very similar, -1 => very dissimilar
}

func (l *WordLink) Validate() {
	if l.Strength > 0 {
		l.Strength = 1
		return
	}
	if l.Strength < -1 {
		l.Strength = -1
		return
	}
}

func (w *Word) IntoBson() bson.M {
	return bson.M{
		"_id":         w.ID,
		"inner":       w.Inner,
		"type":        w.Type,
		"language":    w.Language,
		"source":      w.Source,
		"definitions": w.Definitions,
	}
}

func (w *Word) Header(formatter formatting.FormatType) string {
	var header string = ""
	switch formatter {
	case formatting.FormatType_Plain:
		header = "%s <%s>"
	case formatting.FormatType_Markdown:
		header = "**%s** <*%s*>"
	}
	return fmt.Sprintf(header, strings.Title(strings.ToLower(w.Inner)), w.Type)
}
