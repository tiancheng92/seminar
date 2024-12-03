package repository

import (
	"github.com/Yostardev/gf"
	"github.com/tiancheng92/seminar/pkg/errors"
	"github.com/tiancheng92/seminar/pkg/errors/ecode"
	"github.com/tiancheng92/seminar/pkg/log"
	"github.com/tiancheng92/seminar/store"
	"github.com/tiancheng92/seminar/store/model"
	"github.com/tiancheng92/seminar/types/paginate"
	"gorm.io/gorm"
	"net/url"
	"strings"
)

type genericRepository[M model.Interface] struct {
	db           *gorm.DB
	paginateData *paginate.Data[M]
}

func newGenericRepository[M model.Interface]() *genericRepository[M] {
	return &genericRepository[M]{
		db:           store.GetDefaultDB(),
		paginateData: new(paginate.Data[M]),
	}
}

func (r *genericRepository[M]) Create(attributes M) (*M, error) {
	err := r.db.Model(new(M)).Create(&attributes).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil, errors.WithCode(ecode.ErrDuplicateKey, err)
	}
	return &attributes, errors.WithCode(ecode.ErrCreate, err)
}

func (r *genericRepository[M]) Update(pk any, attributes M) (*M, error) {
	err := r.db.Model(new(M)).
		Select("*").Omit("id", "created_at").Where(gf.StringJoin("`", (*new(M)).GetPrimaryKeyName(), "` = ?"), pk).Updates(&attributes).
		Session(&gorm.Session{NewDB: true}).Where(gf.StringJoin("`", (*new(M)).GetPrimaryKeyName(), "` = ?"), pk).First(&attributes).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil, errors.WithCode(ecode.ErrDuplicateKey, err)
	}
	return &attributes, errors.WithCode(ecode.ErrUpdate, err)
}

func (r *genericRepository[M]) Delete(pk any) error {
	var ent M
	err := r.db.Model(new(M)).Where(gf.StringJoin("`", (*new(M)).GetPrimaryKeyName(), "` = ?"), pk).Delete(&ent).Error
	return errors.WithCode(ecode.ErrDelete, err)
}

func (r *genericRepository[M]) Get(pk any) (*M, error) {
	var ent M
	err := r.db.Model(new(M)).Where(gf.StringJoin("`", (*new(M)).GetPrimaryKeyName(), "` = ?"), pk).First(&ent).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.WithCode(ecode.ErrDataNotFound, err)
	}
	return &ent, errors.WithCode(ecode.ErrGet, err)
}

func (r *genericRepository[M]) Distinct(field string) ([]string, error) {
	var ent []string
	err := r.db.Model(new(M)).Distinct().Pluck(field, &ent).Error
	return ent, errors.WithCode(ecode.ErrGet, err)
}

func (r *genericRepository[M]) List(pq *paginate.Query) (*paginate.Data[M], error) {
	r.paginateData.Init(pq)
	err := r.db.Model(new(M)).Scopes(Paginate[M](pq)).Find(&r.paginateData.Items).Offset(-1).Limit(-1).Count(&r.paginateData.Total).Error
	return r.paginateData, errors.WithCode(ecode.ErrGet, err)
}

func (r *genericRepository[M]) All(params url.Values) ([]*M, error) {
	var ent []*M
	db := r.db.Model(new(M))
	db = handleFilter(db, params, getFieldList[M](db))
	if search := params.Get("search"); search != "" {
		db = handleFuzzySearch[M](db, search)
	}
	err := db.Order(gf.StringJoin("`", (*new(M)).GetDefaultOrderBy(), "` ", strings.ToUpper((*new(M)).GetDefaultOrder()))).Find(&ent).Error
	return ent, errors.WithCode(ecode.ErrGet, err)
}

func (r *genericRepository[M]) Count(params url.Values) (int64, error) {
	var count int64
	db := r.db.Model(new(M))
	db = handleFilter(db, params, getFieldList[M](db))
	if search := params.Get("search"); search != "" {
		db = handleFuzzySearch[M](db, search)
	}
	err := db.Count(&count).Error
	return count, errors.WithCode(ecode.ErrGet, err)
}

