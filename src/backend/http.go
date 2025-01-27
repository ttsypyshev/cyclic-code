package backend

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/errors", ErrorPage)
	r.GET("/syndromes", SyndromePage)
	r.GET("/syndromes/array", SyndromeArrayPage)
	r.GET("/results/", ResultsPage)
}

func ErrorPage(c *gin.Context) {
	errorsView := GetErrorsByClassesString(ErrorClasses)

	tmpl, _ := template.ParseFiles("./templates/errors.html")
	tmpl.Execute(c.Writer, errorsView)
}

func SyndromePage(c *gin.Context) {
	syndromeView := SyndromeTableToString(*GetSyndromeTableVar())

	tmpl, _ := template.ParseFiles("./templates/syndromes.html")
	tmpl.Execute(c.Writer, syndromeView)
}

func SyndromeArrayPage(c *gin.Context) {
	errorMap := GetSyndromeArrayStr(N, GenPolynomial)

	tmpl, _ := template.ParseFiles("./templates/syndromesArray.html")
	tmpl.Execute(c.Writer, errorMap)
}

func ResultsPage(c *gin.Context) {
	tmpl, _ := template.ParseFiles("./templates/results.html")
	tmpl.Execute(c.Writer, Result)
}
