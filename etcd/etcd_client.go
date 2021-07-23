package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

var (
	config clientv3.Config
	client *clientv3.Client
	err    error
	kv     clientv3.KV
	lease  clientv3.Lease
)

func init() {
	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"168.138.69.26:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 初始化实例
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)
}

// Set 设值
func Set(k string, x string) error {
	_, err := kv.Put(context.TODO(), k, x)
	if err != nil {
		return fmt.Errorf("set etcd value error")
	}
	return nil
}

// Get 取值
func Get(k string) (value string, found bool) {
	resp, err := kv.Get(context.TODO(), k)
	if err != nil {
		return "", false
	}
	if resp.Kvs == nil {
		return "", false
	}
	return string(resp.Kvs[0].Value), true
}

// SetExpire 设置值和过期时间（s）
func SetExpire(k string, x string, seconds int64) error {
	grantResp, _ := lease.Grant(context.TODO(), seconds)
	_, err := kv.Put(context.TODO(), k, x, clientv3.WithLease(grantResp.ID))
	if err != nil {
		return fmt.Errorf("set etcd value error")
	}
	return nil
}
