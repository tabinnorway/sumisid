package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tabinnorway/sumisid/go/internal/services"
)

func convertRowToPerson(pr PersonRow) services.Person {
	return services.Person{
		Id:          pr.Id,
		Email:       pr.Email,
		FirstName:   pr.First_name.String,
		MiddleName:  pr.Middle_name.String,
		LastName:    pr.Last_name.String,
		BirthDate:   pr.Birth_date.Time,
		IsAdmin:     pr.Is_admin.Bool,
		PhoneNumber: pr.Phone_number.String,
		MainClubId:  int(pr.Main_club_id.Int32),
	}
}

type PersonRow struct {
	Id           int
	Email        string
	First_name   sql.NullString
	Middle_name  sql.NullString
	Last_name    sql.NullString
	Birth_date   sql.NullTime
	Is_admin     sql.NullBool
	Phone_number sql.NullString
	Main_club_id sql.NullInt32
}

func (d *Database) GetAllPerson(ctx context.Context) ([]services.Person, error) {
	res := []PersonRow{}

	err := d.Client.Select(&res, `select * from people`)
	if err != nil {
		fmt.Println("Could not retrieve people ", err.Error())
		return []services.Person{}, err
	}

	retval := []services.Person{}
	for i := range res {
		person := convertRowToPerson(res[i])
		person.MainClub, err = d.GetClub(ctx, person.MainClubId)
		if err != nil {
			person.MainClub = services.Club{}
		}
		retval = append(retval, person)
	}
	return retval, nil
}

func (d *Database) GetPerson(ctx context.Context, id int) (services.Person, error) {
	var clubRow PersonRow
	row := d.Client.QueryRowContext(ctx, `select * from people where id = $1 order by last_name`, id)
	err := row.Scan(
		&clubRow.Id,
		&clubRow.Email,
		&clubRow.First_name,
		&clubRow.Middle_name,
		&clubRow.Last_name,
		&clubRow.Birth_date,
		&clubRow.Is_admin,
		&clubRow.Phone_number,
		&clubRow.Main_club_id,
	)
	if err != nil {
		return services.Person{}, fmt.Errorf("error fetching person %w", err)
	}
	retval := convertRowToPerson(clubRow)
	retval.MainClub, err = d.GetClub(ctx, retval.MainClubId)
	return retval, nil
}

func (d *Database) UpdatePerson(ctx context.Context, id int, p services.Person) (services.Person, error) {
	rows, err := d.Client.ExecContext(
		ctx,
		`update people
			set email = $1,
				first_name = $2,
				middle_name = $3,
				last_name = $4,
				birth_date = $5,
				is_admin = $6,
				phone_number = $7,
				main_club_id = $8
			where id = $9`,
		p.Email,
		p.FirstName,
		p.MiddleName,
		p.LastName,
		p.BirthDate,
		p.IsAdmin,
		p.PhoneNumber,
		p.MainClubId,
		id,
	)
	if err != nil {
		return services.Person{}, fmt.Errorf("error updating person %w", err)
	}
	numRows, err := rows.RowsAffected()
	if err != nil {
		return services.Person{}, fmt.Errorf("error updating person %w", err)
	}
	if numRows != 1 {
		return services.Person{}, fmt.Errorf("error updating person got %d rows affected, expected 1", err)
	}
	return d.GetPerson(ctx, id)
}

func (d *Database) DeletePerson(ctx context.Context, id int) error {
	_, err := d.Client.ExecContext(ctx, `delete from people where id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting person %w", err)
	}
	return nil
}

func (d *Database) CreatePerson(ctx context.Context, p services.Person) (services.Person, error) {
	lastInsertedId := 0
	err := d.Client.QueryRow(
		`insert into people (email, first_name, middle_name, last_name, birth_date, is_admin, phone_number, main_club_id)
		 values ($1, $2, $3, $4, $5, $6, $7, $8)
		 returning id`,
		p.Email,
		sql.NullString{String: p.FirstName, Valid: true},
		sql.NullString{String: p.MiddleName, Valid: true},
		sql.NullString{String: p.LastName, Valid: true},
		sql.NullTime{Time: p.BirthDate, Valid: true},
		sql.NullBool{Bool: p.IsAdmin, Valid: true},
		sql.NullString{String: p.PhoneNumber, Valid: true},
		sql.NullInt32{Int32: int32(p.MainClubId), Valid: p.MainClubId > 0},
	).Scan(&lastInsertedId)

	if err != nil {
		return services.Person{}, fmt.Errorf("failed to insert person: %w", err)
	}

	return d.GetPerson(ctx, lastInsertedId)
}
