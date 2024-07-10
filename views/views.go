package views

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paulormbsd/hadggar/controllers"
	"github.com/paulormbsd/hadggar/middleware"
	"github.com/paulormbsd/hadggar/pkg/zabbix"
)

//var templates = template.Must(template.ParseGlob("templates/*.html"))

// Função para manipular as rotas das páginas, no final ele inicia o servidor na porta 8000
func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/node/zabbix", zabbix.GetZabbixAPI).Methods("Get")
	r.HandleFunc("/node/zabbix/get", zabbix.GetZabbixHost).Methods("Get")
	r.HandleFunc("/node/zabbix/{id}", zabbix.GetZabbixHostByHost).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
