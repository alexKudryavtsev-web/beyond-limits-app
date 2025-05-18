package entity

type Genre struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type Author struct {
	ID       uint64 `json:"id"`
	FullName string `json:"full_name"`
}

type Dimension struct {
	ID     uint64 `json:"id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type WorkTechnique struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
