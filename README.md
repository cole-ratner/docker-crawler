#Docker Crawler# <br> 
Docker crawler is a linux cli tool that enables users to search for publicly registered container images and then perform concurrent enumeration opertions on the discovered images.<br> 

**Build** <br>
*run:* **scripts/build**

This will build a binary called dockercrawler in the bin directory.
It will also build a docker container for the mockworker tagged as "cole-ratner/mockworker".

**Usage** <br>
*run:* **bin/dockercrawler args...** <br>
<br>
  **-h** *string* "The container registry that you would like to search. (default "hub.docker.com")" <br>
  **-searchterm** *string* "The term that you would like to search within the container registry. (Required)" <br>
  **-w** *string* "The name of the worker app to use for enumerating the container. (Required)" <br>

  Please note: dockercrawler is by default configured to write enumeration logs to /home/data/output/