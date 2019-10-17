# bash

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