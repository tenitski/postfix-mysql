# postfix-mysql

Proof of concept of using Mysql to store SASL users and allowed sender domains.

# Moving parts
- Postfix
- Cyrus SASL library
- Mysql as a storage for users configuration

# Known issues

- Mysql IP has to be used, hostname does not work (Cyrus seem to treat it as socket name rather than hostname)

# (Manually) tested scenarios

- Auth: user unable to send a email not using auth
- Auth: user unable to auth using incorrect username or password
- Auth: user able to auth using username & password from the Mysql table containing users
- Sender: user allowed to send mail only as `*@example.com` can send an email from `test@example.com`
- Sender: user allowed to send mail only as `*@example.com` can not send an email from `test@mail.example.com`
- Sender: user allowed to send mail as `*@example.net` and `*@mail.example.net` can send an email from `test@example.net` and `test@mail.example.net`
- Sender: user allowed to send mail as `test@gmail.com` can send an email from `test@gmail.com`
- Sender: user allowed to send mail as `test@gmail.com` can NOT send emails from any other address

# Helpful commands

```bash
# Look up sender in sender map
postmap -q john@example.com mysql:/etc/postfix/virt/senders.cf

# Get SASL config
saslfinger -s

```

```sql
-- Enable sql log to see what queries Postfix is making
SET global log_output = 'FILE';
SET global general_log_file='/var/log/mysql/general.log';
SET global general_log = 1;
```