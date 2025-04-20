package models

type Reaction int

const (
	Like Reaction = iota
	Dislike
	Funny
	Support
)

var reactionName = map[Reaction]string{
	Like:    "like",
	Dislike: "dislike",
	Funny:   "Funny",
	Support: "support",
}

func (r Reaction) String() string {
	return reactionName[r]
}

type Post struct {
	Id        string
	CreatedAt int64
	AuthorId  string
	Content   string
	Reactions map[Reaction][]string
}
