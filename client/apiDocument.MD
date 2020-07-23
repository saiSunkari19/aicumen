API Documentation
===

### AddEmployee
- #### Request
```json
{
    "base_req": {
        "chain_id": "chain-1",
        "gas": "auto"
    },
    "password": "1234567890",
    "name":"Alice",
    "department":"network",
    "skills":["c","c++","java"]
}
```
- Response 

## Query

### QueryList of Employees
- Request
```json
    /employee/list
```
- Reponse
```json

{
    "height": "0",
    "result": [
        {
            "person": {
                "name": "Krishna",
                "address": "",
                "skills": [
                    "c",
                    "c++",
                    "java"
                ]
            },
            "id": "employee6",
            "department": "tech",
            "status": 1
        },
        {
            "person": {
                "name": "Bob",
                "address": "",
                "skills": [
                    "c",
                    "c++",
                    "java"
                ]
            },
            "id": "employee8",
            "department": "mg",
            "status": 1
        }
    ]
}
```
#### Query Employee By ID
- Request
```json
    /employee/list?id=employee8
```
- Response
```json
{
    "height": "0",
    "result": [
        {
            "person": {
                "name": "Bob",
                "address": "",
                "skills": [
                    "c",
                    "c++",
                    "java"
                ]
            },
            "id": "employee8",
            "department": "mg",
            "status": 1
        }
    ]
}
```
#### QueryList of Employees By Name

- Request
```json
    /employee/list?name=krishna
```

- Response
```json
{
    "height": "0",
    "result": [
        {
            "person": {
                "name": "Krishna",
                "address": "",
                "skills": [
                    "c",
                    "c++",
                    "java"
                ]
            },
            "id": "employee5",
            "department": "tech",
            "status": 1
        },
        {
            "person": {
                "name": "Krishna",
                "address": "",
                "skills": [
                    "c",
                    "c++",
                    "java"
                ]
            },
            "id": "employee6",
            "department": "tech",
            "status": 1
        }
    ]
}
```

### QueryList of Employee By Dept
- Request
```json
    /employee/list?dept=mg
```

- Reponse
```json
{
    "height": "0",
    "result": [
        {
            "person": {
                "name": "Bob",
                "address": "",
                "skills": [
                    "c",
                    "c++",
                    "java"
                ]
            },
            "id": "employee8",
            "department": "mg",
            "status": 1
        }
    ]
}

```