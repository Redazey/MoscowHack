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

alter table users
    owner to hack;

grant delete, insert, select, update on users to hack;

create table roles
(
    id   serial
        primary key,
    name text not null
);

alter table roles
    owner to hack;

grant delete, insert, select, update on roles to hack;

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

alter table userroles
    owner to hack;

grant delete, insert, select, update on userroles to hack;

create table categories
(
    id   serial
        primary key,
    name text not null
);

alter table categories
    owner to hack;

grant delete, insert, select, update on categories to hack;

create table "categoriesVacancies"
(
    id   serial
        primary key,
    name text not null
);

alter table "categoriesVacancies"
    owner to hack;

grant delete, insert, select, update on "categoriesVacancies" to hack;

create table educations
(
    id               serial
        primary key,
    name             text not null,
    "placeEducation" text
);

alter table educations
    owner to hack;

grant delete, insert, select, update on educations to hack;

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

alter table "workExperience"
    owner to hack;

grant delete, insert, select, update on "workExperience" to hack;

create table requirements
(
    id                 serial
        primary key,
    "educationID"      integer not null
        references educations,
    "workExperienceID" integer not null
        references "workExperience"
);

alter table requirements
    owner to hack;

grant delete, insert, select, update on requirements to hack;

create table skills
(
    id   serial
        primary key,
    name text not null
);

alter table skills
    owner to hack;

grant delete, insert, select, update on skills to hack;

create table "requirementsWorkExperience"
(
    id                 serial
        primary key,
    "requirementsID"   integer not null
        references requirements,
    "workExperienceID" integer not null
        references "workExperience"
);

alter table "requirementsWorkExperience"
    owner to hack;

grant delete, insert, select, update on "requirementsWorkExperience" to hack;

create table "requirementsSkills"
(
    id               serial
        primary key,
    "requirementsID" integer not null
        references requirements,
    "skillsID"       integer not null
        references skills
);

alter table "requirementsSkills"
    owner to hack;

grant delete, insert, select, update on "requirementsSkills" to hack;

create table specializations
(
    id                    serial
        primary key,
    name                  text    not null,
    "categoryVacanciesID" integer not null
        references "categoriesVacancies"
);

alter table specializations
    owner to hack;

grant delete, insert, select, update on specializations to hack;

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

alter table vacancies
    owner to hack;

grant delete, insert, select, update on vacancies to hack;

create table news
(
    id       serial
        primary key,
    title    text      not null,
    text     text      not null,
    datetime timestamp not null
);

alter table news
    owner to hack;

grant delete, insert, select, update on news to hack;

create table "categoriesNews"
(
    id           serial
        primary key,
    "newsID"     integer not null
        references news,
    "categoryID" integer not null
        references categories
);

alter table "categoriesNews"
    owner to hack;

grant delete, insert, select, update on "categoriesNews" to hack;

create table genders
(
    id   serial
        primary key,
    name text not null
);

alter table genders
    owner to hack;

grant delete, insert, select, update on genders to hack;

create table positions
(
    id   serial
        primary key,
    name text not null
);

alter table positions
    owner to hack;

grant delete, insert, select, update on positions to hack;

create table portfolios
(
    id   serial
        primary key,
    name text not null,
    link text not null
);

alter table portfolios
    owner to hack;

grant delete, insert, select, update on portfolios to hack;

create table socials
(
    id   serial
        primary key,
    link text not null
);

alter table socials
    owner to hack;

grant delete, insert, select, update on socials to hack;

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

alter table resumes
    owner to hack;

grant delete, insert, select, update on resumes to hack;

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

alter table "jobSeekers"
    owner to hack;

create unique index jobseekers_userid_uindex
    on "jobSeekers" (userid);

grant delete, insert, select, update on "jobSeekers" to hack;

create table recruiters
(
    id         serial
        primary key,
    userid             integer UNIQUE REFERENCES users(id)
);

alter table recruiters
    owner to hack;

grant delete, insert, select, update on recruiters to hack;

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

alter table interviews
    owner to hack;

grant delete, insert, select, update on interviews to hack;

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

alter table "workSchedules"
    owner to hack;

grant delete, insert, select, update on "workSchedules" to hack;

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

alter table appointments
    owner to hack;

grant delete, insert, select, update on appointments to hack;

create table "daysOff"
(
    id            serial
        primary key,
    "recruiterID" integer not null
        references recruiters,
    date          date    not null
);

alter table "daysOff"
    owner to hack;

grant delete, insert, select, update on "daysOff" to hack;

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

