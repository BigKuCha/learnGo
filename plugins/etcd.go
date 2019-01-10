package plugins

import (
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/embed"
	"log"
)

func RunEtcd() {
	cfg := embed.NewConfig()
	cfg.Dir = "default.etcd"
	config := clientv3.Config{
		Endpoints: []string{"http://localhost:2379"},
	}
	cli, err := clientv3.New(config)
	if err != nil {
		log.Fatalf("client err %+v", err)
	}
	_ = cli
	//ctx := cli.Ctx()
	//fmt.Println(ctx.Deadline())
	//res, err := cli.Get(ctx, "foo")
	//res.OpResponse()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Println(res.OpResponse().Get())
}
