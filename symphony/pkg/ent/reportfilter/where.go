// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package reportfilter

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Filters applies equality check predicate on the "filters" field. It's identical to FiltersEQ.
func Filters(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFilters), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.ReportFilter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReportFilter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.ReportFilter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReportFilter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.ReportFilter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReportFilter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.ReportFilter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReportFilter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ReportFilter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReportFilter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ReportFilter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReportFilter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// EntityEQ applies the EQ predicate on the "entity" field.
func EntityEQ(v Entity) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntity), v))
	})
}

// EntityNEQ applies the NEQ predicate on the "entity" field.
func EntityNEQ(v Entity) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntity), v))
	})
}

// EntityIn applies the In predicate on the "entity" field.
func EntityIn(vs ...Entity) predicate.ReportFilter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReportFilter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEntity), v...))
	})
}

// EntityNotIn applies the NotIn predicate on the "entity" field.
func EntityNotIn(vs ...Entity) predicate.ReportFilter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReportFilter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEntity), v...))
	})
}

// FiltersEQ applies the EQ predicate on the "filters" field.
func FiltersEQ(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFilters), v))
	})
}

// FiltersNEQ applies the NEQ predicate on the "filters" field.
func FiltersNEQ(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFilters), v))
	})
}

// FiltersIn applies the In predicate on the "filters" field.
func FiltersIn(vs ...string) predicate.ReportFilter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReportFilter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFilters), v...))
	})
}

// FiltersNotIn applies the NotIn predicate on the "filters" field.
func FiltersNotIn(vs ...string) predicate.ReportFilter {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ReportFilter(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFilters), v...))
	})
}

// FiltersGT applies the GT predicate on the "filters" field.
func FiltersGT(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFilters), v))
	})
}

// FiltersGTE applies the GTE predicate on the "filters" field.
func FiltersGTE(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFilters), v))
	})
}

// FiltersLT applies the LT predicate on the "filters" field.
func FiltersLT(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFilters), v))
	})
}

// FiltersLTE applies the LTE predicate on the "filters" field.
func FiltersLTE(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFilters), v))
	})
}

// FiltersContains applies the Contains predicate on the "filters" field.
func FiltersContains(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFilters), v))
	})
}

// FiltersHasPrefix applies the HasPrefix predicate on the "filters" field.
func FiltersHasPrefix(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFilters), v))
	})
}

// FiltersHasSuffix applies the HasSuffix predicate on the "filters" field.
func FiltersHasSuffix(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFilters), v))
	})
}

// FiltersEqualFold applies the EqualFold predicate on the "filters" field.
func FiltersEqualFold(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFilters), v))
	})
}

// FiltersContainsFold applies the ContainsFold predicate on the "filters" field.
func FiltersContainsFold(v string) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFilters), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.ReportFilter) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.ReportFilter) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ReportFilter) predicate.ReportFilter {
	return predicate.ReportFilter(func(s *sql.Selector) {
		p(s.Not())
	})
}
