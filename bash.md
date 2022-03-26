# bash

This dir

```bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd $DIR
```

`ls` sort by date

```bash
ls -lt
```

`ls` sort by size

```bash
ls -lS
```

20 largest files/dirs in current directory

```bash
du -ah . | sort -nr | head -n 20
```

Replace spaces with underscores in all files in current directory

```bash
for file in *; do mv "$file" `echo $file | tr ' ' '_'` ; done
```

Find symlinks

```bash
find . -type l -ls
```
