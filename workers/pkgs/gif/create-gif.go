package gif

import (
	"encoding/json"
	"errors"
	"log"

	"flaq.club/workers/models"
	"github.com/google/uuid"
	"github.com/hwnprsd/shared_types"
	"github.com/streadway/amqp"
	"gorm.io/gorm"

	"bytes"
	"crypto/rand"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"

	"flaq.club/workers/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ericpauley/go-quantize/quantize"
	"github.com/fogleman/gg"
)

type CreateGifHandler struct {
	DB       *gorm.DB
	NftQueue *amqp.Channel
}

func NewCreateGifHandler(nftQueue *amqp.Channel, db *gorm.DB) *CreateGifHandler {
	return &CreateGifHandler{NftQueue: nftQueue, DB: db}
}

func (h *CreateGifHandler) HandleMessages(payload *amqp.Delivery) {
	log.SetPrefix("GIF_HANDLER: ")

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
	case shared_types.WORK_TYPE_CREATE_GIF:
		message := shared_types.CreateGifMessage{}
		json.Unmarshal(payload.Body, &message)
		log.Println("Asking to mint POAP when disabled")
		// Enable when live
		h.CreateGif(message)
		break
	}
	payload.Ack(false)
}

func (h *CreateGifHandler) CreateGif(message shared_types.CreateGifMessage) error {
	event := models.Web3Event{}
	result := h.DB.Model(&models.Web3Event{}).Where("id = ?", message.EventID).First(&event)
	if result.Error != nil {
		log.Println("Error finding event with ID ", message.EventID)
		return errors.New("Event ID invalid")
	}

	// ----

	awsAccountId := os.Getenv("AWS_ACCOUNT_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_KEY")

	// -----
	fileType := "image"
	fileExtension := ".gif"

	// Open the input GIF file
	f, err := os.Open("final-poap.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Decode the input GIF file
	gifImage, err := gif.DecodeAll(f)
	if err != nil {
		panic(err)
	}

	// Create a new GIF that will hold the modified frames
	newGif := &gif.GIF{}
	q := quantize.MedianCutQuantizer{}

	initialFrame := gifImage.Image[0]
	p := q.Quantize(make([]color.Color, 0, 256), initialFrame)
	imageContext := gg.NewContextForImage(initialFrame)
	if err := imageContext.LoadFontFace("SpaceMono-Regular.ttf", float64(event.FontSize)); err != nil {
		panic(err)
	}
	imageContext.SetColor(color.RGBA{255, 206, 189, 100})
	// Iterate over each frame in the input GIF
	for _, frame := range gifImage.Image {
		newImage := image.NewPaletted(frame.Bounds(), p)
		imageContext.DrawImage(frame, 0, 0)
		imageContext.SetColor(color.White)
		imageContext.DrawString(message.UserName, float64(event.TextX), float64(event.TextY))
		draw.Draw(newImage, frame.Bounds(), imageContext.Image(), frame.Bounds().Min, draw.Src)
		// addLabel(frame, 52, 620, "Ashwin Prasad")
		newGif.Image = append(newGif.Image, newImage)
		newGif.Delay = append(newGif.Delay, gifImage.Delay[0])
	}

	buf := new(bytes.Buffer)
	gif.EncodeAll(buf, newGif)
	creds := credentials.NewStaticCredentials(awsAccountId, awsSecretKey, "")

	// Create an S3 session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-south-1"), // replace with your own region
		Credentials: creds,
	})

	svc := s3.New(sess)
	length := 16 // specify the length of the random string
	randomBytes := make([]byte, length)

	_, _ = rand.Read(randomBytes)

	keyBase := uuid.New()
	keyName := fmt.Sprintf("%s/%s/%s-%s%s", event.EventKey, fileType, keyBase, message.UserName, fileExtension) // replace with your own key name

	_, err = svc.PutObject(&s3.PutObjectInput{
		Body:        bytes.NewReader(buf.Bytes()),
		Bucket:      aws.String(event.BucketName),
		Key:         aws.String(keyName),
		ContentType: aws.String("image/gif"),
	})

	if err != nil {
		log.Println("Error: Failed to create session", err)
		return nil
	}
	gifUrl := fmt.Sprintf("https://%s.s3.ap-south-1.amazonaws.com/%s", event.BucketName, keyName)
	// Reset the typs for json
	fileType = "json"
	fileExtension = ".json"
	keyName = fmt.Sprintf("%s/%s/%s-%s%s", event.EventKey, fileType, keyBase, message.UserName, fileExtension) // replace with your own key name

	//{
	//      "description": "Proof that Aayush attended the FLAQ x Rotaract Bangalore Junction session on 'Dive into Web3'",
	//      "external_url": "https://flaq.club",
	//      "image": "https://flaq-assets.s3.ap-south-1.amazonaws.com/poap1/images/Aayush.gif",
	//      "name": "Flaq x Rotaract Bangalore Junction",
	//      "attributes": [
	//           {
	//                "trait_type": "Level",
	//                "value": "Dive into web3"
	//           }
	//      ]
	// }
	jsonData := map[string]any{
		"description":  event.EventDescription,
		"external_url": "https://flaq.club",
		"image":        gifUrl,
		"name":         event.EventName,
		"attributes": []map[string]string{
			{
				"trait_type": "Event",
				"value":      event.EventName,
			},
		},
	}
	jsonBytes, _ := json.Marshal(jsonData)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Body:        bytes.NewReader(jsonBytes),
		Bucket:      aws.String(event.BucketName),
		Key:         aws.String(keyName),
		ContentType: aws.String("application/json"),
	})

	if err != nil {
		log.Println("Failed to create session", err)
		return nil
	}
	jsonUrl := fmt.Sprintf("https://%s.s3.ap-south-1.amazonaws.com/%s", event.BucketName, keyName)
	log.Println("JSON URL - ", jsonUrl)
	log.Println("GIF Created - ", jsonUrl)

	nftMintMessage := shared_types.NewMintPoapMessage(
		message.UserEmail,
		message.UserWalletAddress,
		message.UserName,
		jsonUrl,
		message.EmailTemplateId,
	)

	// Publish a message asking to mint the NFT
	utils.PublishMessage(h.NftQueue, "nft", nftMintMessage)
	return nil
}
