package main

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

// this file here is where we put all the data stuff for clients

// a struct like this meant to contain all the data we might care about related specifically to a certain client
// it's just dummy stuff for now, really
type client struct {
	ID    int    `json:"id" db:"client_id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

// this is a namespace struct, which i can explain more later if that doesn't make sense
type clientModel struct {
	*sqlx.DB
}

// a global function that's just for convenience
func getClientModel() clientModel {
	return clientModel{db}
}

//here is a very simplistic "get" type function, just to show what models are intended for, really
func (cm clientModel) getClient(clientID int) (client, error) {
	var c client
	err := db.Get(&c,
		`
			SELECT 
				name,
				email
			FROM 
				clients
			WHERE 
				client_id = ?`,
		clientID)
	if err != nill && err != sql.ErrNoRows {
		return c, err
	}
	return c, nil
}

func (cm clientModel) updateClient(c client) error {
	_, err := db.NamedExec(
		`UPDATE 
					clients
				SET 
					name = :name, 
					email = :email 
				WHERE
					client_id = :client_id`,
		c)
}

// very simplistic update function, NOT production ready in any way
func (cm clientModel) deleteClient(clientID int) error {
	_, err := db.Exec(
		`UPDATE
					clients
				SET 
					deleted = true 
				WHERE
					client_id = ?`,
		clientID)
	return err
}
