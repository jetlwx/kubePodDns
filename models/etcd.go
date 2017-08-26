package models

import (
	"context"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/coreos/etcd/client"
)

type KeyValue struct {
	Key   string
	Value string
}

//etcd new client
func NewCli(conn []string) client.Client {

	cfg := client.Config{
		//Endpoints: []string{conn},
		Endpoints: conn,
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: 3 * time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		log.Println("error at create new etcd client:", err)
		return nil
	}
	return c
}

//etcd new key opertion api
func NewKeyAPI(c client.Client) client.KeysAPI {
	kapi := client.NewKeysAPI(c)
	return kapi
}

//create a directory
func CreateDirectory(kapi client.KeysAPI, dirname string) (*client.Response, error) {
	o := client.SetOptions{Dir: true}
	resp, err := kapi.Set(context.Background(), dirname, "", &o)
	return resp, err
}

//list an directory
func ListDirector(kapi client.KeysAPI, dirname string) (kv []KeyValue, err error) {
	resp, err := kapi.Get(context.Background(), dirname, nil)
	if err != nil {
		return nil, err
	}
	// print directory keys
	sort.Sort(resp.Node.Nodes)
	for _, n := range resp.Node.Nodes {
		K := KeyValue{}
		K.Key = n.Key
		K.Value = n.Value
		kv = append(kv, K)

	}
	return kv, nil
}

//add an key
func SetKey(kapi client.KeysAPI, key, value string) (*client.Response, error) {
	resp, err := kapi.Set(context.Background(), key, value, nil)
	return resp, err
}

//get key
func GetKey(kapi client.KeysAPI, key string) (value string, e error) {
	// Get key "/foo"
	resp, err := kapi.Get(context.Background(), key, nil)
	return resp.Node.Value, err
}

//delete key
func DelKey(kapi client.KeysAPI, key string) (*client.Response, error) {
	resp, err := kapi.Delete(context.Background(), key, nil)
	return resp, err
}

/*
	ErrorCodeKeyNotFound  = 100
	ErrorCodeTestFailed   = 101
	ErrorCodeNotFile      = 102
	ErrorCodeNotDir       = 104
	ErrorCodeNodeExist    = 105
	ErrorCodeRootROnly    = 107
	ErrorCodeDirNotEmpty  = 108
	ErrorCodeUnauthorized = 110

	ErrorCodePrevValueRequired = 201
	ErrorCodeTTLNaN            = 202
	ErrorCodeIndexNaN          = 203
	ErrorCodeInvalidField      = 209
	ErrorCodeInvalidForm       = 210

	ErrorCodeRaftInternal = 300
	ErrorCodeLeaderElect  = 301

	ErrorCodeWatcherCleared    = 400
	ErrorCodeEventIndexCleared = 401
*/
func EtcdErrorType(e error) int {
	es := e.Error()
	str := strings.Split(es, ":")
	if len(str) >= 2 {
		no := str[0]
		log.Println("no=", no)
		switch no {
		// key not found
		case "100", "102", "104", "105", "107", "108", "110":
			return 100
		}
	}

	return 0
}
