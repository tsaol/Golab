package main

import (
	"fmt"
    "net/http"
	"github.com/Shopify/sarama"
    "github.com/gorilla/mux"
    "github.com/spf13/viper"
)
type conf struct {
    Host  string `yaml:"host"`
    Topic string `yaml:"topic"`
    Port  string `yaml:"port"`
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
	config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll          // need leader, follow ack
    config.Producer.Partitioner = sarama.NewRandomPartitioner // choose a new partition
    config.Producer.Return.Successes = true                   // deiver message will retunr at success channel

    // a message
    msg := &sarama.ProducerMessage{}
    msg.Topic = "produdct_request"
    msg.Value = sarama.StringEncoder("this is a test log")

    // connect kafka
    client, err := sarama.NewSyncProducer([]string{"192.168.3.209:9082"}, config)
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
func (c *conf) getConf() *conf {
    yamlFile, err := ioutil.ReadFile("conf.yaml")
    if err != nil {
        fmt.Println(err.Error())
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        fmt.Println(err.Error())
    }
    return c
}
