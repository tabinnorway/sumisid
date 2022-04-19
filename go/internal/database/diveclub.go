package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tabinnorway/sumisid/go/internal/diveclub"
)

func convertRowToDiveClub(dcr DiveClubRow) diveclub.DiveClub {
	return diveclub.DiveClub{
		Id:              dcr.Id,
		Name:            dcr.Name,
		StreetAddress:   dcr.StreetAddress.String,
		StreetNumber:    dcr.StreetNumber.String,
		ZipCode:         dcr.ZipCode.String,
		PhoneNumber:     dcr.PhoneNumber.String,
		ContactPersonId: int(dcr.ContactPersonId.Int32),
		ExtraInfo:       dcr.ExtraInfo.String,
	}
}

type DiveClubRow struct {
	Id              int
	Name            string
	StreetAddress   sql.NullString
	StreetNumber    sql.NullString
	ZipCode         sql.NullString
	PhoneNumber     sql.NullString
	ContactPersonId sql.NullInt32
	ExtraInfo       sql.NullString
}

func (d *Database) GetDiveClub(ctx context.Context, id int) (diveclub.DiveClub, error) {
	var clubRow DiveClubRow
	row := d.Client.QueryRowContext(
		ctx,
		`select id, club_name, street_address, street_number, zip, phone_number, contact_person_id, extra_info
		 from diveclubs
		 where id = $1`,
		id,
	)
	err := row.Scan(
		&clubRow.Id,
		&clubRow.Name,
		&clubRow.StreetAddress,
		&clubRow.StreetNumber,
		&clubRow.ZipCode,
		&clubRow.PhoneNumber,
		&clubRow.ContactPersonId,
		&clubRow.ExtraInfo,
	)
	if err != nil {
		return diveclub.DiveClub{}, fmt.Errorf("error fetching dive club %w", err)
	}
	return convertRowToDiveClub(clubRow), nil
}

func (d *Database) UpdateDiveClub(ctx context.Context, id int, dc diveclub.DiveClub) (diveclub.DiveClub, error) {
	row := DiveClubRow{
		Id:              id,
		Name:            dc.Name,
		StreetAddress:   sql.NullString{String: dc.StreetAddress, Valid: true},
		StreetNumber:    sql.NullString{String: dc.StreetAddress, Valid: true},
		ZipCode:         sql.NullString{String: dc.StreetAddress, Valid: true},
		PhoneNumber:     sql.NullString{String: dc.StreetAddress, Valid: true},
		ContactPersonId: sql.NullInt32{Int32: int32(dc.ContactPersonId), Valid: dc.ContactPersonId > 0},
		ExtraInfo:       sql.NullString{String: dc.StreetAddress, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`update diveclubs
			set club_name = name,
				street_address = :streetaddress,
				street_number = :streetnumber,
				zip = :zipcode,
				phone_number = :phonenumber,
				contact_person_id = :contactpersonid,
				extra_info = extrainfo
			where id = :id`,
		row,
	)
	if err != nil {
		return diveclub.DiveClub{}, fmt.Errorf("error deleting dive club %w", err)
	}
	if err := rows.Close(); err != nil {
		return diveclub.DiveClub{}, fmt.Errorf("error closing diveclubs row %w", err)
	}
	return d.GetDiveClub(ctx, dc.Id)
}

func (d *Database) DeleteDiveClub(ctx context.Context, id int) error {
	_, err := d.Client.ExecContext(ctx, `delete from diveclubs where id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting dive club %w", err)
	}
	return nil
}
func (d *Database) CreateDiveClub(ctx context.Context, dc diveclub.DiveClub) (diveclub.DiveClub, error) {
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
		return diveclub.DiveClub{}, fmt.Errorf("failed to insert diveclub: %w", err)
	}

	return d.GetDiveClub(ctx, lastInsertedId)
}
