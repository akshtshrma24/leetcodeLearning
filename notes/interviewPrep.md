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
    * S
        * New Role as VP of SCE, Club facing declining new sign ups, sent out a survey to members, wanted messaging service and hardware project
        * Thought of the project of messaigng service on website with esp32s and lcd screens where users can see message, and presented to president who approved of the project idea.
    * T
        * I had only 4 weeks left until the end of the semester, needed to take action fast, but there were multiple ways of going about the back end that affect the implementation of the code on the esp32s.
        * I needed to dive deep into each of the possible implementations, and pick the one that is the best combination of easier to implement, simple, matching the skillset of my team and the one that would deliver the best results to the members
        * I also had to institute meetings for development checkups, better communication, to earn the trust of my team and president. 
    * A
        * I dove deep into each of the possible solutions, I focused on the 4 criteria most important to me and I studied each solution weighing their positives and negatives.
        * I then presented my solution to the President and Dev team, who after testing some test files I had written, were excited to begin developing.
        * I delegated work corresponding to their interests, and began working along side them as well, holding workshops for new skills, developing along side them to write the SSE on the back end and the embedded software for the esp32s.
    * R
        * My team and I were able to meet the deadline of the 4 week semester, delivered the quality product to the members who loved the project and my team and president who now trusted me to lead and develop future projects. e
        * Interest was back in the club, and this showed the following semester where our sign ups went from the 2, all the way to 10. 
    * R
        * Doing this project I did learn SSEs and some embedded software, but most importantly I learned how to choose from multiple possible solutions.
        * By diving deep and strategizing off of my parameters I was able to choose the most effective solution out of the several options for my team and members.
<hr >

* When did you take a risk, make a mistake, or fail? How did you respond? How did you grow from it? 
    * S
        * When I was a dev team member during the summer, I noticed slow development due to people not being able to develop outside of the room.
        * People were really only developing 2-3 hours a week due to this. 
    * T
        * I gave myself the task of developing a VPN that the dev team could connect to and develop outside of the room. 
    * A
        * I began development immediately, grabbed a raspberry pi flashed it with openWRT, set up an EC2 instance, put openVPN on it, and tried "connecting" the raspberry pi to it, and was stumped as to why it wouldnt work. 
        * I Spent the next month trying this solution before I realised that I needed to take a step back, and delve deeper into each of the components of my project and make sure each one was working as intended or was correct in general. 
        * I found many errors in my original solution, which I was only able to fix after going back and looking more into my project and expanding my knowledge.
        * I fixed my errors, which were in the ccd file, created a test networking and showed some alumni and president, who after testing my solution further, were excited to implement this into the club.
    * R
        * The club dev team were now able to develop as many hours as they wanted to, averageing now about 6-7 hours a week 3x what they were before. 
    * R
        * I did learn about networking and linux, but what I learned more about that was I learned how to overcome setbacks, and how to effectively respond which was to take steps back and reevaluate the solutions. 
        * Is the solution valid or should I scrap it?
        * I grew into a person that is not afraid of minor setbacks because that just means its an oppurtunity for a major come back. 
<hr >
 
* Describe a time you took the lead on a project. 
    * S
        * In college Teacher assigned groups of 4 at random to create a class project, could be whatever tech stack, it just needed to be dockerized, database, front end, and back end
        * I had a short 5 weeks to do this project, and my group was unfamiliar with docker, most databases, but had skills in front end development. 
    * T
        * I knew that I needed to take the lead of this project, the projects success or failure would be on me, and that I needed to institute meetings with my team to teach skills, develop effectively, and most importantly earn their trust. 
    * A
        * I meet with them once a week for overall development progress and meet with them individually multiple times across the week to develop along side them, check in with them.
        * In the event that conflicts did arise, I would make sure that they were heard, and understood the solution and why it is like that.
    * R
        * I ended up earning the trust of my group, finishing the project on time, and delivering a quality product that the teacher was impressed with. 
    * R
        * Although I did learn a little more on docker, database, and full stack development, the soft skills I picked up were more valuable than this. 
        * I learned how to earn trust of a new team, how to keep members motivated, and ensure consistent quality development through frequent meetings with members. 
