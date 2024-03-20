package repository

import (
	"fmt"
	"github.com/fouched/go-contact-app/internal/models"
	"time"
)

func SelectContacts(q string) ([]models.Contact, error) {
	s := "SELECT * FROM contacts c "
	if q != "" {
		s += "WHERE UPPER(c.first) LIKE UPPER('%" + q + "%')" +
			" OR UPPER(c.last) LIKE UPPER('%" + q + "%') "
	}
	s += "ORDER BY c.last, c.first"

	rows, err := db.Query(s)
	// close the rows when function exists
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var c models.Contact
		err := rows.Scan(&c.ID, &c.First, &c.Last, &c.Phone, &c.Email, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			fmt.Println("DB query error, ignoring row")
			fmt.Println(err.Error())
		}
		contacts = append(contacts, c)
	}

	return contacts, err
}

func InsertContact(c models.Contact) (int, error) {
	var id int
	stmt := `INSERT INTO contacts (first, last, phone, email, created_at, updated_at)
    			VALUES($1, $2, $3, $4, $5, $6) returning id`

	err := db.QueryRow(stmt,
		c.First,
		c.Last,
		c.Phone,
		c.Email,
		time.Now(),
		time.Now(),
	).Scan(&id)
	fmt.Println(fmt.Sprintf("Inserted contact with id %d", id))

	return id, err
}

func SelectContactById(id int) (models.Contact, error) {
	row := db.QueryRow("SELECT * FROM contacts WHERE id = $1", id)
	var c models.Contact
	err := row.Scan(&c.ID, &c.First, &c.Last, &c.Phone, &c.Email, &c.CreatedAt, &c.UpdatedAt)

	return c, err
}

func UpdateContactById(contact models.Contact) error {
	stmt := `UPDATE contacts SET 
                    first = $2, 
                    last = $3, 
                    phone = $4, 
                    email = $5,
                    updated_at = $6
    		WHERE id = $1`

	_, err := db.Exec(stmt, contact.ID, contact.First, contact.Last, contact.Phone, contact.Email, time.Now())

	return err
}

func DeleteContactById(id int) error {
	stmt := `DELETE FROM contacts WHERE id = $1`
	_, err := db.Exec(stmt, id)

	return err
}
