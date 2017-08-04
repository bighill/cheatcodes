# mysql cheatsheet #

## Database ##

```sql
SHOW DATABASES;
```

```sql
CREATE DATABASE my_database;
```

```sql
DROP DATABASE my_database;
```

_Always use underscores instead of hyphens in database names_

## Table ##

```sql
CREATE TABLE `my_table` (
  `id` int(10) unsigned NOT NULL auto_increment,
  `title` varchar(50) NOT NULL,
  `url` varchar(100) NOT NULL,
  `description` varchar(255) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `bool` tinyint(1) unsigned NOT NULL default '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `url` (`url`)
);
```

```sql
DESCRIBE my_table;
```

## Column ##

Add column
```sql
ALTER TABLE my_table ADD my_new_column VARCHAR(255);
```

Delete column
```sql
ALTER TABLE my_table DROP COLUMN my_column;
```

Change column datatype
```sql
ALTER TABLE my_table MODIFY my_column VARCHAR(111);
```

Change column name
```sql
ALTER TABLE my_table CHANGE old_column_name new_column_name VARCHAR(255);
```

Add column after a particular column
```sql
ALTER TABLE my_table ADD new_column_name VARCHAR(255) AFTER other_column_name;
```

## Data ##

Insert
```sql
INSERT INTO my_table (my_column1, my_column2)
VALUES (my_value1, my_value2);
```

Update
```sql
UPDATE my_table
SET my_column1='my_value1', my_column2='my_value2'
WHERE my_column3='my_value3';
```

Delete
```sql
DELETE FROM my_table
WHERE my_column=my_value;
```

## User ##

_Typically you're going to want to do these in this order..._

1) Make user
```sql
CREATE USER 'mynewuser'@'localhost' IDENTIFIED BY 'mypassword';
```

2) Give user access to database
```sql
GRANT ALL PRIVILEGES ON my_database.* TO 'mynewuser'@'localhost';
```
```sql
GRANT [type of permission] ON [database name].[table name] TO ‘[username]’@'localhost’;
```

3) Flush when done
```sql
FLUSH PRIVILEGES;
```
