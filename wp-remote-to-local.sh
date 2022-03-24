#!/bin/bash

REMOTE_USER=foo
REMOTE_HOST=foo
REMOTE_WEBDIR=/var/www/foo
REMOTE_DOMAIN=foo.xyz

DB_NAME=foo
DB_USER=foo
DB_PASS='foo'

LOCAL_WEBDIR=/var/www/foo
LOCAL_TMPDIR=/foo
LOCAL_TMPSQL=foo.sql
LOCAL_DOMAIN=foo.xyz

OPTIONS_TABLE=wp_options

#

echo "-- sync files..."

mkdir -p $LOCAL_WEBDIR
rsync -az -v --delete $REMOTE_USER@$REMOTE_HOST:$REMOTE_WEBDIR/ $LOCAL_WEBDIR
chown -R www-data: $LOCAL_WEBDIR

#

echo "-- sync database..."

DUMP="mysqldump --user=${DB_USER} --password='${DB_PASS}' ${DB_NAME}"
mkdir -p $LOCAL_TMPDIR
ssh $REMOTE_USER@$REMOTE_HOST $DUMP > $LOCAL_TMPDIR/$LOCAL_TMPSQL

#

echo "-- load database locally..."

mysql --user=$DB_USER --password=$DB_PASS $DB_NAME < $LOCAL_TMPDIR/$LOCAL_TMPSQL

#

echo "-- search replace in db..."

cd $LOCAL_WEBDIR
wp search-replace $REMOTE_DOMAIN $LOCAL_DOMAIN --skip-columns=guid --allow-root

#

echo "-- remove search engine visibility..."

mysql --user=$DB_USER --password=$DB_PASS $DB_NAME -e "UPDATE ${OPTIONS_TABLE} SET option_value=0 WHERE option_name='blog_public'"

#

echo "-- cleanup..."

rm $LOCAL_TMPDIR/$LOCAL_TMPSQL

echo "done."
