package server

import (
	"github.com/duanchuansong/stk/web"
	"github.com/duanchuansong/stk/xlog"
	"net/http"
)

func StartStaticServer(cfg *web.Static) {
	xlog.Infof("listening on %s,path:%s", cfg.Port, cfg.Dir)
	err := http.ListenAndServe(cfg.Port, http.FileServer(http.Dir(cfg.Dir)))
	if err != nil {
		xlog.Errorf("http.ListenAndServe(%s, http.FileServer(http.Dir(%s))),err:%v", cfg.Port, cfg.Dir, err)
	}
}
