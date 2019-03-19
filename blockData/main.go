package main

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"io/ioutil"
	"log"
	"path/filepath"
	"fmt"
	"database/sql"
)

const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "blockchain"
)


func createConnection()(*sql.DB, error){
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	db,err := sql.Open("mysql",dsn)
	if err != nil{
		fmt.Printf("Open mysql failed,err:%v\n",err)
		return nil,err
	}
	return db,nil
}
func main() {
	// Load the certificate for the TLS connection which is automatically
	// generated by btcd when it starts the RPC server and doesn't already
	// have one.
	btcdHomeDir := btcutil.AppDataDir("btcd", false)
	certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))
	if err != nil {
		log.Fatal(err)
	}

	// Create a new RPC client using websockets.  Since this example is
	// not long-lived, the connection will be closed as soon as the program
	// exits.
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:18556",
		Endpoint:     "ws",
		User:         "hht",
		Pass:         "hht",
		Certificates: certs,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Query the RPC server for the current block count and display it.
	blockHash,_ := client.GetBlockHash(1)
	block,_ := client.GetBlockVerbose(blockHash)
	txHash,_ := chainhash.NewHashFromStr(block.Tx[0])
	out,_ := client.GetTxOut(txHash,0,true)
	fmt.Println(out.ScriptPubKey.Addresses[0])

}