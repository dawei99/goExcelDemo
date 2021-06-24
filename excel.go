package main
/**
 * 基于excelize的Excel操作
 */
import (
    "fmt"
    "github.com/360EntSecGroup-Skylar/excelize/v2"
    "github.com/tidwall/gjson"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"
    "os"
)
// 已处理任务个数
var (
    task_now int64
    total_number int64  // 任务总数
)
func main () {

    parseData := os.Args[1] // 解析json

    var excel_path string = gjson.Get(parseData, "file").String()   // 操作的Excel文件绝对路径

    file, err := excelize.OpenFile(excel_path) // 打开excel
    throwErr(err, "Excel文件不存在!") // 检测打开失败

    var firstSheetName string = file.GetSheetMap()[1] // 工作薄名称

    // 设置参数
    var (
        addImageParam AddImageParams // 插入图片需传递参数
        addTextParam AddTextParams // 插入文本需传递参数
        addSignatureParam addSignatureParams // 签章操作参数
    )
    addImageParam.fileHandle = file
    addTextParam.fileHandle = file

    addImageParam.firstSheetName = firstSheetName
    addTextParam.firstSheetName = firstSheetName

    addSignatureParam.firstSheetName = firstSheetName
    addSignatureParam.fileHandle = file

    // 任务总数
    total_number = gjson.Get(parseData, "images.#").Int() + gjson.Get(parseData, "text.#").Int()
    signaturePath := gjson.Get(parseData, "signature").String()// 签章图片
    // 添加签章任务
    //if signaturePath != "" {
    //    total_number++
    //}
    fmt.Println(total_number)

    //// 全部完成通知管道
    //finishNoticeHandle := make(chan int)

    // 插入图片操作
    gjson.Get(parseData, "images").ForEach(func (i, gResult gjson.Result) bool {
       addImageParam.path = gResult.Get("path").String()
       addImageParam.pos = gResult.Get("pos").String()
       addImageParam.height = gResult.Get("height").String()
       addImageParam.width =  gResult.Get("width").String()
       addImageParam.x = gResult.Get("x").Int()
       addImageParam.y = gResult.Get("y").Int()
       // 验证参数
       //if addImageParam.valid() {
           //go addImageHandle(addImageParam)
            addImageHandle(addImageParam)
       //}
       return true
    })

    // 插入图片操作
    gjson.Get(parseData, "text").ForEach(func (i, gResult gjson.Result) bool {
       addTextParam.text = gResult.Get("text").String()
       addTextParam.pos = gResult.Get("pos").String()
       addTextParam.cellStyle = gResult.Get("cellStyle").String()
       // 验证参数
       if addTextParam.valid() {
           addTextHandle(addTextParam)
       }
       return true
    })

    // 处理签章图片
    addSignatureParam.path = signaturePath
    //if (total_number > 0) {
    //    <-finishNoticeHandle // 等待任务完成
    //}

    if signaturePath != "" {
        addSignatureHandle(addSignatureParam)
    }

    // 保存图像
    errSave := file.Save()
    throwErr(errSave, "Excel文件保存失败!") // 检测打开失败

}