func Paginate[M model.Interface](pq *paginate.Query) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fieldList := getFieldList[M](db)
		db = db.Offset((pq.Page - 1) * pq.PageSize).Limit(pq.PageSize)
		db = handleFilter(db, pq.Params, fieldList)
		db = handleFuzzySearch[M](db, pq.Search)

		if !gf.ArrayContains(fieldList, pq.OrderBy) {
			pq.OrderBy = (*new(M)).GetDefaultOrderBy()
		}

		if !gf.ArrayContains([]string{"desc", "asc"}, strings.ToLower(pq.Order)) {
			pq.Order = (*new(M)).GetDefaultOrder()
		}

		return db.Order(gf.StringJoin("`", pq.OrderBy, "` ", strings.ToUpper(pq.Order)))
	}
}

func getFieldList[M model.Interface](db *gorm.DB) []string {
	stmt := &gorm.Statement{DB: db}
	if err := stmt.Parse(new(M)); err != nil {
		log.Errorf("parse model error: %+v", err)
		return []string{}
	}
	fieldList := make([]string, 0, len(stmt.Schema.Fields))
	for i := range stmt.Schema.Fields {
		if stmt.Schema.Fields[i].DBName != "" {
			fieldList = append(fieldList, stmt.Schema.Fields[i].DBName)
		}
	}
	return fieldList
}

func handleFuzzySearch[M model.Interface](db *gorm.DB, search string) *gorm.DB {
	fuzzySearchFieldList := (*new(M)).GetFuzzySearchFieldList()
	if search != "" && len(fuzzySearchFieldList) > 0 {
		searchField := make([]string, 0, len(fuzzySearchFieldList))
		for i := range fuzzySearchFieldList {
			searchField = append(searchField, gf.StringJoin("IFNULL(`", strings.TrimSpace(fuzzySearchFieldList[i]), "`, '')"))
		}
		db = db.Where(gf.StringJoin("CONCAT(", strings.Join(searchField, ", "), ") LIKE ?"), gf.StringJoin("%", search, "%"))
	}
	return db
}

func handleFilter(db *gorm.DB, params url.Values, fieldList []string) *gorm.DB {
	for k, v := range params {
		if gf.ArrayContains(fieldList, k) {
			if len(v) == 1 {
				if v[0] != "" {
					db = db.Where(gf.StringJoin("`", k, "` = ?"), v[0])
				}
			} else {
				db = db.Where(gf.StringJoin("`", k, "` IN ?"), v)
			}
		}

		fieldSlice := strings.Split(k, "__")
		if len(fieldSlice) == 2 {
			suffix := fieldSlice[len(fieldSlice)-1]
			field := strings.ReplaceAll(k, gf.StringJoin("__", suffix), "")
			if gf.ArrayContains(fieldList, field) {
				for i := range v {
					if v[i] != "" {
						switch suffix {
						case "gte":
							db = db.Where(gf.StringJoin("`", field, "` >= ?"), v[i])
						case "gt":
							db = db.Where(gf.StringJoin("`", field, "` > ?"), v[i])
						case "lte":
							db = db.Where(gf.StringJoin("`", field, "` <= ?"), v[i])
						case "lt":
							db = db.Where(gf.StringJoin("`", field, "` < ?"), v[i])
						case "ne":
							db = db.Where(gf.StringJoin("`", field, "` != ?"), v[i])
						case "sw":
							db = db.Where(gf.StringJoin("`", field, "` LIKE ?"), gf.StringJoin(v[i], "%"))
						case "ew":
							db = db.Where(gf.StringJoin("`", field, "` LIKE ?"), gf.StringJoin("%", v[i]))
						case "like":
							db = db.Where(gf.StringJoin("`", field, "` LIKE ?"), gf.StringJoin("%", v[i], "%"))
						}
					}
				}
			}
		}
	}
	return db
}
