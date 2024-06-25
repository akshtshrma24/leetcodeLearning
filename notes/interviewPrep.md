# Interview Prep 

STAR Method

Psychology
Situation 
Task 
Action 
Result
Reflection
Summary 

<hr>

* Tell me about a time when you were faced with a problem that had a number of possible solutions. What was the problem and how did you decide what to do? What was the outcome? 
    * P 
        * So yea whenever I am faced with an issue or given a task that has multiple possible solutions I like to take a step back and dive deep into each of the possible solutions. Which one is more scalable, which is easier to implement, which is more simple, and more importantly just which one would the user overall like better. So this reminds me of a time when:
    * S 
        * So yea, Good question in my highest role of being VP at SCE, the largest engineering club at my university, I was leading the dev team  to create a brand-new messaging service for our members. I took this initiative because I noticed that I wasnt seeing new membeers in the club as much as I used to, so I sent out a survey to the members and foudn that my thoughts were in fact true. The members were dissasitisfied with the "cool code" that we promised them to be delivered or lack there of the code. I also found that we had 40% less new member sign ups from the year prior. 
    * T
        * After I found that out, I knew that I needed to take action, I could not let our club members, basically our customers down, I needed good feedback from them. nd being the new VP I also needed to do this while to earn the trust of my new team. I decided it would be best to move quick but still dive deep into all the possible solutions to make sure we selected the most simple, scalable, and portable solution. As we needed the solution to be able to run on an ESP32 
    * A
        * Firstly, before I start coding, I like to take my whiteboard out and draw what I have, and what needed to be implemented. Like a simple diagram. I saw that our current infustructure could not handle real time messaging so I could go multiple routes, I could implement SSEs or I could have just another API that can be scraped for new messages. I did research on which one would be better for our use case, messages sent on between users and will be displayed on an lcd screen with an esp32. After looking at my requirements, userbase, and server loads, I chose SSEs as my go to for the server side, and the client would connect via ESP32 and messages would be displayed on the lcd screen. I wrote some test files for both the SSE and esp32 and showed my findings to the dev team and the president who both liked my solution, and were excited to begin development. I worked closely with my team, developing the SSEs necessary and integrating it with our current infrustructure on our website and I worked with some other dev team on writing the embedded software for the esp32s that were to be given to our members. Despite my team having little experience with SSEs, networking, and embedded software I was able to work alongside them and teach them the skills necessary. 
    * R 
        * By the end of our deadline, My team succesfully delivered the fully functional messaign service to our members. With our new engagement by Members and the fact that they found it "cool" we were able to have an increase of over 35% the following semester! And it looks like that trend is going way on up thanks to our summer internship program which I am leading. The project also enhanced our clubs reputation which lead to even more interest and participation in our events. 
    * R
        * This experience taught me more things than just embedded software and SSEs. It taught me that diving deep into possible solutions rather than picking 1 and rolling with it leads to a much smoother transition from code to production. It taught me the value of choosing our members first, and to earn trust is to consistenly communicate effectively. 
    * S
        So to answer the question in short, as my first project as VP, I took initiative to tackle our declining member sign up and satisfaction and created something new and simple to boost member engagement and interest. I first dove deep into the possible solutions, presented my solutions, tested it, and then began development along side my team, then rolled the code out.   
<hr >

