package main

import (
	"context"

	"github.com/carousell/go-logging"
	"github.com/carousell/go-logging/gokit"
)

func main() {

	ctx := context.Background()

	logging.RegisterLogger(gokit.NewLogger())

	ctx = logging.WithLogger(ctx,
		logging.WithAttr("main1", "main1"),
		logging.WithAttr("main2", 123),
	)

	logging.InfoV2(ctx, "before calling foo()")
	//{"@timestamp":"2023-04-27T00:39:45.573743+08:00","caller":"go-logging/new.go:33","level":"info","main1":"main1","main2":123,"msg":"before calling foo()"}

	foo(ctx)
	logging.InfoV2(ctx, "after called foo()")
	//{"@timestamp":"2023-04-27T00:39:45.574079+08:00","caller":"go-logging/new.go:33","level":"info","main1":"main1","main2":123,"msg":"after called foo()"}
}

func foo(ctx context.Context) {
	ctx = logging.WithLogger(ctx,
		logging.WithAttr("foo1", "foo1"),
		logging.WithAttr("foo2", 456),
	)

	logging.InfoV2(ctx, "before calling bar()", logging.WithAttr("foo3", "foo3"))
	//{"@timestamp":"2023-04-27T00:39:45.574018+08:00","caller":"go-logging/new.go:33","level":"info","main1":"main1","main2":123,"foo1":"foo1","foo2":456,"foo3":"foo3","msg":"before calling bar()"}

	bar(ctx)

	logging.InfoV2(ctx, "after called bar()", logging.WithAttr("result", "bar result"))
	//{"@timestamp":"2023-04-27T00:39:45.574068+08:00","caller":"go-logging/new.go:33","level":"info","main1":"main1","main2":123,"foo1":"foo1","foo2":456,"result":"bar result","msg":"after called bar()"}
}

func bar(ctx context.Context) {
	ctx = logging.WithLogger(ctx,
		logging.WithAttr("bar1", "bar1"),
		logging.WithAttr("bar2", 789),
	)
	logging.InfoV2(ctx, "before do something", logging.WithAttr("bar3", 9876))
	//{"@timestamp":"2023-04-27T00:39:45.574037+08:00","caller":"go-logging/new.go:33","level":"info","main1":"main1","main2":123,"foo1":"foo1","foo2":456,"bar1":"bar1","bar2":789,"bar3":9876,"msg":"before do something"}

	logging.InfoV2(ctx, "after do something", logging.WithAttr("result", "success"))
	//{"@timestamp":"2023-04-27T00:39:45.574055+08:00","caller":"go-logging/new.go:33","level":"info","main1":"main1","main2":123,"foo1":"foo1","foo2":456,"bar1":"bar1","bar2":789,"result":"success","msg":"after do something"}
}
