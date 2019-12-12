package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/andrewesteves/tapagguapi/model"
	"github.com/lib/pq"
)

// ReceiptPostgresRepository struct
type ReceiptPostgresRepository struct {
	Conn *sql.DB
}

// NewReceiptPostgresRepository new repository
func NewReceiptPostgresRepository(Conn *sql.DB) ReceiptContractRepository {
	return &ReceiptPostgresRepository{Conn}
}

// All receipts
func (r ReceiptPostgresRepository) All(user model.User, values map[string]string) ([]model.Receipt, error) {
	var receipts []model.Receipt
	var rs *sql.Rows
	var err error
	args := []interface{}{user.ID}
	i := 2

	query := "SELECT id, category_id, company_id, title, tax, discount, extra, total, url, access_key, issued_at, created_at, updated_at FROM receipts WHERE user_id = $1 "

	if len(values) == 0 {
		query += "AND (EXTRACT (MONTH FROM issued_at) = EXTRACT (MONTH FROM CURRENT_DATE)) AND (EXTRACT (YEAR FROM issued_at) = EXTRACT (YEAR FROM CURRENT_DATE)) "
	} else {
		if ct, ok := values["category"]; ok {
			args = append(args, ct)
			query += fmt.Sprintf("AND category_id = $%d ", i)
			i++
		}
		if mt, ok := values["month"]; ok {
			args = append(args, mt)
			query += fmt.Sprintf("AND (EXTRACT (MONTH FROM issued_at) = $%d) ", i)
			i++
		}
		if yr, ok := values["year"]; ok {
			args = append(args, yr)
			query += fmt.Sprintf("AND (EXTRACT (YEAR FROM issued_at) = $%d) ", i)
		}
	}

	query += "ORDER BY created_at DESC"

	rs, err = r.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	for rs.Next() {
		var receipt model.Receipt
		err = rs.Scan(&receipt.ID, &receipt.Category.ID, &receipt.Company.ID, &receipt.Title, &receipt.Tax, &receipt.Discount, &receipt.Extra, &receipt.Total, &receipt.URL, &receipt.AccessKey, &receipt.IssuedAt, &receipt.CreatedAt, &receipt.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categoryRepo := NewCategoryPostgresRepository(r.Conn)
		receipt.Category, _ = categoryRepo.Find(receipt.Category.ID)
		companyRepo := NewCompanyPostgresRepository(r.Conn)
		receipt.Company, _ = companyRepo.Find(receipt.Company.ID)
		itemRepo := NewItemPostgresRepository(r.Conn)
		receipt.Items, _ = itemRepo.All(receipt.ID)
		if len(receipt.Items) < 1 {
			receipt.Items = make([]model.Item, 0)
		}
		receipts = append(receipts, receipt)
	}

	return receipts, nil
}

// Find a receipt
func (r ReceiptPostgresRepository) Find(id int64) (model.Receipt, error) {
	var receipt model.Receipt
	err := r.Conn.QueryRow("SELECT id, category_id, company_id, title, tax, discount, extra, total, url, access_key, issued_at, created_at, updated_at FROM receipts WHERE id = $1", id).Scan(&receipt.ID, &receipt.Category.ID, &receipt.Company.ID, &receipt.Title, &receipt.Tax, &receipt.Discount, &receipt.Extra, &receipt.Total, &receipt.URL, &receipt.AccessKey, &receipt.IssuedAt, &receipt.CreatedAt, &receipt.UpdatedAt)
	categoryRepo := NewCategoryPostgresRepository(r.Conn)
	receipt.Category, _ = categoryRepo.Find(receipt.Category.ID)
	companyRepo := NewCompanyPostgresRepository(r.Conn)
	receipt.Company, _ = companyRepo.Find(receipt.Company.ID)
	itemRepo := NewItemPostgresRepository(r.Conn)
	receipt.Items, _ = itemRepo.All(receipt.ID)
	if len(receipt.Items) < 1 {
		receipt.Items = make([]model.Item, 0)
	}
	if err != nil {
		return model.Receipt{}, err
	}
	return receipt, nil
}

// Store a receipt
func (r ReceiptPostgresRepository) Store(receipt model.Receipt) (model.Receipt, error) {
	var company model.Company
	var category model.Category
	var err error
	lastInsertID := 0

	receipt.Company.User = receipt.User
	receipt.Category.User = receipt.User

	companyRepo := NewCompanyPostgresRepository(r.Conn)
	company, err = companyRepo.FindBy(receipt.Company, "title", receipt.Company.Title)
	if company.Title == "" {
		receipt.Company, err = companyRepo.Store(receipt.Company)
		if err != nil {
			return model.Receipt{}, err
		}
	} else {
		receipt.Company.ID = company.ID
	}

	categoryRepo := NewCategoryPostgresRepository(r.Conn)
	category, err = categoryRepo.FindBy(receipt.Category, "title", receipt.Category.Title)
	if category.Title == "" {
		receipt.Category, err = categoryRepo.Store(model.Category{
			User:  receipt.User,
			Title: receipt.Category.Title,
			Icon:  receipt.Category.Icon,
		})
		if err != nil {
			return model.Receipt{}, err
		}
	} else {
		receipt.Category.ID = category.ID
	}

	err = r.Conn.QueryRow("INSERT INTO receipts (category_id, company_id, user_id, title, tax, discount, extra, total, url, access_key, issued_at, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,now(),now()) RETURNING id", receipt.Category.ID, receipt.Company.ID, receipt.User.ID, receipt.Title, receipt.Tax, receipt.Discount, receipt.Extra, receipt.Total, receipt.URL, receipt.AccessKey, receipt.IssuedAt).Scan(&lastInsertID)
	if err != nil {
		return model.Receipt{}, err
	}
	receipt.ID = int64(lastInsertID)

	if len(receipt.Items) > 0 {
		txn, err := r.Conn.Begin()
		if err != nil {
			return model.Receipt{}, err
		}
		stmt, err := txn.Prepare(pq.CopyIn("items", "receipt_id", "title", "price", "quantity", "total", "tax", "measure", "created_at", "updated_at"))
		if err != nil {
			return model.Receipt{}, err
		}
		for _, item := range receipt.Items {
			_, err = stmt.Exec(receipt.ID, item.Title, item.Price, item.Quantity, item.Total, item.Tax, item.Measure, time.Now(), time.Now())
			if err != nil {
				return model.Receipt{}, err
			}
		}
		_, err = stmt.Exec()
		if err != nil {
			return model.Receipt{}, err
		}
		err = stmt.Close()
		if err != nil {
			return model.Receipt{}, err
		}
		err = txn.Commit()
		if err != nil {
			return model.Receipt{}, err
		}
	}

	return receipt, nil
}

// Update a receipt
func (r ReceiptPostgresRepository) Update(receipt model.Receipt) (model.Receipt, error) {
	rcpt, err := r.Find(receipt.ID)
	if err != nil {
		return model.Receipt{}, err
	}
	if rcpt.ID < 1 {
		return model.Receipt{}, errors.New("Cant't find this receipt id")
	}
	rs, err := r.Conn.Prepare("UPDATE receipts SET title = $1, tax = $2, discount = $3, extra = $4, total = $5, url = $6, access_key = $7, issued_at = $8, updated_at = now() WHERE id = $9")
	if err != nil {
		return model.Receipt{}, err
	}
	rs.Exec(receipt.Title, receipt.Tax, receipt.Discount, receipt.Extra, receipt.Total, receipt.URL, receipt.AccessKey, receipt.IssuedAt, receipt.ID)
	return receipt, nil
}

// Destroy an receipt
func (r ReceiptPostgresRepository) Destroy(id int64) (model.Receipt, error) {
	rcpt, err := r.Find(id)
	if err != nil {
		return model.Receipt{}, err
	}
	if rcpt.ID < 1 {
		return model.Receipt{}, errors.New("Cant't find this receipt id")
	}
	rs, err := r.Conn.Prepare("DELETE FROM receipts WHERE id = $1")
	if err != nil {
		return model.Receipt{}, err
	}
	rs.Exec(id)
	return rcpt, nil
}

// FindManyBy receipt by field name
func (r ReceiptPostgresRepository) FindManyBy(field string, value interface{}) ([]model.Receipt, error) {
	var receipts []model.Receipt
	rs, err := r.Conn.Query(fmt.Sprintf("SELECT id, title, tax, discount, extra, total, url, access_key, issued_at, created_at, updated_at FROM receipts WHERE %s = $1", field), value)
	if err != nil {
		return nil, err
	}

	for rs.Next() {
		var receipt model.Receipt
		err = rs.Scan(&receipt.ID, &receipt.Title, &receipt.Tax, &receipt.Discount, &receipt.Extra, &receipt.Total, &receipt.URL, &receipt.AccessKey, &receipt.IssuedAt, &receipt.CreatedAt, &receipt.UpdatedAt)
		if err != nil {
			return nil, err
		}
		itemRepo := NewItemPostgresRepository(r.Conn)
		receipt.Items, _ = itemRepo.All(receipt.ID)
		if len(receipt.Items) < 1 {
			receipt.Items = make([]model.Item, 0)
		}
		receipts = append(receipts, receipt)
	}

	return receipts, nil
}

// RetrieveStore a receipt
func (r ReceiptPostgresRepository) RetrieveStore(receipt model.Receipt) (model.Receipt, error) {
	var company model.Company
	var category model.Category
	var err error
	lastInsertID := 0

	receipt.Company.User = receipt.User
	receipt.Category.User = receipt.User

	companyRepo := NewCompanyPostgresRepository(r.Conn)
	company, err = companyRepo.FindBy(receipt.Company, "cnpj", receipt.Company.CNPJ)
	if company.CNPJ == "" {
		receipt.Company, err = companyRepo.Store(receipt.Company)
		if err != nil {
			return model.Receipt{}, err
		}
	} else {
		receipt.Company.ID = company.ID
	}

	categoryRepo := NewCategoryPostgresRepository(r.Conn)
	category, err = categoryRepo.FindBy(receipt.Category, "title", "Geral")
	if category.Title == "" {
		receipt.Category, err = categoryRepo.Store(model.Category{
			User:  receipt.User,
			Title: "Geral",
			Icon:  "all",
		})
		if err != nil {
			return model.Receipt{}, err
		}
	} else {
		receipt.Category.ID = category.ID
	}

	err = r.Conn.QueryRow("INSERT INTO receipts (category_id, company_id, user_id, title, tax, discount, extra, total, url, access_key, issued_at, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,now(),now()) RETURNING id", receipt.Category.ID, receipt.Company.ID, receipt.User.ID, receipt.Title, receipt.Tax, receipt.Discount, receipt.Extra, receipt.Total, receipt.URL, receipt.AccessKey, receipt.IssuedAt).Scan(&lastInsertID)
	if err != nil {
		return model.Receipt{}, err
	}
	receipt.ID = int64(lastInsertID)

	if len(receipt.Items) > 0 {
		txn, err := r.Conn.Begin()
		if err != nil {
			return model.Receipt{}, err
		}
		stmt, err := txn.Prepare(pq.CopyIn("items", "receipt_id", "title", "price", "quantity", "total", "tax", "measure", "created_at", "updated_at"))
		if err != nil {
			return model.Receipt{}, err
		}
		for _, item := range receipt.Items {
			_, err = stmt.Exec(receipt.ID, item.Title, item.Price, item.Quantity, item.Total, item.Tax, item.Measure, time.Now(), time.Now())
			if err != nil {
				return model.Receipt{}, err
			}
		}
		_, err = stmt.Exec()
		if err != nil {
			return model.Receipt{}, err
		}
		err = stmt.Close()
		if err != nil {
			return model.Receipt{}, err
		}
		err = txn.Commit()
		if err != nil {
			return model.Receipt{}, err
		}
	}

	return receipt, nil
}

// Count receipts
func (r ReceiptPostgresRepository) Count(user model.User) (int, error) {
	var count int
	err := r.Conn.QueryRow("SELECT COUNT(*) FROM receipts WHERE user_id = $1", user.ID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
