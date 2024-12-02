package paginate

import "net/url"

type Data[M any] struct {
	*Info
	Items []*M
}

func (p *Data[M]) GetPaginate() *Info {
	info := new(Info)
	if p != nil {
		info = p.Info
	}
	if info.PageSize == 0 {
		info.PageSize = int(info.Total)
	}
	return info
}

func (p *Data[M]) GetItems() any {
	if p != nil {
		return p.Items
	}
	return nil
}

func (p *Data[M]) Init(q *Query) {
	p.Info = &Info{
		Page:     q.Page,
		PageSize: q.PageSize,
	}
}

// Query 分页查询
type Query struct {
	Page     int        // 页数
	PageSize int        // 每页数据量
	Order    string     // 排序方式
	OrderBy  string     // 排序字段
	Search   string     // 关键字搜索
	Params   url.Values // 其他参数
}

type Info struct {
	Total    int64 `json:"total"`     // 数据总数
	Page     int   `json:"page"`      // 页数
	PageSize int   `json:"page_size"` // 每页数据量
}
