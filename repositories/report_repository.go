package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetDailySalesReport() (*models.DailySalesReport, error) {
	var report models.DailySalesReport

	// Get total revenue and total transactions for today
	err := repo.db.QueryRow(`
		SELECT 
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(*) as total_transaksi
		FROM transactions 
		WHERE DATE(created_at) = CURRENT_DATE
	`).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// Get best selling product for today
	var bestSeller models.BestSellerItem
	err = repo.db.QueryRow(`
		SELECT 
			p.name,
			SUM(td.quantity) as qty_terjual
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`).Scan(&bestSeller.Nama, &bestSeller.QtyTerjual)

	if err == sql.ErrNoRows {
		// No transactions today, produk_terlaris will be nil
		report.ProdukTerlaris = nil
	} else if err != nil {
		return nil, err
	} else {
		report.ProdukTerlaris = &bestSeller
	}

	return &report, nil
}
