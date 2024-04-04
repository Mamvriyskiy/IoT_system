package pkg

type Home struct {
	Name    string `json:"name"`
	OwnerID int    `json:"ownerId"`
	ID      int    `json:"-"`
}
