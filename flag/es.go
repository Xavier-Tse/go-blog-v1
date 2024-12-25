package flag

import (
	"Backend/models"
)

func ESCreateIndex() {
	models.ArticleModel{}.CreateIndex()
	models.FullTextModel{}.CreateIndex()
}