alter table users
    owner to hack;

grant delete, insert, select, update on users to hack;

create table roles
(
    id   serial
        primary key,
    name text not null
);

alter table roles
    owner to hack;

grant delete, insert, select, update on roles to hack;

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

alter table userroles
    owner to hack;

grant delete, insert, select, update on userroles to hack;

create table categories
(
    id   serial
        primary key,
    name text not null
);

alter table categories
    owner to hack;

grant delete, insert, select, update on categories to hack;

create table "categoriesVacancies"
(
    id   serial
        primary key,
    name text not null
);

alter table "categoriesVacancies"
    owner to hack;

grant delete, insert, select, update on "categoriesVacancies" to hack;

create table educations
(
    id               serial
        primary key,
    name             text not null,
    "placeEducation" text
);

alter table educations
    owner to hack;

grant delete, insert, select, update on educations to hack;

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

alter table "workExperience"
    owner to hack;

grant delete, insert, select, update on "workExperience" to hack;

create table requirements
(
    id                 serial
        primary key,
    "educationID"      integer not null
        references educations,
    "workExperienceID" integer not null
        references "workExperience"
);

alter table requirements
    owner to hack;

grant delete, insert, select, update on requirements to hack;

create table skills
(
    id   serial
        primary key,
    name text not null
);

alter table skills
    owner to hack;

grant delete, insert, select, update on skills to hack;

create table "requirementsWorkExperience"
(
    id                 serial
        primary key,
    "requirementsID"   integer not null
        references requirements,
    "workExperienceID" integer not null
        references "workExperience"
);

alter table "requirementsWorkExperience"
    owner to hack;

grant delete, insert, select, update on "requirementsWorkExperience" to hack;

create table "requirementsSkills"
(
    id               serial
        primary key,
    "requirementsID" integer not null
        references requirements,
    "skillsID"       integer not null
        references skills
);

alter table "requirementsSkills"
    owner to hack;

grant delete, insert, select, update on "requirementsSkills" to hack;

create table specializations
(
    id                    serial
        primary key,
    name                  text    not null,
    "categoryVacanciesID" integer not null
        references "categoriesVacancies"
);

alter table specializations
    owner to hack;

grant delete, insert, select, update on specializations to hack;

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

alter table vacancies
    owner to hack;

grant delete, insert, select, update on vacancies to hack;

create table news
(
    id       serial
        primary key,
    title    text      not null,
    text     text      not null,
    datetime timestamp not null
);

alter table news
    owner to hack;

grant delete, insert, select, update on news to hack;

create table "categoriesNews"
(
    id           serial
        primary key,
    "newsID"     integer not null
        references news,
    "categoryID" integer not null
        references categories
);

alter table "categoriesNews"
    owner to hack;

grant delete, insert, select, update on "categoriesNews" to hack;

create table genders
(
    id   serial
        primary key,
    name text not null
);

alter table genders
    owner to hack;

grant delete, insert, select, update on genders to hack;

create table positions
(
    id   serial
        primary key,
    name text not null
);

alter table positions
    owner to hack;

grant delete, insert, select, update on positions to hack;

create table portfolios
(
    id   serial
        primary key,
    name text not null,
    link text not null
);

alter table portfolios
    owner to hack;

grant delete, insert, select, update on portfolios to hack;

create table socials
(
    id   serial
        primary key,
    link text not null
);

alter table socials
    owner to hack;

grant delete, insert, select, update on socials to hack;

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

alter table resumes
    owner to hack;

grant delete, insert, select, update on resumes to hack;

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

alter table "jobSeekers"
    owner to hack;

create unique index jobseekers_userid_uindex
    on "jobSeekers" (userid);

grant delete, insert, select, update on "jobSeekers" to hack;

create table recruiters
(
    id         serial
        primary key,
    userid             integer UNIQUE REFERENCES users(id)
);

alter table recruiters
    owner to hack;

grant delete, insert, select, update on recruiters to hack;

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

alter table interviews
    owner to hack;

grant delete, insert, select, update on interviews to hack;

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

alter table "workSchedules"
    owner to hack;

grant delete, insert, select, update on "workSchedules" to hack;

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

alter table appointments
    owner to hack;

grant delete, insert, select, update on appointments to hack;

create table "daysOff"
(
    id            serial
        primary key,
    "recruiterID" integer not null
        references recruiters,
    date          date    not null
);

alter table "daysOff"
    owner to hack;

grant delete, insert, select, update on "daysOff" to hack;