<hr >

* What did you do when you needed to motivate a group or promote collaboration on a project? 
    * S
        * During my collegient career, I lead an internship program for 2 years for SJSU where me and another Alumni would interview candidates, allow them to choose a project and then lead them through from start to finish in that project and then deliver a presentation on their work infront of faculty and staff to teach them industry level skills to jumpstart their career.
        * This was my first time running this internship and this summer the project I was leading was development on a video streamer on a raspberry pi using RTMP.
    * T
        * I had 7 interns this summer eager to learn who were extremley motivated at first but their motivation began to dwindle after running into issues, and I needed to keep them motivated and earn their trust as the project lead.
    * A
        * I then understood that I needed to institute frequent meetings with the interns to get them through the bugs, to keep them motivated through the bugs. 
        * I organized workshops to keep my interns learning, I taught them the skills necessary and also sometimes random skills they were just interested in.  
    * R
        * I was able to keep my interns motivated who finished the project and delivered a presentation that impressed the faculty and staff, and my interns also ended up getting internships at major companies one who is interning at amazon.
    * R
        * Managing this internship is an experience I looked forward to every time I did it, it taught me so much on how to keep a group motivated, it taught me that inordor to keep them motivated I needed to keep them first and foremost interested on their bottom line. 
        * Why is the project x good for their interest y and then guide them through project 
<hr >

* How have you used data to develop a strategy? 
    * P 
        * I think that data is very important into building a strategy. Data such as, customer feedback and customer behaviour data based off of metrics gathered by products are all extremely important into developing a strategy. Without this youd be blindly creating a project with little to no direction.
    * S
        * In my time as a student at SJSU, I was very involved with the faculty and staff, I would always ask them questions and be eager to develop software they needed. One of my professors knew that I was very into networking, so he gave me the task of figuring out why the subnet in his office, he had his own little network in there, kept dying at what he thought was random times, with a 15% total daily downtime.
    * T
        * My task was to first gather data on his wifi network answering several questions such as, when does it go down and when do the speeds decrease. Then I could use this data to further investigate at the times at which they went down to figure out why it kept going down to offer a solution.
    * A
        * I planned out a C program to gather network speeds and uptime which was then monitored using metrics exported by the C program such as upload, download speeds, and ping. I then found out that at around the same time at every day the internet would just cut off. I used this data to investigate further and found that IT had been automatically temporarily cutting internet to that room due to the the router raising flags. I then developed a strategy to propose a new check for the IT department to use that would not cut internet. I proposed the solution to the professor who approved it and then presented it to IT, who then after further testing implemented it.  
    * R
        * Thanks to the C program and the data that I had gathered data to filter out possible causes, I was able to narrow down the root causes and develop a strategy based off of that. I then presented the solution after testing and approval from the professor to the IT department who implemented it and then were able to bring my professors up time to match the networks uptime of SJSU.
    * R
        * doing this project for my professor it taught me that being curious and gathering data to develop a strategy is an effective way about going things. 
    * S
        * 

<hr> 

* Why Amazon? 
    * I beilieve that Amazon is great for people who like to build things. This is me. I love to build things, showcase my enginuinity, and take ownership of the code that I write. Ive heard from people who work there that this is company is not for people who want to sit back and let things happen. I want to make things happen. I want to meet smart people, I want to learn new things, I want to build new things with impact, and I believe amazon will give me the perfect platform to do this. 





<hr>


* Taking initiative 
* Learning and Being curious 
* Backbone - Disagreeing and committing 
* Ownership


Introduction
Im Akshit Sharma, you can call me Andy
I graduated frmo San Jose State University in May 2024, but I have been working part time as a software engineer at Windsor, its a small care facility. 
I lead the Engineer Club Last Year as the Vice President, and year prior as Dev Team Lead, during my leadership there I initiated number of projects like RFID door sensors and messaging with esp32s and creating a Road Warrior VPN.
I mentored interns for a Summer internship program by SJSU to teach industry relevant skills such as embedded software and networking. I collobrated on and posted multiple videos regarding docker, and vpns on youtube. 
As part time/hobby on the weekends I cut hair, and people think im good. 
And I can also talk about my past projects in details if youd like. 