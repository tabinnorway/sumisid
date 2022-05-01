create table if not exists clubs (
    id serial primary key,
    club_name varchar(100) not null,
    email varchar(255),
    phone_number varchar(64),
    main_contact int,
    extra_info varchar
);

create table if not exists people (
    id serial primary key,
    email varchar(255),
    first_name varchar(50) not null,
    middle_name varchar(50),
    last_name varchar(50) not null,
    birth_date date default null,
    is_admin boolean not null default false,
    phone_number varchar(64),
    main_club_id int
);

create table if not exists club_contacts (
    id serial primary key,
    club_id int not null,
    contact_id int not null
);

create table if not exists club_members (
    id serial primary key,
    club_id int not null,
    member_id int not null
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

alter table club_members
    add constraint fk_club
    foreign key(club_id)
    references clubs(id);

alter table club_members
    add constraint fk_member
    foreign key(member_id)
    references people(id);
