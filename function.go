package main

type function struct {
	name         string
	isMiddleware bool
	tags         []*tag
}

func newFunction(name string, comments []string) *function {
	tags := make([]*tag, 0)
	isMiddleware := false
	for _, comment := range comments {
		if t := extractToTag(comment); t != nil {
			if t.typ == middleware {
				t.value = name
				isMiddleware = true
				tags = []*tag{t}
				break
			}
			tags = append(tags, t)
		}
	}
	return &function{
		name:         name,
		tags:         tags,
		isMiddleware: isMiddleware,
	}
}
