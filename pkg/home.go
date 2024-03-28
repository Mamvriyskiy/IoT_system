package pkg

type Home struct {
	Id      int    `json:"-"`
	Name    string `json:"name"`
	OwnerID int    `json:"ownerId"`
}
