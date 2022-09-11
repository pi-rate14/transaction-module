-- drop entries and transfers first because they have foreign key constraints from accounts table
DROP TABLE IF EXISTS entries;
DROP TABLE IF EXISTS transfers;
DROP TABLE IF EXISTS accounts;