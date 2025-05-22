# Blog-Aggregator (Gator)
Gator :- An RSS blog Aggregator in Go.

Requirements :- 
You need to install Go and Postgress inroder to use this.

Installation :- 
    Use go install to install the program locally
SetUp:-
In you home directory create a file called `.gatorconfig.json` and paste the following contents in your json file
`
    {"db_url":"postgres://<user_name>:@localhost:5432/gator?sslmode=disable","current_user_name":"<user_name>"}
`
"user_name" :- It can be your name or anyname you like.
"db_url" :- this will be the endpoint for you DB created locally

Where can you find the Home Directory?? :-
For Mac/Linux :- 
On your terminal type `cd ~` and it will take you to the home directory.

For Windows :- 
C:\Users\username, where "username" is your Windows username


Functions:- 

Register (register <username>): Registers/Adds a new user into database.

Login(login <username>) := Logs in the user as the current user.


Reset (reset) := cleans everyting and resets the database(removes all the users from db)

Users (users) := Displays all the users that are registered in the DB.



AddFeed (addfeed <feed_name> <feed_url>) :- adds a feed to the current user(connects the feed to that user).

Feeds (feeds) :- prints all the feeds in the DB to the console.

Follow (follow <feed_url>) :-  It takes a single url argument and creates a new feed follow record for the current user.

Following (following) :=  print all the names of the feeds the current user is following.

Unfollow (unfollow <feed_url>) := accepts a feed's URL as an argument and unfollows it for the current user 

Aggregator(agg) :- fetch the RSS feeds, parse them, and print the posts to the console
Browse (browse):-  view all the posts from the feeds the user follows, right in the terminal!