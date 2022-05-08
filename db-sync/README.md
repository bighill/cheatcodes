# DB Sync

Sync database from one remote server to another remote server.
Use a local computer to run the script and act as a bridge between the two remote servers.

Assumptions:

- The database being sync'd is MySQL
- The local computer can connect via ssh to both remote servers using password-less ssh key pair

## Usage

Configure variables at the top of the script, `run-db-sync.sh`

Run the script

```bash
./run-db-sync.sh
```
