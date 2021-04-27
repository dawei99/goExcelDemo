#### 向Excel批量添加图片
    
    go run excel.go -f /home/rjh/www/angular/go/test.xlsx  -i '[{"path":"/home/rjh/www/angular/go/sign.png","height":"45","pos":"B20"},{"path":"/home/rjh/www/angular/go/sign.png","height":"45","pos":"B22"}]'

效果图：
![avatar](C976DEA8-5C5F-43E2-B544-87D4210559C9.png)

参数：

    -f   excel文件路径
    
    -i   要插入图片信息json

    [{
	    "path": "[图片路径]",
        "pos": "[插入位置，例如：E4]",
        "height": "[设置单元格高度（非必填）]",
        "width": [设置单元格高度（非必填）],
    },...]