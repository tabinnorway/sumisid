alter table diveclubs
    drop constraint fk_contact_person;

alter table people
    drop constraint fk_main_club;

drop table if exists diveclubs;
drop table if exists people;