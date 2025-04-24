package models

type ReactionType int

const (
	Like ReactionType = iota
	Dislike
	Funny
	Support
)

var reactionName = map[ReactionType]string{
	Like:    "like",
	Dislike: "dislike",
	Funny:   "Funny",
	Support: "support",
}

func (r ReactionType) String() string {
	return reactionName[r]
}

type Reaction struct {
	Id        string
	UserId    string
	PostId    string
	Type      string
	CreatedAt int64
	EditedAt  int64
}

type Post struct {
	Id        string
	AuthorId  string
	Content   string
	CreatedAt int64
	EditedAt  int64
}
