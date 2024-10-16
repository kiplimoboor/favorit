package database

import (
	"errors"
	"fmt"

	"github.com/kiplimoboor/favorit/models"
)

func (db *SQLiteDB) CreateGuest(gr models.GuestRequest) error {
	stmt := "SELECT email,phone FROM guests WHERE email=? OR phone=?"
	row := db.db.QueryRow(stmt, gr.Email, gr.Phone)
	var existingEmail, existingPhone string
	row.Scan(&existingEmail, &existingPhone)
	if existingEmail == gr.Email {
		return fmt.Errorf("guest with email %s already registered", gr.Email)
	}
	if existingPhone == gr.Phone {
		return fmt.Errorf("guest with phone %s already registered", gr.Phone)
	}
	g := models.NewGuest(gr)
	stmt = "INSERT INTO guests(first_name,last_name,email,phone,address,created_at) VALUES(?,?,?,?,?,?)"
	_, err := db.db.Exec(stmt, g.FirstName, g.LastName, g.Email, g.Phone, g.Address, g.CreatedAt)
	return err
}

func (db *SQLiteDB) GetGuestBy(key, value string) (*models.Guest, error) {
	stmt := fmt.Sprintf("SELECT * FROM guests WHERE %s=?;", key)
	row := db.db.QueryRow(stmt, value)
	g := new(models.Guest)
	err := row.Scan(&g.Id, &g.FirstName, &g.LastName, &g.Email, &g.Phone, &g.Address, &g.CreatedAt)
	if err != nil {
		return nil, errors.New("guest not found")
	}
	return g, nil
}

func (db *SQLiteDB) UpdateGuest(email, field, newVal string) error {
	if _, err := db.GetGuestBy("email", email); err != nil {
		return err
	}
	stmt := fmt.Sprintf("UPDATE guests SET %s=? WHERE email=?", field)
	_, err := db.db.Exec(stmt, newVal, email)
	return err
}
