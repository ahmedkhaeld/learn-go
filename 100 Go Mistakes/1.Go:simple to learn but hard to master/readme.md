# Go: simple to learn but hard to master

---
**This chapter covers**
* what makes go an efficient, scalable, and productive lag
* Exploring why Go is simple to learn but hard to master
* Presenting the common types of mistakes made by developers


---

### 1.1 Go outline
Feature-wise, Go has no type inheritance, no exceptions, no macros, no partial functions, no support for lazy variable evaluation or immutability, no operator overloading, no pattern matching, and on and on

Why does Go not have feature X? Your favorite feature may be missing because it doesn’t fit, because it affects compilation speed or clarity of design, or because it would make the fundamental system model too difficult.

* Stability: Even though Go receives frequent updates (including improvements and security patches), it remains a stable language. Some may even consider this one of the best features of the language.
* Expressivity: We can define expressivity in a programming language by how naturally and intuitively we can write and read code. A reduced number of keywords and limited ways to solve common problems make Go an expressive language for large codebases.
* Compilation—As developers, what can be more exasperating than having to wait for a build to test our application? Targeting fast compilation times has always been a conscious goal for the language designers. This, in turn, enables productivity.
* Safety—Go is a strong, statically typed language. Hence, it has strict compile-time rules, which ensure the code is type-safe in most cases.

---
### 1.2 Simple doesn't mean easy

Let’s take concurrency, for example. In 2019, a study focusing on concurrency bugs was published: “Understanding Real-World Concurrency Bugs in Go.”1 This study was the first systematic analysis of concurrency bugs. It focused on multiple popular Go repositories such as Docker, gRPC, and Kubernetes. One of the most important takeaways from this study is that most of the blocking bugs are caused by inaccurate use of the message-passing paradigm via channels, despite the belief that message passing is easier to handle and less error-prone than sharing memory.

What should be an appropriate reaction to such a takeaway? Should we consider that the language designers were wrong about message passing? Should we reconsider how we deal with concurrency in our project? Of course not.

It’s not a question of confronting message passing versus sharing memory and determining the winner. However, it’s up to us as Go developers to thoroughly understand how to use concurrency, its implications on modern processors, when to favor one approach over the other, and how to avoid common traps. This example highlights that although a concept such as channels and goroutines can be simple to learn, it isn’t an easy topic in practice.


---
### 1.3 100 Go mistakes
This book presents seven main categories of mistakes
* Bugs
* Needles complexity
* Weaker readability
* Suboptimal or unidiomatic organization
* Lack of API convenience
* Under-optimized code
* Lack of productivity

#### 1.3.1 Bugs
software bugs aren’t only about money. As developers, we should remember how impactful our jobs are.

various software bugs, including data races, leaks, logic errors, and other defects. Although accurate tests should be a way to discover such bugs as early as possible

#### 1.3.2 Needles complexity
A significant part of software complexity comes from the fact that, as developers, we strive to think about imaginary futures. Instead of solving concrete problems right now, it can be tempting to build evolutionary software that could tackle whatever future use case arises. However, this leads to more drawbacks than benefits in most cases because it can make a codebase more complex to understand and reason about.

#### 1.3.3 Weaker readability
software engineering is programming with a time dimension: making sure we can still work with and maintain an application months, years, or perhaps even decades later.

When programming in Go, we can make many mistakes that can harm readability. These mistakes may include nested code, data type representations, or not using named result parameters in some cases. 

####1.3.4 Suboptimal or unidiomatic organization
organizing our code and a project suboptimally and unidiomatically. Such issues can make a project harder to reason about and maintain.

####1.3.5 Lack of API convenience
Making common mistakes that weaken how convenient an API is for our clients is another type of mistake. If an API isn’t user-friendly, it will be less expressive and, hence, harder to understand and more error-prone.

We can think about many situations such as overusing any types, using the wrong creational pattern to deal with options, or blindly applying standard practices from object-oriented programming that affect the usability of our APIs.

####1.3.6 Under-optimized code
It can happen for various reasons, such as not understanding language features or even a lack of fundamental knowledge. Performance is one of the most obvious impacts of this mistake, but not the only one.

####1.3.7 Lack of productivity
In most cases, what’s the best language we can choose when working on a new project? The one we’re the most productive with. Being comfortable with how a language works and exploiting it to get the best out of it is crucial to reach proficiency.
