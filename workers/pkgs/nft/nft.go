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
	"strings"

	"flaq.club/workers/pkgs/nft/FlaqInsignia"
	"flaq.club/workers/pkgs/nft/FlaqPoap"
	"flaq.club/workers/pkgs/nft/SolaceSCWFactory"
	"flaq.club/workers/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hwnprsd/shared_types"
	"github.com/hwnprsd/shared_types/status"
	"github.com/streadway/amqp"
	"gorm.io/gorm"

	rcom "github.com/athanorlabs/go-relayer/common"
	contracts "github.com/athanorlabs/go-relayer/impls/mforwarder"
	"github.com/athanorlabs/go-relayer/relayer"
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
	case shared_types.WORK_TYPE_CREATE_SC_WALLET:
		message := shared_types.CreateSmartContractWallet{}
		json.Unmarshal(payload.Body, &message)
		ownerAddress := common.HexToAddress(message.Address)
		h.CreateSmartContractWallet(ownerAddress)
	case shared_types.WORK_TYPE_RELAY_TX:
		message := shared_types.RelayTxMessage{}
		json.Unmarshal(payload.Body, &message)
		userAddress := common.HexToAddress(message.UserAddress)
		contractAddress := common.HexToAddress(message.ContractAddress)
		h.RelayTransaction(contractAddress, userAddress, message.Data, message.Signature, message.Nonce)
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

func (h *NftMintHandler) CreateSmartContractWallet(ownerAddress common.Address) error {
	log.SetPrefix("SCW_WALLET_CREATOR: ")
	chainIdString, exists := os.LookupEnv("CHAIN_ID")
	failIfFalse(exists)
	rpcUrl, exists := os.LookupEnv("ETH_RPC_URL")
	failIfFalse(exists)
	factoryContractAddressPartial, exists := os.LookupEnv("FACTORY_CONTRACT_ADDRESS")
	failIfFalse(exists)
	privateKeyHex, exists := os.LookupEnv("PRIVATE_KEY")
	failIfFalse(exists)

	contractAddressString := fmt.Sprintf("0x%s", factoryContractAddressPartial)
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
	instance, err := SolaceSCWFactory.NewSolaceSCWFactory(contractAddress, client)
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

	tx, err := instance.CreateSCW(auth, ownerAddress)
	if err != nil {
		log.Println("ERROR Creating SCW", err)
	} else {
		log.Printf("tx sent: %s", tx.Hash().Hex())
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Println("Confirming")
	r, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Println("ERROR Confirming SCW Creation", err)
	} else {
		log.Println("Confirmed:", r.TxHash.String())
	}

	return nil
}

func (h *NftMintHandler) RelayTransaction(contractAddress common.Address, userAddress common.Address, data, userSignature string, userNonce int64) error {
	chainIdString, exists := os.LookupEnv("CHAIN_ID")
	failIfFalse(exists)
	rpcUrl, exists := os.LookupEnv("ETH_RPC_URL")
	failIfFalse(exists)
	privateKeyHex, exists := os.LookupEnv("PRIVATE_KEY")
	failIfFalse(exists)

	log.Println("ChainID", chainIdString)
	log.Printf("Data (hash): %s", data)
	log.Printf("User signature: %s", userSignature)

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Println("Error converting Hex to ECDSA")
		panic(err)
	}
	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	// }
	//
	// instance, err := SolaceSCW.NewSolaceSCW(contractAddress, client)
	// if err != nil {
	// 	panic(err)
	// }
	//
	forwarder, err := contracts.NewIMinimalForwarder(contractAddress, client)
	r, err := relayer.NewRelayer(&relayer.Config{
		Ctx:       context.Background(),
		EthClient: client,
		Forwarder: contracts.NewIMinimalForwarderWrapped(forwarder),
		Key:       rcom.NewKeyFromPrivateKey(privateKey),
		ValidateTransactionFunc: func(_ *rcom.SubmitTransactionRequest) error {
			// Note: an actual application will likely want to set this
			return nil
		},
	})
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	dataString := strings.TrimPrefix(data, "0x")
	dataBytes := common.Hex2Bytes(dataString)

	signatureString := strings.TrimPrefix(userSignature, "0x")
	signatureBytes := common.Hex2Bytes(signatureString)

	resp, err := r.SubmitTransaction(&rcom.SubmitTransactionRequest{
		From:      userAddress,
		To:        contractAddress,
		Value:     big.NewInt(0),
		Signature: signatureBytes,
		Data:      dataBytes,
		Nonce:     big.NewInt(int64(userNonce)),
		Gas:       gasPrice,
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp.TxHash)

	}
	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// chainId, _ := strconv.ParseInt(chainIdString, 10, 64)
	//
	// auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0) // in wei
	// auth.GasLimit = uint64(300000)
	// auth.GasPrice = gasPrice
	//
	// log.Println(contractAddress, userAddress)
	// // dataString := strings.TrimPrefix(data, "0x")
	// // dataBytes := common.Hex2Bytes(dataString)
	//
	// req := SolaceSCW.MinimalForwarderForwardRequest{
	// 	From:  userAddress,
	// 	To:    contractAddress,
	// 	Value: big.NewInt(0),
	// 	Data:  common.Hex2Bytes(data),
	// 	Nonce: big.NewInt(int64(userNonce)),
	// 	Gas:   gasPrice,
	// }
	//
	// signatureString := strings.TrimPrefix(userSignature, "0x")
	// signatureBytes := common.Hex2Bytes(signatureString)
	//
	// verified, err := instance.Verify(&bind.CallOpts{}, req, signatureBytes)
	// log.Println("V", verified)
	// if !verified {
	// 	return nil
	// }
	// if err != nil {
	// 	log.Println("FAILURE", err)
	// 	return err
	// }
	//
	// tx, err := instance.Execute(auth, req, signatureBytes)
	// if err != nil {
	// 	log.Println("ERROR Creating SCW", err)
	// } else {
	// 	log.Printf("tx sent: %s", tx.Hash().Hex())
	// }
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// log.Println("Confirming")
	// r, err := bind.WaitMined(ctx, client, tx)
	// if err != nil {
	// 	log.Println("ERROR Confirming SCW Creation", err)
	// } else {
	// 	log.Println("Confirmed:", r.TxHash.String())
	// }
	//
	return nil
}
