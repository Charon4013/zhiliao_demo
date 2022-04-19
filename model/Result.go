/**
 * @Author: Pan
 * @Date: 2022/3/18 14:03
 */

package model

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResult(data interface{}, c int, m ...string) *Result {
	r := &Result{Data: data, Code: c}

	if e, ok := data.(error); ok {
		if m == nil {
			r.Msg = e.Error()
		}
	} else {
		r.Msg = "SUCCESS"
	}
	if len(m) > 0 {
		r.Msg = m[0]
	}

	return r
}
