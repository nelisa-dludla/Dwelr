package models

type Pagination struct {
	NextPage int
	PrevPage int
	CurrentPage int
	TotalPages int
	TwoBelow int
	TwoAfter int
	ThreeAfter int
	Offset int
	BaseURL string
}

