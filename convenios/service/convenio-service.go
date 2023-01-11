package service

import (
	"bytes"
	"convenios/entidades"
	"convenios/model"
	"convenios/repository"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/dranikpg/dto-mapper"
	"github.com/yookoala/realpath"
	gomail "gopkg.in/mail.v2"
	"html/template"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

const SERVER_SMTP = "convenio-uis-notificaciones@hotmail.com"
const PASS_SMTP = "convenios-uis-notificaciones"
const EMAIL_SECRETARIA = ""

type IConvenioService interface {
	GuardarConvenio(convenio *model.Convenio) (*model.Convenio, error)
	GetConvenios() ([]model.Convenio, error)
	GetConvenio() (*model.Convenio, error)
	ActualizarConvenio(convenio *model.Convenio) error
	GenerarPDF(id string) ([]byte, error)
	FirmarConvenio(id string) error
	CambiarEstadoConvenio(id string, cambio model.CambiarEstadoConvenio) error
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

func FirmarConvenio(id string, file multipart.File, header *multipart.FileHeader) error {

	convenioRespo, err := GetConvenio(id)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	archivo := "upload/firma/" + id + "." + strings.Split(header.Filename, ".")[1]
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	myRealpath, err := realpath.Realpath(archivo)
	fmt.Println(myRealpath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	_, err = io.Copy(f, file)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	convenioRespo.FirmaUrl = myRealpath
	convenioRespo.Estado = model.Firmado

	if err := ActualizarConvenio(convenioRespo); err != nil {
		fmt.Println(err)
		return err
	}

	return sendEmail(convenioRespo)

}

func sendEmail(convenioRespo *model.Convenio) error {

	m := gomail.NewMessage()
	m.SetHeader("From", SERVER_SMTP)
	m.SetHeader("To", EMAIL_SECRETARIA)
	m.SetHeader("Subject", "Nuevo convenio creado")
	m.SetBody("text/html", "Hola se informa que se ha creado el convenio con nombre <b>"+convenioRespo.NombreConvenio+"</b> y id <b>"+convenioRespo.ID.Hex()+"</b><br>Por favor validar desde el portal.")
	d := gomail.NewDialer("smtp-mail.outlook.com", 587, SERVER_SMTP, PASS_SMTP)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func CambiarEstadoConvenio(id string, cambio model.CambiarEstadoConvenio) error {

	convenioRespo, err := GetConvenio(id)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if convenioRespo.Estado != model.Firmado {
		return errors.New("Estado no válido")
	}

	if cambio.CambioEstado {
		convenioRespo.Estado = model.Aprobado_Secretaria
	} else {
		convenioRespo.Estado = model.Rechazado_Secretaria

		if len(cambio.Observacion) < 1 {
			return errors.New("Observacion no válida")
		}

		convenioRespo.Observaciones = cambio.Observacion
	}

	if err := ActualizarConvenio(convenioRespo); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
