docker run -d `
  --name pfs-postgresql `
  -p 5432:5432 `
  --env-file .\backend\database.env `
  postgres