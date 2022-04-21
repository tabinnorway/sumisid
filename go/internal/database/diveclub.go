package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tabinnorway/sumisid/go/internal/services"
)

func convertRowToDiveClub(dcr DiveClubRow) services.DiveClub {
	return services.DiveClub{
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

type DiveClubRow struct {
	Id                int
	Club_name         string
	Street_address    sql.NullString
	Street_number     sql.NullString
	Zip               sql.NullString
	Phone_number      sql.NullString
	Contact_person_id sql.NullInt32
	Extra_info        sql.NullString
}

func (d *Database) GetAllDiveClub(ctx context.Context) ([]services.DiveClub, error) {
	res := []DiveClubRow{}

	err := d.Client.Select(&res, `select * from diveclubs`)
	if err != nil {
		fmt.Println("Could not retrieve dive clubs ", err.Error())
		return []services.DiveClub{}, err
	}

	retval := []services.DiveClub{}
	for i := range res {
		retval = append(retval, convertRowToDiveClub(res[i]))
	}
	return retval, nil
}

func (d *Database) GetDiveClub(ctx context.Context, id int) (services.DiveClub, error) {
	var clubRow DiveClubRow
	row := d.Client.QueryRowContext(ctx, `select * from diveclubs where id = $1 order by id`, id)
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
		return services.DiveClub{}, fmt.Errorf("error fetching dive club %w", err)
	}
	return convertRowToDiveClub(clubRow), nil
}

func (d *Database) UpdateDiveClub(ctx context.Context, id int, dc services.DiveClub) (services.DiveClub, error) {
	fmt.Println("Going to update diveclub")
	row := DiveClubRow{
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
		`update diveclubs
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
		return services.DiveClub{}, fmt.Errorf("error updating dive club %w", err)
	}
	numRows, err := rows.RowsAffected()
	if err != nil {
		return services.DiveClub{}, fmt.Errorf("error updating dive club %w", err)
	}
	fmt.Println("Number of rows affected by update ", numRows)
	if err != nil {
		return services.DiveClub{}, fmt.Errorf("error updating dive club %w", err)
	}
	return d.GetDiveClub(ctx, id)
}

func (d *Database) DeleteDiveClub(ctx context.Context, id int) error {
	_, err := d.Client.ExecContext(ctx, `delete from diveclubs where id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting dive club %w", err)
	}
	return nil
}

func (d *Database) CreateDiveClub(ctx context.Context, dc services.DiveClub) (services.DiveClub, error) {
	lastInsertedId := 0
	err := d.Client.QueryRow(
		`insert into diveclubs (club_name, street_address, street_number, zip, phone_number, contact_person_id, extra_info)
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
		return services.DiveClub{}, fmt.Errorf("failed to insert diveclub: %w", err)
	}

	return d.GetDiveClub(ctx, lastInsertedId)
}
