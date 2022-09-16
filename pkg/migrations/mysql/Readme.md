## Description

Database Migrations for Majoo test backend.

## Requirements

- [SQL-Migrate](https://github.com/rubenv/sql-migrate)

## Usage

1. Create the `dbconfig.yml` first. You can make one or ask another dev. Put it in the parent of `migrations` directory (`majoo`)
   
2. To create a new file, use `sql-migrate new -env='{your_env}' [file-name-here]`, where `[file-name-here]` is the name of the migration file. The env can also be changed to the values on `dbconfig.yml`.

3. To run the migration, use `sql-migrate up -env={your_env}` in the parent directory (`majo`).

4. To down a migration, simply use `sql-migrate down -env={your_env}`

5. In case if you want to delete a migration file, make sure that file is deleted in `migrations` table in your database too

## Known Issue

1. Windows users may encounter issues with the `sql-migrate` command, especially when installing it. If the issue is about `gcc`, you can try to download the gcc from [here](https://jmeubank.github.io/tdm-gcc/).

2. If you create a table from `SQL WORKBENCH` and copy it here, sometimes `VISIBLE` command in foreign keys doesn't work. You can delete it and it should run with no problem

