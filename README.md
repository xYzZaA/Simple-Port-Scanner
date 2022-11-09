# 使用说明

```
  -a string                                                                 
        Enter the address you want to search (default "127.0.0.1")
        可以接收多个地址，输入方式有
          10.10.10.*
          10.10.10.10-15
          10.10.10.0/24
          10.10.10.1
  -m string
        Select scan mode, syn or connect, syn for linux systems only (default "connect")
        syn扫描只能在linux系统下使用
  -n int
        Enter an integer to control the number of goroutine (default 300)
        限制协程数量，尽量不要搞太多，数量越多扫描错误的可能性越大，500是极限了
        (本人不太会性能调优，性能浪费应该挺离谱的)
  -p string
        port
        输入扫描的端口，不指定就默认扫描留下的1000个
        可以输入一个也可以输入多个(0-1000)
```

# 演示说明

![](C:\Users\20477\AppData\Roaming\marktext\images\2022-11-09-21-57-38-C8M0A9MPDW6UAX9]THJSSUG.png)

不要在意背景
