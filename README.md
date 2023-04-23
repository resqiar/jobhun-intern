# IMPORTANT NOTE

- **NOT 100% FINISHED**. This project is not 100% meet the requirement. I know it is a real bummer, but it's still worth checking since the most important features are done (create & read). Edit and delete can be easily be extended.
I do apologize for not completing the task since the date of the task is really inside the **Mudik** state, so I got incredibly busy preparing back to hometown and such. Indeed it is only my defense, so **if it is not meeting the requirement, then I will face the consequences**.  The project even only done in less than half a day, you can check the commit history.
- **Database Backup Included**. To fasten the testing and stuff, I also include the backup of the database inside the `__backup__` folder
- **ENV Included**. The environtment variables usually not the best idea to include in public repo, but for this project case, it is better to include for faster testing/config purposes.

# Technology Used
These are the tools and tech used inside of the repo and the rational reason behind it.

**1. Go Fiber**

As someone who currently switching from NodeJS to Golang, I picked the Go Fiber framework because it provides a development experience similar to the popular NodeJS Express framework.

For me, the conversion to Golang has been a learning experience, and I needed a framework that would make the process as easy as possible. Fiber's Express-like syntax and middleware architecture made it simple for me to get started rapidly constructing API endpoints in Golang.

Another advantage of adopting Fiber is its simplicity. Because Fiber has a tiny codebase and few dependencies, my applications have a smaller attack surface and are less vulnerable to security flaws.

Overall, adopting Fiber has made my shift from NodeJS to Golang easier and more pleasurable. Its Express-like syntax and middleware stack have made writing clean, organized code a pleasure, and its efficiency and minimalist approach have given me confidence in the scalability of my projects.

# How to Install
1. Clone the Project
	```bash
	git clone https://github.com/resqiar/jobhun-intern.git
	```

2. Install all required Go modules
	```bash
	go mod download
	```
3. Run (it will run in `localhost:3000` by default)
	```bash
	go run server.go
	```

> If you prefer to run from the build binary, you can run `./bin/jobhun-intern`
