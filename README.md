# 🐊 Gator - RSS Blog Aggregator in Go

**Gator** is a command-line based RSS blog aggregator built in **Go**. It allows users to register, follow RSS feeds, and view aggregated blog posts directly from the terminal.

---

## 🚀 Features

- ✅ User registration and login system
- 📥 Add and follow RSS feeds
- 🗞️ View blog posts from feeds you're following
- 🔄 Reset and initialize the database
- 📚 View all registered users and available feeds
- 🧹 Unfollow feeds
- 📡 RSS feed parsing and aggregation

---

## 🛠️ Requirements

- [Go](https://golang.org/doc/install) (v1.18 or higher recommended)
- [PostgreSQL](https://www.postgresql.org/download/)

---

## 📦 Installation and Setup

Clone the repository and install the program locally:

```bash
git clone https://github.com/your-username/gator.git
cd gator
go install
```
> [!IMPORTANT]
> Create a config file in your home directory named ***.gatorconfig.json*** with the following content:\
`
{
  "db_url": "postgres://<user_name>:@localhost:5432/gator?sslmode=disable",
  "current_user_name": "<user_name>"
}
`

Replace <user_name> with your PostgreSQL username or any name of your choice.

Ensure a PostgreSQL database named gator exists locally.

📂 Where is your Home Directory?
Mac/Linux: Open a terminal and run cd ~

Windows: Navigate to C:\Users\your_username

## 📖 CLI Commands

| Command                          | Description                                            |
| -------------------------------- | ------------------------------------------------------ |
| `register <username>`            | Register a new user in the database                    |
| `login <username>`               | Log in as an existing user                             |
| `reset`                          | Reset the database and remove all users                |
| `users`                          | List all registered users                              |
| `addfeed <feed_name> <feed_url>` | Add an RSS feed and associate it with the current user |
| `feeds`                          | View all feeds stored in the database                  |
| `follow <feed_url>`              | Follow a specific feed as the current user             |
| `following`                      | Show all feeds the current user is following           |
| `unfollow <feed_url>`            | Unfollow a specific feed                               |
| `agg`                            | Fetch, parse, and display all RSS posts from all feeds |
| `browse`                         | Display posts only from feeds the current user follows |

### 🧪 Example Feeds to Try
Here are some test RSS feeds you can use:

https://xkcd.com/rss.xml

https://blog.golang.org/feed.atom

https://hnrss.org/frontpage

### 🧑‍💻 Author
[Harsha Vardhan Mirthinti](https://www.linkedin.com/in/harshavardhanmirthinti/)

## 💡 Contributing
Pull requests are welcome! Feel free to open issues or suggest improvements.
