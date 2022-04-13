create table if not exists diveclubs (
    id serial primary key,
    club_name varchar(100) not null,
    street_address varchar(255),
    street_number varchar(10),
    zip varchar(10),
    phone_number varchar(64),
    contact_person_id integer,
    extra_info varchar
);

create table if not exists people (
    id serial primary key,
    first_name varchar(50),
    middle_name varchar(50),
    last_name varchar(50),
    birth_date datetime default null,
    is_admin boolean not null default false,
    phone_number varchar(64),
    main_club_id integer
);

alter table diveclubs
    add constraint fk_contact_person
    foreign key (contact_person_id)
    references people(id);

alter table people
    add constraint fk_main_club
    foreign key(main_club_id)
    references diveclubs(id);
