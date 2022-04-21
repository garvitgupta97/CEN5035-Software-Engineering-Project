# Sprint 3

## Project - Overview

A social media platform for students to collaborate and help each other by sharing job openings, informative articles and interview experiences. The application allows students and job seekers to maintain a list of the their job applications privately. Using the list users can easily track the post and companies they have applied to and share it with other users present in their network if they want. Users can also share other types of posts like articles or interview experiences. These post will be seggregated based on different job positions and companies, that will allow users to apply various filters so that they can find the exact type of content they are looking for.

## Useful links of the project
- [Student Universe Repo Link](https://github.com/garvitgupta97/CEN5035-Software-Engineering-Project/)
- [Sprint 1 User stories progress board link](https://github.com/garvitgupta97/CEN5035-Software-Engineering-Project/projects/2)
- [All user stories link](https://github.com/garvitgupta97/CEN5035-Software-Engineering-Project/issues)

## Frontend Tasks achieved - (React js)

- Integrated new post api
- Integrated get all post api
- Added unit test cases for various pages
- Changed frontend layout

## Backend Tasks achieved - (Go)
- Created Post API
- Created get all posts api 
- Added unit tests for various apis

## Demo of the app

### Home Page Demo

  ![This is an image](https://raw.githubusercontent.com/garvitgupta97/CEN5035-Software-Engineering-Project/main/Resources/HomePageDemo.png)


### Logged In Home Page Demo

  ![This is an image](https://raw.githubusercontent.com/garvitgupta97/CEN5035-Software-Engineering-Project/main/Resources/LoggedInHomePageDemo.png)


### Log in Page Demo

  ![This is an image](https://raw.githubusercontent.com/garvitgupta97/CEN5035-Software-Engineering-Project/main/Resources/LoginPageDemo.png)

### New Post on Home Page Demo

  ![This is an image](https://raw.githubusercontent.com/garvitgupta97/CEN5035-Software-Engineering-Project/main/Resources/NewPostOnHomePageDemo.png)

### Post Page Demo

  ![This is an image](https://raw.githubusercontent.com/garvitgupta97/CEN5035-Software-Engineering-Project/main/Resources/PostPageDemo.png)

### Register Page Demo

  ![This is an image](https://raw.githubusercontent.com/garvitgupta97/CEN5035-Software-Engineering-Project/main/Resources/RegisterPageDemo.png)

## Api documentation of backend services
### Api for user sign up

- [http://localhost:8080/api/signup/] (http://localhost:8080/api/signup/)

### Fail Case

  ![This is an image](https://github.com/garvitgupta97/CEN5035-Software-Engineering-Project/blob/5f935876f1f276d8eea001294df14f5c4760e650/Resources/signin_fail.png)
  
### Success Case

  ![This is an image](https://github.com/garvitgupta97/CEN5035-Software-Engineering-Project/blob/5f935876f1f276d8eea001294df14f5c4760e650/Resources/signup_successful.png)
  
  
### Api for user sign in

- [http://localhost:8080/api/signin/](http://localhost:8080/api/signin/)

### Fail Case

  ![This is an image](https://github.com/garvitgupta97/CEN5035-Software-Engineering-Project/blob/main/Resources/signin_fail.png)
  
### Success Case

  ![This is an image](https://github.com/garvitgupta97/CEN5035-Software-Engineering-Project/blob/main/Resources/signin_successful.png)


### Api to fetch all the user's emails

- [api/getUsers/](api/getUsers/)

![This is an image](https://github.com/garvitgupta97/CEN5035-Software-Engineering-Project/blob/main/Resources/getUsers.png)

### API to upsert user profile
- [api/getUsers/](api/updateProfile/)
![This is an image](https://github.com/garvitgupta97/CEN5035-Software-Engineering-Project/blob/main/Resources/updateProfile.png)

### API to create a new post
- [api/getUsers/](api/post/create/)
![This is an image](https://github.com/garvitgupta97/CEN5035-Software-Engineering-Project/blob/main/Resources/createPost.png)


### API to create a fetch all posts
- [api/getUsers/](api/post/allPosts/)
![This is an image](https://github.com/garvitgupta97/CEN5035-Software-Engineering-Project/blob/main/Resources/allPosts.png)


### API to delete a post
- [api/post/delete/](api/post/delete/)
![This is an image](https://github.com/garvitgupta97/CEN5035-Software-Engineering-Project/blob/be_userPostChanges/Resources/deletePost.png)

