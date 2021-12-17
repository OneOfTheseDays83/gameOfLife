package api

type Request struct {
	Rows       uint64   `json:"rows,omitempty"`
	Cols       uint64   `json:"cols,omitempty"`
	Iterations uint64   `json:"iterations,omitempty"`
	Grid       [][]bool `json:"grid,omitempty"`
}
