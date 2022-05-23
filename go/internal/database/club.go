package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tabinnorway/sumisid/go/internal/services"
)

func convertRowToClub(dcr ClubRow) services.Club {
	return services.Club{
		Id:              dcr.Id,
		Name:            dcr.Club_name,
		StreetAddress:   dcr.Street_address.String,
		StreetNumber:    dcr.Street_number.String,
		ZipCode:         dcr.Zip.String,
		PhoneNumber:     dcr.Phone_number.String,
		ContactPersonId: int(dcr.Contact_person_id.Int32),
		ExtraInfo:       dcr.Extra_info.String,
	}
}

type ClubRow struct {
	Id                int
	Club_name         string
	Street_address    sql.NullString
	Street_number     sql.NullString
	Zip               sql.NullString
	Phone_number      sql.NullString
	Contact_person_id sql.NullInt32
	Extra_info        sql.NullString
}

func (d *Database) GetAllClub(ctx context.Context) ([]services.Club, error) {
	res := []ClubRow{}

	err := d.Client.Select(&res, `select * from clubs`)
	if err != nil {
		fmt.Println("Could not retrieve clubs ", err.Error())
		return []services.Club{}, err
	}

	retval := []services.Club{}
	for i := range res {
		retval = append(retval, convertRowToClub(res[i]))
	}
	return retval, nil
}

func (d *Database) GetClub(ctx context.Context, id int) (services.Club, error) {
	var clubRow ClubRow
	row := d.Client.QueryRowContext(ctx, `select * from clubs where id = $1 order by id`, id)
	err := row.Scan(
		&clubRow.Id,
		&clubRow.Club_name,
		&clubRow.Street_address,
		&clubRow.Street_number,
		&clubRow.Zip,
		&clubRow.Phone_number,
		&clubRow.Contact_person_id,
		&clubRow.Extra_info,
	)
	if err != nil {
		return services.Club{}, fmt.Errorf("error fetching club %w", err)
	}
	return convertRowToClub(clubRow), nil
}

func (d *Database) UpdateClub(ctx context.Context, id int, dc services.Club) (services.Club, error) {
	row := ClubRow{
		Id:                id,
		Club_name:         dc.Name,
		Street_address:    sql.NullString{String: dc.StreetAddress, Valid: true},
		Street_number:     sql.NullString{String: dc.StreetNumber, Valid: true},
		Zip:               sql.NullString{String: dc.ZipCode, Valid: true},
		Phone_number:      sql.NullString{String: dc.PhoneNumber, Valid: true},
		Contact_person_id: sql.NullInt32{Int32: int32(dc.ContactPersonId), Valid: dc.ContactPersonId > 0},
		Extra_info:        sql.NullString{String: dc.ExtraInfo, Valid: true},
	}
	rows, err := d.Client.ExecContext(
		ctx,
		`update clubs
			set club_name = $1,
				street_address = $2,
				street_number = $3,
				zip = $4,
				phone_number = $5,
				contact_person_id = $6,
				extra_info = $7
			where id = $8`,
		row.Club_name,
		row.Street_address,
		row.Street_number,
		row.Zip,
		row.Phone_number,
		row.Contact_person_id,
		row.Extra_info,
		id,
	)
	if err != nil {
		return services.Club{}, fmt.Errorf("error updating club %w", err)
	}
	numRows, err := rows.RowsAffected()
	if err != nil {
		return services.Club{}, fmt.Errorf("error updating club %w", err)
	}
	if numRows != 1 {
		return services.Club{}, fmt.Errorf("error updating club %d rows updated, expected 1", numRows)
	}
	return d.GetClub(ctx, id)
}

func (d *Database) DeleteClub(ctx context.Context, id int) error {
	_, err := d.Client.ExecContext(ctx, `delete from clubs where id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting club %w", err)
	}
	return nil
}

func (d *Database) CreateClub(ctx context.Context, dc services.Club) (services.Club, error) {
	lastInsertedId := 0
	err := d.Client.QueryRow(
		`insert into clubs (club_name, street_address, street_number, zip, phone_number, contact_person_id, extra_info)
		 values ($1, $2, $3, $4, $5, $6, $7)
		 returning id`,
		dc.Name,
		sql.NullString{String: dc.StreetAddress, Valid: true},
		sql.NullString{String: dc.StreetNumber, Valid: true},
		sql.NullString{String: dc.ZipCode, Valid: true},
		sql.NullString{String: dc.PhoneNumber, Valid: true},
		sql.NullInt32{Int32: int32(dc.ContactPersonId), Valid: dc.ContactPersonId > 0},
		sql.NullString{String: dc.ExtraInfo, Valid: true},
	).Scan(&lastInsertedId)

	if err != nil {
		return services.Club{}, fmt.Errorf("failed to insert club: %w", err)
	}

	return d.GetClub(ctx, lastInsertedId)
}
