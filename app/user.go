package app

import (
	"encoding/json"
	"fmt"
	_ "github.com/alenstar/nanoweb/config"
	"github.com/alenstar/nanoweb/controller"
	"github.com/alenstar/nanoweb/log"
	"github.com/alenstar/nanoweb/modules"
	"strconv"
	"strings"
)

type UserController struct {
	controller.CommonController
}

func (u *UserController) Get() {
	if id, ok := u.Ctx.Params["id"]; ok {
		log.Info("UserInfo ", id)
		uid, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Error("strconv.ParseUint", err.Error())
			uid = 0
		}
		uinfo := &UserInfo{Id: 0, UserId: uid}
		if uid != 0 {
			has, err := modules.DefaultEngine().Get(uinfo)
			if err != nil || !has {
				uinfo.Id = 0
				log.Warn("Get ", err.Error())
			} else {
				out := struct {
					Status  int       `json:"status"`
					Content *UserInfo `json:"content"`
				}{
					Status:  200,
					Content: uinfo,
				}
				// b, _ := json.Marshal(&out)
				// u.Ctx.Write(b)
				u.WriteJSON(&out)
				return
			}
		}
	} else {
		var err error
		uid := uint64(0) // use uid to page
		if lid, ok := u.Ctx.Params["lastid"]; ok {
			uid, err = strconv.ParseUint(lid, 10, 64)
			if err != nil {
				log.Error("strconv.ParseUint", err.Error())
				uid = 0
			}
		}

		// get id by ArticleId
		uinfo := &UserInfo{Id: 0, UserId: uid}
		if uid != 0 {
			has, err := modules.DefaultEngine().Get(uinfo)
			if err != nil || !has {
				uinfo.Id = 0
				log.Warn("Get ", err.Error())
			}
		}

		var userIds []uint64
		results, _ := modules.DefaultEngine().Query(fmt.Sprintf("select user_id from tb_user_info where id > %d order by id limit 50", uinfo.Id)) //  limit 2;

		for _, v := range results {
			for k, vv := range v {
				log.Debug("Query: ", k, string(vv))
				ids := strings.Split(string(vv), ",")
				for _, i := range ids {
					iid, err := strconv.ParseUint(i, 10, 64)
					if err == nil {
						userIds = append(userIds, iid)
					}
				}
				log.Debug("ID ", ids)
			}
		}

		log.Debug("Article Get ", userIds)
		out := struct {
			Status int      `json:"status"`
			Ids    []uint64 `json:"ids"`
		}{
			Status: 200,
			Ids:    userIds,
		}
		b, _ := json.Marshal(&out)
		u.Ctx.Write(b)
		return
	}
	u.NotFound()
}

func (u *UserController) Post() {
	log.Debug("User Post")
	uinfo := &UserInfo{}
	err := json.Unmarshal(u.Ctx.CopyBody(), uinfo)
	if err == nil {
		if len(uinfo.Name) > 0 && len(uinfo.Email) > 0 {
			uinfo.UserId = uinfo.CalcUserId()
			_, err = modules.DefaultEngine().Insert(uinfo)
			if err != nil {
				log.Error("orm Insert ", err.Error())
			}
			u.Ctx.WriteString(fmt.Sprintf(`{"status":200, "content":"%d"}`, uinfo.UserId))
			return
		} else {
			log.Debug("Bad UserInfo: ", uinfo)
		}
	} else {
		log.Error("json.Unmarshal ", err.Error())

	}
	u.Ctx.WriteString(fmt.Sprintf(`{"status":404, "info":"%s"}`, err.Error()))

}

func (u *UserController) Delete() {
	log.Debug("User Delete")
	if id, ok := u.Ctx.Params["id"]; ok {
		iid, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Error("strconv.ParseUint", err.Error())
			u.NotFound()
			return
		}
		uinfo := &UserInfo{UserId: iid}
		_, err = modules.DefaultEngine().Delete(uinfo)
		if err != nil {
			log.Error("orm Delete ", err.Error())
			u.NotFound()
		}
	} else {
		log.Debug("Not found")
		u.NotFound()
	}
}

func (u *UserController) Put() {
	log.Debug("User Put", u.Ctx.Params)
	if id, ok := u.Ctx.Params["id"]; ok {
		iid, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Error("strconv.ParseUint", err.Error())
			u.NotFound()
			return
		}
		uinfo := &UserInfo{UserId: iid}
		err = json.Unmarshal(u.Ctx.CopyBody(), uinfo)
		if err == nil {
			if len(uinfo.Name) > 0 && len(uinfo.Email) > 0 {
				uinfo.UserId = iid
				_, err = modules.DefaultEngine().Update(uinfo)
				if err != nil {
					log.Error("orm Insert ", err.Error())
				}
				u.Ctx.WriteString(fmt.Sprintf(`{"status":200, "content":"%d"}`, uinfo.UserId))
				return
			} else {
				log.Debug("Bad UserInfo: ", uinfo)
			}
		} else {
			log.Error("json.Unmarshal ", err.Error())

		}
	}
	u.Ctx.WriteString(`{"status":404, "info":"xxx"}`)
}
