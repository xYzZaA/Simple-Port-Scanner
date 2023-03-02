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
        限制协程数量，尽量不要搞太多，数量越多扫描错误的可能性越大，500最好
  -p string
        port
        输入扫描的端口，不指定就默认扫描留下的1000个
        可以输入一个也可以输入多个(0-1000)
  -oi 
        不需要给出参数，调用此参数停止端口扫描，单纯探测ip存活性
  -nobt
        不需要给出参数，用于取消ftp、mysql的弱口令爆破
  -bt ftp,mysql,ssh,redis
        给出需要爆破的协议，用逗号隔开，给出哪个爆破哪个，没有就爆破全部
        因为ftp没加时间控制，会默认跑完全部弱口令，时间会比较长
        目前只支持ftp mysql ssh redis
  -o
        不用给出参数，会在flag目录下创建一个result.txt，存放结果
  -dict string
        指定字典的路径，填写文件名，不要用result.txt作为字典名
  -ru   string
        高版本的redis支持多用户，默认用户名为default，如果有需要爆破的用户
        用此命令输入，支持多个输入，用逗号隔开
```

# 演示说明

![](C:\Users\20477\AppData\Roaming\marktext\images\2022-11-09-21-57-38-C8M0A9MPDW6UAX9]THJSSUG.png)
