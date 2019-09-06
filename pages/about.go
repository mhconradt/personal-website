package pages

import "net/http"

type Usage struct {
	Name, Description string
}

type Language struct {
	Name string
	SkillLevel string
	Usages []Usage
	Opinion string
}

type AboutInfo struct {
	Languages []Language
}

func GetGolang() Language {
	rest := Usage{
		Name: "REST APIs",
		Description: "The concurrency model of the language is built into the language's HTTP package, allowing me to build multi-threaded HTTP servers with no extra effort. ",
	}
	ws := Usage{
		Name: "WebSockets",
		Description: "The language's concurrency model makes having long-running processes, such as WebSocket connections, simple to manage. " +
			"A WebSocket connection only takes about 16KB of RAM when using the most popular library. ",
	}
	dm := Usage{
		Name: "Data Mining",
		Description: "The language's concurrency model is great for starting up thousands of processes and pipelining the results of one stage to the next. " +
			"This makes the language an obvious choice for the data collection phase of data-mining. ",
	}
	site := Usage{
		Name: "This Website",
		Description: "This website is entirely written in Go. I chose to use Go in order to reduce time-to-interactive. ",
	}
	return Language{
		Name: "Golang",
		SkillLevel: "Very Proficient",
		Usages: []Usage{rest, ws, dm, site},
		Opinion: "Golang is my favorite language. " +
			"Its particular balance between efficiency and ease of development is incredibly attractive. " +
			"The language's concurrency model and simplicity are also fun to work with. " +
			"However, Go is missing a few things, notably generics, and I am looking forward to Go 2. ",
	}
}

func GetPython() Language {
	automation := Usage{
		Name: "Automation",
		Description: `If a task takes as long to complete manually as it does to automate it, I automate it.` +
		`The ease of writing Python makes many more tasks "automatable" with Python than say, C++.` +
			"I have automated the migration of thousands of documents from S3 Glacier to S3 Standard and taking inventory of all of the functions we have deployed. ",
	}
	nlp := Usage{
		Name: "Natural Language Processing",
		Description: "I have used the nltk library to train a sentiment analysis classifier on a corpus of movie reviews, achieving a surprisingly good F1 score. ",
	}
	cv := Usage{
		Name: "Computer Vision",
		Description: "I have used OpenCV 4 to draw bounding boxes around relevant paragraphs in my recipe book, in order to form a convenient digital index of them. " +
			"I achieved surprisingly good results using a variety of techniques, including Canny edge detection and component dilation. " +
			"I want to put some examples online, but need to think of a way to do so without breaking copyright laws. ",
	}
	algorithms := Usage{
		Name: "Algorithms and Data Structures",
		Description: "Python's expressivity makes it easy to think about problems in the context of the language. " +
			"If I am learning about an algorithm or data structure, such as DFS and BFS and radix-tries, respectively, Python is my first choice. ",
	}
	return Language{
		Name: "Python",
		SkillLevel: "Proficient",
		Usages: []Usage{automation, nlp, cv, algorithms},
		Opinion: "Python is very easy to work with due to its expressiveness. " +
			"However, its performance is untenable in a web-scale environment if it's anything but a wrapper around C/C++ or the JVM (NumPy and PySpark come to mind, respectively). " +
			"For building programs whose usage must be manually triggered, such as training a neural network or automating a task, it is useful. ",
	}
}

func GetJS() Language {
	react := Usage{
		Name: "React",
		Description: "I was primarily working on React for awhile at Tile Five, so gained significant expertise in building React applications. " +
			"I have spent time in the trenches working on Tile Five's main application and built out an e-commerce platform for Tile Five. " +
			`I'm good at thinking in React, as they say, but the CSS part of things can be a bit tedious and does not leverage my natural aptitude.`,
	}
	rn := Usage{
		Name: "React Native",
		Description: "I built a cross-platform mobile application at work that is used for checking into a gym, profile creation and document signing. " +
			"Once again, I find the styling aspect of front-end programming tedious. " +
			"However, I build maintainable, functioning code and am familiar with the often challenging environment of mobile development. ",
	}
	node := Usage{
		Name: "Node.js",
		Description: "My primary focus at work is building microservices deployed on AWS Lambda with the Serverless framework. " +
			"I have built a calendar system with booking, a document signing platform, an automated email system and a bunch of arbitrary CRUD microservices. " +
			"I have an eye for increasing performance of serverless (and other stateless) architectures, in ways such as reducing the time required for retrieving database credentials by 70ms on each call. " +
			"It takes a level of skill and meta-cognition while writing business logic in JavaScript without shooting yourself in the foot by writing poorly organized, unmaintainable code. " +
			"The next phase of my Node.js development will be learning how to integrate C++ modules into a program to speed up compute-bound tasks. ",
	}
	return Language{
		Name: "JavaScript",
		SkillLevel: "Expert",
		Usages: []Usage{node, react, rn},
		Opinion: "Pure JavaScript (not TypeScript) is like the Wild West because it's easy to do some stuff you should not do. " +
			"I enjoy using functional programming in JavaScript, as it makes code more maintainable and easy to keep track of. ",
	}
}

func GetCPP() Language {
	learning := Usage{
		Name: "Learning",
		Description: "I am in the early phases of scaling the learning curve of C++, but am attracted to its level of performance. " +
			"I want to have it in my back pocket for use in scenarios where the benefits of its additional performance outweigh the difficulty of development. " +
			"Programming in C++ feels like being at the helm of an aircraft carrier in both the positive and negative senses. ",
	}
	return Language{
		Name: "C++",
		SkillLevel: "Beginner",
		Usages: []Usage{learning},
		Opinion: "I am attracted to the efficiency and steep learning-curve of C++, and am excited to learn more about the language. ",
	}
}

func GetLua() Language {
	redis := Usage{
		Name: "Programming Redis Servers",
		Description: "I have used Lua for writing scripts to be run on my Redis servers for this site's blog and other projects. " +
			"As you have picked up, I am a freak for efficiency. " +
			"Thus, the prospect of eliminating a round-trip between my application server and the Redis server is highly attractive. ",
	}
	return Language{
		Name: "Lua",
		SkillLevel: "Enough to get the job done. No more. No less. ",
		Usages: []Usage{redis},
		Opinion: "I would not have heard of Lua were it not for Redis' support for the language, but love this use case. ",
	}
}

func GetScala() Language {
	migration := Usage{
		Name: "Database Migration Tool",
		Description: "I wrote a tool that uses the Flyway Java API to create database schemas and bring them up to date with the current schema version. ",
	}
	return Language{
		Name: "Scala",
		SkillLevel: "Beginner",
		Usages: []Usage{migration},
		Opinion: "I love functional programming in Scala. " +
			"Due to a mix of static-typing and performance, it seems like the functional programmer's best language option at the time of me writing this. " +
			"I am excited to use and learn about Scala more going forward. ",
	}
}

func FetchAbout(_ *http.Request) (ai AboutInfo, code int) {
	golang := GetGolang()
	python := GetPython()
	js := GetJS()
	cpp := GetCPP()
	lua := GetLua()
	scala := GetScala()
	return AboutInfo{
		Languages: []Language{
			golang,
			js,
			python,
			cpp,
			scala,
			lua,
		},
	}, 200
}
