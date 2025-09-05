## `tinytweet` - cli tool to post to X from terminal

<img width="1193" height="636" alt="Image" src="https://github.com/user-attachments/assets/10dfe02e-dba0-4378-9cd2-c7f1e497f694" />

### How to use

1. Clone the repository:
```bash
git clone git@github.com:0xnaman1/tinytweet.git
```

2. Create a **.envrc** file in the root of the project and add the following variables. I used `direnv` pkg for the env variables. In case you are using some other env pkg, you might have to modify the `main.go` file to read the env variables.

```.env
export CONSUMER_KEY="YOUR_CONSUMER_KEY"
export CONSUMER_SECRET="YOUR_CONSUMER_SECRET"
export ACCESS_TOKEN="YOUR_ACCESS_TOKEN"
export ACCESS_SECRET="YOUR_ACCESS_SECRET"
```

I used [https://developer.x.com/](https://developer.x.com/) to get the consumer key, consumer secret, access token, and access secret. Twitter uses OAuth 1.0a for authentication.

3. Install the dependencies:

```bash
go mod tidy
```

4. Build the binary:

```bash
go build .
```

5. Run the binary:

```bash
tinytweet
```
