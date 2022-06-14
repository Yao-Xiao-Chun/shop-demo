package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"shop-demo/apps/app/api/internal/logic"
	"shop-demo/apps/app/api/internal/svc"
	"shop-demo/apps/app/api/internal/types"
)

func LoginApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginApiLogic(r.Context(), svcCtx)
		resp, err := l.LoginApi(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
