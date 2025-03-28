package constants

// HandlerResult 处理结果 1. 未处理，2. 处理， 3. 拒绝 4.撤销
type HandlerResult int

const (
	NoHandlerResult     HandlerResult = iota + 1 // 未处理
	PassHandlerResult                            // 通过
	RefuseHandlerResult                          // 拒绝
	CancelHandlerResult
)
