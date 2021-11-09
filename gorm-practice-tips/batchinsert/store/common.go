package store

import (
	"errors"
	"fmt"
	"strings"

	"git.woa.com/bigfishchen/go-toolkit/pkg/convert"
	"git.woa.com/bigfishchen/go-toolkit/pkg/types"
	"github.com/qqzeng/gorm-practice-tips/batchinsert/proto"
)

// ExactMatch ...
type ExactMatch struct {
	Value string
}

// FuzzyMatch ...
type FuzzyMatch struct {
	Value string
}

// RangeMatch ...
type RangeMatch struct {
	Start string
	End   string
}

// InMatch ...
type InMatch struct {
	Values []string
}

// Query ...
type Query struct {
	// Matches defines how to match values.
	Matches types.StringAnyMap
}

// CommonParams ...
type CommonParams struct {
	Page     int
	PageSize int
	Query    Query
	Sort     []string
	Group    string
}

// Parse parse a api query param to db query param
func Parse(source proto.QueryCommonParams, fieldMapping map[string]string) (*CommonParams, error) {
	dst := CommonParams{
		Page:     source.Offset,
		PageSize: source.Limit,
	}

	// parseKey convert model Field to db column
	parseKey := func(s string) string {
		d := convert.BigToUnderline(s)
		if fieldMapping == nil {
			return d
		}

		if _, ok := fieldMapping[s]; ok {
			d = fieldMapping[s]
		}

		return d
	}

	// sort
	if len(source.Sorters) > 0 {
		for i := range source.Sorters {
			s := source.Sorters[i].OrderBy
			d := "+" + parseKey(s)
			if source.Sorters[i].Desc {
				d = "-" + parseKey(s)
			}

			dst.Sort = append(dst.Sort, d)
		}
	}

	// query
	query := Query{Matches: make(map[string]interface{})}
	for _, m := range source.Filters {
		key := parseKey(m.Name)
		switch m.Type {
		case proto.MatchExact:
			query.Matches[key] = ExactMatch{Value: m.Value}
		case proto.MatchFuzzy:
			query.Matches[key] = FuzzyMatch{Value: "%" + m.Value + "%"}
		case proto.MatchIn:
			query.Matches[key] = InMatch{Values: strings.Split(m.Value, proto.MatchSeparator)}
		case proto.MatchRange:
			if !strings.Contains(m.Value, proto.MatchSeparator) {
				return nil, errors.New("range match should given two number")
			}
			ranges := strings.Split(m.Value, proto.MatchSeparator)
			query.Matches[key] = RangeMatch{Start: ranges[0], End: ranges[1]}
		default:
			return nil, fmt.Errorf("invalid filter type %s", m.Type)
		}
	}

	if source.Group != "" {
		dst.Group = parseKey(source.Group)
	}

	dst.Query = query

	return &dst, nil
}
