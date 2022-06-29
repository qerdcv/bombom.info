package domain

type ClubMember struct {
	Icon struct {
		ID int `json:"id"`
	} `json:"icon"`
	Tag       string `json:"tag"`
	Name      string `json:"name"`
	Trophies  int    `json:"trophies"`
	Role      string `json:"role"`
	NameColor string `json:"nameColor"`
}

type Club struct {
	Tag              string       `json:"tag"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	Trophies         int          `json:"trophies"`
	RequiredTrophies int          `json:"requiredTrophies"`
	Members          []ClubMember `json:"members"`
	Type             string       `json:"string"`
	BadgeID          int          `json:"badgeID"`
}
