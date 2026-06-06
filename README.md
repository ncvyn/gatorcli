# gator

### NOTE

This is a very basic RSS blog aggre*gator* CLI as part of a boot.dev guided project. Please don't throw tomatoes *(to-mah-toes?)* at me if the program crashes and burns your computer /s 

### Installation

You'll need Postgres and Go installed to run this program. (Any modern version should do!)

Then, run `go install https://github.com/ncvyn/gatorcli` followed by `go run .` or build the program with `go build`. 

### Setup

At your `HOME` directory, create a file named `.gatorconfig.json` (configurable in config.go if desired)

You can run the following subcommands:
- **login** [name]
- **register** [name]
- **reset** *(This will RESET the program's database!)*
- **users**
- **agg** [time]
- **addfeed** [feed name] [feed url]
- **feeds**
- **follow** [feed url]
- **following**
- **unfollow** [feed url]
- **browse** [limit (default: 2)]

### Modifying

If you're using NixOS (like me), you'll have to install goose and sqlc to run SQL migrations.
