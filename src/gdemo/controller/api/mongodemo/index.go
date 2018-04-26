package mongodemo

import (
	"gdemo/dao"
	"gdemo/errno"
	"gdemo/svc"

	"github.com/goinbox/exception"
	"github.com/goinbox/gohttp/query"
)

type indexActionParams struct {
	Status int `bson:"status"`

	offset int
	cnt    int
}

var indexQueryConditions map[string]string = map[string]string{
	"status": dao.MONGO_COND_EQUAL,
}

func (d *MongoDemoController) IndexAction(context *MongoDemoContext) {
	ap, exists, e := d.parseIndexActionParams(context)
	if e != nil {
		context.ApiData.Err = e
		return
	}

	mqp := &svc.MongoQueryParams{
		ParamsStructPtr: ap,
		Exists:          exists,
		Conditions:      indexQueryConditions,

		OrderBy: "id",
		Offset:  ap.offset,
		Cnt:     ap.cnt,
	}

	entities, err := context.demoSvc.SelectAll(mqp)
	if err != nil {
		context.ApiData.Err = exception.New(errno.E_SYS_MONGO_ERROR, err.Error())
		return
	}

	context.ApiData.Data = entities
}

func (d *MongoDemoController) parseIndexActionParams(context *MongoDemoContext) (*indexActionParams, map[string]bool, *exception.Exception) {
	ap := new(indexActionParams)

	qs := query.NewQuerySet()
	qs.IntVar(&ap.Status, "status", false, errno.E_API_DEMO_INVALID_STATUS, "invalid status", nil)
	qs.IntVar(&ap.offset, "offset", false, errno.E_COMMON_INVALID_QUERY_OFFSET, "invalid offset", nil)
	qs.IntVar(&ap.cnt, "cnt", false, errno.E_COMMON_INVALID_QUERY_CNT, "invalid cnt", nil)
	e := qs.Parse(context.QueryValues)
	if e != nil {
		return ap, nil, e
	}

	if ap.Status < 0 {
		return ap, nil, exception.New(errno.E_API_DEMO_INVALID_STATUS, "invalid status")
	}
	if ap.offset < 0 {
		return ap, nil, exception.New(errno.E_COMMON_INVALID_QUERY_OFFSET, "invalid offset")
	}
	if ap.cnt < 0 {
		return ap, nil, exception.New(errno.E_COMMON_INVALID_QUERY_CNT, "invalid cnt")
	}

	return ap, qs.ExistsInfo(), nil
}