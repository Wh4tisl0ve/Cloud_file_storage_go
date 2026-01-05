# Cloud_file_storage_go
Многопользовательское файловое облако. Пользователи сервиса могут использовать его для загрузки и хранения файлов.

## Migrations
**Library** - golang-migrate/migrate/v4  
**Location** - /migrations  
**Name pattern** - `<version>_<name>.<up/down>.sql`  
**Name example** - `001_init.down.sql`

### Action:
```bash copy 
go run cmd/migrate/main.go up
```

```bash copy 
go run cmd/migrate/main.go down
```