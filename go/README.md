# GO

## Usage
in the `src` folder, there is a script that can be run to generate skeleton folders for each day with the problems.

### To generate new folder for a year (this will also sync each day problems, if they are available)
```shell
go run create_aoc_folders.go 2022
```

### To sync the day folders 

```shell 
go run create_aoc_folders.go 2022 --sync-only
```
