Gator: RSS Feed aggregator

-You will need Postgres and Go installed to run this program.

-You can download and compile the program by running: 
go install

-You will need to create a config file within your home directory (cd ~) called .gatorconfig.json, with contents:
{"db_url":"postgres://your-postgres-db-key","current_user_name":""} 

xXCOMMANDSXx:

-register: register a username, the program will then create a row for this user within the users table.

-login: enter an existing username to set the current to that username.

-users: prints to the console a list of all users, indicating which user is currently logged in.

-addfeed: provide a single word name or hyphenated name and corresponding url for a RSS feed. This feed will be added to the feeds table. The current user will automatically follow this feed.

-feeds: prints a list of all feeds, their urls and username of uploaders to the console.

-follow: has the current user follow provided url of existing feed.

-unfollow: has the current user unfollow the provided feed url.

-following: prints to the console every feed that the current user is following.

-agg: begins retrieving posts from each feed the current user is following, and stores them within the posts db. You must provide an interval for the posts to be retrieved at.

-browse: displays all posts from feeds the current user is following. You may supply a limit to the # of posts displayed as a single integer.

-reset: resets entire db.


