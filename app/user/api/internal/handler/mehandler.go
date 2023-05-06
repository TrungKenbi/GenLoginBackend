package handler

import (
	"net/http"

	"genlogin/app/user/api/internal/logic"
	"genlogin/app/user/api/internal/svc"
	"genlogin/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func meHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMeLogic(r.Context(), svcCtx)
		resp, err := l.Me(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
