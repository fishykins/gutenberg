package lang

type Definition struct {
	Theme int32  `json:"theme" bson:"theme,omitempty"`
	Obs   bool   `json:"obs" bson:"obs,omitempty"`
	Text  string `json:"text" bson:"text,omitempty"`
}
