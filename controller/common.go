package controller

type CommonController struct {
	BaseController
}

func (c *CommonController) Get() {
	c.WriteJSON(struct {
		Status int    `json:"status"`
		Info   string `json:"info"`
	}{
		Status: -1,
		Info:   "not found",
	})
}
