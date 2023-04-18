package gpt

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/hwnprsd/shared_types"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/validate"
)

func (handler *GptHandler) handleAnalysis(b []byte, taskId uint) error {
	text, err := parsePdfData(b)
	if err != nil {
		log.Println("Error extracting data", err)
		handler.ApiQueue.PublishMessage(shared_types.NewApiCallbackMessage(
			taskId,
			"CV_EXTRACTION_FAILED_2",
			err.Error(),
		))
		return err
	}
	handler.ApiQueue.PublishMessage(shared_types.NewApiCallbackMessage(
		taskId,
		"CV_EXTRACTION_COMPLETE",
		"awaiting analysis",
	))
	analysis, err := GetCvAnalysis(*text)
	if err != nil {
		log.Println("Error analysing CV data", err)
		handler.ApiQueue.PublishMessage(shared_types.NewApiCallbackMessage(
			taskId,
			"CV_ANALYSIS_FAILED",
			err.Error(),
		))
		return err
	}
	handler.ApiQueue.PublishMessage(shared_types.NewApiCallbackMessage(
		taskId,
		"CV_ANALYSIS_COMPLETE",
		analysis,
	))

	return nil
}

func (handler *GptHandler) ParsePdfCvBytes(data shared_types.PdfParseCVBytesMessage) error {
	return handler.handleAnalysis(data.B, data.TaskId)
}

func (handler *GptHandler) ParsePdfCv(data shared_types.PdfParseCVMessage) error {
	// 1. Parse PDF and get text data
	// 2. Use GPT to extract the required information
	// 3. Call back API Queue with the data and store it in the database

	b, err := extractBytesFromUrl(data.Url)
	if err != nil {
		log.Println("[API Error] Error fetching data from URL", err)
		handler.ApiQueue.PublishMessage(shared_types.NewApiCallbackMessage(
			data.TaskId,
			"CV_EXTRACTION_FAILED_1",
			err.Error(),
		))
		return err
	}
	return handler.handleAnalysis(b, data.TaskId)
}

func parsePdfData(b []byte) (*string, error) {
	re := regexp.MustCompile(`\(([^)]+)\)`)
	conf := model.NewDefaultConfiguration()
	rs := bytes.NewReader(b)
	ctx, err := readValidateAndOptimize(rs, conf)
	if err != nil {
		return nil, err
	}

	if err := ctx.EnsurePageCount(); err != nil {
		return nil, err
	}

	pages, err := api.PagesForPageSelection(ctx.PageCount, []string{}, true)
	if err != nil {
		return nil, err
	}

	pdfText := ""

	for p, v := range pages {
		if !v {
			continue
		}
		r, err := pdfcpu.ExtractPageContent(ctx, p)
		if err != nil {
			return nil, err
		}
		if r == nil {
			continue
		}

		data, _ := ioutil.ReadAll(r)
		str := string(data)
		matches := re.FindAllStringSubmatch(str, -1)
		for _, match := range matches {
			text := match[1]
			pdfText += text + ""
		}
	}

	// fileName = strings.TrimSuffix(filepath.Base(fileName), ".pdf")
	return &pdfText, nil
}

func extractBytesFromUrl(pdfUrl string) ([]byte, error) {
	resp, err := http.Get(pdfUrl)
	log.Println("Extracing text from PDF")
	if err != nil {
		log.Println("Error fetching the PDF from URL", err)
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func readAndValidate(rs io.ReadSeeker, conf *model.Configuration) (ctx *model.Context, err error) {
	if ctx, err = api.ReadContext(rs, conf); err != nil {
		return nil, err
	}

	if conf.ValidationMode == model.ValidationNone {
		// Bypass validation
		return ctx, nil
	}

	if err = validate.XRefTable(ctx.XRefTable); err != nil {
		return nil, err
	}

	return ctx, nil
}

func readValidateAndOptimize(rs io.ReadSeeker, conf *model.Configuration) (ctx *model.Context, err error) {
	ctx, err = readAndValidate(rs, conf)
	if err != nil {
		return nil, err
	}

	if err = pdfcpu.OptimizeXRefTable(ctx); err != nil {
		return nil, err
	}

	return ctx, nil
}
