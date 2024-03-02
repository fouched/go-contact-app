package repo

import (
	"fmt"
	"github.com/fouched/go-contact-app/app/models"
)

func SelectContacts() (error, []models.Contact) {
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

	return err, contacts
}

func AddContact(c models.Contact) (error, int) {
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

	return err, id
}
