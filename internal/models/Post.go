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

type Reactions struct {
	Id        string
	UserId    string
	Type      string
	Timestamp int64
}

type Post struct {
	Id         string
	AuthorId   string
	CreatedAt  int64
	Content    string
	ReactionId *string
}
