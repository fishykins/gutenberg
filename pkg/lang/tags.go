package lang

func GetTags(source string, open rune, close rune) []string {
	var tags []string
	var tag string
	var inTag bool
	for i, c := range source {
		if c == open {
			inTag = true
			tag = ""
			continue
		}
		if c == close {
			inTag = false
			if tag != "" {
				tags = append(tags, tag)
			}
		}
		if inTag {
			tag += string(c)
		}
		if i == len(source)-1 && inTag {
			tags = append(tags, tag)
		}
	}
	return tags
}
