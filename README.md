#### 向Excel批量添加图片
    
    测试命令：go run excel.go -f $PWD/test.xlsx  -i '[{"path":"/home/rjh/www/angular/go/src/sign.png","height":"70","pos":"B23","y":1,"x":1},{"path":"/home/rjh/www/angular/go/src/sign.png","height":"70","pos":"B23","y":45,"x":1}]'
    打包命令：go build  -o dist/exec  excel.go

效果图：
![avatar](src/C976DEA8-5C5F-43E2-B544-87D4210559C9.png)

参数：

    -f   excel文件路径
    
    -i   要插入图片信息json

    [{
	    "path": "[图片路径]",
        "pos": "[插入位置，例如：E4]",
        "height": "[设置单元格高度（非必填）]",
        "width": [设置单元格高度（非必填）],
        "x" : [单元格左内边距],
        "y" : [单元格上内边距],
    },...]