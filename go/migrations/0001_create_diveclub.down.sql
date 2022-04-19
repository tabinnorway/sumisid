alter table diveclubs
    drop constraint fk_contact_person;

alter table people
    drop constraint fk_main_club;

alter table competitions
    crop constraint fk_competition_location;

alter table locations
    crop constraint fk_location_contact_person;

drop table if exists diveclubs;
drop table if exists people;
drop table if exists competitions;
drop table if exists locations;