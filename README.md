# Info Harvester/API broker app in Go. 

Still a work in progress, but at some point the app will be:
- getting info from one API endpoint
- storing it in a Postgres DB
- posting it to a dashboard on another API endpoint.

The end goal would be a customer-facing CI/CD pipeline vizualizer, that will show relevant info to different teams such as - last successful build, available deploy enviornments, plans in progress, last failed build/deploy, etc.