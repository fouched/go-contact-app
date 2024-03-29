package repo

import (
	"fmt"
	"github.com/fouched/go-contact-app/internal/models"
	"strconv"
	"time"
)

func SelectContacts(q string, offset int, pageSize int) ([]models.Contact, error) {

	s := "SELECT * FROM contacts c "
	if q != "" {
		s += "WHERE UPPER(c.first) LIKE UPPER('%" + q + "%')" +
			" OR UPPER(c.last) LIKE UPPER('%" + q + "%') "
	}
	s += "ORDER BY c.last, c.first "
	s += "LIMIT " + strconv.Itoa(pageSize)
	s += " OFFSET " + strconv.Itoa(offset)

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

func SelectContactCount(q string) (int, error) {

	var c int
	s := "SELECT COUNT(c.id) FROM contacts c "
	if q != "" {
		s += "WHERE UPPER(c.first) LIKE UPPER('%" + q + "%')" +
			" OR UPPER(c.last) LIKE UPPER('%" + q + "%') "
	}
	err := db.QueryRow(s).Scan(&c)

	return c, err
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

func EmailExists(email string, id int) (bool, error) {
	rows, err := db.Query("SELECT id FROM contacts WHERE id <> $1 AND UPPER(email) = UPPER($2)", id, email)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
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
