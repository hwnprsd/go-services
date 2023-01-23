package nft

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/streadway/amqp"

	"flaq.club/workers/pkgs/nft/FlaqPoap"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func HandleMessages(payload *amqp.Delivery) {
	payload.Ack(false)
}

func Mint(address common.Address, uri string) {
	rpcUrl := os.Getenv("RPC_URL")
	contractAddressString := os.Getenv("CONTRACT_ADDRESS")
	privateKeyHex := os.Getenv("PRIVATE_KEY")

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
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
