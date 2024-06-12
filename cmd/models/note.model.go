package models

type Note struct {
	ID         int      `bson:"_id,omitempty" json:"id,omitempty"`
	Title      string   `bson:"title" json:"title"`
	Content    string   `bson:"content" json:"content"`
	Categories []string `bson:"categories" json:"categories"`
	SharedWith []string `bson:"sharedWith" json:"sharedWith"`
	CreatedAt  int      `bson:"createdAt" json:"createdAt"`
	UpdatedAt  int      `bson:"updatedAt" json:"updatedAt"`
}
