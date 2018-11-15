```
func GetConfigEtcd(unique_name string) (string, error) {
	dialTimeout := beego.AppConfig.DefaultInt("etcd::etcd_dialtimeout", 5)
	endPoint := beego.AppConfig.String("etcd::etcd_host")
	endPoints := []string{endPoint}

	var dataType string

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endPoints,
		DialTimeout: time.Duration(dialTimeout) * time.Second,
	})

	if err != nil {
		beego.Warn("connect ETCD error, the error:", err)
		return "", err
	}

	defer cli.Close()

	resp, err := cli.Get(context.TODO(), unique_name)
	if err != nil {
		beego.Warn("get config from ETCD error, the error:", err)
		return "", err
	}

	for _, ev := range resp.Kvs {
		dataType = string(ev.Value)
	}

	return dataType, nil
}

```
