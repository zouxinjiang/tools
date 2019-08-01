# clog
## Introduce
This component is a logger for go

## Install
```bash
go get github.com/zouxinjiang/clog
```
## Usage
1. internal format func

Name | Description | Example
--- | --- | ---
fn | the short function name | clog.TestCommonClog
FN | the long function name | github.com/zouxinjiang/clog.TestCommonClog
ln | the line number | 100
t | the 12 value time | 1970-01-01 01:11:11
T | the 24 value time | 1970-01-01 13:11:11
f | the short file name | clog_test.go
F | the long file name | D:/xxx/github.com/zouxinjiang/clog/clog_test.go
l | the log level | INFO
${data} | special sign for print data | -

the default log format is
```
[$l] $T file:$f line:$ln func:$fn ${data}

result example:
[WARNING] 2019-08-01 10:29:13 file:clog_test.go line:32 func:clog.TestCommonClog warning info
```

2. json format log
    - set the format is `{"fn":"$fn","line":$ln,"data":${data}}`
    - call object.SetDataFormat(FMT_Json)
    - result example `{"fn":"clog.TestClog","line":19,"data":[{"aaa":"ccc"}]}`
```go
    lg := Clog{
		w:      os.Stdout,
		level:  Lvl_Info | Lvl_Debug | Lvl_Error | Lvl_Warning,
		format: `{"fn":"$fn","line":$ln,"data":${data}}`,
	}
	lg.SetDataFormat(FMT_Json)
	lg.Error(map[string]string{"aaa": "ccc"})
	lg.Debug("aaa")
```
3. set print level
```go
lg.SetShowLevel(Lvl_Info | Lvl_Warning | Lvl_Error | Lvl_Debug)
```

4. add custom format function
    - add function
    - modify the format string
    - example `[INFO] 2019-08-01 10:45:05 file:clog_test.go line:45 func:clog.TestAddFunc custom:my func ssssssssssssssss`
```go
    lg := NewClog()
	lg.SetFormat(`[$l] $T file:$f line:$ln func:$fn custom:$my ${data}`)
	lg.AddCustomFormatFunc("my", func(level LogLevel, skip int) string {
		return "my func"
	})

	lg.Info("ssssssssssssssss")
```


    
    