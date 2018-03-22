package engine
// 程序的设计准则，自顶向下， 可以先定义Scheduler的接口，再去实现
type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface{
	Submit(Request)
	ConfigChannelMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request){
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigChannelMasterWorkerChan(in)
	for i := 0; i < e.WorkerCount; i++{
		createWorker(in ,out)
	}
	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}
	for{
		result := <-out
		for _, request := range result.Requests{
			e.Scheduler.Submit(request)
		}

		for _, item := range result.Items{
			println("Got item:", item.(string))
		}
	}
}


func createWorker(in chan Request, out chan ParseResult){
	go func() {
		for{
			r := <-in
			parseResult,err := Worker(r)
			if err != nil{
				continue
			}
			out <- parseResult
		}
	}()
}
