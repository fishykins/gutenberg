package lang

type WordType int8
type NounType int8

const (
	WordTypeNone WordType = iota
	WordTypeNoun
	WordTypePronoun
	WordTypeVerb
	WordTypeAdjective
	WordTypeAdverb
	WordTypePreposition
	WordTypeConjunction
)

const (
	PropperNoun NounType = iota
	ConcreteNoun
	AbstractNoun
	CollectiveNoun
)

func (s WordType) String() string {
	switch s {
	case WordTypeNoun:
		return "noun"
	case WordTypePronoun:
		return "pronoun"
	case WordTypeVerb:
		return "verb"
	case WordTypeAdjective:
		return "adjective"
	case WordTypeAdverb:
		return "adverb"
	case WordTypePreposition:
		return "preposition"
	case WordTypeConjunction:
		return "conjunction"
	}
	return ""
}

func (s NounType) String() string {
	switch s {
	case PropperNoun:
		return "propper"
	case ConcreteNoun:
		return "concrete"
	case AbstractNoun:
		return "abstract"
	case CollectiveNoun:
		return "collective"
	}
	return ""
}

func (s NounType) IsCommon() bool {
	return s != PropperNoun
}
