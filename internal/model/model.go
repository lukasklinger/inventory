package model

// model for inventory entries
type Entry struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Note        string      `json:"note"`
	Rectangles  []Rectangle `json:"rectangles"`
}

// model tracking map rectangles for entries
type Rectangle struct {
	X int `json:"x"`
	Y int `json:"y"`
}
