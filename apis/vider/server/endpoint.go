package server

import (
	"io"
	"log"
	"net/http"
)

//POST 新建一个课程 Params title desc
//GET 课程详情 Params
//PUT 修改课程信息
//DELETE 删除课程
type PostCourseJson struct {
	Title string `json:"title" valid:"required"`
	Desc  string `json:"desc" valid:"required"`
}

func (api *VideoApi) PostCourse(w http.ResponseWriter, reader io.Reader) {
	jsonData := PostCourseJson{}
	err := ParseJson(reader, &jsonData)
	if err != nil {
		log.Printf("ParseJson Err:%v", err)
		JsonResp(w, ErrInterResp("服务器错误", err))
		return
	}
	JsonResp(w, FineResp("新建成功", nil))
}

func (api *VideoApi) DeleteCourse(w http.ResponseWriter, r *http.Request) {

}
func (api *VideoApi) PutCourse(w http.ResponseWriter, r *http.Request) {

}
func (api *VideoApi) GetCourse(w http.ResponseWriter, r *http.Request) {

}
