package main

import (
	"fmt"
    "log"
    "net/http"
    "io/ioutil"
	"github.com/Shopify/sarama"
    "github.com/gorilla/mux"
    "gopkg.in/yaml.v2"
)
type Conf struct {
    host  string `yaml:"host"`
    port  string `yaml:"port"`
}


func main() {
    r := mux.NewRouter()
    r.HandleFunc("/product/{product_id}", productDetail).Methods("GET")
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, I am gateway: %s\n", r.URL.Path)
    })
    http.ListenAndServe(":8081", r)
}

func productDetail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "product id: is %v\n", vars["product_id"])
    SendMessage()
}


func SendMessage(){
    kfkconf := getConf("httpserver.yaml")
    
	config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll          // need leader, follow ack
    config.Producer.Partitioner = sarama.NewRandomPartitioner // choose a new partition
    config.Producer.Return.Successes = true                   // deiver message will retunr at success channel

    // a message
    msg := &sarama.ProducerMessage{}
    msg.Topic = "produdct_request"
    msg.Value = sarama.StringEncoder("this is a test log")
    
    kfkurl := kfkconf.host +":"+kfkconf.port
    fmt.Printf("kfk url is :%v\n", kfkurl)

    // connect kafka
    client, err := sarama.NewSyncProducer([]string{kfkurl}, config)
    if err != nil {
        fmt.Println("producer closed, err:", err)
        return
    }
    defer client.Close()
	
    // send message
    pid, offset, err := client.SendMessage(msg)
    if err != nil {
        fmt.Println("send msg failed, err:", err)
        return
    }
    fmt.Printf("pid:%v offset:%v\n", pid, offset)
}

//read config
//reference https://golangdocs.com/golang-yaml-package
func getConf(path string) (c *Conf) {
    yfile, err := ioutil.ReadFile(path)
    
    if err != nil {
         log.Fatal(err)
    }
    // confData := make(map[string]Conf)
    var confData Conf
    err2 := yaml.Unmarshal(yfile, confData)
    if err2 != nil {
         log.Fatal(err2)
    }
    return &confData
}
