package server

import "github.com/gin-gonic/gin"

type Server struct {
	*gin.Engine
	opt *Option
}

type Option struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}

func New(opt *Option) *Server {
	gin.SetMode(opt.Env)
	return &Server{
		gin.New(),
		opt,
	}
}

// Run 启动一个应用程序
func (s *Server) Run() error {
	err := s.Engine.Run(s.opt.Host + ":" + s.opt.Port)
	if err != nil {
		return err
	}

	return err
}
