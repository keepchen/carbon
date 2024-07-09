# carbon  
[![Go](https://github.com/keepchen/go-sail/actions/workflows/go.yml/badge.svg)](https://github.com/keepchen/go-sail/actions/workflows/go.yml)  [![CodeQL](https://github.com/keepchen/go-sail/actions/workflows/codeql.yml/badge.svg)](https://github.com/keepchen/go-sail/actions/workflows/codeql.yml)  [![Go Report Card](https://goreportcard.com/badge/github.com/keepchen/carbon)](https://goreportcard.com/report/github.com/keepchen/carbon)    
An easy-use datetime toolkit written in Go.

#### Requirement  
> go version >= 1.18

#### Installation  
> go get -u github.com/keepchen/carbon

#### Features  
- [x] Calculate diff  
- [x] Transform to points  
- [x] Easy format  
- [x] Absolute  
- [x] Date time nth  

#### Examples  
- New  
```go
carbon.NewFromString("2006-01-02 15:04:05", "2024-10:10 12:00:00")

carbon.NewFromTime(time.Now())

carbon.NewFromSeconds(time.Now().Unix())

carbon.NewFromMilliSeconds(tim.UnixMilli())
```  
- Diff
```go
c := carbon.NewFromString("2006-01-02 15:04:05", "2024-10:10 12:00:00")
u := carbon.NewFromString("2006-01-02 15:04:05", "2025-10:10 12:00:00")
c.Diff(u).Day().Int64()
c.Diff(u).Day().Float64()
c.Diff(u).Day().Float64(3)

c.Diff(u).Minute().Int64()
c.Diff(u).Minute().Float64()
c.Diff(u).Minute().Float64(3)

c.Diff(u).Hour().Int64()
c.Diff(u).Hour().Float64()
c.Diff(u).Hour().Float64(3)
```  
- Format  
```go
// layout is case insensitive

carbon.NewFromTime(time.Now()).Format("y-m-d h:i:s")

//or 
carbon.NewFromTime(time.Now()).Format("Y-M-D H:I:S")

//or 
carbon.NewFromTime(time.Now()).Format("Y-m-D H:i:S")
```  
- Nth  
```go
carbon.NewFromTime(time.Now()).WeekNth()
```  
- Absolute  
```go
c := carbon.NewFromString("2006-01-02 15:04:05", "2024-10:10 12:00:00")
u := carbon.NewFromString("2006-01-02 15:04:05", "2025-10:10 12:00:00")
c.Diff(u).Day().Abs().Int64()
```

