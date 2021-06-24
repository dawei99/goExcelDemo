package main

// addImage函数参数
type addSignatureParams struct {
    ExcelInfoStruct
    path string
}

// excel添加图片操作
func addSignatureHandle(info addSignatureParams) {
    //defer (func() {
    //    err := recover()
    //    if err != nil {
    //        fmt.Println(err)
    //    }
    //    task_now++
    //    if task_now == total_number {
    //        finishNotice<-1
    //    }
    //})()

    var (
        path = info.path // 图片地址
        firstSheetName = info.firstSheetName  // 工作薄名称
        fileHandle = info.fileHandle  // excel操作句柄
    )

    // 查找空列
    cols, err := fileHandle.Cols(firstSheetName)
    if err != nil {
        panic("签章添加失败，" + err.Error())
    }
    colTotals := 0 // 总列数
    for cols.Next() {
        colTotals++
    }
    pos := num2Abc(colTotals) + "3" // 位置

    // 添加图片
    if errAddPic := fileHandle.AddPicture(firstSheetName, pos, path, `{ "x_scale": 1,"y_scale": 1,"positioning": "absolute","x_offset": 1,"y_offset": 1}`) ; errAddPic != nil {
       panic("签章添加失败，"+"path："+path+"，pos："+pos+"，错误："+errAddPic.Error())
    }

}

