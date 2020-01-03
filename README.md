# qovery-go-sdk

Get Qovery instance
```$go
qovery := qovery.New()
db := qovery.GetDatabaseConfigurationByName("my-pql")

host := db.Host
port := db.Port
username := db.Username
password := db.Password
```
