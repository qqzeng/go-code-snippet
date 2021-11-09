package proto

const (
	FilterMust   = "Must"   // and
	FilterShould = "Should" // or

	MatchExact = "Exact" // exact match
	MatchFuzzy = "Fuzzy" // fuzzy match
	MatchRange = "Range" // range match, v=a,b
	MatchIn    = "In"    // in match, v=a,b,c,d

	MatchSeparator = "," // MatchSeparator
)

const (
	TimeDefaultLayout = "2006-01-02 15:04:05"
)

// Filters ...
type Filters []Filter

// Filter ...
type Filter struct {
	Name  string `json:"Name" binding:"required"`
	Type  string `json:"Type" binding:"required"`
	Value string `json:"Value" binding:"required"`
}

// Sorters ...
type Sorters []Sorter

// SortInfo ...
type Sorter struct {
	OrderBy string `json:"OrderBy" binding:"required"`
	Desc    bool   `json:"Desc" binding:"omitempty"`
}

// QueryCommonParams ...
type QueryCommonParams struct {
	Limit   int     `json:"Limit" binding:"required,min=1"`
	Offset  int     `json:"Offset" binding:"min=0"`
	Filters Filters `json:"Filters" binding:"omitempty"`
	Sorters Sorters `json:"Sorters" binding:"omitempty"`
	Group   string  `json:"Group" binding:"omitempty"`
}
