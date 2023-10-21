package service

import (
	"bytes"
	"convenios/entidades"
	"convenios/model"
	"convenios/repository"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/dranikpg/dto-mapper"
	gomail "gopkg.in/mail.v2"
)

const SERVER_SMTP = "convenio-uis-notificaciones@hotmail.com"
const PASS_SMTP = "convenios-uis-notificaciones"

type flujoRoles struct {
	estado        model.EstadoConvenio
	siguienteRole model.Role
}

type filtroConvenio struct {
	filtro bool
	estado model.EstadoConvenio
}

var filtroMacro = map[model.Role]filtroConvenio{
	model.Directo_Juridico: {
		filtro: false,
		estado: model.Aprobado_Director_Relex,
	},
	model.Consejo_Academico: {
		filtro: true,
		estado: model.Aprobado_Director_Relex,
	},
}

var filtroInvestigacion = map[model.Role]filtroConvenio{
	model.Directo_Juridico: {
		filtro: false,
		estado: model.Aprobado_Consejo_Academico,
	},
	model.Vicerectoria: {
		filtro: true,
		estado: model.Aprobado_Consejo_Academico,
	},
}

// manejo los permisos
var permisos = map[model.Role][]model.EstadoConvenio{
	model.Secretaria:        {model.Creado, model.Aprobado_Rectoria},
	model.Director_Relex:    {model.Aprobado_Secretaria},
	model.Consejo_Academico: {model.Aprobado_Director_Relex},
	model.Vicerectoria:      {model.Aprobado_Consejo_Academico},
	model.Directo_Juridico:  {model.Aprobado_Vicerectoria, model.Aprobado_Director_Relex, model.Aprobado_Consejo_Academico},
	model.Rectoria:          {model.Aprobado_Director_Juridico},
}

// acá manejo el flujo
var flujo = map[model.Role]map[bool]flujoRoles{
	model.Secretaria: {
		true: flujoRoles{
			estado:        model.Aprobado_Secretaria,
			siguienteRole: model.Director_Relex,
		},
		false: flujoRoles{
			estado:        model.Rechazado_Secretaria,
			siguienteRole: model.Gestor,
		},
	},
	model.Director_Relex: {
		true: flujoRoles{
			estado:        model.Aprobado_Director_Relex,
			siguienteRole: model.Directo_Juridico,
		},
		false: flujoRoles{
			estado:        model.Rechazado_Director_Relex,
			siguienteRole: model.Gestor,
		},
	},
	model.Director_Relex_Macro: {
		true: flujoRoles{
			estado:        model.Aprobado_Director_Relex,
			siguienteRole: model.Consejo_Academico,
		},
		false: flujoRoles{
			estado:        model.Rechazado_Director_Relex,
			siguienteRole: model.Gestor,
		},
	},
	model.Consejo_Academico: {
		true: flujoRoles{
			estado:        model.Aprobado_Consejo_Academico,
			siguienteRole: model.Directo_Juridico,
		},
		false: flujoRoles{
			estado:        model.Rechazado_Consejo_Academico,
			siguienteRole: model.Gestor,
		},
	},
	model.Consejo_Academico_Investigacion: {
		true: flujoRoles{
			estado:        model.Aprobado_Consejo_Academico,
			siguienteRole: model.Vicerectoria,
		},
		false: flujoRoles{
			estado:        model.Rechazado_Consejo_Academico,
			siguienteRole: model.Gestor,
		},
	},
	model.Vicerectoria: {
		true: flujoRoles{
			estado:        model.Aprobado_Vicerectoria,
			siguienteRole: model.Directo_Juridico,
		},
		false: flujoRoles{
			estado:        model.Rechazado_Vicerectoria,
			siguienteRole: model.Gestor,
		},
	},
	model.Directo_Juridico: {
		true: flujoRoles{
			estado:        model.Aprobado_Director_Juridico,
			siguienteRole: model.Rectoria,
		},
		false: flujoRoles{
			estado:        model.Rechazado_Director_Juridico,
			siguienteRole: model.Gestor,
		},
	},
	model.Rectoria: {
		true: flujoRoles{
			estado:        model.Aprobado_Rectoria,
			siguienteRole: model.Secretaria,
		},
		false: flujoRoles{
			estado:        model.Rechazado_Rectoria,
			siguienteRole: model.Gestor,
		},
	},
}

