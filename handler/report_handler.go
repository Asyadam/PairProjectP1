package handler

import (
	"fmt"

	"github.com/Asyadam/PairProjectP1/repository"
)

type ReportHandler struct {
	ReportRepo *repository.ReportRepository
}

func NewReportHandler(reportRepo *repository.ReportRepository) *ReportHandler {
	return &ReportHandler{ReportRepo: reportRepo}
}

func (h *ReportHandler) UserReport() {
	err := h.ReportRepo.UserReport()

	if err != nil {
		fmt.Println("User Report Failed:", err)
		return
	}
}

func (h *ReportHandler) SalesReport() {
	err := h.ReportRepo.SalesReport()

	if err != nil {
		fmt.Println("Sales Report Failed:", err)
		return
	}
}

func (h *ReportHandler) StockReport() {
	err := h.ReportRepo.StockReport()

	if err != nil {
		fmt.Println("Stock Report Failed:", err)
		return
	}
}
