package service

import (
	"bytes"
	"convenios/entidades"
	"convenios/model"
	"convenios/repository"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/dranikpg/dto-mapper"
	"html/template"
)

type IConvenioService interface {
	GuardarConvenio(convenio *model.Convenio) (*model.Convenio, error)
	GetConvenios() ([]model.Convenio, error)
	GetConvenio() (*model.Convenio, error)
	ActualizarConvenio(convenio *model.Convenio) error
	GenerarPDF(id string) ([]byte, error)
}

func GuardarConvenio(convenio *model.Convenio) (*model.Convenio, error) {

	entity := &entidades.Convenio{}

	if err := dto.Map(&entity, convenio); err != nil {
		fmt.Println(err)
		return nil, err
	}

	entity, err := repository.SaveConvenio(entity)

	if err != nil {
		return nil, err
	}

	if err := dto.Map(&convenio, entity); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return convenio, nil
}

func GetConvenios() ([]model.Convenio, error) {

	var convenioModel []model.Convenio

	entityList, err := repository.GetConvenios()

	if err != nil {
		return nil, err
	}

	if err := dto.Map(&convenioModel, entityList); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return convenioModel, nil
}

func GetConvenio(id string) (*model.Convenio, error) {
	var convenioModel model.Convenio
	entity, err := repository.GetConvenio(id)

	if err != nil {
		return nil, err
	}

	if err := dto.Map(&convenioModel, entity); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &convenioModel, nil
}

func ActualizarConvenio(convenio *model.Convenio) error {

	entity := &entidades.Convenio{}

	if err := dto.Map(&entity, convenio); err != nil {
		fmt.Println(err)
		return err
	}

	return repository.ActualizarConvenio(entity)

}

func GenerarPDF(id string) ([]byte, error) {
	var err error

	convenioRespo, err := GetConvenio(id)

	if err != nil {
		return nil, err
	}

	var pdf = model.ConvenioPDF{
		NumeroConvenio: id,
		Convenio:       *convenioRespo,
	}

	var templ *template.Template

	if templ, err = template.ParseFiles("convenio.html"); err != nil {
		fmt.Println("Error")
	}

	var body bytes.Buffer

	if err = templ.Execute(&body, pdf); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewReader(body.Bytes()))

	page.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)

	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	err = pdfg.Create()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return pdfg.Bytes(), nil

}
