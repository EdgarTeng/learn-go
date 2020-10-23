# compatibilities in go



## 1. Shift operation

```go
for i := 0; i < len(merged); i++ {
		var tmp = merged[i] >> ((i%8)+1)
		merged[i] = merged[i] ^ tmp ^ 17
}
```

**Issue:**

shift count type int, must be unsigned integer

**Fix:**

1) update go to latest(>=1.13)

2) convert to `uint` 

```go
var tmp = merged[i] >> uint((i%8)+1)
```



## 2. Time method

```go
duration := 5*time.Minute.Milliseconds()
```

**Issue:**

time.Minute.Milliseconds undefined (type time.Duration has no field or method Milliseconds)

note: module requires Go 1.13

**Fix:**

1) update go to latest(>=1.13)

2) remove time method `Milliseconds` instead of `5*60*1000`