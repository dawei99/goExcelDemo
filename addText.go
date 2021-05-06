package main

import (
    "fmt"
)
// addText函数参数
type AddTextParams struct {
    ExcelInfoStruct
    text string
    pos string
    cellStyle string
}
func (a *AddTextParams) valid () bool {
    if len(a.text) == 0 || len(a.pos) == 0 {
        return false
    }
    return true
}
// excel添加文本
func addTextHandle(info AddTextParams, finishNotice chan int) {
    defer (func() {
        err := recover()
        if err != nil {
            fmt.Println(err)
        }
        task_now++
        if task_now == total_number {
            finishNotice<-1
        }
    })()

    var (
       fileHandle = info.fileHandle  // excel操作句柄
       firstSheetName = info.firstSheetName  // 工作薄名称

       text = info.text  // 宽度
       pos = info.pos  // 位置
       cellStyle = info.cellStyle // 高度
    )

    fileHandle.SetCellValue(firstSheetName, pos, text)
    style, err := fileHandle.NewStyle(cellStyle)
    if err == nil {
       err = fileHandle.SetCellStyle(firstSheetName, pos, pos, style)
    }

}