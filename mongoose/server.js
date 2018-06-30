var express=require("express");
var mongoose=require("mongoose");
var bodyParser=require("body-parser");
var User=require("./models/User.js");
var app=express();

//Connecting to the MongoDb
mongoose.connect("mongodb://localhost/MyDatab");

//Writig the routes
app.get("/insertdata",(req,res)=>{
  var user=new User({
    username:"madhukar",
    email:"madhukarpateld@gmail.com",
    phonenumber:837676230
  })
  user.save((error)=>{
    if(error)
      res.status(500).send("Server Error");
    else {
      res.status(200).send("inserted succesfully");
    }
  })
})
app.get("/users",(req,res)=>{
  User.find({})
  .exec((error,users)=>{
    if(error)
      res.status(500).send("Server error");
    else {
      res.json(users);
    }
  })
})

//Listening at the port 9999
app.listen(8088,function(){
  console.log("Server listenig at port 8088");
})
