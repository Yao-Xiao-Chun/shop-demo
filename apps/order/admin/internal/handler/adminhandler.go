package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"shop-demo/apps/order/admin/internal/logic"
	"shop-demo/apps/order/admin/internal/svc"
	"shop-demo/apps/order/admin/internal/types"
)

func AdminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAdminLogic(r.Context(), svcCtx)
		resp, err := l.Admin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
