# TCL - Telegram Client

- default user account info
- parse public chats
- send messages over many user
  Not presented for free
- (schedule messages)
- (managing many accounts)


**Description:**  
TCL is a Telegram client built with the [go-tdlib](https://github.com/zelenin/go-tdlib) library in Go. This client facilitates seamless interaction with Telegram’s API, allowing users to perform various actions such as logging in, searching for users, and managing conversations.

**Setup:**  
To configure the client, create a `.env` file in the root directory with the following content:

app id & hash from [here](https://my.telegram.org/apps)

```
APP_ID="XXXXXXXX"
 APP_HASH="XXXXXXXXXXXXXXXXXXXXXXX"
```


**Login Process:**  
To log in to your account, run the client and follow the prompts:

```bash
./Hook_TCL 
start
Enter phone number: 
+XXXXXXXX
Enter code: 
XXXXXX
Enter password: 
XXXX(local tg pass if present)
```

## User Search Interface

![Screenshot from 2024-11-01 02-43-50](https://github.com/user-attachments/assets/1e8a81ae-f815-4ba1-b886-49c7da9cd943)


## Public chat parse interface, save to file option

![Screenshot from 2024-11-01 02-49-15](https://github.com/user-attachments/assets/dd69a61e-5154-44d0-8f3e-762d1141f299)

## Sending messages to may user, one per line or from txt file

![Screenshot from 2024-11-01 02-47-46](https://github.com/user-attachments/assets/c4002f86-1772-4b25-9da2-0afb85025e84)




