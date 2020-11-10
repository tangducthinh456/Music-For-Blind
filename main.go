package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"fmt"
	"os/exec"
	"strings"
)
//"\"C:\\Program Files (x86)\\AnthemScore\\AnthemScore.exe\""
const AudioSavePath = "C:\\Users\\Dell\\Desktop\\musictranscription\\Music-Transcription\\Audio\\"
const ExecuteProgramPath = "C:\\Program Files (x86)\\AnthemScore\\AnthemScore.exe"
const SaveXMLPath = "C:\\Users\\Dell\\Desktop\\musictranscription\\Music-Transcription\\MusicXML\\"
func main(){
	fmt.Println(ExecuteProgramPath)
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, err := c.FormFile("file")
		if err != nil{
			c.String(http.StatusInternalServerError, fmt.Sprintf("error : %s", err))
			return
		}
		start := c.PostForm("start")
		end := c.PostForm("end")

		log.Println(file.Filename, start, end)
		
		// Upload the file to specific dst.
		err = c.SaveUploadedFile(file, AudioSavePath + file.Filename)
        if err != nil{
			c.String(http.StatusInternalServerError, fmt.Sprintf("error : %s", err))
			return
		}
		filename := strings.Split(file.Filename, ".")
		//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		param := []string{}
		param = append(param, AudioSavePath + file.Filename)
		param = append(param, AudioSavePath + file.Filename)
		param = append(param, "-a")
		param = append(param, "-x" + SaveXMLPath + filename[0] + ".xml")
		if start != "" && end != ""{
			param = append(param, "-s " + start)
			param = append(param, "-e " + end)
		}
		//cmd := exec.Command(ExecuteProgramPath, AudioSavePath + file.Filename,"-a" ,"-x" + SaveXMLPath + filename[0] + ".xml")
		cmd := exec.Command(ExecuteProgramPath, param...)
		err = cmd.Run()
		if err != nil{
			c.String(http.StatusInternalServerError, fmt.Sprintf("error : %s", err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' success!", file.Filename))
	})
	router.Run(":8080")
}