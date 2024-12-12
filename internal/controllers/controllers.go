package controllers

import (
	"encoding-restoration/internal/models"
	"html/template"

	"github.com/gin-gonic/gin"
)

func ErrorPage(c *gin.Context) {
	errorsView := models.GetErrorsByClassesString(models.ErrorClasses)

	tmpl, _ := template.ParseFiles("./templates/errors.html")
	tmpl.Execute(c.Writer, errorsView)
}

func SyndromePage(c *gin.Context) {
	syndromeView := models.SyndromeTableToString(*models.GetSyndromeTableVar())

	tmpl, _ := template.ParseFiles("./templates/syndromes.html")
	tmpl.Execute(c.Writer, syndromeView)
}

func SyndromeArrayPage(c *gin.Context) {
	errorMap := models.GetSyndromeArrayStr(models.N, models.GenPolynomial)

	tmpl, _ := template.ParseFiles("./templates/syndromesArray.html")
	tmpl.Execute(c.Writer, errorMap)
}

func ResultsPage(c *gin.Context) {
	tmpl, _ := template.ParseFiles("./templates/results.html")
	tmpl.Execute(c.Writer, models.Result)
}
