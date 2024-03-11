package repository

import (
	"fmt"
	"github.com/fouched/go-contact-app/internal/models"
)

func SelectContacts() ([]models.Contact, error) {
	rows, err := db.Query("SELECT * FROM contacts")
	// close the rows when function exists
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var c models.Contact
		err := rows.Scan(&c.ID, &c.First, &c.Last, &c.Phone, &c.Email)
		if err != nil {
			fmt.Println("DB query error, ignoring row")
		}
		contacts = append(contacts, c)
	}

	return contacts, err
}

func InsertContact(c models.Contact) (int, error) {
	var id int
	stmt := `INSERT INTO contacts (first, last, phone, email)
    			VALUES($1, $2, $3, $4) returning id`

	err := db.QueryRow(stmt,
		c.First,
		c.Last,
		c.Phone,
		c.Email,
	).Scan(&id)
	fmt.Println(fmt.Sprintf("Inserted contact with id %d", id))

	return id, err
}

func SelectContactById(id int) (models.Contact, error) {
	row := db.QueryRow("SELECT * FROM contacts WHERE id = $1", id)
	var c models.Contact
	err := row.Scan(&c.ID, &c.First, &c.Last, &c.Phone, &c.Email)

	return c, err
}

func UpdateContactById(contact models.Contact) error {
	stmt := `UPDATE contacts SET 
                    first = $2, 
                    last = $3, 
                    phone = $4, 
                    email = $5
    		WHERE id = $1`

	_, err := db.Exec(stmt, contact.ID, contact.First, contact.Last, contact.Phone, contact.Email)

	return err
}
