```
GET /:file
Returns the contents of the file
```

```
GET /:file/info
Returns the statistics of the file
	returns:
		expires: date
		size: integer
		downloads: integer
```

```
POST /:file
Saves the contents of the file
  needs:
		file: binary string
		expires: date, default: now + 3 days
		auto_delete: boolean, default: false
		max_downloads: integer, default: 0 (infinite)
		key?: string
	returns:
		admin_key: string
		download_url: string
```

```
DELETE /:file
  needs:
		admin_key: string
	returns:
		success: boolean
		message: string
```