// Función para obtener el estado y el siguiente rol después de la aprobación
func ObtenerEstadoYSiguienteRol(role model.Role, aprobo bool) (flujoRoles, error) {
	if transiciones, ok := flujo[role]; ok {
		if estadoSiguiente, ok := transiciones[aprobo]; ok {
			return estadoSiguiente, nil
		}
	}

	return flujoRoles{}, errors.New("Error de role para cambiar estado")
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

func GetConvenios(roleUser string, idGestor string) ([]model.Convenio, error) {

	if roleUser != model.Gestor.String() {
		idGestor = ""
	}

	entityList, err := repository.GetConvenios(idGestor, permisos[model.Role(roleUser)])

	if err != nil {
		return nil, err
	}

	var convenioModel []model.Convenio

	if err := dto.Map(&convenioModel, entityList); err != nil {
		fmt.Println(err)
		return nil, err
	}

	filtroMacro, okMacro := filtroMacro[model.Role(roleUser)]
	var convenioModelFiltrado []model.Convenio

	for _, convenio := range convenioModel {
		convenioModelFiltrado = append(convenioModelFiltrado, convenio)
	}

	if okMacro {
		convenioModelFiltrado = filtrarConvenioMacro(filtroMacro, convenioModelFiltrado)
		if roleUser != model.Directo_Juridico.String() {
			return convenioModelFiltrado, nil
		}
	}

	filtroInvestigacion, okInv := filtroInvestigacion[model.Role(roleUser)]

	if okInv {
		convenioModelFiltrado = filtrarConvenioInvestigacion(filtroInvestigacion, convenioModelFiltrado)
		return convenioModelFiltrado, nil
	}

	return convenioModel, nil

}

func filtrarConvenioMacro(filtroMacro filtroConvenio, convenioModel []model.Convenio) []model.Convenio {
	var convenioModelFiltrado []model.Convenio

	for _, convenio := range convenioModel {

		if convenio.Estado != filtroMacro.estado {
			convenioModelFiltrado = append(convenioModelFiltrado, convenio)
		}

		if convenio.Estado == filtroMacro.estado && ((filtroMacro.filtro && convenio.TipologiaConvenio == "Macro") ||
			(!filtroMacro.filtro && convenio.TipologiaConvenio != "Macro")) {
			convenioModelFiltrado = append(convenioModelFiltrado, convenio)
		}

	}
	return convenioModelFiltrado
}

func filtrarConvenioInvestigacion(filtroInv filtroConvenio, convenioModel []model.Convenio) []model.Convenio {
	var convenioModelFiltrado []model.Convenio

	for _, convenio := range convenioModel {

		if convenio.Estado != filtroInv.estado {
			convenioModelFiltrado = append(convenioModelFiltrado, convenio)
		}

		if convenio.Estado == filtroInv.estado && ((filtroInv.filtro && convenio.Caracterizacion == "Investigacion") ||
			(!filtroInv.filtro && convenio.Caracterizacion != "Investigacion")) {
			convenioModelFiltrado = append(convenioModelFiltrado, convenio)
		}

	}
	return convenioModelFiltrado
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

	type Datos struct {
		Users []string `json:"users"`
	}

	requestBody, err := json.Marshal(&Datos{Users: convenioRespo.HistorialFirma})

	if err != nil {
		return nil, err
	}

	// Realizar la solicitud POST
	resp, err := http.Post("http://localhost:8081/api/usuario/id", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("Error al realizar la solicitud POST:")
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New("Error al realizar la solicitud POST")
	}

	var firmaInfo []model.FirmaInfo

	json.Unmarshal(responseBody, &firmaInfo)

	var pdf = model.ConvenioPDF{
		NumeroConvenio: id,
		Convenio:       *convenioRespo,
		FirmaInfo:      firmaInfo,
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

func sendEmail(convenioRespo *model.Convenio, role string) error {

	idGestor := ""

	if role == model.Gestor.String() {
		idGestor = "/" + convenioRespo.IdGestorCreador
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8081/api/usuario/correo/"+role+idGestor, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var usuario model.Usuario
	json.Unmarshal(bodyBytes, &usuario)

	fmt.Println(usuario.Email)

	m := gomail.NewMessage()
	m.SetHeader("From", SERVER_SMTP)

	m.SetHeader("To", usuario.Email)

	switch role {
	case model.Secretaria.String(), model.Director_Relex.String(), model.Consejo_Academico.String(), model.Vicerectoria.String(), model.Directo_Juridico.String():
		m.SetHeader("Subject", "Nuevo convenio creado")
		m.SetBody("text/html", "Hola se informa que se ha creado el convenio con nombre <b>"+convenioRespo.NombreConvenio+"</b> y id <b>"+convenioRespo.ID.Hex()+"</b><br>Por favor validar desde el portal.")
	case model.Gestor.String():
		m.SetHeader("Subject", "Convenio rechazado")
		m.SetBody("text/html", "Hola se informa que se ha rechazado el convenio con nombre <b>"+convenioRespo.NombreConvenio+"</b> y id <b>"+convenioRespo.ID.Hex()+"</b><br>Por favor validar desde el portal.")
	default:
		return errors.New("Error de role para enviar mail")
	}

	d := gomail.NewDialer("smtp-mail.outlook.com", 587, SERVER_SMTP, PASS_SMTP)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func CambiarEstadoConvenio(id string, cambio model.CambiarEstadoConvenio, role string, idAprueba string) error {

	convenioRespo, err := GetConvenio(id)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if convenioRespo.Estado == model.Aprobado_Rectoria && role == string(model.Secretaria) {
		convenioRespo.Estado = model.Convenio_Aprobado
		if err := ActualizarConvenio(convenioRespo); err != nil {
			return err
		}

		return nil
	}

	if convenioRespo.TipologiaConvenio == "Macro" && convenioRespo.Estado == model.Rechazado_Consejo_Academico {
		return errors.New("Error No se puede cambiar el estado ya que se encuentra en estado " + string(convenioRespo.Estado))
	}

	if convenioRespo.Caracterizacion == "Investigacion" && role == model.Consejo_Academico.String() {
		role = role + " " + "investigacion"
	}

	if convenioRespo.TipologiaConvenio == "Macro" && role == model.Director_Relex.String() {
		role = role + " " + "macro"
	}

	flujoRole, err := ObtenerEstadoYSiguienteRol(model.Role(role), cambio.CambioEstado)

	if err != nil {
		return err
	}

	if !cambio.CambioEstado {

		if len(cambio.Observacion) < 1 {
			return errors.New("Observacion no válida")
		}

		convenioRespo.Observaciones = cambio.Observacion
		convenioRespo.HistorialFirma = []string{convenioRespo.IdGestorCreador}
	}

	if cambio.CambioEstado {
		convenioRespo.HistorialFirma = append(convenioRespo.HistorialFirma, idAprueba)
	}

	convenioRespo.Estado = flujoRole.estado
	sendEmail(convenioRespo, string(flujoRole.siguienteRole))

	if err := ActualizarConvenio(convenioRespo); err != nil {
		return err
	}

	return nil
}
