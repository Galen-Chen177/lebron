package handler

import (
	"net/http"

	"lebron/apps/app/api/internal/logic"
	"lebron/apps/app/api/internal/svc"
	"lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewApiLogic(r.Context(), svcCtx)
		resp, err := l.Api(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
