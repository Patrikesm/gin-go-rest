package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/patrike-miranda/gin-go-rest/controllers"
	"github.com/patrike-miranda/gin-go-rest/database"
	"github.com/patrike-miranda/gin-go-rest/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	// esse trecho apenas deixa a estrutura de teste mais fácil de ler
	gin.SetMode(gin.ReleaseMode)

	// aqui define o modelo padrão de routes para o gin
	routes := gin.Default()
	return routes
}

func CreateMockStudent() {
	var student = models.Student{
		Name: "Mock Student",
		CPF:  "12345678910",
		RG:   "123456789",
	}

	database.DB.Create(&student)

	ID = int(student.ID)
}

func DeleteMockStudent() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

// todo teste em go vai receber a palavra "Test" na frente assim o go reconhece
// além disso ele vai utilizar sempre esse "t *testing.T" para ter as ferramentas
func TestVerifyWelcomeStatusCode(t *testing.T) {
	// como estamos realizando um teste do zero para não pegar todas as rotas definimos aqui do zero
	r := SetupTestRoutes()

	// com isso criamos uma mesma rota com o que queremos testar
	r.GET("/:name", controllers.Welcome)

	//para realizar o teste precisamos definir um request de acordo com a rota
	req, _ := http.NewRequest("GET", "/patrike", nil)

	// e um response que vai armazenar o resultado
	res := httptest.NewRecorder()

	// e utilizando o gin executamos a rota
	r.ServeHTTP(res, req)

	// por fim definimos uma verificação com o testify
	assert.Equal(t, http.StatusOK, res.Code, "Should be equals")

	// aqui é a validação do body
	// para isso será necessário criar um mock de comparação antes
	mockDaResposta := `{"API says":"Olá patrike, tudo bem?"}`

	// depois foi necessário ler todo conteúdo da resposta da requisição utilizando "ioutil"
	responseBody, _ := ioutil.ReadAll(res.Body)

	// por fim mais uma verificação usando testify
	assert.Equal(t, mockDaResposta, string(responseBody))
}

func TestCheckAllStudentsReturn(t *testing.T) {
	database.ConnectDB()

	CreateMockStudent()

	r := SetupTestRoutes()

	r.GET("/alunos", controllers.GetAll)

	req, _ := http.NewRequest("GET", "/alunos", nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	defer DeleteMockStudent()
}

func TestSearchByDocument(t *testing.T) {
	database.ConnectDB()

	CreateMockStudent()

	r := SetupTestRoutes()

	r.GET("/alunos/document/:document", controllers.GetByDocument)

	req, _ := http.NewRequest("GET", "/alunos/document/12345678910", nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var body []models.Student

	rawBody, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(rawBody, &body)

	for _, b := range body {
		assert.Equal(t, "Mock Student", b.Name)
	}

	assert.Equal(t, http.StatusOK, res.Code)

	defer DeleteMockStudent()
}

func TestGetOneStudent(t *testing.T) {
	var body models.Student

	database.ConnectDB()
	CreateMockStudent()

	r := SetupTestRoutes()
	r.GET("/alunos/:id", controllers.GetOne)

	req, _ := http.NewRequest("GET", "/alunos/"+strconv.Itoa(ID), nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	rawBody, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(rawBody, &body)

	assert.Equal(t, "12345678910", body.CPF)
	assert.Equal(t, "123456789", body.RG)

	defer DeleteMockStudent()
}

func TestDeleteMockStudent(t *testing.T) {
	database.ConnectDB()

	CreateMockStudent()

	r := SetupTestRoutes()

	r.DELETE("/alunos/:id", controllers.Delete)

	req, _ := http.NewRequest("DELETE", "/alunos/"+strconv.Itoa(ID), nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	defer DeleteMockStudent()
}

func TestUpdateMockStudent(t *testing.T) {
	var student = models.Student{
		Name: "Patrike",
		CPF:  "09887654321",
		RG:   "123654789",
	}

	database.ConnectDB()

	CreateMockStudent()

	r := SetupTestRoutes()

	r.PATCH("/alunos/:id", controllers.Update)

	encondedStudent, _ := json.Marshal(student)

	req, _ := http.NewRequest("PATCH", "/alunos/"+strconv.Itoa(ID), bytes.NewBuffer(encondedStudent))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var body models.Student

	rawBody, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(rawBody, &body)

	assert.Equal(t, "09887654321", body.CPF)
	assert.Equal(t, "123654789", body.RG)

	defer DeleteMockStudent()
}
