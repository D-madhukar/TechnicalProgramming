const mongoose=require("mongoose");
var Schema=mongoose.Schema;
var UserSchema=new Schema({
  username:{type:String,required:true,default:"abc@gamil.com"},
  email:{type:String,required:true},
  phonenumber:{type:Number,require:true},
  address:{
    city:String,
    street:String,
    houseno:String,
    pincode:Number
  }
})
module.exports=mongoose.model("User",UserSchema);
