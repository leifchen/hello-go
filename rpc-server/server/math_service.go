package server

// MathService 服务
type MathService struct {
}

// Args 参数
type Args struct {
	A, B int
}

// Add 加法
func (m *MathService) Add(args Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}
