# Payment service

## Features

 * Account creation with the initial balance
 * Payment creation which leads to money transfer between accounts (multicurrency support)
 * Review current exchange rates
 * Review payment history per account
 * Review list of all accounts in the system

## Requirements

To start application or run integration tests you need only docker.

## Installation

To start application you can run:
`docker-compose up`.  
It will listen on port 9090.

Application run required database migrations on startup. Migrations are stored in `migrations` folder.

## Configuration

Application configured via ENV variables:

 * POSTGRESQL_HOST - database host
 * POSTGRESQL_PORT - database port
 * POSTGRESQL_USER - database user
 * POSTGRESQL_PASSWORD - database password
 * POSTGRESQL_DATABASE - database name
 * HTTP_HOST - host listen on (default=0.0.0.0)
 * HTTP_PORT - port listen on (default=80)

## Integration testing

To run integration tests, which are located in `integration_tests` folder, run `./scripts/integration_tests.sh`.

## API

### Errors

| codes | messages |
| :---: | :--- |
| 1 | INVALID_CURRENCY_CODE |
| 2 | INVALID_ACCOUNT_ID |
| 3 | NEGATIVE_BALANCE |
| 4 | NEGATIVE_PAYMENT_AMOUNT |
| 5 | NOT_ENOUGH_MONEY |
| 6 | SAME_ACCOUNT_TRANSFER | 
| 100 | INTERNAL_ERROR |
| 101 | INVALID_OFFSET_VALUE |
| 102 | INVALID_LIMIT_VALUE |

**Content example**

```json
{
  "code": 1,
  "msg": "INVALID_CURRENCY_CODE"
}
```

### Accounts

#### List of accounts

Get list of all accounts with balance and currency code. Limit and offset are optional, by default offset=0 and limit=100

**URL**: `/accounts?offset=0&limit=20`

**Method**: `GET`

##### Response

**Code**: `200 OK`

**Content example**

```json
[
  {
    "id": "6cd061e4-4a0c-4991-85fd-4825702f873b",
    "currency_numeric_code": 980,
    "balance": "10000.5"
  }
]
```

#### Account by ID

Get account by ID.

**URL**: `/accounts/6cd061e4-4a0c-4991-85fd-4825702f873b`

**Method**: `GET`

##### Response

**Code**: `200 OK`

**Content example**

```json
{
  "id": "6cd061e4-4a0c-4991-85fd-4825702f873b",
  "currency_numeric_code": 980,
  "balance": "10000.5"
}
```

#### Create account

Create new account with positive balance

**URL**: `/accounts`

**Method**: `POST`

##### Request

**Content example**

```json
{
  "currency_numeric_code": 980,
  "balance": "1500.50"
}
```

##### Response

**Code**: `201 CREATED`

**Content example**

```json
{
  "id": "6cd061e4-4a0c-4991-85fd-4825702f873b",
}
```

### Currencies

#### List of currencies

**URL**: `/currencies`

**Method**: `GET`

##### Response

**Code**: `200 OK`

**Content example**

```json
[
  {
    "numeric_code": 643,
    "alpha_code": "RUB",
    "minor": 2
  },
  {
    "numeric_code": 840,
    "alpha_code": "USD",
    "minor": 2
  },
  {
    "numeric_code": 933,
    "alpha_code": "BYN",
    "minor": 2
  },
  {
    "numeric_code": 978,
    "alpha_code": "EUR",
    "minor": 2
  },
  {
    "numeric_code": 980,
    "alpha_code": "UAH",
    "minor": 2
  }
]
```

### Payments

#### List of payments per account

Get list of payment history per account. Limit and offset are optional, by default offset=0 and limit=100

**URL**: `/accounts/6cd061e4-4a0c-4991-85fd-4825702f873b/payments?offset=0&limit=20`

**Method**: `GET`

##### Response

**Code**: `200 OK`

**Content example**

```json
[
  {
    "id": "6cd061e4-4a0c-4991-85fd-4825702f873b",
    "currency_numeric_code": 980,
    "balance": "10000.5"
  }
]
```

#### Create payment

Create new payment and tranfer money between accounts.

**URL**: `/payments`

**Method**: `POST`

##### Request

**Content example**

```json
{
  "from_account": "6a261c31-7009-4082-80c4-6abde77082bb",
  "to_account": "2d47cfa8-d092-4bc3-a30c-93e8c9aaf6d6",
  "currency_numeric_code": 980,
  "amount": "950.55"
}
```

##### Response

**Code**: `201 CREATED`

**Content example**

```json
{
  "id": "6cd061e4-4a0c-4991-85fd-4825702f873b",
}
```
