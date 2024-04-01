package pkg

type Home struct {
	ID      int    `json:"-"`
	Name    string `json:"name"`
	OwnerID int    `json:"ownerId"`
}
