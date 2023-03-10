package nft

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"flaq.club/workers/pkgs/nft/FlaqInsignia"
	"flaq.club/workers/pkgs/nft/FlaqPoap"
	"flaq.club/workers/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hwnprsd/shared_types"
	"github.com/hwnprsd/shared_types/status"
	"github.com/streadway/amqp"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NftMintHandler struct {
	DB                    *gorm.DB
	ApiQueue, MailerQueue *utils.Queue
}

func NewNftMintHandler(apiQueue, mailerQueue *utils.Queue, db *gorm.DB) *NftMintHandler {
	return &NftMintHandler{DB: db, MailerQueue: mailerQueue, ApiQueue: apiQueue}
}

func (h *NftMintHandler) HandleMessages(payload *amqp.Delivery) {
	baseMessage := shared_types.MessagingBase{}
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
			payload.Reject(false)
		}
	}()
	err := json.Unmarshal(payload.Body, &baseMessage)
	if err != nil {
		log.Printf("Error parsing JSON message. Please check what the sender sent! QUEUE - %s", payload.Body)
		payload.Reject(false)
		return
	}
	switch baseMessage.WorkType {
	case shared_types.WORK_TYPE_MINT_POAP:
		message := shared_types.MintPoapMessage{}
		json.Unmarshal(payload.Body, &message)
		log.Println("Asking to mint POAP when disabled")
		// Enable when live
		h.MintPoap(message)
		break
	case shared_types.WORK_TYPE_MINT_QUIZ_NFT:
		message := shared_types.MintQuizNFTMessage{}
		json.Unmarshal(payload.Body, &message)
		ownerAddress := common.HexToAddress(message.Address)
		h.MintInsignia(ownerAddress, message.TokenURI)
		break
	}
	payload.Ack(false)
}

func failIfFalse(exists bool) {
	if !exists {
		panic("Invalid ENV")
	}
}

func (h *NftMintHandler) MintInsignia(address common.Address, uri string) {
	chainIdString, exists := os.LookupEnv("CHAIN_ID")
	failIfFalse(exists)
	chainId, _ := strconv.ParseInt(chainIdString, 10, 64)
	rpcUrl, exists := os.LookupEnv("ETH_RPC_URL")
	failIfFalse(exists)
	contractAddressStringPartial, exists := os.LookupEnv("CONTRACT_ADDRESS_QUIZ")
	failIfFalse(exists)
	privateKeyHex, exists := os.LookupEnv("PRIVATE_KEY")
	failIfFalse(exists)

	// Remember to remove "0x" form the address provided in the ENV, to ensure that docker compose doesn't end up parsing the hex
	contractAddressString := fmt.Sprintf("0x%s", contractAddressStringPartial)

	log.Printf("Minting on RPC := %s and chain ID := %d and Owner Address := %s", rpcUrl, chainId, contractAddressString)

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

	instance, err := FlaqInsignia.NewFlaqInsignia(contractAddress, client)
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

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := instance.MintInsignia(auth, address, uri)
	if err != nil {
		log.Println("ERROR MINTING POAP", err)
	} else {
		log.Printf("tx sent: %s", tx.Hash().Hex())
	}
}

func (h *NftMintHandler) MintPoap(message shared_types.MintPoapMessage) {

	startMessage := shared_types.NewApiCallbackMessage(message.TaskId, status.POAP_MINT_START, "")
	h.ApiQueue.PublishMessage(startMessage)

	log.SetPrefix("NFT_MINTER: ")
	chainIdString, exists := os.LookupEnv("CHAIN_ID")
	failIfFalse(exists)
	rpcUrl, exists := os.LookupEnv("ETH_RPC_URL")
	failIfFalse(exists)
	contractAddressStringPartial, exists := os.LookupEnv("CONTRACT_ADDRESS")
	failIfFalse(exists)
	privateKeyHex, exists := os.LookupEnv("PRIVATE_KEY")
	failIfFalse(exists)
	address := common.HexToAddress(message.Address)

	contractAddressString := fmt.Sprintf("0x%s", contractAddressStringPartial)
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

	chainId, _ := strconv.ParseInt(chainIdString, 10, 64)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := instance.MintCollectionNFT(auth, address, message.TokenURI)
	if err != nil {
		log.Println("ERROR MINTING POAP", err)
	} else {
		log.Printf("tx sent: %s", tx.Hash().Hex())
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Println("Confirming")
	r, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Println("ERROR Confirming POAP", err)
	} else {
		log.Println("Confirmed:", r.TxHash.String())
	}

	log.Println("Sending mail message")
	mailMessage := shared_types.NewSendMailMessage(
		message.TaskId,
		message.Email,
		"[NFT Inside] Thank you for attending an event with Flaq!",
		message.EmailTemplateId,
		map[string]string{
			"EVENT_NAME":  "Web3 Gorkha Siliguri Event",
			"USER_NAME":   message.Name,
			"RARIBLE_URL": fmt.Sprintf("https://rarible.com/user/%s/owned", message.Address),
			"TX_URL":      fmt.Sprintf("https://polygonscan.com/tx/%s", r.TxHash.String()),
		},
	)
	endMessage := shared_types.NewApiCallbackMessage(message.TaskId, status.POAP_MINT_COMPLETE, "")
	h.ApiQueue.PublishMessage(endMessage)
	if message.Email != "" {
		h.MailerQueue.PublishMessage(mailMessage)
	}

}
