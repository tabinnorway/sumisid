create table if not exists clubs (
    id uuid primary key,
    club_name varchar(100) not null,
    email varchar(255),
    phone_number varchar(64),
    main_contact uuid,
    extra_info varchar
);

create table if not exists people (
    id uuid primary key,
    email varchar(255),
    first_name varchar(50) not null,
    middle_name varchar(50),
    last_name varchar(50) not null,
    birth_date date default null,
    is_admin boolean not null default false,
    phone_number varchar(64),
    main_club_id uuid
);

create table if not exists club_contacts (
    id serial primary key,
    club_id uuid not null,
    contact_id uuid not null
);

alter table clubs
    add constraint fk_contact_person
    foreign key (main_contact)
    references people(id);

alter table people
    add constraint fk_main_club
    foreign key(main_club_id)
    references clubs(id);

alter table club_contacts
    add constraint fk_club
    foreign key(club_id)
    references clubs(id);

alter table club_contacts
    add constraint fk_contact
    foreign key(contact_id)
    references people(id);