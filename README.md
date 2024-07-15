#### Run init for each chapter

```bash
cd code
for i in $(\ls -1); do cd $i; go mod init $i; cd ..; done
```

#### Configure VSCode

```json
"go.testFlags": ["-v"]
```
