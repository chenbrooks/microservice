package upload

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"io"
	"log"
	"microservice/utils"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func SaveImageURL(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := utils.GetUrlParam("image_url", r)
	file := downURLFile(url)
	fileUrl := upload2alioss(file)
	fmt.Fprint(w, utils.RespOK("图片上传成功", fileUrl))

}

func upload2alioss(localFile *os.File) string {

	sourceFileName := localFile.Name()
	aliyunFileName := strings.Replace(localFile.Name(), "/tmp/", "images/", 1)

	ossClient, err := oss.New(viper.GetString("alioss.bucket_domain"), viper.GetString("alioss.secret.key"), viper.GetString("alioss.secret.value"))
	utils.CheckErr(err)
	bucket, err := ossClient.Bucket(viper.GetString("alioss.bucket_name"))
	utils.CheckErr(err)

	err = bucket.PutObjectFromFile(aliyunFileName, sourceFileName)
	//bucket.PutObject(aliyunFileName, localFile)
	utils.CheckErr(err)
	return getAliossFileURL(aliyunFileName)
}

func getAliossFileURL(filePath string) string {
	return "https://" + viper.GetString("alioss.bucket_name") + "." + viper.GetString("alioss.bucket_domain") + "/" + filePath
}

func downURLFile(url string) *os.File {
	fmt.Println("start to download url file")
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	defer res.Body.Close()
	fileName := getFileName(path.Ext(url))
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	file.Close()
	fmt.Println("downfile done")
	return file
}

func getFileName(url string) string {
	layout := "200601/02"
	dir := "/tmp/" + (time.Now().Format(layout))
	os.MkdirAll(dir, 0755)
	return dir + "/" + getMD5(url+"salt_string") + path.Ext(url)
}
func getMD5(imageID string) string {
	timeInt := imageID + strconv.Itoa(time.Now().Nanosecond())
	h := md5.New()
	h.Write([]byte(timeInt))
	return hex.EncodeToString(h.Sum(nil))
}
