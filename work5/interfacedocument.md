# 项目分层

项目采用了mvc架构

另：项目的所有SQL语句都是原生

##### dao:

sqlinit.go (初始化db) 

##### controller: 

coursecontroller.go  (课程的控制器)

usercontroller.go(学生的控制器)

user_coursecontrooler.go(选课的控制器)

##### models:

courses.go (课程模型和数据库操作)

users.go(学生模型和数据库操作)

users_courses.go(选课模型和数据库操作)

##### router:

router.go(三个路由组，一共11个接口)

##### view:

users.go(学生的视觉层)

courses.go(课程的视觉层)

users_courses.go(选课系统的视觉层)

##### main.go

##### go.mod

## 接口：

###### post /user/

返回值：1.学号不为10位，返回

```
"code":200,
"msg": "学号不为10位",
```

2.密码小于6位，返回

```
"code": 200,
"msg":  "密码不能少于6位",
```

3.用户名为空，返回

```
"code": 200,
"msg": "用户名不能为空",
```

4.最大学分小于等于0，返回

```
"code": 200,
"msg": "最大学分不能小于1",
```

5.调用接口成功，返回

```
c.JSON(http.StatusOK,gin.H{
   "code": 200,
   "msg": "注册成功",
   "username": username,
   "password": password,
   "number": number,
   "maxcredit": maxcredit,  //用户名，密码，学号和最大学分数
})
```

6.调用接口失败，返回

```
"code" : 404,
"msg": "注册失败",
```

###### get /user/

返回值：1.学号不为10位，返回

```
"code":200,
"msg": "学号不为10位",
```

2.密码小于6位，返回

```
"code": 200,
"msg":  "密码不能少于6位",
```

3.调用接口成功，返回

```
"code" : 200,
"msg":  "登陆成功",
"username": username,
"password": password,
"telephone": number, //账号，密码，学号
```

4.调用接口失败，返回

```
"code": 404,
"msg": "登陆失败",
```

###### put /user/

返回值：1.新密码与旧密码一致，返回

```
 c.JSON(http.StatusOK,gin.H{
          "code": 200,
          "msg": "新密码不能与原密码一致",
})
```

2.接口调用成功

```
c.JSON(http.StatusOK,gin.H{
       "code":200,
       "msg": "修改密码成功",
       "username": username,
       "newpassword": password,
       "number": number,
})
```

3.接口调用失败

```
    c.JSON(http.StatusNotFound,gin.H{
      "code" : 404,
      "msg": "修改密码失败",
})
```

###### delete /user/

1.接口调用成功

```
c.JSON(http.StatusOK,gin.H{
   "code" :200,
   "msg": "删除名为:"+username+"的用户成功",
})
```

2.接口调用失败

```
c.JSON(http.StatusNotFound,gin.H{
   "code" : 404,
   "msg": "删除用户失败",
})
```



###### post/course

1.课程名为空

```
    c.JSON(http.StatusOK,gin.H{
      "code": 200,
      "msg": "课程名不能为空",
})
```

2.学分小于等于0

```
 c.JSON(http.StatusOK,gin.H{
      "code":200 ,
      "msg": "学分不能小于1",
})
```

3.最大选课人数小于1

```
c.JSON(http.StatusOK,gin.H{
      "code":200,
      "msg": "最大选课人数不能小于1",
})
```

4.调用接口成功

```
c.JSON(http.StatusOK,gin.H{
   "code":200,
   "msg": "增加课程成功",
   "coursename": CourseName,
   "credit": Credit,
   "maxnumber":MaxNumber,  //课程名 学分和最大选课人数
})
```

5.调用接口失败

```
"code":404,
"msg": "增加课程失败",
```

###### get/course

1.调用接口成功

```
c.JSON(http.StatusOK,gin.H{
   "code":200,
   "msg": "查询课程成功",
   "coursename": CourseName,
   "credit": Credit,
   "maxnumber":MaxNumber,
}) 
```

2.调用接口失败

```
c.JSON(http.StatusNotFound,gin.H{
   "code": 404,
   "msg": "查询失败",
})
```

###### put/course

1.调用接口成功

```
c.JSON(http.StatusOK,gin.H{
   "code": 200,
   "msg": "修改课程成功",
   "coursename": CourseName,
   "newcredit": Credit,
   "newmaxnumber": MaxNumber,
})
```

2.调用接口失败

```
c.JSON(http.StatusNotFound,gin.H{
      "code": 404,
      "msg": "修改课程信息失败",
})
```

###### delete/course

1.调用接口成功

```
c.JSON(http.StatusOK,gin.H{
   "code":200,
   "msg": "删除名为:"+CourseName+"的课程成功",
})
```

2.调用接口失败

```
c.JSON(http.StatusNotFound,gin.H{
      "code": 404,
      "msg": "删除课程信息失败",
})
```

###### post/userCourse

1.调用接口成功

```
c.JSON(http.StatusOK,gin.H{
   "code": 200,
   "msg": "选课成功",
   "username": Username,
   "coursename": Coursename,
   "number": Number, //学生姓名，课程名，和学号
}) 
```

2.调用接口失败

```
c.JSON(http.StatusNotFound,gin.H{
   "code": 404,
   "msg": "选课失败",
})
```

###### get/userCourse

1.调用接口成功

```
c.JSON(http.StatusOK,gin.H{
   "code": 200,
   "msg": "查询信息成功",
   "number": Number,
   "course": u, //学号与所选课程
})
```

2.调用接口失败

```
c.JSON(http.StatusNotFound,gin.H{
   "code": 404,
   "msg": "查询学生选课信息失败",
})
```



###### delete/userCourse

1.调用接口成功

```
c.JSON(http.StatusOK,gin.H{
   "code": 200,
   "msg": "删除学生选课信息成功",
})
```

2.调用接口失败

```
c.JSON(http.StatusNotFound,gin.H{
   "code": 404,
   "msg": "删除学生选课信息失败",
})
```