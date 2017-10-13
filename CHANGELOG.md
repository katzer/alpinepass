# alpinepass

## 1.4.4 - UNRELEASED

## 1.4.3 - 2017-10-06

1. Support multiple users for each system configuration.
2. Improve error messages, do not show the help text when an error occurs.

## 1.4.2 - 2017-09-29

1. Adjust the release script.
2. Introduce the "tool" configuration type.
3. Improved error messages. When an error occurs an error message is shown which indicates the error reason followed by the help text.
4. Add the "--debug" flag which prints the stacktrace when an error occurs. No help text is shown.
5. New filter logic:
    * "Exact" filters "key=value" match when the key's content matches exactly the filter value. "Earth" matches "Earth" but not "Earths".
    * "Contains" filters "key:value" match then the filter value is contained anywhere in the key's content. "Bar" matches "FooBar" and "BarFoo" but not "BazFoo".

## 1.0.0 - 2017-03-20

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

5. New "readable" flag for formatting the output. It works with both output file and console output!
    ```
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

8. Input verification checks that certain properties are present for the different configuration types.

9. New "skip" flag for disabling input verification.
    ```
	$ alpinepass -d -r -s
    [
        {
            "id": "B01.prod.server.PROD-App",
            "title": "PROD DB",
            "location": "B01",
            "environment": "prod",
            "user": "prodUserB01",
			"host": "prodDatabase.B01"
        }
    ]
	$ alpinepass -d -r
    The host "prodDatabase.B01" does not match the naming convention "[environment]Db.[location]".
	```
