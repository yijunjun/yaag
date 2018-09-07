package martiniyaag

import (
	"github.com/yijunjun/yaag/middleware"
	"github.com/yijunjun/yaag/yaag"
	"github.com/yijunjun/yaag/yaag/models"
	"github.com/go-martini/martini"
	"net/http"
)

func Document(c martini.Context, w http.ResponseWriter, r *http.Request) {
	if !yaag.IsOn() {
		c.Next()
		return
	}
	apiCall := models.ApiCall{}
	writer := middleware.NewResponseRecorder(w)
	c.MapTo(writer, (*http.ResponseWriter)(nil))
	middleware.Before(&apiCall, r)
	c.Next()
	middleware.After(&apiCall, writer,  r)
}
