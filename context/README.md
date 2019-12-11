# Context
context 包主要是用来简化处理单个请求的多个 gorouting 之间与请求域的数据、取消信号、截止时间等相关操作。


### 接口定义
```go
type Context interface {
    
    // 获取设置的截止时间，如果未设置截止时间，ok == false，需调用取消函数取消；当到达截止时间，Context 会自动发起取消请求。
    Deadline() (deadline time.Time, ok bool) 
    
    // 返回一个只读的 chan， 类型为 struct {}; 在 gorouting 中，如果该方法返回的 chan 可以读取，就意味着 parent context 发起了取消请求，应该进行清理操作。
    Done() <-chan struct{}
  
    // 返回取消的错误原因，Context 因为什么被取消。
    Err() error
    
    // 获取该 Context 上绑定的值，是一个键值对；线程安全。
    Value(key interface{}) interface{}
}
```

###  基本实现类型
context 包定义 emptyCtx 类型，并实现了 Context 接口的方法。

```go
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}
```

并通过 new 方法创建两个类型。
```go
var (
  background = new(emptyCtx)
  todo = new(emptyCtx)
)

func Background() Context() {
  return background
}

func TODO() Context() {
  return todo
}
```

### Context 继承
Context 通过下面四个方法实现创建子 Context，[源码实现](https://tip.golang.org/src/context/context.go)。
```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

func WithValue(parent Context, key, val interface{}) Context
```

### 使用技巧
* 常用 Background() 进行创建
* Context 不建议放在结构体中，要以参数的方式传递
* 建议把 Context 作为第一个参数进行，变量名统一命名为 ctx
* Context 是线程安全的
* 对某一 Context 执行取消操作，其他 goroutine 都将会收到取消信号
