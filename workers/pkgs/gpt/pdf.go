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

func (handler *GptHandler) ParsePdfCv(data shared_types.PdfParseCVMessage) error {
	// 1. Parse PDF and get text data
	// 2. Use GPT to extract the required information
	// 3. Call back API Queue with the data and store it in the database

	text, err := extractText(data.Url)
	if err != nil {
		log.Println("Error extracting data", err)
		handler.ApiQueue.PublishMessage(shared_types.NewApiCallbackMessage(
			data.TaskId,
			"CV_EXTRACTION_FAILED",
			err.Error(),
		))
		return err
	}
	analysis, err := GetCvAnalysis(*text)
	if err != nil {
		log.Println("Error analysing CV data", err)
		handler.ApiQueue.PublishMessage(shared_types.NewApiCallbackMessage(
			data.TaskId,
			"CV_ANALYSIS_FAILED",
			err.Error(),
		))
		return err
	}
	handler.ApiQueue.PublishMessage(shared_types.NewApiCallbackMessage(
		data.TaskId,
		"CV_ANALYSIS_COMPLETE",
		analysis,
	))

	return nil
}

func extractText(pdfUrl string) (*string, error) {
	resp, err := http.Get(pdfUrl)
	log.Println("Extracing text from PDF")
	re := regexp.MustCompile(`\(([^)]+)\)`)
	if err != nil {
		log.Println("Error fetching the PDF from URL", err)
		return nil, err
	}
	defer resp.Body.Close()

	conf := model.NewDefaultConfiguration()
	b, _ := ioutil.ReadAll(resp.Body)
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
