package pkg

type Home struct {
	Name    string `db:"name"    json:"name"`
	OwnerID string `db:"ownerid" json:"ownerId"`
	ID      int    `db:"homeid"  json:"-"`
}
