


CREATE KEYSPACE auth1 WITH REPLICATION = {
  'class': 'SimpleStrategy',
  'replication_factor':1
};

CREATE TABLE IF NOT EXISTS auth1.Users (
    Email text,
    verifiedemail boolean,
    GivenName text,
    FamilyName text,
    Phone text,
    Picture       text,
    Locale        text,
    Password      text,
    Status        text,
    Provider text,
    Id text,
    CreatedAt TIMESTAMP,
    UpdatedAt TIMESTAMP,
    PRIMARY KEY(Email)
);

CREATE TABLE IF NOT EXISTS auth1.Logins (
    Username         text,
    SignedToken      text,
    SignedFreshToken text,
    Provider         text,
    Updated          TIMESTAMP,
    ExpiredAt        TIMESTAMP,
    PRIMARY KEY(Username)
);

