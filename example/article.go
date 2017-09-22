package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/alenstar/nanotask/config"
	"github.com/alenstar/nanotask/controller"
	"github.com/alenstar/nanotask/log"
	"github.com/alenstar/nanotask/modules"
	"strconv"
	"strings"
)

type ArticleController struct {
	controller.CommonController
}

// id is md5(article-title)
func (a *ArticleController) Get() {
	if id, ok := a.Ctx.Params["id"]; ok {
		log.Info("Article ", id)
		aid, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Error("strconv.ParseUint", err.Error())
			aid = 0
		}
		ainfo := &ArticleInfo{Id: 0, ArticleId: aid}
		if aid != 0 {
			has, err := modules.DefaultEngine().Get(ainfo)
			if err != nil || !has {
				ainfo.Id = 0
				log.Warn("Get ", err.Error())
			} else {
				out := struct {
					Status  int          `json:"status"`
					Content *ArticleInfo `json:"content"`
				}{
					Status:  200,
					Content: ainfo,
				}
				// b, _ := json.Marshal(&out)
				// a.Ctx.Write(b)
				a.WriteJSON(&out)
				return
			}
		}
	} else {
		var err error
		aid := uint64(0) // use aid to page
		if lid, ok := a.Ctx.Params["lastid"]; ok {
			aid, err = strconv.ParseUint(lid, 10, 64)
			if err != nil {
				log.Error("strconv.ParseUint", err.Error())
				aid = 0
			}
		}

		// get id by ArticleId
		ainfo := &ArticleInfo{Id: 0, ArticleId: aid}
		if aid != 0 {
			has, err := modules.DefaultEngine().Get(ainfo)
			if err != nil || !has {
				ainfo.Id = 0
				log.Warn("Get ", err.Error())
			}
		}

		var articleIds []uint64
		results, _ := modules.DefaultEngine().Query(fmt.Sprintf("select article_id from tb_article_info where id > %d order by id limit 50", ainfo.Id)) //  limit 2;

		for _, v := range results {
			for k, vv := range v {
				log.Debug("Query: ", k, string(vv))
				ids := strings.Split(string(vv), ",")
				for _, i := range ids {
					iid, err := strconv.ParseUint(i, 10, 64)
					if err == nil {
						articleIds = append(articleIds, iid)
					}
				}
				log.Debug("ID ", ids)
			}
		}

		log.Debug("Article Get ", articleIds)
		out := struct {
			Status int      `json:"status"`
			Ids    []uint64 `json:"ids"`
		}{
			Status: 200,
			Ids:    articleIds,
		}
		b, _ := json.Marshal(&out)
		a.Ctx.Write(b)
		return
	}
	a.NotFound()
}

func (a *ArticleController) Post() {
	log.Debug("Article Post", string(a.Ctx.CopyBody()))
	ainfo := &ArticleInfo{}
	err := json.Unmarshal(a.Ctx.CopyBody(), ainfo)
	if err == nil {
		if len(ainfo.Title) > 0 && len(ainfo.Author) > 0 {
			ainfo.ArticleId = ainfo.CalcArticleId()
			_, err = modules.DefaultEngine().Insert(ainfo)
			if err != nil {
				log.Error("orm Insert ", err.Error())
			}
			a.Ctx.WriteString(fmt.Sprintf(`{"status":200, "content":"%d"}`, ainfo.ArticleId))
			return
		} else {
			log.Debug("Bad ArticleInfo: ", ainfo)
		}
	} else {
		log.Error("json.Unmarshal ", err.Error())

	}
	a.Ctx.WriteString(fmt.Sprintf(`{"status":404, "info":"%s"}`, err.Error()))
}

func (a *ArticleController) Delete() {
	log.Debug("Article Delete")
	if id, ok := a.Ctx.Params["id"]; ok {
		iid, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Error("strconv.ParseUint", err.Error())
			a.NotFound()
			return
		}
		ainfo := &ArticleInfo{ArticleId: iid}
		_, err = modules.DefaultEngine().Delete(ainfo)
		if err != nil {
			log.Error("orm Delete ", err.Error())
			a.NotFound()
		}
	} else {
		log.Debug("Not found")
		a.NotFound()
	}
}

func (a *ArticleController) Put() {
	log.Debug("Article Put", a.Ctx.Params["id"])
	if id, ok := a.Ctx.Params["id"]; ok {
		iid, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Error("strconv.ParseUint", err.Error())
			a.NotFound()
			return
		}
		ainfo := &ArticleInfo{ArticleId: iid}
		err = json.Unmarshal(a.Ctx.CopyBody(), ainfo)
		if err == nil {
			if len(ainfo.Title) > 0 && len(ainfo.Author) > 0 {
				ainfo.ArticleId = iid
				_, err = modules.DefaultEngine().Update(ainfo)
				if err != nil {
					log.Error("orm Insert ", err.Error())
				}
				a.Ctx.WriteString(fmt.Sprintf(`{"status":200, "content":"%d"}`, ainfo.ArticleId))
				return
			} else {
				log.Debug("Bad ArticleInfo: ", ainfo)
			}
		} else {
			log.Error("json.Unmarshal ", err.Error())

		}
	}
	a.Ctx.WriteString(`{"status":404, "info":"xxx"}`)
}
