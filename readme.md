Q1. How to use your program?

    Step 1 . Create docker image from the provided DockerFile running command
             docker build -t camunda-autoscalar .

    step 2 . Once image is created, we can use the image to create a "camunda-autoscalar" deployments in Kubernetes cluster
          kubectl apply -f camunda-autoscalar-deployment.yaml


Q2.  How long did it take you to solve the exercise?
     
     It took me 12-15 hours to complete the full exercise. I gave around 2 hours every day after work.


Q3.  Which additional steps would you take in order to make this code production ready? Why?
   
     N/A

Q4.  Which step took most of the time? Why?
    
     Coding the logic in Golang, took me most time. To be specfic, the step to unmarshall json to extract process count took me most time.
     Learning and implementing this step was most cumbersome after work :(


**Comment:  It would be great if the best solution can be shared with me.

