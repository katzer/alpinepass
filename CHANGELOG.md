# alpinepass

## 0.0.1 (UNRELEASED)

1. Basic functionality. Read "input.yml" and write "output.json".

    ```
	$ ls
	input.yml
	$ alpinepass
	$ ls
	input.yml output.json
    ```

2. New "input" flag for specifying the input file.

    ```
	$ alpinepass -i /path/to/input.yml
    ```

3. New "output" flag for specifying the output file.

    ```
	$ alpinepass -o /path/to/output.json
    ```

4. New "display" flag for previewing the output in the console. An output file will not be written.

    ```
	$ alpinepass -d
	[{"id": "B01.prod.server.PROD-App","title": "PROD App","location": "B01","environment": "prod","user": "prodUserB01"}]
    ```

5. New "readable" flag for formatting the output.

    ```
	$ alpinepass -d -r
    [
        {
            "id": "B01.prod.server.PROD-App",
            "title": "PROD App",
            "location": "B01",
            "environment": "prod",
            "user": "prodUserB01"
        }
    ]
	$ alpinepass -r
	$ cat output.yml
    [
        {
            "id": "B01.prod.server.PROD-App",
            "title": "PROD App",
            "location": "B01",
            "environment": "prod",
            "user": "prodUserB01"
        }
	]
    ```

6. New "passwords" flag for including passwords in the output.

    ```
	$ alpinepass -d -p
	[
        {
            "id": "B01.prod.server.PROD-App",
            "title": "PROD App",
            "location": "B01",
            "environment": "prod",
            "user": "prodUserB01",
			"password": "prod_pw"
        }
	]
	```

7. New "filter" flag for filtering the input.

    ```
	$ alpinepass -d -p
	[
        {
            "id": "B01.prod.server.PROD-App",
            "title": "PROD App",
            "location": "B01",
            "environment": "prod",
            "user": "prodUserB01",
			"password": "prod_pw"
        },
		{
            "id": "D02.prod.server.PROD-App",
            "title": "PROD App",
            "location": "D02",
            "environment": "prod",
            "user": "prodUserD02",
			"password": "prod_pw"
        }
	]
	$ alpinepass -d -p -f location:D02
	[
		{
            "id": "D02.prod.server.PROD-App",
            "title": "PROD App",
            "location": "D02",
            "environment": "prod",
            "user": "prodUserD02",
			"password": "prod_pw"
        }
	]
	```

8. New "skip" flag for disabling input verification.

    ```
	$ alpinepass -d
	$ alpinepass -d s
	```
