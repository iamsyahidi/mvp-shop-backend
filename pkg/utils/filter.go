package utils

import (
	"regexp"
	"strconv"
	"strings"
)

type Pagination struct {
	Limit         int
	Page          int
	SortField     string
	SortDirection string
}

func GeneratePaginationFromRequest(filter map[string][]string) (pagination Pagination, search map[string]string) {

	var limit int
	if paramLimit, ok := filter["limit"]; ok && (len(paramLimit) > 0) {
		// string to int
		limit, _ = strconv.Atoi(paramLimit[0])
	}

	var page int
	if paramPage, ok := filter["page"]; ok && (len(paramPage) > 0) {
		// string to int
		page, _ = strconv.Atoi(paramPage[0])
	}

	var sortField string
	if paramSortField, ok := filter["sort_field"]; ok && (len(paramSortField) > 0) {
		sortField = paramSortField[0]
	}

	var sortDirection string
	if paramSortDirection, ok := filter["sort_direction"]; ok && (len(paramSortDirection) > 0) {
		sortDirection = paramSortDirection[0]
	}

	search = make(map[string]string)
	if paramSearch, ok := filter["search"]; ok && (len(paramSearch) > 0) {
		entries := strings.Split(strings.TrimSpace(filter["search"][0]), ",")
		for _, e := range entries {
			parts := strings.Split(e, "=")
			search[parts[0]] = parts[1]
		}
	}

	isValidLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	if !isValidLetter(sortField) {
		sortField = "created_at"
	}

	switch {
	case strings.ToLower(sortDirection) == "asc":
		sortDirection = "ASC"
	case strings.ToLower(sortDirection) == "desc":
		sortDirection = "DESC"
	default:
		sortDirection = "ASC"
	}

	if limit == 0 {
		limit = 5
	}
	if page == 0 {
		page = 1
	}

	pagination = Pagination{
		Limit:         limit,
		Page:          page,
		SortField:     sortField,
		SortDirection: sortDirection,
	}
	return pagination, search
}
