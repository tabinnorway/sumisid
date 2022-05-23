create table if not exists clubs (
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
    email varchar(128),
    first_name varchar(50) not null,
    middle_name varchar(50),
    last_name varchar(50) not null,
    birth_date date default null,
    is_admin boolean not null default false,
    phone_number varchar(64),
    main_club_id integer
);

create table if not exists locations (
    id serial primary key,
    name varchar(50) not null,
    street_address varchar(255),
    street_number varchar(10),
    zip varchar(10),
    phone_number varchar(64),
    contact_person_id integer
);

create table if not exists competitions (
    id serial primary key,
    name varchar(50) not null,
    location_id int not null,
    start_date date not null,
    end_date date not null,
    arranging_club_id int not null,
    comments text
);

alter table clubs
    add constraint fk_contact_person
    foreign key (contact_person_id)
    references people(id);

alter table people
    add constraint fk_main_club
    foreign key(main_club_id)
    references clubs(id);

alter table competitions
    add constraint fk_competition_location
    foreign key(location_id)
    references locations(id);

alter table competitions
    add constraint fk_competition_arranging_club
    foreign key(arranging_club_id)
    references clubs(id);

alter table locations
    add constraint fk_location_contact_person
    foreign key(contact_person_id)
    references people(id);
    