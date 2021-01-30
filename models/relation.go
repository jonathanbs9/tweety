package models

// Relation => relacion para seguir otro user
type Relation struct {
	UserID         string `bson:"userid" json:user_id`
	UserRelationID string `bson:"userRelationId" json:user_id`
}
