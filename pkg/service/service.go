package service

import (
	"github.com/god-jason/bucket/config"
	"github.com/kardianos/service"
	"log"
)

var svc service.Service
var logger service.Logger

func Register(startup, shutdown func()) (err error) {
	var serviceConfig = &service.Config{
		Name:        config.GetString(MODULE, "name"),
		DisplayName: config.GetString(MODULE, "display"),
		Description: config.GetString(MODULE, "description"),
		Arguments:   config.GetStringSlice(MODULE, "arguments"),
	}

	p := &Program{
		Startup:  startup,
		Shutdown: shutdown,
	}

	svc, err = service.New(p, serviceConfig)
	if err != nil {
		return err
	}

	logger, err = svc.Logger(nil)
	if err != nil {
		return err
	}

	return nil
}

func Run() error {
	return svc.Run()
}

func Start() error {
	return svc.Start()
}

func Restart() error {
	return svc.Restart()
}

func Install() error {
	return svc.Install()
}

func Uninstall() error {
	return svc.Install()
}

func Error(v ...interface{}) {
	if logger != nil {
		err := logger.Error(v...)
		if err != nil {
			log.Println(v...)
		}
	} else {
		log.Println(v...)
	}
}

func Warn(v ...interface{}) {
	if logger != nil {
		err := logger.Warning(v...)
		if err != nil {
			log.Println(v...)
		}
	} else {
		log.Println(v...)
	}
}

func Info(v ...interface{}) {
	if logger != nil {
		err := logger.Info(v...)
		if err != nil {
			log.Println(v...)
		}
	} else {
		log.Println(v...)
	}
}

type Program struct {
	Startup  func()
	Shutdown func()
}

func (p *Program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	p.Shutdown()
	return nil
}

func (p *Program) run() {
	//
	//// 此处编写具体的服务代码
	//hup := make(chan os.Signal, 2)
	//signal.Notify(hup, syscall.SIGHUP)
	//quit := make(chan os.Signal, 2)
	//signal.Notify(quit, os.Interrupt, os.Kill)
	//
	//go func() {
	//	for {
	//		select {
	//		case <-hup:
	//		case <-quit:
	//			//_ = p.Shutdown() //全关闭两次
	//			//os.Exit(0)
	//		}
	//	}
	//}()

	//内部启动
	p.Startup()
}
