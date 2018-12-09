package server

import (
	"MCDaemon-go/command"
	"MCDaemon-go/config"
	parser "MCDaemon-go/parsers"
	plugin "MCDaemon-go/plugins"
	"bufio"
	"io"
	"log"
	"os/exec"
	"strconv"
	"sync"
)

//单例模式
var (
	is_set bool
	svr    *Server
	err    error
)

type Server struct {
	Stdout        *bufio.Reader    //子进程输出
	Cmd           *exec.Cmd        //子进程实例
	stdin         io.WriteCloser   //用于关闭输入管道
	stdout        io.ReadCloser    //用于关闭输出管道
	lock          sync.Mutex       //输入管道同步锁
	pulginPool    chan interface{} //插件池
	maxRunPlugins int              //插件最大并发数
	SubServers    []*Server        //保存存档后的镜像（用于之后保存并开启镜像服务器的插件需要）
}

//单例模式
func init() {
	is_set = false
}

//获取实例
func GetServerInstance() *Server {
	if !is_set {
		svr = &Server{}
	}
	return svr
}

//根据参数初始化服务器
func (svr *Server) Init(argv []string) {
	//创建子进程实例
	svr.Cmd = exec.Command("java", argv...)
	svr.Cmd.Dir = "minecraft"
	svr.stdout, err = svr.Cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	//接管子进程输入输出
	svr.Stdout = bufio.NewReader(svr.stdout)
	svr.stdin, err = svr.Cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	//初始化插件执行池参数
	svr.maxRunPlugins, _ = strconv.Atoi(config.Cfg.Section("MCDeamon").Key("maxRunPlugins").String())
	svr.pulginPool = make(chan interface{}, svr.maxRunPlugins)
}

//运行子进程
func (svr *Server) run_process() {
	svr.Cmd.Start()
}

//运行所有语法解析器
func (svr *Server) RunParsers(word string) {
	for _, val := range parser.ParseList {
		cmd, ok := val.Parsing(word)
		if ok && plugin.PluginsList[cmd.Cmd] != nil {
			//异步运行插件
			svr.pulginPool <- 1
			go svr.RunPlugin(cmd)
		}
	}
}

//运行插件
func (svr *Server) RunPlugin(cmd *command.Command) {
	plugin.PluginsList[cmd.Cmd].Handle(cmd, svr)
	<-svr.pulginPool
}

//等待现有插件的完成并停止后面插件的运行，在执行相关操作
func (svr *Server) RunUniquePlugin(handle func()) {
	<-svr.pulginPool
	for i := 0; i < 10; i++ {
		svr.pulginPool <- 1
	}
	handle()
	for i := 0; i < 10; i++ {
		<-svr.pulginPool
	}
	svr.pulginPool <- 1
}

//重启服务器
func (svr *Server) Restart() {
	svr.Close()
	//获取所有启动项配置
	MCDeamon := config.GetStartConfig()
	//初始化
	svr.Init(MCDeamon)
	//等待加载地图
	svr.WaitEndLoading()
	//正式运行MCD
	svr.Run()
}

//关闭服务器
func (svr *Server) Close() {
	svr.Execute("/stop")
}
