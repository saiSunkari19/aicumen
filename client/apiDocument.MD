API Documentation
===

### AddEmployee

- EndPoint: `/employee/add`
- Request
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
- Response when I use broadcast mode `async`
```json
{
    "height": "0",
    "txhash": "6B20F16965594D30F0AB505C2A0BF9F8B633215318B4982D61C962C0F2184330"
}
```
- Resonse when I use broadcast mode `block`
```json
{
    "height": "37",
    "txhash": "09F55A4BC1CE36E5A4FA875056FB436745C95C9E0BC51038CB04912E6AC5E09D",
    "raw_log": "[{\"msg_index\":0,\"log\":\"\",\"events\":[{\"type\":\"events_msg_add_employee\",\"attributes\":[{\"key\":\"employee_id\",\"value\":\"employee0\"},{\"key\":\"sender\",\"value\":\"cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"msg_add_employee\"}]}]}]",
    "logs": [
        {
            "msg_index": 0,
            "log": "",
            "events": [
                {
                    "type": "events_msg_add_employee",
                    "attributes": [
                        {
                            "key": "employee_id",
                            "value": "employee0"
                        },
                        {
                            "key": "sender",
                            "value": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
                        }
                    ]
                },
                {
                    "type": "message",
                    "attributes": [
                        {
                            "key": "action",
                            "value": "msg_add_employee"
                        }
                    ]
                }
            ]
        }
    ],
    "gas_wanted": "49584",
    "gas_used": "48506"
}
```

### UpdateEmployee
- EndPoint: `/employee/update`
- Request
```json
{
    "base_req": {
        "chain_id": "chain-1",
        "gas": "auto"
    },
    "password": "1234567890",
    "id":"employee5",
    "address":"HMP",
    "department":"mg",
    "skills":["c","c++","java","rust"]
}
```
- Response
```json
{
    "height": "0",
    "txhash": "56D6F90030CF8F8EECEEB6442E54287B6AF74002B93C8CBF29FCBECB46148316"
}
```


### DeleteEmployee
- EndPoint: `/employee/delete`
- We have `remove` filed`false`, where it will deactivate the employee info
- Where `remove` is `true`, it will permanently delete the employee info

- Request with remove field `false`
```json
{
    "base_req": {
        "chain_id": "chain-1",
        "gas": "auto"
    },
    "password": "1234567890",
    "id":"employee3",
    "remove": false
}
```
- Request with remove field `true`
```json
{
    "base_req": {
        "chain_id": "chain-1",
        "gas": "auto"
    },
    "password": "1234567890",
    "id":"employee4",
    "remove": true
}

```
- Response
```json
{
    "height": "0",
    "txhash": "56D6F90030CF8F8EECEEB6442E54287B6AF74002B93C8CBF29FCBECB46148316"
}
```
### Restore Employee
- EndPoint: `/employee/restore`
- Where we `active` the employee info, where it's partially deleted. `remove` is `false`. 
- Request
```json
{
    "base_req": {
        "chain_id": "chain-1",
        "gas": "auto"
    },
    "password": "1234567890",
    "id":"employee3"
}

```
- Response
```json
{
    "height": "0",
    "txhash": "E1B7041B6789D2E097CFC0E41D4F90BA53FE895AC3E02C423F3B0F726ED521AD"
}
```

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
                ],
                "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
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
                ],
                 "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
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
                ],
                 "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
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
                ],
                 "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
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
                ],
                 "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
            },
            "id": "employee8",
            "department": "mg",
            "status": 1
        }
    ]
}

```

### Query DeActivated Employees
- Request
```json
    /employee/viewdeleted
```
- Resonse
```json
{
    "height": "0",
    "result": [
        {
            "person": {
                "name": "Bob",
                "address": "HMP",
                "skills": [
                    "c",
                    "c++",
                    "java"
                ],
                 "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
            },
            "id": "employee3",
            "department": "tech",
            "status": 2
        }
    ]
}
```

### Search Employees

#### Search Employee By ID
- Request
```json
  /employee/search?id=employee3
```
- Response
```json
{
    "height": "0",
    "result": [
        {
            "person": {
                "name": "Bob",
                "address": "HMP",
                "skills": [
                    "c",
                    "c++",
                    "java"
                ],
                 "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
            },
            "id": "employee3",
            "department": "tech",
            "status": 2
        }
    ]
}
```

#### Search Employees By Department
- Request
```json
    /employee/search?dept=mg
```
- Resonse
```json

{
    "height": "0",
    "result": [
        {
            "person": {
                "name": "Alice",
                "address": "KNR",
                "skills": [
                    "rust"
                ],
                 "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
            },
            "id": "employee4",
            "department": "mg",
            "status": 1
        },
        {
            "person": {
                "name": "Alice",
                "address": "HMP",
                "skills": [
                    "c",
                    "java",
                    "c++",
                    "rust"
                ],
                 "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
            },
            "id": "employee5",
            "department": "mg",
            "status": 1
        }
    ]
}
```

#### Search Employees By Name
- Request
```json
  /employee/search?name=bob
```
- Response
```json
{
    "height": "0",
    "result": [
        {
            "person": {
                "name": "Bob",
                "address": "HMP",
                "skills": [
                    "c",
                    "c++",
                    "java"
                ],
                 "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
            },
            "id": "employee3",
            "department": "tech",
            "status": 2
        }
    ]
}
```
#### Search Employees By Address
- Request
```json
  employee/search?addr=hmp
```

- Response
```json
{
    "height": "0",
    "result": [
        {
            "person": {
                "name": "Bob",
                "address": "HMP",
                "skills": [
                    "c",
                    "c++",
                    "java"
                ],
                 "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
            },
            "id": "employee3",
            "department": "tech",
            "status": 2
        },
        {
            "person": {
                "name": "Alice",
                "address": "HMP",
                "skills": [
                    "c",
                    "java",
                    "c++",
                    "rust"
                ],
                 "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
            },
            "id": "employee5",
            "department": "mg",
            "status": 1
        }
    ]
}
```

#### Search Employees By Skills
- Request
```json
  /employee/search?skills=c,rust
```
- Response
```json
{
    "height": "0",
    "result": [
        {
            "person": {
                "name": "Alice",
                "address": "KNR",
                "skills": [
                     "c",
                    "rust"
                ],
                 "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
            },
            "id": "employee4",
            "department": "mg",
            "status": 1
        },
        {
            "person": {
                "name": "Alice",
                "address": "HMP",
                "skills": [
                    "c",
                    "java",
                    "c++",
                    "rust"
                ],
                 "owner": "cosmos1yddf3fz00v35m7newqnfq6cqjvsnzq60356awz"
            },
            "id": "employee5",
            "department": "mg",
            "status": 1
        }
    ]
}
```