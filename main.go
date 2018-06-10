package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/tidwall/gjson"
	"crypto/sha1"
	"crypto/hmac"
	"strings"
	"os/exec"

)

const secret="itfanr.cc"


func sha1Str(data string) string {
	h:=sha1.New()
	h.Write([]byte(data))
	my_sign:=h.Sum(nil)
	// 使用`%x`来将散列结果格式化为16进制的字符串
	return fmt.Sprintf("%x",my_sign)
}

func hmacsha1Str(key string,data string) string {
	h_key:=[]byte(key)
	mac:=hmac.New(sha1.New,h_key)
	mac.Write([]byte(data))
	my_mac:=mac.Sum(nil)
	// 使用`%x`来将散列结果格式化为16进制的字符串
	return fmt.Sprintf("%x",my_mac)
}


func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// 测试请求是否正常访问
	router.GET("/", func(c *gin.Context){
		c.String(http.StatusOK, "Wellcome to use GitHub Webhook Service!")
	})

	// 接收webhook请求
	router.POST("/deploy", func(c *gin.Context) {
		// 获取请求头信息
		x_Event:= c.Request.Header.Get("X-GitHub-Event")
		x_Signature:=c.Request.Header.Get("X-Hub-Signature")//sha1=e3d6a3c570a74df03741db3a048c81b46867a9db
		x_UserAgent:=c.Request.Header.Get("User-Agent") //GitHub-Hookshot/70e2f11
		x_Delivery:=c.Request.Header.Get("X-GitHub-Delivery") //81e4c220-6c01-11e8-97dd-0e2035b75e57

		fmt.Println(x_Event)
		fmt.Println(x_UserAgent)
		fmt.Println(x_Delivery)
		fmt.Println(x_Signature)

		// 获取sign
		x_sign_str_arr:=strings.Split(x_Signature,"=")
		if len(x_sign_str_arr)==2{
			x_Signature=x_sign_str_arr[1]
		}
		fmt.Println(x_Signature)

		// 获取Body
		req_body, _ := ioutil.ReadAll(c.Request.Body)
		req_str:=string(req_body)

		// 生成本地 Sign
		my_sign:=hmacsha1Str(secret,req_str)
		fmt.Println(my_sign)

		// 判断sign是否一致
		if x_Signature == my_sign{
			// sign验证成功，进行操作

			// 只对 Event =push 进行操作
			if x_Event == "push" {
				fmt.Println("**********")
				fmt.Println(req_str)

				// 从 body中获取参数
				cmt_ref:=gjson.Get(req_str,"ref")// 分支
				cmt_id:=gjson.Get(req_str,"commits.id")
				cmt_message:=gjson.Get(req_str,"commits.message")
				cmt_timestamp:=gjson.Get(req_str,"commits.timestamp")
				cmt_added_list:=gjson.Get(req_str,"commits.added")
				cmt_removed_list:=gjson.Get(req_str,"commits.removed")
				cmt_modified_list:=gjson.Get(req_str,"commits.modified")

				fmt.Println(cmt_ref.String())
				fmt.Println(cmt_id.String())
				fmt.Println(cmt_message.String())
				fmt.Println(cmt_timestamp.Time())
				fmt.Println(cmt_added_list.Array())
				fmt.Println(cmt_removed_list.Array())
				fmt.Println(cmt_modified_list.Array())

				// 检查是否为特定分支


				// 执行clone部署操作
				exec.Command("sh","/app/build.sh")

				c.JSON(http.StatusOK, "Success")
			}else{
				c.JSON(http.StatusOK, "No Support Event")
			}
		}else{
			// sign验证失败，返回错误
			c.JSON(http.StatusOK, "Sign Error")
		}
	})
	
	router.Run(":8080")
}