* When did you take a risk, make a mistake, or fail? How did you respond? How did you grow from it? 
    * P
        * So I believe that making mistakes and failing is going to happen to anyone and everyone, no matter the experience level but I believe that mistakes and failures are only failures if you do not learn or grow from them. Although making mistakes can sometimes be costly it matters a lot to go back into your previous thought process dive deep into where you failed, and take ownership of that to change it do deliver a quality result.
    * S
        * I actually have a really good story for this, when I was a up coming junior in college, I was a part of the dev team at the engineering club, and I was and still am eager to learn. I noticed that the development on current projets in the room was awfully slow due to people not being able to work on projects outside of the room, which meant usually for only a couple of hours between monday - thursday. So in total people could only work on projects on average 5-6 hours a week.
    * T
        * I thought simple enough I am going to make a VPN, that will solve that issue and bring access to the room remortely. It needed to be invisible to SJSU as they would take it down if they found it. 
    * A
        * I immediately began to drawing a solution, where I would have openVPN on an EC2 instance and have a router as a client to that. I drew the solution and presented it to some alumni who, approved. This is where my mistake lies, I did not dive deep into the solution and each of its parts to figure out if this is actually right or if there is another simpler solution I could implement. I went head first instead of brain first. Because of these mistakes, I spent a month trying to implement something that was just wrong. It was only when I went back really thought about each one of my components, each file, each command, expanded my knowledge and thought about other factors too such as scalability and frugality I came to the right solution. I redrew my solutions, I represented it with the alumni who was impressed that I was able to figure it out, and then approved it again. I began to work on a test network where I tested my work on with multiple factors, then after vigurous testing put it in production in the engineering room.
    * R
        * It has been running problem free for about 6 months now, and has greatly boosted our productivity. Our existing memberes are now able to put double the hours into their projects and it has increased our membership sign up as new members can work on projects they never could before. Overall it has boosted our interest in the club by over 20%.
    * R
        * I wouldnt count this as a complete failure, I was able to take my mistakes, learn not only how to configure a VPN but how and why to dive deep into the project before working on it. Why I it is important to focus on simplicity of a solution. I most importantly learned to take ownership of my mistakes and how to turn failures into future succeses.
    * S
        * So to answer that question I made a mistake when diving too quick into a project I wanted to do and learned important principles from it such as taking ownership, diving deep, and simplifying.
<hr >

* Describe a time you took the lead on a project. 
    * P
        * I took the lead on a project in 
    * S
    * T
    * A
    * R
    * R
    * S
<hr >

* What did you do when you needed to motivate a group or promote collaboration on a project? 
    * P 
        * Motivating a group and increasing collaboration on a project is something that is crucial to the quality of the product. Without a motivated team you get an equally good project. Firstly taking ownership of the project is a must, from there you need to earn the trust of the team and be ok with disagreeing. When disagreements do happen it is important to listen to the other person and then come to a decision and commit to tit then. 
    * S
        * In one of my classes we had to make a project where I was assigned a group of 4 (including me), and I had to build a project from start to finish, with scrum meetings weekly, implementing agile methods and getting everyone to contribute. I was given freedom of the project in terms of tech stack, but it needed to be a fullstack website with a database deployed with docker, but some of my team members were not familiar with docker or dbs. 
    * T
        * I had to take ownership of the proejct, dive deep with my team to explain inner workings of the tech stack that I chose with my team. And effectively and consistently communicate to earn trust and build a quality project. All this needed to be done in a tight time frame without sacrificing the quality of the project. 
    * A
        * I setup weekly meetings with my group, and always asked the person with the least experience to open the meetings because that way his ideas werent influenced by others. This really motivated him to always come prepared to the meetings. I worked alongside my group during the meetings and outside to dive deep into topics and explain why things were working the way they were and if they had any disagreements to always listen to them, and make sure they are heard. Doing this I was able to earn the trust of my team, and deliver quality updates to the teacher (my boss in this situation). 
    * R
        * I was able lead my team into delivering an outstanding final presentation of our project and along the way I was able to teach the skills they were not afluent in yet, such as docker and dbs. We earned a very good grade on the project and impressed the teacher with our presentation. 
    * R
        * Leading this project, sure I learned more about docker and dbs and the tech stack we used, but the more important lesson that I learned was how to lead and take ownership of a project while keeping a team involved and motivated and how to use consistent effective communication to overcome conflicts and earn trust of the team.
    * S
        * So in short, I simply gave people a voice, worked along side them, effectively communicated and dove deep into the whys of the projects to clear confusion and iron out details and overcome conflicts to keep my team motivated and my project at the highest standard.
<hr >

* How have you used data to develop a strategy? 
    * P 
        * I think that data is very important into building a strategy. Data such as, customer feedback and customer behaviour data based off of metrics gathered by products are all extremely important into developing a strategy. Without this youd be blindly creating a project with little to no direction.
    * S
        * In my OS class for our final project, we were tasked to 
    * T
    * A
    * R
    * R
    * S
