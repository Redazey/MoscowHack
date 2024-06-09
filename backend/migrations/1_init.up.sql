create table users
(
    id         serial
        primary key,
    surname    text not null,
    name       text not null,
    patronymic text,
    age        integer,
    photourl   text,
    email      text not null,
    password   text not null,
    push       boolean
);

create table roles
(
    id   serial
        primary key,
    name text not null
);

create table userroles
(
    id     serial
        primary key,
    userid integer
        constraint userroles_users_id_fk
            references users,
    roleid integer
        constraint userroles_roles_id_fk
            references roles
);

create table categories
(
    id   serial
        primary key,
    name text not null
);

create table "categoriesVacancies"
(
    id   serial
        primary key,
    name text not null
);

create table educations
(
    id               serial
        primary key,
    name             text not null,
    "placeEducation" text
);

create table "workExperience"
(
    id                 serial
        primary key,
    "startDate"        date not null,
    "endDate"          date,
    post               text not null,
    company            text not null,
    "projectName"      text,
    "projectRole"      text,
    "structureCommand" text,
    "completedTasks"   text,
    environment        text not null,
    instruments        text not null,
    technologies       text not null
);

create table requirements
(
    id                 serial
        primary key,
    "educationID"      integer not null
        references educations,
    "workExperienceID" integer not null
        references "workExperience"
);

create table skills
(
    id   serial
        primary key,
    name text not null
);

create table "requirementsWorkExperience"
(
    id                 serial
        primary key,
    "requirementsID"   integer not null
        references requirements,
    "workExperienceID" integer not null
        references "workExperience"
);

create table "requirementsSkills"
(
    id               serial
        primary key,
    "requirementsID" integer not null
        references requirements,
    "skillsID"       integer not null
        references skills
);

create table specializations
(
    id                    serial
        primary key,
    name                  text    not null,
    "categoryVacanciesID" integer not null
        references "categoriesVacancies"
);

create table vacancies
(
    id                    serial
        primary key,
    name                  text    not null,
    "departmentCompany"   text    not null,
    description           text,
    "categoryVacanciesID" integer not null
        references "categoriesVacancies",
    "requirementsID"      integer not null
        references requirements,
    "workingConditionsID" integer,
    "geolocationCompany"  text
);

create table news
(
    id       serial
        primary key,
    title    text      not null,
    text     text      not null,
    datetime timestamp not null
);

create table "categoriesNews"
(
    id           serial
        primary key,
    "newsID"     integer not null
        references news,
    "categoryID" integer not null
        references categories
);

create table genders
(
    id   serial
        primary key,
    name text not null
);

create table positions
(
    id   serial
        primary key,
    name text not null
);

create table portfolios
(
    id   serial
        primary key,
    name text not null,
    link text not null
);

create table socials
(
    id   serial
        primary key,
    link text not null
);

create table resumes
(
    id                 serial
        primary key,
    "positionID"       integer not null
        references positions,
    "experienceNumber" integer not null,
    "portfolioID"      integer
        references portfolios,
    "softSkills"       text,
    "specializationID" integer not null
        references specializations,
    "workExperienceID" integer
        references "workExperience",
    "skillsID"         integer
        references skills,
    "educationID"      integer
        references educations
);

create table "jobSeekers"
(
    id                 serial
        primary key,
    userid             integer UNIQUE REFERENCES users(id),
    "genderID"         integer not null
        references genders,
    "socialID"         integer
        references socials,
    "resumeID"         integer
        references resumes,
    "specializationID" integer not null
        references specializations
);

create unique index jobseekers_userid_uindex
    on "jobSeekers" (userid);

create table recruiters
(
    id         serial
        primary key,
    userid             integer UNIQUE REFERENCES users(id)
);

create table interviews
(
    id            serial
        primary key,
    date          date    not null,
    "recruiterID" integer not null
        references recruiters,
    "jobSeekerID" integer not null
        references "jobSeekers",
    link          text    not null
);

create table "workSchedules"
(
    id            serial
        primary key,
    "recruiterID" integer not null
        references recruiters,
    "startTime"   time    not null,
    "endTime"     time    not null,
    "dayOfWeek"   integer
);

create table appointments
(
    id              serial
        primary key,
    "recruiterID"   integer                  not null
        references recruiters,
    title           text                     not null,
    description     text,
    "startDateTime" timestamp with time zone not null,
    "endDateTime"   timestamp with time zone not null
);

create table "daysOff"
(
    id            serial
        primary key,
    "recruiterID" integer not null
        references recruiters,
    date          date    not null
);