package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	models "PT.ADVANCE-FINANCE/Models"
	DataPerusahaan "PT.ADVANCE-FINANCE/Models/DataPerusahaan"
	KeuanganRepository "PT.ADVANCE-FINANCE/repository"
)

// InsertJurusanController adalah handler untuk menyisipkan data keuangan baru.
func GetAdvanceFinanceControllersByID(c *gin.Context) {
	var request DataPerusahaan.Keuangan
	var response models.BaseResponseModels

	response = KeuanganRepository.GetAdvanceFinanceByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertAdvanceFinanceControllers(c *gin.Context) {
	var request DataPerusahaan.Keuangan
	var response models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = KeuanganRepository.InsertAdvanceFinance(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateAdvanceFinanceControllers(c *gin.Context) {
	var request DataPerusahaan.Keuangan
	var response models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = KeuanganRepository.UpdateAdvanceFinance(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteAdvanceFinanceControllers(c *gin.Context) {
	var request DataPerusahaan.Keuangan
	var response models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = KeuanganRepository.DeleteAdvanceFinance(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}