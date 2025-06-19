package persistence

import (
	"fmt"

	filterModel "github.com/test-tzs/nomraeite/internal/domain/model/filter"
	"gorm.io/gorm"
)

// GormFilterBuilder applies domain filters to GORM queries
type GormFilterBuilder struct{}

// ApplyBaseFilter applies base filter conditions to a GORM query
func (b *GormFilterBuilder) ApplyBaseFilter(query *gorm.DB, baseFilter *filterModel.BaseFilter) *gorm.DB {
	// Apply joins first to ensure conditions can reference joined tables
	for _, join := range baseFilter.Joins {
		joinStmt := fmt.Sprintf("%s %s ON %s", join.Type, join.Table, join.Condition)
		if len(join.Parameters) > 0 {
			query = query.Joins(joinStmt, join.Parameters...)
		} else {
			query = query.Joins(joinStmt)
		}
	}

	// Apply standard conditions
	for _, condition := range baseFilter.Conditions {
		switch condition.Operator {
		case filterModel.Equal:
			query = query.Where(fmt.Sprintf("%s = ?", condition.Field), condition.Value)
		case filterModel.NotEqual:
			query = query.Where(fmt.Sprintf("%s != ?", condition.Field), condition.Value)
		case filterModel.GreaterThan:
			query = query.Where(fmt.Sprintf("%s > ?", condition.Field), condition.Value)
		case filterModel.GreaterThanOrEqual:
			query = query.Where(fmt.Sprintf("%s >= ?", condition.Field), condition.Value)
		case filterModel.LessThan:
			query = query.Where(fmt.Sprintf("%s < ?", condition.Field), condition.Value)
		case filterModel.LessThanOrEqual:
			query = query.Where(fmt.Sprintf("%s <= ?", condition.Field), condition.Value)
		case filterModel.Like:
			query = query.Where(fmt.Sprintf("%s LIKE ?", condition.Field), condition.Value)
		case filterModel.In:
			query = query.Where(fmt.Sprintf("%s IN (?)", condition.Field), condition.Value)
		}
	}

	// Apply OR conditions
	for _, orGroup := range baseFilter.OrConditions {
		if len(orGroup.Conditions) > 0 {
			var conditions []string
			var values []any

			for _, condition := range orGroup.Conditions {
				switch condition.Operator {
				case filterModel.Equal:
					conditions = append(conditions, fmt.Sprintf("%s = ?", condition.Field))
				case filterModel.NotEqual:
					conditions = append(conditions, fmt.Sprintf("%s != ?", condition.Field))
				case filterModel.GreaterThan:
					conditions = append(conditions, fmt.Sprintf("%s > ?", condition.Field))
				case filterModel.GreaterThanOrEqual:
					conditions = append(conditions, fmt.Sprintf("%s >= ?", condition.Field))
				case filterModel.LessThan:
					conditions = append(conditions, fmt.Sprintf("%s < ?", condition.Field))
				case filterModel.LessThanOrEqual:
					conditions = append(conditions, fmt.Sprintf("%s <= ?", condition.Field))
				case filterModel.Like:
					conditions = append(conditions, fmt.Sprintf("%s LIKE ?", condition.Field))
				case filterModel.In:
					conditions = append(conditions, fmt.Sprintf("%s IN (?)", condition.Field))
				}
				values = append(values, condition.Value)
			}

			// Only apply if we have valid conditions
			if len(conditions) > 0 {
				// Build the OR clause
				orClause := "(" + conditions[0] + ")"
				orArgs := []any{values[0]}

				for i := 1; i < len(conditions); i++ {
					orClause = orClause + " OR (" + conditions[i] + ")"
					orArgs = append(orArgs, values[i])
				}

				query = query.Where(orClause, orArgs...)
			}
		}
	}

	// Apply date filters
	for _, dateFilter := range baseFilter.DateFilters {
		if dateFilter.Date != nil {
			switch dateFilter.Operator {
			case filterModel.Equal:
				query = query.Where("DATE("+dateFilter.Field+") = DATE(?)", dateFilter.Date)
			case filterModel.NotEqual:
				query = query.Where("DATE("+dateFilter.Field+") != DATE(?)", dateFilter.Date)
			case filterModel.GreaterThan:
				query = query.Where("DATE("+dateFilter.Field+") > DATE(?)", dateFilter.Date)
			case filterModel.GreaterThanOrEqual:
				query = query.Where("DATE("+dateFilter.Field+") >= DATE(?)", dateFilter.Date)
			case filterModel.LessThan:
				query = query.Where("DATE("+dateFilter.Field+") < DATE(?)", dateFilter.Date)
			case filterModel.LessThanOrEqual:
				query = query.Where("DATE("+dateFilter.Field+") <= DATE(?)", dateFilter.Date)
			}
		}
	}

	// Apply sorting
	if len(baseFilter.Sort) > 0 {
		for _, sort := range baseFilter.Sort {
			query = query.Order(fmt.Sprintf("%s %s", sort.Field, sort.Direction))
		}
	} else {
		// Default sort by ID if not specified
		query = query.Order("id DESC")
	}

	return query
}

// ApplyPagination applies pagination to a GORM query
func (b *GormFilterBuilder) ApplyPagination(query *gorm.DB, pagination filterModel.Pagination) *gorm.DB {
	offset := (pagination.Page - 1) * pagination.PageSize
	return query.Offset(offset).Limit(pagination.PageSize)
}

// NewGormFilterBuilder creates a new GormFilterBuilder
func NewGormFilterBuilder() *GormFilterBuilder {
	return &GormFilterBuilder{}
}
