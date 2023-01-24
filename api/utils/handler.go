package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RequestBody struct {
	Data             interface{}
	Query            interface{}
	User             interface{}
	ParseQueryParams bool
}

type RequestError struct {
	StatusCode int
	Err        error
	Message    string
}

func (re *RequestError) Error() string {
	return fmt.Sprintf("Request Errored Out %s", re.Err)
}

type PostHandler func(data RequestBody) (interface{}, error)
type GetHandler func(data RequestBody) (interface{}, error)

func handlePanic(c *fiber.Ctx) {
	// Panic handler
	if err := recover(); err != nil {
		fmt.Println("We survived a panic")
		fmt.Println(err)
		c.Status(500).JSON(fiber.Map{
			"status_code": 500,
			"message":     "A server panic has occured",
			"error":       err,
		})
	}
}

func PostRequestHandler(c *fiber.Ctx, actualHandler PostHandler, body RequestBody) error {
	// Handle Panics
	defer handlePanic(c)

	if b, err := ValidateBody(body.Data, c); b {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status_code": 400,
			"error":       err,
		})
	}
	res, err := actualHandler(body)
	if err != nil {
		if reqError, ok := err.(*RequestError); ok {
			return c.Status(reqError.StatusCode).JSON(fiber.Map{
				"status_code": reqError.StatusCode,
				"error":       reqError.Err.Error(),
				"message":     reqError.Message,
			})
		} else {
			panic(err)
		}

	}
	return c.Status(200).JSON(fiber.Map{
		"status_code": 200,
		"data":        res,
	})
}

func GetRequestHandler(c *fiber.Ctx, actualHandler GetHandler, body RequestBody) error {
	// Handle Panics
	defer handlePanic(c)

	res, err := actualHandler(body)

	if err != nil {
		if reqError, ok := err.(*RequestError); ok {
			return c.Status(reqError.StatusCode).JSON(fiber.Map{
				"status_code": reqError.StatusCode,
				"error":       reqError.Err.Error(),
				"message":     reqError.Message,
			})
		} else {
			panic(err)
		}

	}
	return c.Status(200).JSON(fiber.Map{
		"status_code": 200,
		"data":        res,
	})
}
