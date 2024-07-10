package zabbix

import (
	"net/http"
)

func GetZabbixAPI(w http.ResponseWriter, r *http.Request) {
	//user := "Admin"
	//pass := "zabbix"
	//api := zabbix.NewAPI("http://192.168.10.105/api_jsonrpc.php")
	//api.Login(user, pass)

	//res, err := api.Version()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Fprint(w, "Connected to zabbix api v", res)
}

func GetZabbixHost(w http.ResponseWriter, r *http.Request) {
	// user := "Admin"
	// pass := "zabbix"
	// host := zabbix.NewAPI("http://192.168.10.105/api_jsonrpc.php")
	// host.Login(user, pass)
	// res, err := host.HostsGet(zabbix.Params{"output": "extend"})
	// if err != nil {
	// 	panic(err)
	// }
	// for _, server := range res {
	// 	fmt.Fprint(w, "Name:", server.Name, " ID:", server.HostID)
	// }
}

func GetZabbixHostByHost(w http.ResponseWriter, r *http.Request) {
	// user := "Admin"
	// pass := "zabbix"
	// host := zabbix.NewAPI("http://192.168.10.105/api_jsonrpc.php")
	// host.Login(user, pass)
	// vars := mux.Vars(r)
	// id := vars["id"]
	// res, err := host.HostGetByHost(id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Fprint(w, "Name:", res)
}
