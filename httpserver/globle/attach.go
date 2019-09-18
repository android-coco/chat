package globle

import (
	"chat/model"
	"chat/util"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {
	err := os.MkdirAll("./mnt", os.ModePerm)
	if err != nil {
		util.Logger.Errorf("[resp ok] %+v", err)
	}
}

func Upload(c *gin.Context) {
	UploadLocal(c)
}

//1.存储位置 ./mnt,需要确保已经创建好
//2.url格式 /mnt/xxxx.png  需要确保网络能访问/mnt/
func UploadLocal(c *gin.Context) {
	//todo 获得上传的源文件s
	srcFile, head, err := c.Request.FormFile("file")
	if err != nil {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusBadRequest,
			ErrorMsg: "parameter violation!",
		}, err)
		return
	}

	//todo 创建一个新文件d
	suffix := ".png"
	//如果前端文件名称包含后缀 xx.xx.png
	ofileName := head.Filename
	tmp := strings.Split(ofileName, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	//如果前端指定filetype
	//formdata.append("filetype",".png")
	fileType := c.Request.FormValue("filetype")
	if len(fileType) > 0 {
		suffix = fileType
	}
	//time.Now().Unix()
	filename := fmt.Sprintf("%d%04d%s",
		time.Now().Unix(), rand.Int31(),
		suffix)
	dstFile, err := os.Create("./mnt/" + filename)
	if err != nil {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusInternalServerError,
			ErrorMsg: err.Error(),
		}, err)
		return
	}

	// 将源文件内容copy到新文件
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusInternalServerError,
			ErrorMsg: err.Error(),
		}, err)
		return
	}
	// 将新文件路径转换成url地址

	url := "/mnt/" + filename
	// 响应到前端
	util.RespOK(c, model.ApiResp{
		ErrorNo:  util.SuccessCode,
		ErrorMsg: "",
		Data:     url,
	})
}

//即将删掉,定期更新
const (
	AccessKeyId="5p2RZKnrUanMuQw9"
	AccessKeySecret="bsNmjU8Au08axedV40TRPCS5XIFAkK"
	EndPoint="oss-cn-shenzhen.aliyuncs.com"
	Bucket="winliondev"
)

//权限设置为公共读状态
//需要安装
func UploadOss(c *gin.Context){
	//todo 获得上传的文件
	srcfile,head,err:=c.Request.FormFile("file")
	if err!=nil{
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusBadRequest,
			ErrorMsg: "parameter violation!",
		}, err)
		return
	}


	//todo 获得文件后缀.png/.mp3

	suffix := ".png"
	//如果前端文件名称包含后缀 xx.xx.png
	ofilename := head.Filename
	tmp := strings.Split(ofilename,".")
	if len(tmp)>1{
		suffix = "."+tmp[len(tmp)-1]
	}
	//如果前端指定filetype
	//formdata.append("filetype",".png")
	filetype := c.Request.FormValue("filetype")
	if len(filetype)>0{
		suffix = filetype
	}

	// 初始化ossclient
	client,err:=oss.New(EndPoint,AccessKeyId,AccessKeySecret)
	if err!=nil{
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusInternalServerError,
			ErrorMsg: err.Error(),
		}, err)
		return
	}
	// 获得bucket
	bucket,err := client.Bucket(Bucket)
	if err!=nil{
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusInternalServerError,
			ErrorMsg: err.Error(),
		}, err)
		return
	}
	// 设置文件名称
	//time.Now().Unix()
	filename := fmt.Sprintf("mnt/%d%04d%s",
		time.Now().Unix(), rand.Int31(),
		suffix)
	// 通过bucket上传
	err=bucket.PutObject(filename,srcfile)
	if err!=nil{
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusInternalServerError,
			ErrorMsg: err.Error(),
		}, err)
		return
	}
	// 获得url地址
	url := "http://"+Bucket+"."+EndPoint+"/"+filename

	// 响应到前端
	util.RespOK(c, model.ApiResp{
		ErrorNo:  util.SuccessCode,
		ErrorMsg: "",
		Data:     url,
	})
}