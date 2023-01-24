package nft

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"log"
	"math/big"
	"os"

	"flaq.club/workers/pkgs/nft/FlaqPoap"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hwnprsd/shared_types"
	"github.com/streadway/amqp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func HandleMessages(payload *amqp.Delivery) {
	message := shared_types.MintPoapMessage{}
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
			payload.Reject(false)
		}
	}()
	err := json.Unmarshal(payload.Body, &message)
	if err != nil {
		log.Printf("Error parsing JSON message. Please check what the sender sent! QUEUE - %s", payload.Body)
		payload.Reject(false)
		return
	}
	MintPoap(message.Address, message.TokenURI)
	payload.Ack(false)
}

func failIfFalse(exists bool) {
	if !exists {
		panic("Invalid ENV")
	}
}

func MintPoap(addressString string, uri string) {
	rpcUrl, exists := os.LookupEnv("RPC_URL")
	failIfFalse(exists)
	contractAddressString, exists := os.LookupEnv("CONTRACT_ADDRESS")
	failIfFalse(exists)
	privateKeyHex, exists := os.LookupEnv("PRIVATE_KEY")
	failIfFalse(exists)
	address := common.BytesToAddress([]byte(addressString))

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Println("Error converting Hex to ECDSA")
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	contractAddress := common.HexToAddress(contractAddressString)
	instance, err := FlaqPoap.NewFlaqPoap(contractAddress, client)
	if err != nil {
		panic(err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := instance.MintCollectionNFT(auth, address, uri)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("tx sent: %s", tx.Hash().Hex())
}
