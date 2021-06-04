#### 快速操作excel

    打包：go build -o exec *.go
    
    开发：go run *.go '{"file":"162130136036533.xlsx","images":[{"path":"sign.png","height":"70","pos":"B23","y":1,"x":1}],"text":[{"text":"asdfsafd","pos":"A1"}],"signature":"sign.png"}'
    生产：./exec '{"file":"test.xlsx","images":[{"path":"sign.png","height":"70","pos":"B23","y":1,"x":1}],"text":[{"text":"asdfsafd","pos":"A1"}]}'

参数说明：

    {
    	"file": "test.xlsx",
    	"images": [{
    		"path": "sign.png",
    		"height": "70",
    		"pos": "B23",
    		"y": 1,
    		"x": 1
    	}],
    	"text": [{
    		"text": "asdfsafd",
    		"pos": "A1"
    	}],
    	"signature":"sign.png"
    }
    
##### 效果图：
![avatar](C976DEA8-5C5F-43E2-B544-87D4210559C9.png)