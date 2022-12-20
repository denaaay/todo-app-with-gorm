## TODO List App with GORM

### Description

Aplikasi bernama todo app yang berfungsi untuk membuat list beserta status pekerjaan dari aktifitas yang kita kerjakan. Fitur dari aplikasi ini adalah:

- Register
- Login
- Create Todo list
- Read Todo list
- Change Status Todo List
- Remove Todo List

Terdapat chain middleware untuk menghandle Method dan Authentication dengan menggunakan metode session based token kemudian menyimpan semua data user dan todo list di database **PostgresSQL** menggunakan **GORM**.

### Penting untuk mengubah koneksi database lokal menjadi milik anda menjadi :

```go
dbCredentials = Credential{
    Host:         "localhost",
    Username:     "postgres", // <- ubah ini
    Password:     "postgres", // <- ubah ini
    DatabaseName: "database", // <- ubah ini
    Port:         5432,
}
```
