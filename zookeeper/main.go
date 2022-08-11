package main

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"sync"
	"time"
)

var conn *zk.Conn

func init() {
	var err error
	conn, _, err = zk.Connect([]string{"localhost:2181", "localhost:2181", "localhost:2181"}, time.Second)
	if err != nil {
		panic(err)
	}
}

func main() {
	defer conn.Close()

	createNodes()

	tree("/hello", "-")

	watch("/hello")

	modify("/hello")

	deleteChildren("/hello")

	tree("/hello", "-")

	lock()

}

func lock() {
	var count = 0
	lock := zk.NewLock(conn, "/hello/lock2", zk.WorldACL(zk.PermAll))
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := lock.LockWithData([]byte("locked"))
			handleErr(err)
			defer func(lock *zk.Lock) {
				err := lock.Unlock()
				handleErr(err)
			}(lock)
			count++
		}()
	}
	wg.Wait()
	fmt.Println("count is", count)
}

func watch(path string) {
	go func() {
		data, _, ch, err := conn.GetW(path)
		if err != nil {
			panic(err)
		}
		fmt.Printf("The current data is: %v.\nI'm watching.\n", string(data[:]))
		for event := range ch {
			switch event.Type {
			case 1:
				fmt.Printf("Node created.")
				// Do somthing
			case 2:
				fmt.Printf("Node deleted!")
				// Do somthing
			case 3:
				fmt.Printf("Node changed.")
				// Do somthing
			case 4:
				fmt.Printf("Node children changed.")
				// Do somthing
			default:
				fmt.Printf("Something happen.")
			}
			// Or just ...
			fmt.Println("event type is ", event.Type.String())

		}
	}()

	modify(path)
}

func deleteChildren(path string) {
	children, _, _ := conn.Children(path)
	for _, child := range children {
		child = path + "/" + child
		exist, s, _ := conn.Exists(child)
		if exist {
			err := conn.Delete(child, s.Version)
			handleErr(err)
		}
	}
}

func modify(path string) {
	old, s, _ := conn.Get(path)
	ns, _ := conn.Set(path, []byte("girls"), s.Version)
	fmt.Printf("data %v -> %v; version %v -> %v \n", string(old), "girls", s.Version, ns.Version)
}

func tree(path string, bar string) {
	v, _, _ := conn.Get(path)

	fmt.Println(bar, path+"["+string(v)+"]")
	paths, _, _ := conn.Children(path)
	for _, p := range paths {
		tree(path+"/"+p, bar+"-")
	}
}

func createNodes() {
	// create normal node  持久 无序
	path, err := conn.Create("/hello", []byte("world"), 0, zk.WorldACL(zk.PermAll))
	handleErr(err)
	fmt.Println("created node path is", path)

	// create 临时节点
	path, _ = conn.Create("/ni_hao", []byte("shi_jie"), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))

	// 持久 时序
	for i := 0; i < 10; i++ {
		path = fmt.Sprintf("/hello/node%v", i)
		_, err = conn.Create(path, []byte("world"), zk.FlagSequence, zk.WorldACL(zk.PermAll))
		handleErr(err)
	}

	// 临时 时序
	for i := 0; i < 20; i++ {
		path = fmt.Sprintf("/EphemeralNode%v", i)
		_, err = conn.Create(path, []byte("shi_jie"), zk.FlagEphemeral|zk.FlagSequence, zk.WorldACL(zk.PermAll))
		handleErr(err)
	}
}

func handleErr(err error) {
	if err == zk.ErrNodeExists {
		fmt.Println(err)
	} else if err != nil {
		panic(err)
	}
}
