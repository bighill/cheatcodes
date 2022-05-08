#!/bin/bash

SOURCE_SSH_USERNAME=foo
SOURCE_SSH_HOSTNAME=foo
SOURCE_DB_DATABASE=foo
SOURCE_DB_USERNAME=foo
SOURCE_DB_PASSWORD='foo' # wrap in single quote to allow special characters

DESTINATION_SSH_USERNAME=foo
DESTINATION_SSH_HOSTNAME=foo
DESTINATION_DB_DATABASE=foo
DESTINATION_DB_USERNAME=foo
DESTINATION_DB_PASSWORD='foo' # wrap in single quote to allow special characters

#
#   Configure variables above this line
#

TMP_FILE=tmp.sql
THIS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

echo "--- Initializing... "
cd $THIS_DIR
if [ -f "$TMP_FILE" ]; then
    rm $TMP_FILE
fi

echo "--- Get database from $SOURCE_SSH_HOSTNAME... "
DUMP_CMD="mysqldump --user=${SOURCE_DB_USERNAME} --password='${SOURCE_DB_PASSWORD}' ${SOURCE_DB_DATABASE}"
ssh $SOURCE_SSH_USERNAME@$SOURCE_SSH_HOSTNAME $DUMP_CMD > $TMP_FILE

echo "--- Push database to $DESTINATION_SSH_HOSTNAME... "
scp $TMP_FILE $DESTINATION_SSH_USERNAME@$DESTINATION_SSH_HOSTNAME:~/$TMP_FILE

echo "--- Import database... "
IMPORT_CMD="mysql --user=${DESTINATION_DB_USERNAME} --password='${DESTINATION_DB_PASSWORD}' ${DESTINATION_DB_DATABASE} < ${TMP_FILE}"
ssh $DESTINATION_SSH_USERNAME@$DESTINATION_SSH_HOSTNAME $IMPORT_CMD

echo "--- Clean up... "
rm $TMP_FILE
ssh $DESTINATION_SSH_USERNAME@$DESTINATION_SSH_HOSTNAME "rm $TMP_FILE"

echo "--- Done"