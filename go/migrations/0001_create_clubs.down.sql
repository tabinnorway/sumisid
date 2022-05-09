alter table clubs
    drop constraint fk_contact_person;

alter table people
    drop constraint fk_main_club;

alter table competitions
    drop constraint fk_competition_location;

alter table locations
    drop constraint fk_location_contact_person;

drop table if exists clubs;
drop table if exists people;
drop table if exists competitions;
drop table if exists locations;