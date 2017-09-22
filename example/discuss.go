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

type DiscussController struct {
	controller.CommonController
}

func (d *DiscussController) Get() {
	log.Debug("Discuss Get")
	if id, ok := d.Ctx.Params["id"]; ok {
		log.Info("Discuss ", id)
		did, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Error("strconv.ParseUint", err.Error())
			did = 0
		}
		dinfo := &DiscussInfo{Id: 0, DiscussId: did}
		if did != 0 {
			has, err := modules.DefaultEngine().Get(dinfo)
			if err != nil || !has {
				dinfo.Id = 0
				log.Warn("Get ", err.Error())
			} else {
				out := struct {
					Status  int          `json:"status"`
					Content *DiscussInfo `json:"content"`
				}{
					Status:  200,
					Content: dinfo,
				}
				//b, _ := json.Marshal(&out)
				//d.Ctx.Write(b)
				d.WriteJSON(&out)
				return
			}
		}
	} else {
		var err error
		did := uint64(0)
		if lid, ok := d.Ctx.Params["lastid"]; ok {
			did, err = strconv.ParseUint(lid, 10, 64)
			if err != nil {
				log.Error("strconv.ParseUint", err.Error())
				did = 0
			}
		}

		// get id by DiscussId
		dinfo := &DiscussInfo{Id: 0, DiscussId: did}
		if did != 0 {
			has, err := modules.DefaultEngine().Get(dinfo)
			if err != nil || !has {
				dinfo.Id = 0
				log.Warn("Get ", err.Error())
			}
		}

		var discussIds []uint64
		results, _ := modules.DefaultEngine().Query(fmt.Sprintf("select discuss_id from tb_discuss_info where id > %d order by id limit 50", dinfo.Id)) //  limit 2;

		for _, v := range results {
			for k, vv := range v {
				log.Debug("Query: ", k, string(vv))
				ids := strings.Split(string(vv), ",")
				for _, i := range ids {
					iid, err := strconv.ParseUint(i, 10, 64)
					if err == nil {
						discussIds = append(discussIds, iid)
					}
				}
				log.Debug("ID ", ids)
			}
		}

		log.Debug("Discuss Get ", discussIds)
		out := struct {
			Status int      `json:"status"`
			Ids    []uint64 `json:"ids"`
		}{
			Status: 200,
			Ids:    discussIds,
		}
		b, _ := json.Marshal(&out)
		d.Ctx.Write(b)
		return
	}
	d.NotFound()
}

func (d *DiscussController) Post() {
	log.Debug("Discuss Post")
	dinfo := &DiscussInfo{}
	err := json.Unmarshal(d.Ctx.CopyBody(), dinfo)
	if err == nil {
		if len(dinfo.Content) > 0 {
			dinfo.DiscussId = dinfo.CalcDiscussId()
			_, err = modules.DefaultEngine().Insert(dinfo)
			if err != nil {
				log.Error("orm Insert ", err.Error())
			}
			d.Ctx.WriteString(fmt.Sprintf(`{"status":200, "content":"%d"}`, dinfo.DiscussId))
			return
		} else {
			log.Debug("Bad DiscussInfo: ", dinfo)
		}
	} else {
		log.Error("json.Unmarshal ", err.Error())

	}
	d.Ctx.WriteString(fmt.Sprintf(`{"status":404, "info":"%s"}`, err.Error()))
}

func (d *DiscussController) Delete() {
	log.Debug("Discuss Delete")
	if id, ok := d.Ctx.Params["id"]; ok {
		iid, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Error("strconv.ParseUint", err.Error())
			d.NotFound()
			return
		}
		dinfo := &DiscussInfo{}
		dinfo.DiscussId = iid
		_, err = modules.DefaultEngine().Delete(dinfo)
		if err != nil {
			log.Error("orm Delete ", err.Error())
			d.NotFound()
		}
	} else {
		log.Debug("Not found")
		d.NotFound()
	}
}

func (d *DiscussController) Put() {
	log.Debug("Discuss Put", d.Ctx.Params)
	if id, ok := d.Ctx.Params["id"]; ok {
		iid, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Error("strconv.ParseUint", err.Error())
			d.NotFound()
			return
		}

		dinfo := &DiscussInfo{DiscussId: iid}
		err = json.Unmarshal(d.Ctx.CopyBody(), dinfo)
		if err == nil {
			if len(dinfo.Content) > 0 {
				dinfo.DiscussId = iid
				_, err = modules.DefaultEngine().Update(dinfo)
				if err != nil {
					log.Error("orm Insert ", err.Error())
				}
				d.Ctx.WriteString(fmt.Sprintf(`{"status":200, "content":"%d"}`, dinfo.DiscussId))
				return
			} else {
				log.Debug("Bad DiscussInfo: ", dinfo)
			}
		} else {
			log.Error("json.Unmarshal ", err.Error())

		}
	}
	d.Ctx.WriteString(`{"status":404, "info":"xxx"}`)
}
