package models

// import (
// 	"context"
// 	"time"

// 	"log"
// 	"strings"

// 	"github.com/coreos/etcd/client"
// )

// //conn like  http://127.0.0.1:2379
// func EtcdInit(conn string) (client.KeysAPI, error) {
// 	cfg := client.Config{
// 		Endpoints: []string{conn},
// 		Transport: client.DefaultTransport,
// 		// set timeout per request to fail fast when the target endpoint is unavailable
// 		HeaderTimeoutPerRequest: 3 * time.Second,
// 	}

// 	c, err := client.New(cfg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	kapi := client.NewKeysAPI(c)
// 	return kapi, err
// }

// func EtcdSet(kapi client.KeysAPI, k, v string) (*client.Response, error) {
// 	resp, err := kapi.Set(context.Background(), k, v, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// // print common key info
// 	// log.Printf("Set is done. Metadata is %q\n", resp)

// 	return resp, err
// }

// //get key's value
// func EtcdGet(kapi client.KeysAPI, k string) (string, error) {
// 	resp, err := kapi.Get(context.Background(), k, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	return resp.Node.Value, err
// }

// //update key's,value
// func EtcdUpdate(kapi client.KeysAPI, k, v string) (*client.Response, error) {
// 	resp, err := kapi.Update(context.Background(), k, v)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp, err
// }

// //delete key
// func EtcdDelete(kapi client.KeysAPI, k string) (*client.Response, error) {
// 	resp, err := kapi.Delete(context.Background(), k, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp, err
// }

// func EtcdCreate(kapi client.KeysAPI, k, v string) (*client.Response, error) {
// 	resp, err := kapi.Create(context.Background(), k, v)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp, err
// }

// /*
// 	ErrorCodeKeyNotFound  = 100
// 	ErrorCodeTestFailed   = 101
// 	ErrorCodeNotFile      = 102
// 	ErrorCodeNotDir       = 104
// 	ErrorCodeNodeExist    = 105
// 	ErrorCodeRootROnly    = 107
// 	ErrorCodeDirNotEmpty  = 108
// 	ErrorCodeUnauthorized = 110

// 	ErrorCodePrevValueRequired = 201
// 	ErrorCodeTTLNaN            = 202
// 	ErrorCodeIndexNaN          = 203
// 	ErrorCodeInvalidField      = 209
// 	ErrorCodeInvalidForm       = 210

// 	ErrorCodeRaftInternal = 300
// 	ErrorCodeLeaderElect  = 301

// 	ErrorCodeWatcherCleared    = 400
// 	ErrorCodeEventIndexCleared = 401
// */
// func EtcdErrorType(e error) int {
// 	es := e.Error()
// 	str := strings.Split(es, ":")
// 	if len(str) >= 2 {
// 		no := str[0]
// 		log.Println("no=", no)
// 		switch no {
// 		// key not found
// 		case "100", "102", "104", "105", "107", "108", "110":
// 			return 100
// 		}
// 	}

// 	return 0
// }
