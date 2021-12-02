package lang

type Theme struct {
	ID   int32  `json:"ID" bson:"_id,omitempty"`
	Name string `json:"Name" bson:"name,omitempty"`
}

type ThemeLink struct {
	ID       int32   `json:"ID" bson:"_id,omitempty"`
	Theme    int32   `json:"A" bson:"theme,omitempty"`
	Strength float32 `json:"Strength" bson:"strength,omitempty"` // How strong the link is. 0 => no link, 1 => very similar, -1 => very dissimilar
}

func (l *ThemeLink) Validate() {
	if l.Strength > 0 {
		l.Strength = 1
		return
	}
	if l.Strength < -1 {
		l.Strength = -1
		return
	}
}
