# GoLivyRestAPI
This program has function to run spark apps via livy using golang

<br />

## Prerequisites
### Install & Run Apache Livy
* Must have install docker-ce on yout system
* Download zip package apache livy from the server https://livy.apache.org/download/
<pre>wget https://www-us.apache.org/dist/incubator/livy/0.6.0-incubating/apache-livy-0.6.0-incubating-bin.zi</pre>
* Unzip & move to path your system
<pre>unzip apache-livy-0.6.0-incubating-bin.zip -d /opt/</pre>
* Setting config livy
<pre>vim /opt/apache-livy-0.6.0-incubating-bin/conf/livy.cof</pre>
* And you will see like this
<pre>
.
.
# List of local directories from where files are allowed to be added to user sessions. By
# default it's empty, meaning users can only reference remote URIs when starting their
# sessions.
livy.file.local-dir-whitelist = /home/reja/ #add this line for path of jar spark apps 
.
.
</pre>


## How to compile this project
<pre>docker build --rm -t golivyrest:v1 . </pre>

<br />

## How to run the program

<pre>
docker run -d -p 1212:1212 -v /home/reja/goproject/GoLivyRestAPI/config:/config golivyrest:v1
</pre>

## Example test
<pre>curl -XPOST localhost:1212/runlivy -d '{"table":"winevariety","zooKeeper":"master.research.ph,datanode1.research.ph,datanode2.research.ph","hbaseMaster":"master.research.ph","pathCSV":"file:///home/reja/IdeaProjects/SparkBatchHbase/zomato.csv"}'</pre>